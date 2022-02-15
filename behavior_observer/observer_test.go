package behavior_observer

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"testing"
	"time"
)

func TestSubject_Notify(t *testing.T) {
	var (
		sub  SubjectImpl
		ober ObserverImpl
		err  error
	)
	sub = &Subject{
		mux: sync.Mutex{},
		os:  map[string]ObserverImpl{},
	}
	ober = &Observer{
		id: uuid.NewString(),
	}
	err = sub.Register(ober)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_ = sub.Notify(context.TODO(), &Msg{
		ObserverID: ober.GetID(),
		Msg:        time.Now().String(),
	})
	t.Log("success")
}
