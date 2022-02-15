package behavior_observer

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type Msg struct {
	ObserverID string
	Msg        string
}

type SubjectImpl interface {
	Register(observer ObserverImpl) error
	Remove(ID string) error
	IExist(id string) bool
	Notify(ctx context.Context, msg *Msg) error
}

type ObserverImpl interface {
	Exec(ctx context.Context, msg *Msg)
	GetID() string
}

type Observer struct {
	id string
}

func (o *Observer) GetID() string {
	return o.id
}

func (o *Observer) Exec(ctx context.Context, msg *Msg) {
	fmt.Println("---------exec id: " + o.id + " msg: " + msg.Msg + "----------")
	ctx.Done()
}

type Subject struct {
	mux sync.Mutex
	os  map[string]ObserverImpl
}

func (s *Subject) Register(observer ObserverImpl) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.IExist(observer.GetID()) {
		return errors.New("id: " + observer.GetID() + " already exist.")
	}
	s.os[observer.GetID()] = observer
	return nil
}

func (s *Subject) Remove(id string) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if !s.IExist(id) {
		return errors.New("id: " + id + " not exist.")
	}
	delete(s.os, id)
	return nil
}

func (s *Subject) IExist(id string) bool {
	if _, ok := s.os[id]; !ok {
		return false
	}
	return true
}

func (s *Subject) Notify(ctx context.Context, msg *Msg) error {
	if !s.IExist(msg.ObserverID) {
		return errors.New("id: " + msg.ObserverID + " not exist.")
	}
	s.os[msg.ObserverID].Exec(ctx, msg)
	return nil
}
