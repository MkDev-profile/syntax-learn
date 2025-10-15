package main

func Main_pointer() {
    example_main()
}

// ------------------------------------------------------------------------------
// Example:
// pass pointer to function by value (by copy)

func example_main() {
    v := 5 // v: addr=0xA1 val=5 
    p := &v // p: addr=0xA2 val=ref(to-> 0xA1)
    
    println(*p) // 5
    changePointer(p) // pass <Copy of p>
    println(*p) // <Output>
}

func changePointer(p *int) { // input: <Copy of p> (=> copy of "pointer to v")
    v := 3 // local v: addr=0xA3 val=3
    //p = &v // changes <Copy of p>; p_copy=ref(to-> 0xA3); Output: 5, 5
	*p = v // changes memory value where p refers; val(0xA1) = val(0xA3); Output: 5, 3
}

// ------------------------------------------------------------------------------

// Pass by value - copies entire array
func processArray(arr [1000]int) {
}

// Pass by pointer - avoids copy
func processArrayPtr(arr *[1000]int) {
}


// --------------------------------------------------------------------------------





















