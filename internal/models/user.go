package models

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Name     string
	Password string
	Surname  string
	Email    string
	Addres   string
	Phone    string
	Role     string
}

var (
	ErrWrongPasswordOrEmail = errors.New("wrong email or/and password")
	ErrEmailExists          = errors.New("this email exists")
	ErrWrongPassword        = errors.New("wrong Password")
	ErrEmailNotValid        = errors.New("email not valid")
)

const defaultRole = "user"

func (u *User) Validate(fs ...func() error) error {
	for _, f := range fs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (u *User) ValidatePassword() error {
	if len(u.Password) < 4 || len(u.Password) > 35 {
		return ErrWrongPassword
	}
	return nil
}

func (u *User) ValidateEmail() error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(u.Email) {
		return ErrEmailNotValid
	}
	return nil
}

func (u *User) CleanPassword() {
	u.Password = ""
}

func (u *User) Standardize() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if u.Role != "" {
		u.Role = strings.ToLower(strings.TrimSpace(u.Role))
	} else {
		u.Role = defaultRole
	}
}

func (u *User) ComparePassword(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
