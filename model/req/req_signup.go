package req

type ReqSignUp struct {
	FullName string `validate:"required" json:"fullName"`
	Email string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}