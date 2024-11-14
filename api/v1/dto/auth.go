package dto

type SignUpRequest struct {
	FirstName  string `json:"first_name" form:"first_name" validate:"required"`
	MiddleName string `json:"middle_name" form:"middle_name" validate:"required"`
	LastName   string `json:"last_name" form:"last_name" validate:"required"`
	Email      string `json:"email" form:"email" validate:"required"`
	Username   string `json:"username" form:"username" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
}
