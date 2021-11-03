package foundation

import (
	"fmt"
	"reflect"
)

const (
	first  = iota + 6  // every time iota is called (anywhere) it is incremented
	second = 2 << iota // bit shift 2, iota (1) time = multiplying by 2
	third              // if we specify nothing -> will use the expression from the line above (where iota will have incremented -> 2 << 2 = 2 * 2^2 = 8)
)

const (
	fourth = iota // iota resets to zero for each constant block
)

// primitive data type declaration stuff etc
func Primitives() {
	//// primitive types
	fmt.Println("Primitive Types:")

	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Sophie"
	fmt.Println(firstName)

	b := true
	var ce bool
	fmt.Println(b, ce, reflect.TypeOf(b))

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c) // multiple assignment
	fmt.Println(r, im)

	fmt.Print("---------\n\n")

	//// pointers
	// pointer arithmetic is not allowed
	fmt.Println("Pointers:")

	var firstName2 *string = new(string) // can't assign to an unspecified pointer location because we try to indirect an empty location so we must give it a new intitialization
	*firstName2 = "Jamie"
	fmt.Println(*firstName2)

	name := "Jaime"
	fmt.Println(name)

	ptr := &name
	fmt.Println(ptr, *ptr)

	name = "Wyatt"
	fmt.Println(ptr, *ptr)

	//// constants
	// constants must be assignable at compile time -> can't use a const to capture the return value of a func because of this
	const d = 3 // implicitly typed const - compiler will implicitly type it when it runs into it
	fmt.Println(d + 1)
	fmt.Println(d + 1.2)

	const e int = 3 // explicitly typed const
	fmt.Println(e + 2)
	fmt.Println(float32(e) + 2.1)

	fmt.Println("iotas:", first, second, third, fourth)
}
