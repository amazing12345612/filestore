package mq

import (
	"testing"
)

func TestMain(t *testing.T) {
	initChannel()
	// err := Publish("test", "", []byte("hello world"))
}
