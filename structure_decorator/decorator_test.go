package structure_decorator

import (
	"log"
	"testing"
)

func TestNewPistol(t *testing.T) {

	pistol := NewPistol(&Gun{}, func(options ...interface{}) {
		log.Printf("增强功能\n")
	})

	pistol.Fire()
}
