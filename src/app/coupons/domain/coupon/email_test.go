package coupon

import "testing"

func TestNewEmail(t *testing.T) {
	t.Run("returns error when empty address is passed", func(t *testing.T) {
		_, err := NewEmail("")

		if err != EmailCannotBeEmptyErr {
			t.Errorf("got %q, want %q", err, EmailCannotBeEmptyErr)
		}
	})

	t.Run("returns error when address is less than 3 chars long", func(t *testing.T) {
		var email2CharsLong = "aa"

		_, err := NewEmail(email2CharsLong)

		if err != EmailIsInvalidErr {
			t.Errorf("got %q, want %q", err, EmailIsInvalidErr)
		}
	})

	t.Run("returns error when address is longer than 255 chars long", func(t *testing.T) {
		var email256CharsLong = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com"

		_, err := NewEmail(email256CharsLong)

		if err != EmailIsInvalidErr {
			t.Errorf("got %q, want %q", err, EmailIsInvalidErr)
		}
	})

	t.Run("creates a email", func(t *testing.T) {
		var validEmail = "foo@bar.com"

		email, _ := NewEmail(validEmail)

		if email.address != validEmail {
			t.Errorf("got %q, want %q", validEmail, email.address)
		}
	})
}

func TestEmailAddress(t *testing.T) {
	t.Run("returns email address as lower case string", func(t *testing.T) {
		var input = "FOO@bar.com"
		var expected = "foo@bar.com"

		e, _ := NewEmail(input)

		if e.Address() != expected {
			t.Errorf("got %q, want %q", input, expected)
		}
	})
}
