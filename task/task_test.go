package task

import (
	"fmt"
	"testing"
	"time"
)

func a(i int) (int, int) {
	return i + 1, i + 2
}

func b(s string) {
	fmt.Println("hello", s)
}

type P struct {
	name string
}

func (p *P) Greeting(flag string) {
	fmt.Println("Hello", flag, "I am", p.name)
}

func TestTask(t *testing.T) {

	c := func(ss []string) string {
		return ss[0] + "+" + ss[1]
	}

	// ---------------------------
	//
	// Define Task & Parameters & Returns

	cP4a, cR4a := make(chan interface{}), make(chan interface{})
	Task(a, cP4a, cR4a)

	cP4b, cR4b := make(chan interface{}), make(chan interface{})
	Task(b, cP4b, cR4b)

	cP4p, cR4p := make(chan interface{}), make(chan interface{})
	p := P{name: "Tick"}
	Task(p.Greeting, cP4p, cR4p)

	cP4c, cR4c := make(chan interface{}), make(chan interface{})
	Task(c, cP4c, cR4c)

	// ---------------------------
	//
	// What to do if Task has handled task

	go func() {
		for {
			select {
			case ra := <-cR4a:
				fmt.Println(ra)
			case rb := <-cR4b:
				fmt.Println(rb)
			case rp := <-cR4p:
				fmt.Println(rp)
			case rc := <-cR4c:
				fmt.Println(rc)

				// default:
				// 	time.Sleep(1000 * time.Millisecond)
				// 	fmt.Println("waiting...")
			}
		}
	}()

	// ---------------------------
	//
	// Once Params have been filled, Task will do it

	time.Sleep(1 * time.Second)
	cP4a <- 5

	time.Sleep(1 * time.Second)
	cP4c <- []string{"1", "2"}

	time.Sleep(2 * time.Second)
	cP4b <- "Bob"

	cP4p <- "Tock"

	time.Sleep(3 * time.Second)
	cP4a <- 500

	// ---------------------------
	//
	// Hold main
	time.Sleep(1 * time.Second)
}
