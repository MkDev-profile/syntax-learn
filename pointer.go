package main

import "fmt"

func Main_pointer() {
	check_pointer()
}

func check_pointer() {

	var k1 string = "v1"
	//var k1 int = 1
	//var k1 [3]int = [3]int{1, 0, 0}
	//var k1 map[int]string = map[int]string{0:"1", 1:"0", 2:"0"}
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xA      | k1   | v1
	//fmt.Printf("%p | k1 | %s\n", &k1, k1)
	//fmt.Printf("%p | k1 | %d\n", &k1, k1)
	fmt.Printf("%p | k1 | %v\n", &k1, k1)

	p1 := &k1
	// &k1 это memory address of k1 (hex-адрес в ram-памяти в котором хранится k1)
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xB      | p1    | 0xA
	fmt.Printf("%p | p1 | %p\n", &p1, p1)

	// dereference of pointer: означает get value of "underlying" variable (получить значение of переменной на которую ссылается pointer)
	var dv1 string = *p1
	//var dv1 int = *p1
	//var dv1 [3]int = *p1
	//var dv1 map[int]string = *p1
	fmt.Println(dv1) // "v1"

	var p2_ret *string = func_InputPointer(p1)
	//var p2_ret *int = func_InputPointer(p1)
	//var p2_ret *[3]int = func_InputPointer(p1)
	//var p2_ret *map[int]string = func_InputPointer(p1)
	// assign result to p2_ret; p2_ret это новая переменная, у которой значение равно значению of result-а (т.е. значение копируется).
	// p.s. ram-memory (example)
	// address: | key:      | value:
	// 0xE      | p2_ret    | 0xD
	fmt.Printf("%p | p2_ret | %p\n", &p2_ret, p2_ret)

	fmt.Println(*p2_ret) // "v2_update"

	p1 = p2_ret
	// присваивание(assignment) означает что копируется значение(value).
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xB      | p1    | 0xD

	*p1 = "v2_update_latest"
	//*p1 = 22
	//(*p1)[1] = 2 
	//(*p1)[1] = "2" 
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xD      | k2   | v2_update_latest

	fmt.Println(*p2_ret) // "v2_update_latest"

	var pp1 **string = &p1
	//var pp1 **int = &p1
	//var pp1 **[3]int = &p1
	//var pp1 **map[int]string = &p1
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xF      | pp1    | 0xB

	fmt.Println(**pp1) // "v2_update_latest"
}

func func_InputPointer(p2 *string) *string {
//func func_InputPointer(p2 *int) *int {
//func func_InputPointer(p2 *[3]int) *[3]int {
//func func_InputPointer(p2 *map[int]string) *map[int]string {
	// p2 = "Copy(копия)" of input-a (т.е. в stack-e allocat-ится новая переменная (p2), у которой значение(value) такое-же как у input-переменной, т.е. "значение копируется")
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xC      | p2    | 0xA
	fmt.Printf("%p | p2 | %p\n", &p2, p2)

	*p2 = "v1_update"
	//*p2 = 11
	//(*p2)[1] = 1
	//(*p2)[1] = "1"
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xA      | k1   | v1_update

	var k2 string = "v2"
	//var k2 int = 2
	//var k2 [3]int = [3]int{2, 0, 0}
	//var k2 map[int]string = map[int]string{0:"2", 1:"0", 2:"0"}
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xD      | k2   | v2
	//fmt.Printf("%p | k2 | %s\n", &k2, k2)
	//fmt.Printf("%p | k2 | %d\n", &k2, k2)
	fmt.Printf("%p | k2 | %v\n", &k2, k2)

	p2 = &k2
	// p.s. ram-memory (example)
	// address: | key:  | value:
	// 0xC      | p2    | 0xD

	*p2 = "v2_update"
	//*p2 = 21
	//(*p2)[1] = 1
	//(*p2)[1] = "1"
	// p.s. ram-memory (example)
	// address: | key: | value:
	// 0xD      | k2   | v2_update

	return p2
}

/*

// output for "string":

0xc000026070 | k1 | v1
0xc000060040 | p1 | 0xc000026070
v1
0xc000060050 | p2 | 0xc000026070
0xc0000260a0 | k2 | v2
0xc000060048 | p2_ret | 0xc0000260a0
v2_update
v2_update_latest
v2_update_latest

// output for "int":

0xc00008c060 | k1 | 1
0xc00008e038 | p1 | 0xc00008c060
1
0xc00008e048 | p2 | 0xc00008c060
0xc00008c090 | k2 | 2
0xc00008e040 | p2_ret | 0xc00008c090
21
22
22

// output for "array":

0xc000014168 | k1 | [1 0 0]
0xc000060040 | p1 | 0xc000014168
[1 0 0]
0xc000060050 | p2 | 0xc000014168
0xc0000141b0 | k2 | [2 0 0]
0xc000060048 | p2_ret | 0xc0000141b0
[2 1 0]
[2 2 0]
[2 2 0]

// output for "map":

0xc000060038 | k1 | map[0:1 1:0 2:0]
0xc000060048 | p1 | 0xc000060038
map[0:1 1:0 2:0]
0xc000060058 | p2 | 0xc000060038
0xc000060060 | k2 | map[0:2 1:0 2:0]
0xc000060050 | p2_ret | 0xc000060060
map[0:2 1:1 2:0]
map[0:2 1:2 2:0]
map[0:2 1:2 2:0]

*/



















