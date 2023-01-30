package main

import(
	"log"
	"test/test"
	"test/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	test.OrderDBInit()
	go test.InitProducer(config.GetMode())
	e := test.RouteInit()

	e.Logger.Fatal(e.Start(":8081"))
}
