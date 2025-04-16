package register

import (
	"fmt"
	"os/user"

	"main.go/models"
)

func RegisterService(userDetails UserProfile) error {
	// Example logic
	if(userDetails.password && userDetails.emailAddress)
	fmt.Printf("Registering user with email: %s\n", userDetails.emailAddress)
	return nil
}
