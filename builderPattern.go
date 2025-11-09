package main

import "fmt"

// Problem Statement : Implement Builder Pattern
// Pros: When there are too many arguments in a building of an object
// where there is optional configurations, such that passing nil in most of the fields is just pointless

// Just make the methods of builder service return itself so you can chain the methods
// At the end create a build method that instantiates the actual object

type User struct {
	userId      string // Required Property for example
	userName    string
	email       string
	password    string
	phoneNumber string
}

func createUser(userId string, userName string, password string, email string, phoneNumber string) *User {
	return &User{userId, userName, email, password, phoneNumber}
}

func (user *User) DisplayUserInfo() {
	fmt.Println("------------------")
	fmt.Println("User Id: " + user.userId)
	fmt.Println("User Name: " + user.userName)
	fmt.Println("User Email: " + user.email)
	fmt.Println("User Password: " + user.password)
	fmt.Println("User Phone: " + user.phoneNumber)
	fmt.Println("------------------")
}

type UserBuilder struct {
	userId      string // Required Property for example
	userName    string
	email       string
	password    string
	phoneNumber string
}

func Builder(userId string) *UserBuilder {
	return &UserBuilder{userId: userId}
}

func (ub *UserBuilder) Username(username string) *UserBuilder {
	ub.userName = username
	return ub
}
func (ub *UserBuilder) Email(email string) *UserBuilder {
	ub.email = email
	return ub
}
func (ub *UserBuilder) Password(password string) *UserBuilder {
	ub.password = password
	return ub
}
func (ub *UserBuilder) PhoneNumber(phoneNumber string) *UserBuilder {
	ub.phoneNumber = phoneNumber
	return ub
}
func (ub *UserBuilder) Build() *User {
	return createUser(ub.userId, ub.userName, ub.email, ub.password, ub.phoneNumber)
}

func BuilderPatternExample() {
	user1 := Builder("1").
		Email("gg@gmail.com").
		Username("Depri").
		PhoneNumber("1888888888888").
		Build()
	user1.DisplayUserInfo()

	user2 := Builder("2").
		Email("ez@gmail.com").
		Password("123456").
		Username("Rekt").
		Build()
	user2.DisplayUserInfo()
}
