package test

import (
	"time"
)

type OrderPlaced struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	Testttttttttttttttttttttttt string `json:"testttttttttttttttttttttttt" type:"string"` 
	Name string `json:"name" type:"string"` 
	Price string `json:"price" type:"string"` 
	Qty string `json:"qty" type:"string"` 
	
}

func NewOrderPlaced() *OrderPlaced{
	event := &OrderPlaced{EventType:"OrderPlaced", TimeStamp:time.Now().String()}

	return event
}
