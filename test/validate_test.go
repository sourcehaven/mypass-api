package test

import (
	"github.com/sourcehaven/mypass-api/internal/validate"
	"testing"
)

func TestPositiveInt(t *testing.T) {
	t.Run("Above zero", func(t *testing.T) {
		if validate.PositiveInt(100) != nil {
			t.Fail()
		}
		if validate.PositiveInt(1) != nil {
			t.Fail()
		}
	})
	t.Run("Zero", func(t *testing.T) {
		if validate.PositiveInt(0) == nil {
			t.Fail()
		}
	})
}

func TestNonEmptyStr(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		if validate.NonEmptyStr("") == nil {
			t.Fail()
		}
	})
	t.Run("Non empty string", func(t *testing.T) {
		if validate.NonEmptyStr("text") != nil {
			t.Fail()
		}
	})
}

func TestNonEmptyByteArr(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		if validate.NonEmptyByteArr([]byte("")) == nil {
			t.Fail()
		}
	})
	t.Run("Non empty string", func(t *testing.T) {
		if validate.NonEmptyByteArr([]byte("text")) != nil {
			t.Fail()
		}
	})
}

func TestEmailAddress(t *testing.T) {
	t.Run("Valid email addresses", func(t *testing.T) {
		if validate.IsEmailAddress("xy@example.com") != nil {
			t.Fail()
		}
		if validate.IsEmailAddress("hello@there.com") != nil {
			t.Fail()
		}
	})
	t.Run("Invalid email addresses", func(t *testing.T) {
		if validate.IsEmailAddress("example@gmail.") == nil {
			t.Fail()
		}
		if validate.IsEmailAddress("gmail") == nil {
			t.Fail()
		}
		if validate.IsEmailAddress("xy@gmail") == nil {
			t.Fail()
		}
	})
}

func TestIsStrongEnoughPassword(t *testing.T) {
	t.Run("Valid passwords", func(t *testing.T) {
		if validate.IsStrongPassword("password", 8, 0, 0, 0) != nil {
			t.Error("password is at least 8 character, hence should not have failed")
		}
		if validate.IsStrongPassword("pa$Sw0rd", 8, 1, 1, 1) != nil {
			t.Error("pa$Sw0rd is at least 8 character and has 1 special, numeric and capital character, hence should not have failed")
		}
	})

	t.Run("Invalid passwords", func(t *testing.T) {
		if validate.IsStrongPassword("password", 12, 0, 0, 0) == nil {
			t.Fail()
		}
		if validate.IsStrongPassword("pa$Sw0rd", 8, 2, 1, 1) == nil {
			t.Fail()
		}
		if validate.IsStrongPassword("pa$Sw0rd", 8, 1, 2, 1) == nil {
			t.Fail()
		}
		if validate.IsStrongPassword("pa$Sw0rd", 8, 1, 1, 3) == nil {
			t.Fail()
		}
	})
}
