package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	panic("not implemented") // TODO: Implement
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	panic("not implemented") // TODO: Implement
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	panic("not implemented") // TODO: Implement
}

func (ed *EventDispatcher) Clear() error {
	panic("not implemented") // TODO: Implement
}
