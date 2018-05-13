package themes

import (
	"regexp"
	"strings"
)

func helperSlugify(label string) string {
	re := regexp.MustCompile("[^a-z0-9а-я]+")
	return strings.Trim(re.ReplaceAllString(strings.ToLower(label), "-"), "-")
}
