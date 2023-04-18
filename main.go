package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

/**
* & and * --> & is get the pointers (usually an address in memory), * is Derefrence
*/

// with pointers
func zeroptr(iptr *int) {
	*iptr = 0 // * disini diartikan meng-pointerkan dari memory address ke current value pd address variable
	// dan disini kita akan meng-assign value 0 ke variable pointer address tersebut
}

func main() {
	// 1 ====================================================== //
	i := 1
	fmt.Println("initial:", i) // ini akan ng-print value asli
	var ptr_ *int              // *int artinya pointer variable yg berasal dr address yg memiliki nilai int
	ptr_ = &i

	zeroval(i)
	fmt.Println("zeroval:", i) // ini akan ng-print 1, karena zeroval() mengganti value variable pada fungsi di dalamnya, bukan pointer ke variable asal

	zeroptr(&i)                // &i akan memberikan memory address pada param
	fmt.Println("zeroptr:", i) // ini akan ng-print hasil penggantian value pada method zeroptr()

	fmt.Println("pointer:", &i) // print pointer / address

	fmt.Println("ptr:", *ptr_) // using * to print the actual value that pointer to address i to ptr

	// End of 1 =============================================== //

	// 2 ====================================================== //
	var num int = 5
	// prints the value stored in variable
	fmt.Println("Variable Value:", num)
	// prints the address of the variable
	fmt.Println("Memory Address:", &num)

	// Pointer Variable
	// create the pointer variable, int type
	var ptr *int = &num // *int artinya pointer variable yg berasal dr address yg memiliki nilai int, &num artinya kita assign address dr variable num
	fmt.Println("Pointer of ptr: ", ptr)
	fmt.Println("Actual value of ptr: ", *ptr)

	// We can also create pointer variables of other types
	// pointer variable of string type
	var sampleString string = "lorem ipsum"
	var ptr1 *string // * artinya, ptr1 hanya menerima assign address dr variable yg type string
	ptr1 = &sampleString
	//*ptr1 = "aaaa"
	fmt.Println("Pointer of ptr1: ", ptr1)
	fmt.Println("Actual value of ptr1: ", *ptr1)

	// End of 2 ===================================================== //

	// 3 ============================================================ //
	var nomor int
	var ptrNomor *int // artinya tahu kan? ptrNomor hanya bisa di assign address yg type data int

	nomor = 22
	fmt.Println("\nAddress of num:", &nomor)
	fmt.Println("Value of num:", nomor)

	ptrNomor = &nomor
	fmt.Println("\nAddress of pointer ptrNomor:", ptrNomor)
	fmt.Println("Content of pointer ptrNomor:", *ptrNomor)

	nomor = 11
	fmt.Println("\nAddress of pointer ptrNomor:", ptrNomor)
	fmt.Println("Content of pointer ptrNomor:", *ptrNomor) // ptrNomor akan berubah value 11

	*ptrNomor = 2
	fmt.Println("\nAddress of num:", &nomor)
	fmt.Println("Value of num:", nomor) // value nomor akan berubah disini

	// End of 3 ===================================================== //

}
