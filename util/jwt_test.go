package util

import (
	"testing"
)

func TestJwtGenerate(t *testing.T) {
	a := "test1"
	b := "test2"
	atoken,_ := GenerateToken(a,b)
	c,d,_ := GenToken2(a,b)
	// fmt.Println(atoken, rtoken)
	t.Log(atoken)
	t.Log(c,d)
	t.Log("C")
}
