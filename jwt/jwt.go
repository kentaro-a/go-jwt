package jwt

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	_ "fmt"
	"strings"
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

// const JwtExpiredTime time.Duration = time.Minute * 60 * 24	// 1day
const JwtExpiredTime time.Duration = time.Minute * 1
const Alg string = "HS256"
const SecretKey string = "xxxx"

func Publish(data interface{}) (*Jwt, error) {
	jwt := &Jwt{
		Header: JwtHeader{
			"JWT",
			Alg,
		},
		Payload: JwtPayload{
			time.Now().Add(JwtExpiredTime),
			data,
		},
	}
	err := jwt.setSignature()
	return jwt, err
}

func (jwt *Jwt) Encode() string {
	header, _ := json.Marshal(jwt.Header)
	payload, _ := json.Marshal(jwt.Payload)
	return base64UrlEncode(string(header)) + "." + base64UrlEncode(string(payload)) + "." + jwt.Signature
}
func Decode(jwt_str string) (*Jwt, error) {
	jwt_splits := strings.Split(jwt_str, ".")
	if len(jwt_splits) != 3 {
		return nil, errors.New("Invalid auth code")
	}

	_header, err := base64UrlDecode(jwt_splits[0])
	if err != nil {
		return nil, errors.New("Invalid auth code")
	}
	jwt_header := &JwtHeader{}
	err = json.Unmarshal([]byte(_header), jwt_header)
	if err != nil {
		return nil, errors.New("Invalid auth code")
	}

	_payload, err := base64UrlDecode(jwt_splits[1])
	if err != nil {
		return nil, errors.New("Invalid auth code")
	}
	jwt_payload := &JwtPayload{}
	err = json.Unmarshal([]byte(_payload), jwt_payload)
	if err != nil {
		return nil, errors.New("Invalid auth code")
	}

	_signature := jwt_splits[2]

	return &Jwt{Header: *jwt_header, Payload: *jwt_payload, Signature: _signature}, nil
}

func (jwt *Jwt) setSignature() error {
	s, err := jwt.calcSignature()
	if err == nil {
		jwt.Signature = s
	}
	return err
}

func (jwt *Jwt) calcSignature() (string, error) {
	header, err := json.Marshal(jwt.Header)
	if err != nil {
		return "", err
	}
	payload, err := json.Marshal(jwt.Payload)
	if err != nil {
		return "", err
	}
	in := string(header) + "." + string(payload)
	h := hash(in)
	return base64UrlEncode(h), nil
}

func hash(s string) string {
	s = s + SecretKey
	b := sha256.Sum256([]byte(s))
	return string(b[:])
}

func base64UrlEncode(s string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(s))
}

func base64UrlDecode(s string) (string, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func (jwt *Jwt) Authenticate() (bool, error) {
	sig, err := jwt.calcSignature()
	if err == nil {
		// If Signature is valid and timestamp is not expired then true.
		if jwt.Signature == sig && jwt.Payload.ExpireAt.After(time.Now()) {
			return true, nil
		}
	}
	return false, nil
}
