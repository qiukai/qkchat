package id

import (
	"fmt"
	"testing"
	"time"
)

func TestNext(t *testing.T) {

	for i := 0; i < 1000; i++ {
		go next1()
		time.Sleep(1000000)
	}

}

func next1() {
	s := Next()
	fmt.Println(s)

}
