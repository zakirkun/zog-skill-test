package utils

import (
	"regexp"
	"strings"
)

// Slugify generates a slug from the given title
func Slugify(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Remove all non-alphanumeric characters except for spaces
	reg, _ := regexp.Compile("[^a-z0-9 ]+")
	slug = reg.ReplaceAllString(slug, "")

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	return slug
}
