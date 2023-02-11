package util

import (
	"fmt"
	"testing"
)

func TestSha(m *testing.T) {
	t := Sha1([]byte("abcdg"))
	fmt.Println(t)
	a := Sha1Stream{}
	a.Update([]byte("abcdg"))
	fmt.Println(a.Sum())

}
