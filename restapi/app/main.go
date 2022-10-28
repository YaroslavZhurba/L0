package main

import (
	"flag"
	"fmt"
	"log"
	"restapi/src/apiserver"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	// flag.StringVar(&configPath, "config-path", "../configs/apiserver.toml", "path to config file")
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

var s *apiserver.APIServer

func startServer() {
	flag.Parse()

 	config := apiserver.NewConfig()
 	_, err := toml.DecodeFile(configPath, config)
 	if err != nil {
 		log.Fatal(err)
 	}
	s = apiserver.New(config)
 	if err := s.Start(); err != nil {
 		log.Fatal(err)
 	}
}

func main() {
 	go startServer()
	time.Sleep(5 * time.Second)
	// if ss, err := s.SubscribeChannel("bar"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(ss)
	// }
	for ;; {
		var str string 
		var name string
		fmt.Scanln(&str)
		if str == "exit" {
			break
		}
		if str == "sub" {
			fmt.Scanln(&name)
			if ss, err := s.SubscribeChannel(name); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ss)
			}
		}
		if str == "unsub" {
			fmt.Scanln(&name)
			if ss, err := s.UnsubscribeChannel(name); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ss)
			}
		}
	}
}