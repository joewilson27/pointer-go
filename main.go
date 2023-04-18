package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// with pointers
func zeroptr(iptr *int) {
	*iptr = 0 // * disini diartikan meng-pointerkan dari memory address ke current value pd address variable
	// dan disini kita akan meng-assign value 0 ke variable pointer address tersebut
}

/*
* Pointer Golang
* Source: https://www.youtube.com/watch?v=a4HcEsJ1hIE&ab_channel=TechWithTim
*
* Pointer Golang
* & (ampersand) -> get the pointer / mengambil address dr sebuah variable
* * (asterix) -> derefrence
 */

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

	// 4 =========================================================== //
	x := 7
	//y := &x
	var y *int // ini sama
	y = &x     // dengan y := &x
	// fmt.Println(*x) // tidak bisa print, cannot indirect
	fmt.Println(&x) // value would be the variable address (the reference of x)
	fmt.Println(y)  // would print the reference of x
	fmt.Println(*y) // ini bisa, tidak seperti (*x) di atas, karena ini hanya bisa di lakukan oleh variable yg me-reference pointer dr variable lain

	fmt.Println(x, y) // print before change the value of y which pointer to x. x = 7
	*y = 8
	fmt.Println(x, y) // x = 8
	// End of 4 ================================================== //

	// 5 =========================================================== //
	toChange := "hello"
	fmt.Println(toChange) // before
	changeValue(&toChange)
	fmt.Println(toChange) // after, value of toChange is changed to "changed!"

	toChange2 := "hello2"
	fmt.Println(toChange2) // before
	cantChangeValue(toChange2)
	fmt.Println(toChange2) // after, can't change the value because the param of cantChangeValue is not pointer

	toChange3 := "hello3"
	fmt.Println(toChange3) // before
	toChange3 = changeValue2(toChange3)
	fmt.Println(toChange3) // after, value change by copying a value from variable inside function
	// End of 5 ==================================================== //

}

func changeValue(str *string) { // remember this? the param would be the variable that pointers to variable that has string type
	*str = "changed!"
}

func cantChangeValue(str string) {
	str = "can't changed!"
}

func changeValue2(str string) string {
	str = "changed, but copying value not change the reference"
	return str
}
