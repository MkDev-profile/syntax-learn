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

-- Double-Deadlock example:

var mu_1 sync.Mutex
var mu_2 sync.Mutex

func F1() {
    mu_1.Lock()
    fmt.Println("mu_1 locked by F1")

    fmt.Println("try lock mu_2")
    mu_2.Lock()

    // code
}

func F2() {
    mu_2.Lock()
    fmt.Println("mu_2 locked by F2")

    fmt.Println("try lock mu_1")
    mu_1.Lock() // (blocked = waiting forever for mu_1)

    // code
}

go F1()
go F2()

Deadlock: 
example output:
GR_1: mu_1 locked by F1
GR_2: mu_2 locked by F2
GR_1: try lock mu_2 // p.s. (blocked = waiting forever for mu_2)
GR_2: try lock mu_1 // p.s. (blocked = waiting forever for mu_1)

*/











