package behavior_observer

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"testing"
	"time"
)

func TestSubject_Notify(t *testing.T) {
	var (
		sub      SubjectImpl
		observer ObserverImpl
		err      error
	)
	sub = &Subject{
		mux:      sync.Mutex{},
		os:       map[string]ObserverImpl{},
		cancelCh: make(chan struct{}),
		msgCh:    make(chan *Msg, 1024),
	}
	go sub.loop()

	for i := 1; i < 10; i++ {
		observer = &Observer{
			id: uuid.NewString(),
		}
		if err = sub.Register(observer); err != nil {
			t.Errorf("err: %v", err)
		}
		sub.Notify(&Msg{
			ObserverID: observer.GetID(),
			Msg:        fmt.Sprintf("%v", i),
		})
	}

	time.Sleep(2 * time.Second)
	sub.Stop()
}
