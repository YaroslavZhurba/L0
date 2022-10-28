package store

import (
	"database/sql"
	"fmt"
	"restapi/src/model"
)

// UserRepository ...
// add here cash
//
 type OrderRepository struct {
 	store *Store
	cash map[int]*model.Order
 }

 // Create ...
 func (r *OrderRepository) Create() (*model.Order, error) {
	orderDefault := model.OrderDefault()
	o := &model.Order{}
 	if err := r.store.db.QueryRow(
 		`insert into Orders (id, "order") values ($1, $2)`,
		 orderDefault.Id, orderDefault.OrderJson,
 	).Scan(&o.Id); err != nil {
 		return nil, err
 	}

 	return o, nil
 }

  // Add ...
  func (r *OrderRepository) Add(order *model.Order) (error) {
 	if _, err := r.store.db.Exec(
 		`insert into Orders (id, "order") values ($1, $2)`,
		 order.Id, order.OrderJson,
 	); err != nil {
 		return err
 	}

 	return nil
 }

  // DeleteById ...
  func (r *OrderRepository) DeleteById(id int) (sql.Result, error) {
	var result sql.Result
	if result, err := r.store.db.Exec(
	   `delete from Orderds
	   where id= ($1) `,
	   id,
	); err != nil {
		return result, err
	}
	return result, nil
}

 // FindById ...
 func (r *OrderRepository) FindById(id int) (*model.Order, error) {
 	o := &model.Order{}
	if val, ok := r.cash[id]; ok {
		fmt.Print("Id = ")
		fmt.Print(id)
		fmt.Println(" is taken from cash")
		return val, nil
	}
 	if err := r.store.db.QueryRow(
		`select id, "order" from Orders
		where id = ($1) `,
		id,
 	).Scan(
		&o.Id,
		&o.OrderJson); err != nil {
 		return nil, err
 	}
	r.cash[id] = o

 	return o, nil
 }