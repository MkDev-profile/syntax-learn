// Флаги и битовые маски

package main

import "fmt"

const maxFileSizeInBytes = 1 << 26 // 64MB   
// p.s. (2^26)/1024/1024 = 64 MB   // 1 КилоБайт = 1024 байта; 1 байт = 1/1024 КилоБайт

type Permission uint

const (
    None Permission = 0
    
    Read    Permission = 1 << iota // 0001
    Write                          // 0010
    Execute                        // 0100
    Delete                         // 1000
    
    // Combined permissions
    ReadWrite = Read | Write
    All       = Read | Write | Execute | Delete
)

func (p *Permission) Add(flag Permission) {
    *p |= flag // *p = *p | flag
}

func (p Permission) Has(flag Permission) bool {
    return p&flag == flag
}

func (p *Permission) Toggle(flag Permission) {
    *p ^= flag
}

func (p *Permission) Remove(flag Permission) {
    *p &^= flag // AND NOT operator
}

func Main_bit_flags() {
	fmt.Printf("maxFileSizeInBytes: %d\n", maxFileSizeInBytes)

    fmt.Println("Initial `ReadWrite`")
    var perm Permission = ReadWrite
    fmt.Printf("%05b\n", Read)
    fmt.Printf("%05b\n", Write)
    fmt.Printf("%05b\n", perm)
    
    fmt.Println("Add `Delete`")
    perm.Add(Delete)
    fmt.Printf("%05b\n", Delete)
    fmt.Printf("%05b\n", perm)

    fmt.Println("Toggle `Execute`")
    fmt.Printf("%05b\n", Execute)    
    perm.Toggle(Execute)
    fmt.Printf("%05b\n", perm)  
    
    fmt.Println("Remove `Read`")
    fmt.Printf("%05b\n", Read) 
    perm.Remove(Read)
    fmt.Printf("%05b\n", perm)
}


/*

Побитовый сдвиг:

-- explain example:

1 << 26 
означает битовый сдвиг влево числа 1 на 26 позиций.
p.s.
Число 1 в двоичной системе: 00000001
Сдвигаем все биты на 26 позиций влево
Результат: 100000000000000000000000000 (1 с 26 нулями)
1 << 26 = 2^26 = 67,108,864 (67 миллионов)


-- p.s.
битовые операции очень быстрые



*/








