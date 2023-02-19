package user

import (
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateSignupUser validates signup user request
func (us *UserService) ValidateSignupUser(req *userpb.SignupRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Email, validation.Required, is.Email, validation.Max(255)),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
		validation.Field(&req.Name, validation.Required, validation.Max(255)),
	)
}

// ValidateSignInUser validates signIn user request
func (us *UserService) ValidateSignInUser(req *userpb.SignInRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Email, validation.Required, is.Email, validation.Max(255)),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
