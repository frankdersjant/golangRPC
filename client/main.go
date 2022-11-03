package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Customer struct {
	FirstName string
	LastName  string
}

func main() {
	var reply Customer
	//	var db []Customer

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	Bart := Customer{"First Name", "Last Name"}

	err = client.Call("API.AddCustomer", Bart, &reply)
	if err != nil {
		log.Fatal("AddCustomer error: ", err)
	}
	//client.Call("API.GetCustomerDB", Bart, &db)

	fmt.Println("Called addedCustomer: ", Bart)

	//client.Call("API.EditItem", Item{"Second", "A new second item"}, &reply)

	//client.Call("API.GetDB", "", &db)
	//fmt.Println("Database: ", db)

	//client.Call("API.GetByName", "First", &reply)
	//fmt.Println("first item: ", reply)

}
