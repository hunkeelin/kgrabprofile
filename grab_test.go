package kgrabprofile

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println("testing")
	_, err := GrabProfile("varoinfra")
	if err != nil {
		panic(err)
	}
}
