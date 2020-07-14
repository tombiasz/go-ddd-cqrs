package coupon

import "testing"

var email2CharsLong = "aa"
var email256CharsLong = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com"
var validEmail = "foo@bar.com"

func TestCreateEmail(t *testing.T) {
	t.Run("returns error when empty address is passed", func(t *testing.T) {
		_, err := CreateEmail("")

		if err != EmailCannotBeEmptyErr {
			t.Errorf("got %q, want %q", err, EmailCannotBeEmptyErr)
		}
	})

	t.Run("returns error when address is less than 3 chars long", func(t *testing.T) {
		_, err := CreateEmail(email2CharsLong)

		if err != EmailIsInvalidErr {
			t.Errorf("got %q, want %q", err, EmailIsInvalidErr)
		}
	})

	t.Run("returns error when address is longer than 255 chars long", func(t *testing.T) {
		_, err := CreateEmail(email256CharsLong)

		if err != EmailIsInvalidErr {
			t.Errorf("got %q, want %q", err, EmailIsInvalidErr)
		}
	})

	t.Run("creates a email", func(t *testing.T) {
		email, _ := CreateEmail(validEmail)

		if email.address != validEmail {
			t.Errorf("got %q, want %q", validEmail, email.address)
		}
	})
}
