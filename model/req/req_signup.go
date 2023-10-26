package req

type ReqSignUp struct {
	FullName string `validate:"required" json:"fullName"`
	Email string `validate:"required,email" json:"email"`
	Password string `validate:"required,pwd" json:"password"`
}