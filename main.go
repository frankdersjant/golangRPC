package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

type Customer struct {
	FirstName string
	LastName  string
}

var CustomerDatabase []Customer

// Note that the method is with a CAPITOL LETTER!
func (a *API) GetCustomerDB(empty string, reply *[]Customer) error {
	*reply = CustomerDatabase
	return nil
}

// Note that the method is with a CAPITOL LETTER!
func (a *API) AddCustomer(addedCustomer Customer, reply *Customer) error {

	//do some stuff with a db

	fmt.Println("AddCustomer")
	return nil
}

func main() {

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
