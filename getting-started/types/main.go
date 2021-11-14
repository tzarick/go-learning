package main

import (
	"fmt"

	"github.com/tzarick/go-learning/getting-started/types/organization"
)

func main() {
	// var p organization.Identifiable = organization.Person{
	// 	FirstName: "me",
	// 	LastName:  "last",
	// } // bc we've typed it as Identifiable, we won't be able to access expanded Person attributes that aren't on Identifiable
	// p1 := organization.Person{
	// 	FirstName: "me",
	// 	LastName:  "last",
	// }

	p2 := organization.NewPerson("somebody", "else", organization.NewEuId("123-23-1234", "Italy"))
	err := p2.SetTwitterHandle(organization.TwitterHandle("@something"))
	// fmt.Printf("%T\n", organization.TwitterHandle("test"))
	if err != nil {
		fmt.Printf("err occurred while setting twitter handle: %s\n", err.Error())
	}
	// println(p2.GetTwitterHandle())
	// println(p2.GetTwitterHandle().RedirectUrl())
	// println(p2.ID())
	// println(p2.Country())
	// println(p2.FullName())
	// println(p.ID(), p1.FirstName)

	//// Comparing Types

	// name1 := Name{First: "", Last: ""}
	// name2 := Name{First: "one", Last: "Wilson"}

	// if name1 == (Name{}) {
	// 	println("We match")
	// }

	// Hashable Types (can use in map)

	// portfolio := map[Name][]organization.Person{}
	// portfolio[name1] = []organization.Person{p2}

	// if name1.Equals(name2) {
	// 	println("we match")
	// }

	// ssn := organization.NewSocialSecurityNumber("12340543")
	// eui := organization.NewEuId("111", "france")
	// eui2 := organization.NewEuId("111", "france")

	// fmt.Printf("%T\n", ssn)
	// fmt.Printf("%T\n", eui)
	// if eui == eui2 {
	// 	println("We match")
	// }

	//// Switching on Types
	println(p2.Country())
}

type Name struct {
	First  string
	Last   string
	Middle []string
}

// if we aren't comparable (unpredictable memory layout)
func (n Name) Equals(otherName Name) bool {
	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

type OtherName struct {
	First string
	Last  string
}
