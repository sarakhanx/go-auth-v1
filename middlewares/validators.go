package validate

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	//NOTE -  Simple regex for email validation
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
