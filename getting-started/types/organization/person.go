package organization

import (
	"errors"
	"fmt"
	"strings"
)

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

type Person struct {
	firstName     string // properties that we want to be exported / accessible from other places must start with a capital letter, otherwise they won't be exported
	lastName      string
	twitterHandle TwitterHandle
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

// "method receiver function"
func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// go implicitly inherits interfaces. we don't have to tell it that Person is of type Identifiable, but go understands it because we've satisfied the interface
// types can implement interfaces
func (p *Person) ID() string {
	return "12345"
}

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
