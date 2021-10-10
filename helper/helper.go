package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("required email")
	}
	if email != "" {
		if err := ValidateEmailFormat(email); err != nil {
			return errors.New("invalid email")
		}
	}
	return nil
}

func ValidateEmailFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func ResponseJSON(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, obj)
}
