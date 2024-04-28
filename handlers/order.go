package handlers

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Order created")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Orders listed")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
}
