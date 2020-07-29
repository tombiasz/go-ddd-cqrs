package coupon

import "testing"

func TestNewCode(t *testing.T) {
	t.Run("returns error when empty code is passed", func(t *testing.T) {
		_, err := NewCode("")

		if err != CodeCannotBeEmptyErr {
			t.Errorf("got %q, want %q", err, CodeCannotBeEmptyErr)
		}
	})

	t.Run("returns error when code is less than 8 chars long", func(t *testing.T) {
		var code7CharsLong = "aabbccd"

		_, err := NewCode(code7CharsLong)

		if err != CodeIsInvalidErr {
			t.Errorf("got %q, want %q", err, CodeIsInvalidErr)
		}
	})

	t.Run("returns error when address is longer than 11 chars long", func(t *testing.T) {
		var code12CharsLong = "aabbccddeeff"

		_, err := NewCode(code12CharsLong)

		if err != CodeIsInvalidErr {
			t.Errorf("got %q, want %q", err, CodeIsInvalidErr)
		}
	})

	t.Run("creates a code", func(t *testing.T) {
		var validCode = "aabbccddee"

		code, _ := NewCode(validCode)

		if code.value != validCode {
			t.Errorf("got %q, want %q", validCode, code.value)
		}
	})
}

func TestCodeValue(t *testing.T) {
	t.Run("returns value as lower case string", func(t *testing.T) {
		var input = "AABBccDDee"
		var expected = "aabbccddee"

		code, _ := NewCode(input)

		if code.Value() != expected {
			t.Errorf("got %q, want %q", input, expected)
		}
	})
}
