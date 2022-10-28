package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"

	"restapi/src/model"
	"restapi/src/store"
)

// APIServer ...
 type APIServer struct {
 	config *Config
 	logger *logrus.Logger
 	router *mux.Router
	store  *store.Store
	sc     stan.Conn
	subs   map[string]stan.Subscription
 }

 // New ...
 func New(config *Config) *APIServer {
 	return &APIServer{
 		config: config,
 		logger: logrus.New(),
 		router: mux.NewRouter(),
 	}
 }

 // Start ...
 func (s *APIServer) Start() error {
 	if err := s.configureLogger(); err != nil {
 		return err
 	}

 	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	if err := s.configureStan(); err != nil {
		return err
	}

	s.subs = make(map[string]stan.Subscription)

 	s.logger.Info("starting api server")

 	return http.ListenAndServe(s.config.BindAddr, s.router)
 }

 func (s *APIServer) configureLogger() error {
 	level, err := logrus.ParseLevel(s.config.LogLevel)
 	if err != nil {
 		return err
 	}

 	s.logger.SetLevel(level)

 	return nil
 }

 func (s *APIServer) configureRouter() {
 	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/bye", s.handleBye())
	s.router.HandleFunc("/json", s.handleJson())
 }


 func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) configureStan() error {
	var err error
	s.sc, err = stan.Connect("test-cluster", "sub-1")
	if err != nil {
		return err
	}
	return nil
}

 func (s *APIServer) handleHello() http.HandlerFunc {
 	return func(w http.ResponseWriter, r *http.Request) {
		o, err := s.store.Order().Create()
		if err != nil {
			io.WriteString(w, "Hello, error :(")
		} else {
			io.WriteString(w, "Order: Id = " +
			 strconv.Itoa(o.Id))
		}
 	}
 }

 func (s *APIServer) handleBye() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := s.store.Order().FindById(1)
		if err != nil {
			io.WriteString(w, "Bye, error :(")
		} else {
			str := fmt. Sprintf("%#v", o.OrderJson)
			io.WriteString(w, "Order: Id = " +
			 strconv.Itoa(o.Id) + 
			 ";  " + str)
		}
	}
}

func (s *APIServer) handleJson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			s.logger.Info(err)
		}
		o, err := s.store.Order().FindById(id)
		if err != nil {
			io.WriteString(w, "No element :(")
		} else {
			str := fmt. Sprintf("%#v", o.OrderJson)
			io.WriteString(w, "Order: Id = " +
			 strconv.Itoa(o.Id) + 
			 ";  " + str)
		}
	}
}


func (s *APIServer) SubscribeChannel(name string) (string, error) {
	if _, ok := s.subs[name]; ok {
		return "Channel " + name + " is already exist", nil
	}
	
	sub, err := s.sc.Subscribe(name, func(m *stan.Msg) {
		var order model.Order
		if err := json.Unmarshal([]byte(m.Data), &order); err != nil {
			s.logger.Info("Channel " + name + " error:")
			s.logger.Info(err)
			return
		}
		err := s.store.Order().Add(&order)
		if err != nil {
			s.logger.Info("Channel " + name + " error:")
			s.logger.Info(err)
			return
		} else {
			s.logger.Info("Channel " + name + " added DATA successfuly!")
			s.logger.Info("Added id = " + strconv.Itoa(order.Id))
		}
	})
	s.subs[name] = sub

	return "Channel " + name + " added seccessfuly!", err
}

func (s *APIServer) UnsubscribeChannel(name string) (string, error) {
	if _, ok := s.subs[name]; !ok {
		return "Channel " + name + " is deleted seccessfuly!", nil
	}
	s.subs[name].Unsubscribe()
	delete(s.subs, name)
	return "Channel " + name + " is deleted seccessfuly!", nil
}