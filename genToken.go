package main

import (
	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)
 const (
     APP_ID     =  "1054249344749408"
     APP_SECRET =  "dca087c5d2327b08f82580ff7a358d9d"
)

// GenerateAccessToken is a sample that shows you
// how to generate facebook access_token from server side
func GenerateAccessToken() string {
	// initalize facebook-go-sdk
	fb := facebook.New()

	// set your application id (client_id)
	fb.SetAppID(`{APP_ID}`)

	// set your application secret (client_secret)
	fb.SetAppSecret(`{APP_SECRET}`)

	// Call https://www.facebook.com/v2.10/dialog/oauth?client_id={APP_ID}&redirect_uri={YOUR_REDIRECT_URI}
	// And then, retrieve the `code` value from {YOUR_REDIRECT_URI}?code=....

	// Then call GenerateAccessToken(redirectURI string, code string) method
	data, err := fb.GenerateAccessToken(`{YOUR_REDIRECT_URI}`, `{GENERATED_CODE}`)
	if err != nil {
		panic(err)
	}

	// print data object
       st :=	fmt.Sprintf(`%+v`, data)
       return  st
}
