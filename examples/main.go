package main

import (
	"fmt"

	"github.com/ralstan-vaz/go-errors"
	"github.com/ralstan-vaz/go-errors/grpc"
	"github.com/ralstan-vaz/go-errors/http"
)

var UserNotFound = errors.Error{
	Kind:        errors.NotFound,
	Code:        "USER_NOT_FOUND",
	Message:     "This user was not found",
	Description: "User was not found in the database. It was probably deleted during clean up",
}

func main() {
	err := getUsers()
	if errors.IsNotFound(err) {
		fmt.Println(err.Error())
	}

	fmt.Println(errors.IsNotFound(err))

	fmt.Println(errors.IsUnauthorized(err))

	fmt.Println("HTTP: ", http.StatusCode(err))
	fmt.Println("GRPC: ", grpc.StatusCode(err))

	fmt.Println(errors.Get(err))

}

func getUsers() error {
	srcErr := errors.New(UserNotFound)
	err := errors.NewNotFound("Item was not found").Wrap(srcErr)
	fmt.Printf("%+v\n", err)
	// return errors.NewUnauthorized("Item")
	return err
}
