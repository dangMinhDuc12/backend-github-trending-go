package req


type ReqUpdate struct {
	FullName string `validate:"required" json:"fullName"`
	Email string `validate:"required,email" json:"email"`
}