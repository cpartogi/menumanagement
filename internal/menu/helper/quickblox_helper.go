package helper

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type qbPayload struct {
	ApplicationID int    `json:"application_id"`
	AuthKey       string `json:"auth_key"`
	Signature     string `json:"signature"`
	Timestamp     int64  `json:"timestamp"`
	Nonce         int64  `json:"nonce"`
}

type Body struct {
	Session token `json:"session"`
}

type token struct {
	ApplicationID int       `json:"application_id"`
	CreatedAt     time.Time `json:"created_at"`
	ID            int64     `json:"id"`
	Ts            int64     `json:"ts"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserID        int       `json:"user_id"`
	SecID         string    `json:"_id"`
	Token         string    `json:"token"`
}

func GetSignature(ts, nonce int64) string {
	field := fmt.Sprintf("application_id=81342&auth_key=jtzxLqpQVtjKuBV&nonce=%d&timestamp=%d", nonce, ts)
	key_for_sign := []byte("yTwjZtPPdzBKLNc")
	h := hmac.New(sha1.New, key_for_sign)
	h.Write([]byte(field))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateQBToken() string {
	ts := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 9999
	nonce := int64(rand.Intn(max-min+1) + min)
	var respBody UrlHttpResponse

	signature := GetSignature(ts, nonce)
	data := qbPayload{
		ApplicationID: 81342,
		AuthKey:       "jtzxLqpQVtjKuBV",
		Signature:     signature,
		Timestamp:     ts,
		Nonce:         nonce,
	}

	payload, _ := json.Marshal(data)
	apiCall := APICall{
		URL:       "https://api.quickblox.com/session.json",
		Method:    http.MethodPost,
		FormParam: string(payload),
	}

	resp, err := apiCall.Call()
	if err == nil {
		json.Unmarshal([]byte(resp.Body), &respBody)
	}

	if resp.StatusCode > 299 {
		log.Println(resp.Body)
		return ""
	}

	var ss Body
	err = json.Unmarshal([]byte(resp.Body), &ss)
	if err != nil {
		log.Println(err)
	}

	return ss.Session.Token
}
