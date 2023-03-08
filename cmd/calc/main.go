package main

import (
	"fmt"
	"github.com/aaronthangnguyen/go-calc"
)

func main() {
	f := "store.txt"
	store := event.EventStore{}
	store.Load(f)
	event1 := event.
		NewEventBuilder().
		Operator("plus").
		Value(1).
		Build()
	store.Push(event1)
	event2 := event.
		NewEventBuilder().
		Operator("multiply").
		Value(10).
		Build()
	store.Push(event2)
	fmt.Println(store.String())
	store.Save(f)

}