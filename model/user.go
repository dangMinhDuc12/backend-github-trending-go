package model

import "time"

type User struct {
	UserId string `db:"user_id, omitempty" json:"userId"` //add tag for sqlx
	FullName string `db:"full_name, omitempty" json:"fullName"`
	Email string `db:"email, omitempty" json:"email"`
	Password string `db:"password,omitempty" json:"password,omitempty"`
	Role string `db:"role, omitempty" json:"role"`
	CreatedAt time.Time `db:"created_at, omitempty" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at, omitempty" json:"updatedAt"`
	Token string `json:"token"`
}