package main

type ContB string

type A struct {
	content string
}

type C struct {
	cont int
}

type B struct {
	contentA A
	content  ContB
	contentC C
}

func GetContentA() string {
	return "contentA"
}

func GetContentB() ContB {
	return "contentB"
}

func ContentA(conten string) A {
	return A{content: conten}
}

func ContentB(a A, conten ContB) B {
	return B{contentA: a, content: conten, contentC: C{cont: 6}}
}

// type User struct {
// 	name string
// }

// // NewUser - Creates a new instance of User
// func NewUser(name string) User {
// 	return User{name: name}
// }

// // NewUserName - Returns a string to provide the name of a new user
// func NewUserName() string {
// 	return "James"
// }

// // Get - A method with user as dependency
// func (u *User) Get(message string) string {
// 	return fmt.Sprintf("Hello %s - %s", u.name, message)
// }

// // Run - Depends on user and calls the Get method on User
// func Run(user User) {
// 	result := user.Get("It's nice to meet you!")
// 	fmt.Println(result)
// }

func main() {
	// user := Initialize()
	// Run(user)
}
