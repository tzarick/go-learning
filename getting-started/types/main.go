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

	p2 := organization.NewPerson("somebody", "else", organization.NewSocialSecurityNumber("123-23-1234"))
	err := p2.SetTwitterHandle(organization.TwitterHandle("@something"))
	fmt.Printf("%T\n", organization.TwitterHandle("test"))
	if err != nil {
		fmt.Printf("err occurred while setting twitter handle: %s\n", err.Error())
	}
	println(p2.GetTwitterHandle())
	println(p2.GetTwitterHandle().RedirectUrl())
	println(p2.ID())
	println(p2.FullName())
	// println(p.ID(), p1.FirstName)
}
