package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct{
	secret string
}

type Claims struct{
	UserID string `json:"sub"`
	jwt.RegisteredClaims
}

func NewJWTManager(secret string) *JWTManager{
	return	&JWTManager{secret: secret}
}

func(j *JWTManager) Generate(userID string)(string, error){
	//here jwt.MapClaims is to tell who this token is for, sub means (subject), which means who this token is for, and 
	//exp is expiration time
	claims:=Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
		},
	}
	// claims:=jwt.MapClaims{
	// 	"sub":userID,
	// 	"exp":time.Now().Add(24*time.Hour).Unix(),
	// }

	//jwt.NewWithClaims combine token with claims
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(j.secret))
}

func(j *JWTManager) Parse(tokenSring string)(*Claims,error){
	token,err:=jwt.ParseWithClaims(tokenSring,&Claims{},func (token *jwt.Token)(any,error){
		return []byte(j.secret),nil
	})
	if err!=nil{
		return nil,err
	}
	claims,ok:=token.Claims.(*Claims)
	if !ok{
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims,nil
}