package main

import (
	"fmt"
	"sync"
)

/*

-- Theory: GoLang channel internaly:

channel (канал) в GoLang:
chan

channel alloc-ейтится в heap-e

source code:
type hchan struct
(p.s. h=Heap,chan=канал)
// some fields of hchan:
- buf буфер (p.s. buffer = circular queue),
- qcount: количество элементов в буфере,
- dataqsiz: размерность буфера,
- closed uint32 (for atomic): флаг закрыт/открыт канал(buffer),
- elemsize: alloc size for item,
- elem type: type of item,
- recvq: receive queue = очередь of reader-ов(goroutines which read from channel),
- sendq: send queue = очередь of sender-ов(goroutines which push to channel),
- recvx: current индекс in recvq,
- sendx: current индекс in sendq,
- lock: lockobj of this channel-instance.

*/

//
// Merge several channels into one channel:
//

func genDataForWorkers() <-chan int {
    out := make(chan int)

    go func() {
        for n := range 5 {
            out <- n
        }

        close(out)
    }()

    return out
}

func square(n int) int {
    return n*n
}

func funcWork(inpChan <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        for inpItem := range inpChan {
            res := square(inpItem)
            out <- res
        }

        close(out)
    }()

    return out
}

func merge(inpChans ...<-chan int) <-chan int {
    outChan := make(chan int)

    var wg sync.WaitGroup

    funcRedirect := func(inpChan <-chan int) {
        defer wg.Done()

        for res := range inpChan {
            outChan <- res
        }
    }

    wg.Add(len(inpChans))

    for _, inpChan := range inpChans {
        go funcRedirect(inpChan)
    }

    go func() {
        wg.Wait()
        close(outChan)
    }()
    
    return outChan
}

func Main_channel() {
    // generate input channel = input data for workers.
    commonInpChan := genDataForWorkers()

    // workers. 
    // p.s. input channel is shared between workers.
    w1_out := funcWork(commonInpChan)
    w2_out := funcWork(commonInpChan)

    // read from merged output channel.
    for n := range merge(w1_out, w2_out) {
        fmt.Println(n)
    }
}

















































