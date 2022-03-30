package task

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func add(arr []int) int {
	return arr[0] + arr[1]
}

func b(s string) string {
	return fmt.Sprint("hello ", s)
}

type P struct {
	name string
}

func (p *P) Greeting(flag string) string {
	return fmt.Sprint("Hello ", flag, " I am", p.name)
}

func TestTask(t *testing.T) {

	// let us know where is the panic log task
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := func(ss []string) string {
		return ss[0] + "+" + ss[1]
	}

	// ---------------------------
	//
	// Define Task & Parameters & Returns

	cP4a, cR4a := make(chan any), make(chan any)
	if err := Task(true, add, cP4a, cR4a); err != nil {
		fmt.Println(err)
	}

	cP4b, cR4b := make(chan any), make(chan any)
	if err := Task(true, b, cP4b, cR4b); err != nil {
		fmt.Println(err)
	}

	cP4p, cR4p := make(chan any), make(chan any)
	p := P{name: "Tick"}
	if err := Task(true, p.Greeting, cP4p, cR4p); err != nil {
		fmt.Println(err)
	}

	cP4c, cR4c := make(chan any), make(chan any)
	if err := Task(true, c, cP4c, cR4c); err != nil {
		fmt.Println(err)
	}

	// ---------------------------
	//
	// What to do When Task has handled task

	go func() {
		for {
			select {
			case r := <-cR4a:
				if err, ok := r.(error); ok { // MUST check error & log to see which task is panic
					log.Fatalln(err)
				}
				fmt.Println("cR4a", r)

			case r := <-cR4b:
				if err, ok := r.(error); ok {
					log.Fatalln(err)
				}
				fmt.Println("cR4b", r)

			case r := <-cR4p:
				if err, ok := r.(error); ok {
					log.Fatalln(err)
				}
				fmt.Println("cR4p", r)

			case r := <-cR4c:
				if err, ok := r.(error); ok {
					log.Fatalln(err)
				}
				fmt.Println("cR4c", r)

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
	cP4a <- []int{2, 5}

	time.Sleep(1 * time.Second)
	cP4c <- []string{"1", "2"}

	time.Sleep(1 * time.Second)
	cP4b <- "foo"

	cP4p <- "bar"

	cP4a <- []int{3, 2}

	// ---------------------------
	//
	// Hold main
	time.Sleep(1 * time.Second)
}
