package linelogin

import (
	"encoding/json"
	"io"
	extensions "monkey-in-mountain-pass/extensions"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type VerifyIdTokenRequestDto struct {
	Issuer                string   `json:"iss"`
	Subject               string   `json:"sub"`
	Audience              string   `json:"aud"`
	Expiration            int      `json:"exp"`
	IssuedAt              int      `json:"iat"`
	Nonce                 string   `json:"nonce"`
	AuthenticationMethods []string `json:"amr"`
	Username              string   `json:"name"`
	PictureUrl            string   `json:"picture"`
	Email                 string   `json:"email"`
}

func VerifyIdToken(idToken string) (*VerifyIdTokenRequestDto, error) {
	payload := createVerifyIdTokenPayload(idToken)
	resp, err := requestVerifyIdToken(payload)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var respDto VerifyIdTokenRequestDto
	err = json.Unmarshal(respBody, &respDto)
	if err != nil {
		return nil, err
	}
	return &respDto, nil
}

func createVerifyIdTokenPayload(idToken string) url.Values {
	return url.Values{
		"id_token":  {idToken},
		"client_id": {extensions.GetEnvironment("CLIENT_ID_FOR_LINE_LOGIN")},
	}
}

func requestVerifyIdToken(payload url.Values) (*http.Response, error) {
	requestUrl := "https://api.line.me/oauth2/v2.1/verify"
	requestBodyReader := strings.NewReader(payload.Encode())
	req, err := http.NewRequest(http.MethodPost, requestUrl, requestBodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
