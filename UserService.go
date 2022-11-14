package main

type IUserService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// UserService is a concrete implementation of IUserService
type UserService struct{}
