package apihelper

import "fmt"

//CreateURL formats URL's for the bad api
func CreateURL(baseURL, firstSuffix, secondSuffix string) string {
	result := fmt.Sprintf("%s%s/%s", baseURL, firstSuffix, secondSuffix)
	return result
}
