package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type SMSSenderAPI struct {
}

func NewSMSSenderAPI() *SMSSenderAPI {
	return &SMSSenderAPI{}
}

func (receiver *SMSSenderAPI) SendMessage(params map[string]any) {
	fmt.Println("Sending message:", params["message"].(string))
}

type Adapter interface {
	Send(params map[string]any)
}

type AdapterImplementation struct {
	sender *SMSSenderAPI
}

func NewAdapterImplementation(sender *SMSSenderAPI) Adapter {
	return &AdapterImplementation{sender}
}

func (receiver *AdapterImplementation) Send(params map[string]any) {
	receiver.sender.SendMessage(params)
}

func main() {
	smsSenderAPI := NewSMSSenderAPI()

	adapterImplementation := NewAdapterImplementation(smsSenderAPI)

	params := map[string]any{
		"to":      71000000001,
		"message": "Hello world",
		"sender":  "Tester",
	}

	adapterImplementation.Send(params)
}

/*
Sending message: Hello world
*/
