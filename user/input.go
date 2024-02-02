package user

type RegisterUser struct {
	Name       string `json:"name" validate:"required,omitempty" structs:"omitempty"`
	Occupation string `json:"occupation" validate:"required,omitempty" structs:"omitempty"`
	Email      string `json:"email" validate:"required,omitempty,email"`
	Password   string `validate:"required,min=1,max=100" json:"password"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}
