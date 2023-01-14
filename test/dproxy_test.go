package test

import (
	"filestore/service/dbproxy/client"
	"fmt"
	"testing"
)

func TestDbproxy(t *testing.T) {
	rep, err := client.GetUserInfo("admin123")
	if err != nil {
		fmt.Println(rep)
	}
}
