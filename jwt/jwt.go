package jwt

import (
	"encoding/json"
	"time"
)

type Jwt struct {
	Header    JwtHeader
	Payload   JwtPayload
	Signature string
}
type JwtHeader struct {
	Type      string `json:"typ"`
	Algorithm string `json:"alg"`
}
type JwtPayload struct {
	ExpireAt time.Time   `json:"exp"`
	Data     interface{} `json:"data"`
}

func Publish(alg string, data interface{}) (*Jwt, error) {
	jwt = &Jwt{
		JwtHeader{
			"JWT",
			alg,
		},
		JwtPayload{
			time.Now(),
			data,
		},
	}
	jwt.makeSignature()
	return jwt
}

func (jwt *Jwt) makeSignature() {
	jwt.Signature = jwt.getSignature()
}

func (jwt *Jwt) getSignature() string {

}

func base64UrlEncode(s string) string {

}

func (jwt *Jwt) Authenticate() (bool, error) {

}
