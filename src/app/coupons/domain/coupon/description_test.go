package coupon

import "testing"

func TestCreateDescription(t *testing.T) {
	t.Run("returns error when empty value is passed", func(t *testing.T) {
		_, err := CreateDescription("")

		if err != DescriptionCannotBeEmptyErr {
			t.Errorf("got %q, want %q", err, DescriptionCannotBeEmptyErr)
		}
	})

	t.Run("returns error when value longer than 200 chars long", func(t *testing.T) {
		var desc201CharsLong = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer et urna efficitur, sodales nibh sit amet, dapibus ipsum. Quisque malesuada libero in diam fermentum rutrum. Mauris eget porttitor proin."

		_, err := CreateDescription(desc201CharsLong)

		if err != DescriptionCannotBeLongerThan200CharsErr {
			t.Errorf("got %q, want %q", err, DescriptionCannotBeLongerThan200CharsErr)
		}
	})

	t.Run("creates a description", func(t *testing.T) {
		var validDesc = "Lorem ipsum dolor sit amet."

		desc, _ := CreateDescription(validDesc)

		if desc.value != validDesc {
			t.Errorf("got %q, want %q", validDesc, desc.value)
		}
	})
}
