package event

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Operator int

const (
	Plus Operator = iota
	Minus
	Multiply
)

func (o Operator) String() string {
	return [...]string{"+", "-", "*"}[o]
}

type Event struct {
	Operator Operator `json:"operator"`
	Value    int      `json:"value"`
}

type EventStore []Event

func (s *EventStore) Push(e Event) {
	*s = append(*s, e)
}

func (s *EventStore) Load(f string) {
	d, _ := os.ReadFile(f)
	json.Unmarshal(d, s)
}

func (s *EventStore) Save(f string) {
	d, _ := json.Marshal(s)
	os.WriteFile(f, d, 0644)
}

func (s *EventStore) String() string {
	var b strings.Builder
	sep := ""
	for _, e := range *s {
		b.WriteString(fmt.Sprintf("%s%s %d", sep, e.Operator.String(), e.Value))
		sep = "\n"
	}
	return b.String()
}

type EventBuilder struct {
	event Event
}

func NewEventBuilder() *EventBuilder {
	return &EventBuilder{}
}

func (b *EventBuilder) Operator(o string) *EventBuilder {
	switch o {
	case "plus":
		b.event.Operator = Plus
	case "minus":
		b.event.Operator = Minus
	case "multiply":
		b.event.Operator = Multiply
	}
	return b
}

func (b *EventBuilder) Value(v int) *EventBuilder {
	b.event.Value = v
	return b
}

func (b *EventBuilder) Build() Event {
	return b.event
}
