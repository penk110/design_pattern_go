package behavior_observer

import (
	"errors"
	"fmt"
	"log"
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
	Notify(msg *Msg)
	Stop()
	loop()
}

type ObserverImpl interface {
	Exec(msg *Msg)
	GetID() string
}

type Observer struct {
	id string
}

func (o *Observer) GetID() string {
	return o.id
}

func (o *Observer) Exec(msg *Msg) {
	fmt.Println("---------exec id: " + o.id + " msg: " + msg.Msg + "----------")
}

type Subject struct {
	mux      sync.Mutex
	os       map[string]ObserverImpl
	msgCh    chan *Msg
	cancelCh chan struct{}
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

func (s *Subject) Notify(msg *Msg) {
	s.msgCh <- msg
}

func (s *Subject) loop() {
	for {
		select {
		case m := <-s.msgCh:
			go func(msg *Msg) {
				if !s.IExist(msg.ObserverID) {
					log.Println("id: " + msg.ObserverID + " not exist.")
				}
				s.os[msg.ObserverID].Exec(msg)
			}(m)
		case <-s.cancelCh:
			return
		}
	}
}

func (s *Subject) Stop() {
	s.cancelCh <- struct{}{}
}
