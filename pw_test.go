package pw

import (
	"strings"
	"testing"
)

func contains(chars string, password string) bool {
	for _, char := range strings.Split(chars, "") {
		if strings.Contains(password, char) {
			return true
		}
	}

	return false
}

// containsAll checks if password contains all chars in the charset
func containsAll(password, charset string) bool {
	for _, char := range strings.Split(charset, "") {
		if !strings.Contains(password, char) {
			return false
		}
	}
	return true
}

func TestDefaultNewPassword(t *testing.T) {
	l := 100
	p := NewPassword(l)

	if len(p) != l {
		t.Errorf("expected password of length %v, got %v", l, len(p))
	}

	if !contains(SpecialChars, p) {
		t.Errorf("should contain special characters")
	}

	if !contains(UpperChars, p) {
		t.Errorf("should contain upper characters")
	}

	if !contains(LowerChars, p) {
		t.Errorf("should contain lower characters")
	}

	if !contains(NumberChars, p) {
		t.Errorf("should contain number characters")
	}
}

func TestUpperCaseOnlyNewPassword(t *testing.T) {
	l := 100
	p := NewPassword(l, WithUpperCase())

	if len(p) != l {
		t.Errorf("expected password of length %v, got %v", l, len(p))
	}

	if !contains(p, UpperChars) {
		t.Errorf("should contain upper characters")
	}

	if contains(p, SpecialChars) {
		t.Errorf("should not contain special characters")
	}

	if contains(p, LowerChars) {
		t.Errorf("should not contain lower characters")
	}

	if contains(p, NumberChars) {
		t.Errorf("should not contain number characters")
	}
}

func TestLowerCaseOnlyNewPassword(t *testing.T) {
	l := 100
	p := NewPassword(l, WithLowerCase())

	if len(p) != l {
		t.Errorf("expected password of length %v, got %v", l, len(p))
	}

	if !contains(p, LowerChars) {
		t.Errorf("should contain lower characters")
	}

	if contains(p, SpecialChars) {
		t.Errorf("should not contain special characters")
	}

	if contains(p, UpperChars) {
		t.Errorf("should not contain upper characters")
	}

	if contains(p, NumberChars) {
		t.Errorf("should not contain number characters")
	}
}

func TestNumberOnlyPassword(t *testing.T) {
	l := 100
	p := NewPassword(l, WithNumbers())

	if len(p) != l {
		t.Errorf("expected password of length %v, got %v", l, len(p))
	}

	if !contains(p, NumberChars) {
		t.Errorf("should contain number characters")
	}

	if contains(p, SpecialChars) {
		t.Errorf("should not contain special characters")
	}

	if contains(p, LowerChars) {
		t.Errorf("should not contain lower characters")
	}

	if contains(p, UpperChars) {
		t.Errorf("should not contain upper characters")
	}
}

func TestSpecialCharOnlyPassword(t *testing.T) {
	l := 100
	p := NewPassword(l, WithSpecial())

	if len(p) != l {
		t.Errorf("expected password of length %v, got %v", l, len(p))
	}

	if !contains(p, SpecialChars) {
		t.Errorf("should contain special characters")
	}

	if contains(p, NumberChars) {
		t.Errorf("should not contain number characters")
	}

	if contains(p, LowerChars) {
		t.Errorf("should not contain lower characters")
	}

	if contains(p, UpperChars) {
		t.Errorf("should not contain upper characters")
	}
}

func TestDistribution(t *testing.T) {
	const charset = "12"
	pass := randomFromChars(100, charset)
	if !containsAll(pass, charset) {
		t.Errorf("should contain all the chars from charset %q\n", charset)
	}
}
