package apihelper

import "fmt"

func CreateURL(baseURL, firstSuffix, secondSuffix string) string {
	result := fmt.Sprintf("%s%s/%s", baseURL, firstSuffix, secondSuffix)
	return result
}
