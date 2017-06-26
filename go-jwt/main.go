package main

// In this example, I use the popular jwt-go package.
import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

// customValidation is the function passed to `jwt.Parse` later.  The
// examples all use an anonymous function in the `Parse()` call.  But I
// prefer to split it out for now, so that I can separate thinking about
// the pieces for now.  I can do whatever checks I want to in this
// function.  And then at the end, I need to make sure I return the
// HMAC secret as a `[]byte`.
func customValidation(token *jwt.Token) (interface{}, error) {
	// This snippet is lifted from the example in the docs.  They
	// recommend checking that the correct algorithm was actually
	// used.  I need to review this again later.  I don't know what
	// `token.Method.(*jwt.SigningMethodHMAC)` is doing.
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	// I read my secret out a file and remove the trailing newline.
	// I am ignoring the errors for now.
	secret, _ := ioutil.ReadFile("secret.txt")
	secret = secret[:len(secret)-1]
	return secret, nil
}

func main() {
	// I read my token out a file and remove the trailing newline.
	// I am ignoring the errors for now.
	jwtoken, _ := ioutil.ReadFile("token.txt")
	jwtoken = jwtoken[:len(jwtoken)-1]
	// This will read the token and validate it using the function
	// I defined earlier.
	token, _ := jwt.Parse(string(jwtoken), customValidation)
	fmt.Println(token)
}
