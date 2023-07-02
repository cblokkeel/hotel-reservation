package types

import (
	"errors"
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      int = 12
	minFirstNameLen int = 2
	minLastNameLen  int = 2
	minPasswordLen  int = 6
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func ParamsToUserValidated(params CreateUserParams) (*User, error) {
	encpwd, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpwd),
	}, nil
}

func (params *CreateUserParams) Validate() error {
	errorsSlice := []error{}
	if len(params.FirstName) < minFirstNameLen {
		errorsSlice = append(errorsSlice, fmt.Errorf("firstName should have minimal length of %d\n", minFirstNameLen))
	}
	if len(params.LastName) < minLastNameLen {
		errorsSlice = append(errorsSlice, fmt.Errorf("lastName should have minimal length of %d\n", minLastNameLen))
	}
	if len(params.Password) < minPasswordLen {
		errorsSlice = append(errorsSlice, fmt.Errorf("password should have minimal length of %d\n", minPasswordLen))
	}
	if !isEmailValid(params.Email) {
		errorsSlice = append(errorsSlice, fmt.Errorf("unvalid email"))
	}
	return errors.Join(errorsSlice...)
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
