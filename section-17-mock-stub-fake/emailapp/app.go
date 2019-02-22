package emailapp

import "strings"

type MailgunClient struct {
}

func (mc *MailgunClient) Welcome(name, email string) error {
	return nil
}

type User struct{}
type UserStore struct{}

func (us *UserStore) Create(name, email string) (*User, error) {
	return &User{}, nil
}

type EmailClient interface {
	Welcome(name, email string) error
}

func Signup(name, email string, ec EmailClient, us *UserStore) (*User, error) {
	email = strings.ToLower(email)
	user, err := us.Create(name, email)
	if err != nil {
		return nil, err
	}
	err = ec.Welcome(name, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
