package behavior_observer_eventbus

import (
	"errors"
	"log"
	"reflect"
	"sync"
)

type BusImpl interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlerMp map[string][]reflect.Value
	mux       sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlerMp: map[string][]reflect.Value{},
		mux:       sync.Mutex{},
	}
}

// Subscribe 订阅
func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.mux.Lock()
	defer bus.mux.Unlock()
	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return errors.New("handler is not func")
	}
	handlers, ok := bus.handlerMp[topic]
	if !ok {
		handlers = []reflect.Value{}
	}
	handlers = append(handlers, v)
	bus.handlerMp[topic] = handlers
	return nil
}

// Publish 发布
// 这里异步执行，并且不会等待返回结果
func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlerMp[topic]
	if !ok {
		log.Println("not found handlers in topic: ", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	// params
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}
	// Call func
	for i := range handlers {
		// 订阅者收到主题，进行消费
		go handlers[i].Call(params)
	}
}
