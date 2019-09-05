package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/logrusorgru/aurora"
)

func signToken(data jwt.MapClaims, signedMethod jwt.SigningMethod, secret string) (string, error) {
	// set default iat
	_, iat := data["iat"]
	if !iat {
		data["iat"] = time.Now().Unix()
	}
	// sign token
	token := jwt.NewWithClaims(signedMethod, data)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
func verifySignedMethod(method string, t *jwt.Token) bool {
	signedMethod := jwt.GetSigningMethod(method)
	return signedMethod.Alg() == t.Method.Alg()

}

// parseToken --
func parseToken(tokenString, secret, signedMethod string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		ok := verifySignedMethod(signedMethod, token)
		if !ok {
			return nil, fmt.Errorf("invalid signed method")
		}
		return []byte(secret), nil
	})
	if token == nil {
		fmt.Print(Red(err))
		return
	}
	printToken(token, err)
}

// printToken --
func printToken(t *jwt.Token, err error) {
	// print header
	fmt.Println(Bold(BrightMagenta("HEADER")))
	for k, v := range t.Header {
		fmt.Println(Blue(k+": "), Yellow(v))
	}

	claims, _ := t.Claims.(jwt.MapClaims)
	// payload json
	fmt.Println(Bold(BrightMagenta("PAYLOAD")))
	jsonStr, _ := json.MarshalIndent(claims, "", " ")
	fmt.Println(White(string(jsonStr)))
	// payload info
	fmt.Println(Bold(BrightMagenta("INFO")))
	keys := sortedKey(claims)
	for _, k := range keys {
		v := claims[k]
		switch k {
		case "exp":
			fmt.Println(Blue(k+": "), formatExp(v))
		case "iat", "nbf":
			fmt.Println(Blue(k+": "), White(formatTime(castTime(v))))
		}
	}
	// sig
	fmt.Println(Bold(BrightMagenta("SIG")))
	if err != nil {
		fmt.Println(Red(fmt.Sprintf("%s %v", t.Signature, err)))
	} else {
		fmt.Println(Green(fmt.Sprintf("%s %v", t.Signature, err)))
	}
}
func sortedKey(c jwt.MapClaims) []string {
	var keys []string
	for k := range c {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

/*
## Common attribute
Audience  string `json:"aud,omitempty"`
ExpiresAt int64  `json:"exp,omitempty"`
Id        string `json:"jti,omitempty"`
IssuedAt  int64  `json:"iat,omitempty"`
Issuer    string `json:"iss,omitempty"`
NotBefore int64  `json:"nbf,omitempty"`
Subject   string `json:"sub,omitempty"`
*/
func formatExp(exp interface{}) string {
	tt := castTime(exp)
	formated := formatTime(tt)
	now := time.Now()
	if tt.Before(now) { // expired
		return fmt.Sprintf("%s %s", formated, Red("Expire In: -"+time.Since(tt).String()))
	}
	//
	since := time.Since(tt).String()
	return fmt.Sprintf("%s %s", formated, Green("Expire In: "+since[1:len(since)]))
}

func castTime(e interface{}) time.Time {
	t := e.(float64)
	tt := time.Unix(int64(t), 0)
	return tt
}
func formatTime(t time.Time) string {
	ttFormated := t.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%v %s%s", t.Unix(), Blue("formated:"), ttFormated)
}
