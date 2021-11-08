package interfaces

import (
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Id        int64  `json:"id"`
	UserTitle string `json:"user_title"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	Town      string `json:"town"`
	IsActive  int8   `json:"is_active"`
}

type RegisterReq struct {
	UserTitle string `json:"user_title,required"`
	UserName  string `json:"user_name,required"`
	UserPass  string `json:"user_pass,required"`
	Phone     string `json:"phone,required"`
	City      string `json:"city,required"`
	Town      string `json:"town,required"`
}

type LoginReq struct {
	UserName string `json:"user_name,required"`
	UserPass string `json:"user_pass,required"`
}

type ForgetPasswordReq struct {
	UserName string `json:"user_name,required"`
}

type Verify struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code,ommitempty"`
}

type TokenClaim struct {
	Id        int64  `json:"id"`
	UserTitle string `json:"user_title"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	Town      string `json:"town"`
	Exp       int64  `json:"exp"`
	Iat       int64  `json:"iat"`
	jwt.StandardClaims
}
