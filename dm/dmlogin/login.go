package dmlogin

import "github.com/dgrijalva/jwt-go"

const (
	// jwt token签发者
	ISSUER     = "libo.explorer"
	SIGNED_KEY = "datastructLearningWeb"
)

type LoginReq struct {
	Number   string `json:"number"`
	Password string `json:"password"`
}

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}
