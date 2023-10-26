package req

type ReqSignIn struct {
	Email string `validate:"required,email" json:"email"`
	Password string `validate:"required,pwd" json:"password"`
}