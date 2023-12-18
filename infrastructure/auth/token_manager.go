package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type Token struct{}

type AccessDetails struct {
	TokenUuid string
	UserId    uint64
}

func NewToken() *Token {
	return &Token{}
}

type TokenInterface interface {
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

// Token implements the TokenInterface
var _ TokenInterface = &Token{}

func TokenValid(r *http.Request) (userId uint64, err error) {
	// Verify token
	token, err := VerifyToken(r)
	if err != nil {
		return 0, err
	}
	// Check token validity
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Extract token metadata
	metadata, err := NewToken().ExtractTokenMetadata(r)
	if err != nil {
		return 0, err
	}

	return metadata.UserId, nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return []byte(r.Header.Get("OSP-Secret-Key")), nil
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		logrus.Error("Parsing Token, ", err)
		return nil, err
	}
	return token, nil
}

// get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	// fmt.Println("WE ENTERED METADATA")
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			TokenUuid: accessUuid,
			UserId:    userId,
		}, nil
	}
	return nil, err
}
