package utils

import (
	"errors"
	"gin/basic/model"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToekn(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	myToken, err := token.SignedString([]byte("gotutorial"))
	if err != nil {
		return "", err
	}
	log.Println(myToken)
	return myToken, nil
}
func ValidateToken(userToken string) (bool, error) {
	jwtKeyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Some unexpected error")
		}
		return []byte("gotutorial"), nil
	}
	token, err := jwt.Parse(userToken, jwtKeyfunc)
	if err != nil {
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, errors.New("Token is Invalid")
}
