package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// properties that we want to be exported / accessible from other places must start with a capital letter, otherwise they won't be exported

// we could have just included handle as a string in the Person struct and then have a bunch of logic on Person to take care of the handle
// however, using the new type TwitterHandle allows us to break things into smaller pieces and delegate some responsibility and make things more readable and understandable

// a type definition copies the fields of a type over to a new type whereas a type alias copies the fields and the method sets -> they become that exact type

// type Handle struct {
// 	handle string
// 	name   string
// }

// type TwitterHandle = string // capital letter = public // type alias -> resolves to string exactly
type TwitterHandle string // type declaration -> extend string and make it its own thing

func (th TwitterHandle) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

// in an interface, you can just have functions
// don't need func keyword bc it's the only thing we can put in here.
// anything that implements this type, will also be of type Identifiable
type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type Conflict interface {
	ID() string
}

type socialSecurityNumber string

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States"
}

type euId struct {
	id      string
	country string
}

func (eui euId) ID() string {
	return eui.id
}

// using interface{} isn't the best. Doesn't give us any indication into what type it really is. (kind of like `any` in typescript)
func NewEuId(id interface{}, country string) Citizen {
	switch v := id.(type) { // go takes care of casting v for us
	case string:
		return euId{
			id:      v, // can also do this -> id: id.(string),
			country: country,
		}
	case int:
		return euId{
			id:      strconv.Itoa(v),
			country: country,
		}
	case euId: // can also switch on structs!
		return v
	case Person:
		euId, _ := v.Citizen.(euId)
		return euId
	default:
		panic("using invalid to for EU identifier")
	}

}

func (eui euId) Country() string {
	return eui.country
}

type Name struct {
	first string
	last  string
}

// "method receiver function"
func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name          // embedding this struct by using it directly here
	twitterHandle TwitterHandle
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

// go implicitly inherits interfaces. we don't have to tell it that Person is of type Identifiable, but go understands it because we've satisfied the interface
// types can implement interfaces
// func (p *Person) ID() string {
// 	return "12345"
// }

// in order to *change state* we must use a pointer-based receiver OR we need to return a new version of that type
// when this is called without a pointer receiver, a copy is used - we wouldn't be changing the state like we are intending
// oftentimes people just set all method receivers to be pointer receivers (pointer-based vs value-based), even for read only methods. Also more memory efficient
func (p *Person) SetTwitterHandle(handle TwitterHandle) error {
	if len(handle) == 0 {
		p.twitterHandle = handle // empty handle
	} else if !strings.HasPrefix(string(handle), "@") {
		return errors.New("twitter handle must start with @ symbol")
	}

	p.twitterHandle = handle
	return nil
}

func (p *Person) GetTwitterHandle() TwitterHandle {
	return p.twitterHandle
}
