package utils

import "testing"

func TestSlug(t *testing.T) {
	title := "Hello world!"

	t.Logf("%s", Slugify(title))
}
