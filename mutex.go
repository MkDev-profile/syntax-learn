package main

import (
	"fmt"
	"sync"
)

type Container struct {
    counters map[string]int
    mu       sync.Mutex
}

// func (c *Container) => c is pointer => c.mu is pointer => (not copies mutex)
func (c *Container) inc(name string) { 
	fmt.Printf("%p %p (inside method of container)\n", c, &(c.mu))
	
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}

func Main_mutex() {
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }

	fmt.Printf("%p %p (inside main)\n", &c, &c.mu)

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
		fmt.Printf("%p %p (inside anonym func)\n", &c, &c.mu)	

        for range n {
            (&c).inc(name)
        }
    }

    wg.Go(func() {
		fmt.Printf("%p %p (inside goroutine)\n", &c, &c.mu)

        doIncrement("a", 1)
    })

    wg.Go(func() {
        doIncrement("a", 1)
    })

    wg.Go(func() {
        doIncrement("b", 1)
    })

    wg.Wait()
    fmt.Println(c.counters)
}

/*

output:

0xc000026070 0xc000026078 (inside main)
0xc000026070 0xc000026078 (inside anonym func)
0xc000026070 0xc000026078 (inside method of container)
0xc000026070 0xc000026078 (inside goroutine)
0xc000026070 0xc000026078 (inside anonym func)
0xc000026070 0xc000026078 (inside method of container)
0xc000026070 0xc000026078 (inside anonym func)
0xc000026070 0xc000026078 (inside method of container)
map[a:2 b:1]

*/


/*

Deadlock examples:

// example: 
// "Взаимная блокировка"

func F1() {
    mutA.Lock() // 1
    mutB.Lock() // 3, blocked
    // code
    mutB.Unlock()
    mutA.Unlock()
}

func F2() {
    mutB.Lock() // 2
    mutA.Lock() // 3, blocked
    // code
    mutA.Unlock()
    mutB.Unlock()
}

go F1()
go F2()

// example: "Вложенная блокировка"

func F1() {
    mutA.Lock() // 1
    F2() // 2
    // code
    mutA.Unlock()
}

func F2() {
    mutA.Lock() // 3, blocked
    // code
    mutA.Unlock()
}

go F1()

*/













