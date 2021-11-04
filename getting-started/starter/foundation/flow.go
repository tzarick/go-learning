package foundation

import "github.com/tzarick/go-learning/getting-started/starter/models"

func ControlFlow() {
	var i int
	for i < 5 { // while loop
		println("i", i)
		i++
		if i == 3 {
			break // continue is also valid
		}
	}

	// scoped within the loop
	for j := 0; j < 5; j++ {
		println("j", j)
	}

	var k int // scoped outside the loop
	for ; k < 5; k++ {
		println("k", k)
	}

	i = 0
	for { // infinite loop
		if i == 5 {
			break
		}
		i++
	}

	i = 0
	for { // also an infinite loop
		if i == 5 {
			break
		}
		i++
	}

	slice := []int{1, 2, 3}
	for i = 0; i < len(slice); i++ {
		println("boring way")
	}

	//// range

	for i, v := range slice {
		println("nice way", i, v)
	}

	// can also iterate over keys and values using range construct
	m := map[string]float32{"pi": 3.14, "something": 1234.50}
	for k, v := range m {
		println("slick map stuff", k, v)
	}

	// can ignore the second return value
	for k := range m {
		println(k)
	}

	// can ignore first val with _
	for _, v := range m {
		println(v)
	}

	// when go doesn't know what to do or how to handle something -> a go panic (kind of like exception?) // normally we just use errors but panic is good for really bad situations

	// panic("Something bad!")

	u1 := models.User{
		ID:        1,
		FirstName: "Oi",
		LastName:  "io",
	}
	u2 := models.User{
		ID:        2,
		FirstName: "lo",
		LastName:  "oil",
	}

	if u1.ID == u2.ID {
		println("same")
	} else if u1.FirstName == u2.FirstName {
		println("similar")
	} else {
		println("diff")
	}

	//// switch
	type HTTPRequest struct {
		Method string
	}

	// there is implicit breaking (not fallthrough) -> we have to explicitly specify fallthrough
	r := HTTPRequest{Method: "GET"}
	switch r.Method {
	case "GET":
		println("get request")
	case "POST":
		println("post request")
		fallthrough
	case "PUT":
		println("put request")
	default:
		println("unhandled")
	}
}
