package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // a slice that holds pointers to User objects
	nextID = 1     // at the package level, we can allow the compiler to implicity type the var
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
