package validate

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"unicode"
)

func PositiveInt(num uint) error {
	if num < 1 {
		return errors.New("id must be greater than 0")
	}
	return nil
}

func NonEmptyStr(val string) error {
	if val == "" {
		return errors.New("string is empty")
	}
	return nil
}

func NonEmptyByteArr(val []byte) error {
	if len(val) == 0 {
		return errors.New("byte array is empty")
	}
	return nil
}

func IsEmailAddress(email string) error {
	if govalidator.IsEmail(email) {
		return nil
	}
	return errors.New("invalid email address")
}

func IsStrongPassword(pwd string, minLength, minNumeric, minCapital, minSpecial uint64) error {

	if uint64(len(pwd)) < minLength {
		return fmt.Errorf("length must be at least %d characters", minLength)
	}

	var (
		numericCount uint64
		capitalCount uint64
		specialCount uint64
	)

	for _, char := range pwd {
		switch {
		case unicode.IsNumber(char):
			numericCount++
		case unicode.IsUpper(char):
			capitalCount++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			specialCount++
		}
	}

	if numericCount < minNumeric {
		return fmt.Errorf("at least %d number(s) required", minNumeric)
	}

	if capitalCount < minCapital {
		return fmt.Errorf("at least %d capital letter(s) required", minCapital)
	}

	if specialCount < minSpecial {
		return fmt.Errorf("at least %d special character(s) required", minSpecial)
	}

	return nil
}
