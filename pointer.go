package main

import "fmt"

func Main_pointer() {
	check_string_pointer()
}

func check_string_pointer() {

	var k1 string = "v1"
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xA      | k1   | v1
	fmt.Printf("%p | k1 | %s\n", &k1, k1)

	p1 := &k1
	// &k1 это memory address of k1 (hex-адрес в ram-памяти в котором хранится k1)
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xB      | p1    | 0xA
	fmt.Printf("%p | p1 | %v\n", &p1, p1)

	// dereference of pointer: означает get value of "underlying" variable (получить значение of переменной на которую ссылается pointer)
	var dv1 string = *p1
	fmt.Println(dv1) // "v1"

	var p2_ret *string = func_InputPointer(p1)
	// assign result to p2_ret; p2_ret это новая переменная, у которой значение равно значению of result-а (т.е. значение копируется).
	// p.s. ram-memory (example)
	// address: | key:      | value:
	// 0xE      | p2_ret    | 0xD
	fmt.Printf("%p | p2_ret | %v\n", &p2_ret, p2_ret)

	fmt.Println(*p2_ret) // "v2_update"

	p1 = p2_ret
	// присваивание(assignment) означает что копируется значение(value).
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xB      | p1    | 0xD

	*p1 = "v2_update_latest"
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xD      | k2   | v2_update_latest

	fmt.Println(*p2_ret) // "v2_update_latest"

	var pp1 **string = &p1
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xF      | pp1    | 0xB

	fmt.Println(**pp1) // "v2_update_latest"
}

func func_InputPointer(p2 *string) *string {
	// p2 = "Copy(копия)" of input-a (т.е. в stack-e allocat-ится новая переменная (p2), у которой значение(value) такое-же как у input-переменной, т.е. "значение копируется")
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xC      | p2    | 0xA
	fmt.Printf("%p | p2 | %v\n", &p2, p2)

	*p2 = "v1_update"
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xA      | k1   | v1_update

	var k2 string = "v2"
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xD      | k2   | v2
	fmt.Printf("%p | k2 | %s\n", &k2, k2)

	p2 = &k2
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xC      | p2    | 0xD

	*p2 = "v2_update"
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xD      | k2   | v2_update

	return p2
}

/*

// output:

0xc000026070 | k1 | v1
0xc000060040 | p1 | 0xc000026070
v1
0xc000060050 | p2 | 0xc000026070
0xc0000260a0 | k2 | v2
0xc000060048 | p2_ret | 0xc0000260a0
v2_update
v2_update_latest
v2_update_latest

*/



















