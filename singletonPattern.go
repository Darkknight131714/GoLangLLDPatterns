package main

import (
	"fmt"
	"sync"
)

// Problem Statement: How to implement singletonPattern in golang.
// 1st Method: Have a GetterMethod for an object, which will be executed only once
// Pros: Lazy Initialization, used only when needed

var userManager *UserManagementService
var once sync.Once

type UserManagementService struct {
}

// GetUserManagementService This executes the initialization only once, and other times returns directly
func GetUserManagementService() *UserManagementService {
	once.Do(func() {
		userManager = &UserManagementService{}
	})
	return userManager
}

//2nd Method: If you dont want to have a lazy initialzation as object will be used everytime,
// You can also declare its initialzation in global context
// As this gets executed before any method [Even main function]

var userManager2 = &UserManagementService{}

func GetUserManagementService2() *UserManagementService {
	return userManager2
}

func SingletonPatternExample() {
	firstMethodLocalUserManager1 := GetUserManagementService()
	firstMethodLocalUserManager2 := GetUserManagementService()
	fmt.Println("Checking if both are same from 1st approach")
	fmt.Println(firstMethodLocalUserManager1 == firstMethodLocalUserManager2)

	secondMethodLocalUserManager1 := GetUserManagementService2()
	secondMethodLocalUserManager2 := GetUserManagementService2()
	fmt.Println("Checking if both are same from 2nd approach")
	fmt.Println(secondMethodLocalUserManager1 == secondMethodLocalUserManager2)
}
