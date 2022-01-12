package structure_bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorNotification(t *testing.T) {
	sender := NewEmailSender([]string{"tester@gemail.com", "golang@gemil.com"})

	notify := NewErrorNotification(sender)
	err := notify.Notify("TestNewErrorNotification", nil)

	assert.Nil(t, err)
}

func TestNewEmailSender(t *testing.T) {

}
