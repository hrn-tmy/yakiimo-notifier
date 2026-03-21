package handler

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
	"yakiimo-notifier/internal/constant"
	"yakiimo-notifier/internal/gen"
)

func validateCreateUser(req gen.CreateUserRequest) string {
	errs := make([]string, 0)

	if req.Email == "" {
		errs = append(errs, fmt.Sprintf(constant.Required, constant.Email))
	}
	if utf8.RuneCountInString(string(req.Email)) > 255 {
		errs = append(errs, fmt.Sprintf(constant.MaxLength, constant.Email, 255))
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(string(req.Email)) {
		errs = append(errs, fmt.Sprintf(constant.Format, constant.Email))
	}

	if req.Name == "" {
		errs = append(errs, fmt.Sprintf(constant.Required, constant.Name))
	}
	if utf8.RuneCountInString(req.Name) > 255 {
		errs = append(errs, fmt.Sprintf(constant.MaxLength, constant.Name, 255))
	}

	if req.Password == "" {
		errs = append(errs, fmt.Sprintf(constant.Required, constant.Password))
	}
	if utf8.RuneCountInString(req.Password) < 12 {
		errs = append(errs, fmt.Sprintf(constant.MinLength, constant.Password, 12))
	}
	if utf8.RuneCountInString(req.Password) > 20 {
		errs = append(errs, fmt.Sprintf(constant.MaxLength, constant.Password, 20))
	}
	if !validatePasswordKinds(req.Password) {
		errs = append(errs, constant.PasswordKind)
	}

	if len(errs) > 0 {
		return strings.Join(errs, ",")
	}

	return ""
}

func validateNotifyReady(req gen.PostNotifyReadyJSONRequestBody) string {
	var errs []string
	if req.MachineId == "" {
		errs = append(errs, fmt.Sprintf(constant.Required, constant.MachineID))
	}
	if req.FinishedAt == "" {
		errs = append(errs, fmt.Sprintf(constant.Required, constant.FinishedAt))
	}
	if req.Quantity <= 0 {
		errs = append(errs, fmt.Sprintf(constant.Min, constant.Quantity, 1))
	}

	if len(errs) > 0 {
		return strings.Join(errs, ",")
	}

	return ""
}

func validatePasswordKinds(password string) bool {
	var hasUpper, hasLower, hasDigit, hasSymbol bool
	for _, p := range password {
		switch {
		case unicode.IsUpper(p):
			hasUpper = true
		case unicode.IsLower(p):
			hasLower = true
		case unicode.IsDigit(p):
			hasDigit = true
		default:
			hasSymbol = true
		}
	}

	var count int
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasDigit {
		count++
	}
	if hasSymbol {
		count++
	}

	return count >= 3
}
