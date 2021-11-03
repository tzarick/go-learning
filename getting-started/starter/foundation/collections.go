package foundation

import "fmt"

// Arrays, Slices, Maps, Structs (data side of a class)
func Collections() {
	fmt.Println("")
	fmt.Println("Collections:")

	//// arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[2])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//// slices
	slice := arr[:] // colon operator, from beginning to end (nothing specified = all elements)

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice) // slice is kind of like a pointer to the array - both reference the same memory

	slice2 := []int{1, 2, 3} // compiler will manage the underlying array for us
	fmt.Println(slice2)

	slice2 = append(slice2, 4, 3) // append lets the compiler know to resize the underlying array etc
	fmt.Println(slice2)

	// beginning index inclusive, end index exclusive
	s2 := slice2[1:]  // index 1 until end -> [2 3 4 3]
	s3 := slice2[:2]  // index 0 until 2 -> [1 2]
	s4 := slice2[1:2] // index 1 until 2 -> [2]
	s5 := slice2[1:1] // -> []

	fmt.Println(s2, s3, s4, s5)

	//// maps
	// map[key-type]value-type -> consistently typed pairs, can dynamically update
	fmt.Println("Maps:")
	m := map[string]int{"something": 7}
	fmt.Println(m)
	fmt.Println(m["something"])

	m["something"] = 42
	fmt.Println(m)

	delete(m, "something") // delete a key/value pair
	fmt.Println(m)

	//// structs
	// Structs can be mixed type but are fixed at compile time (can't dynamically change later)
	fmt.Println("Structs:")

	// can define a struct at any scope we deem necessary. Right now it is only available in this func
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Sami"
	u.LastName = "Malya"
	fmt.Println(u)
	fmt.Println(u.LastName)

	u2 := user{
		ID:        2,
		FirstName: "Rod",
		LastName:  "Saud",
	}
	fmt.Println(u2)
}
