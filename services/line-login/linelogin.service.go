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

type AccessTokenResponseDto struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func GetAccessToken(code string) (*AccessTokenResponseDto, error) {
	payload := createRequestAccessTokenPayload(code)
	resp, err := requestAccessTokenToLineLogin(payload)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var respDto AccessTokenResponseDto
	err = json.Unmarshal(respBody, &respDto)
	if err != nil {
		return nil, err
	}
	return &respDto, nil
}

func createRequestAccessTokenPayload(code string) url.Values {
	return url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {extensions.GetEnvironment("CLIENT_ID_FOR_LINE_LOGIN")},
		"redirect_uri":  {extensions.GetEnvironment("REDIRECT_URI_FOR_LINE_LOGIN")},
		"client_secret": {extensions.GetEnvironment("CLIENT_SECRET_FOR_LINE_LOGIN")},
	}
}

func requestAccessTokenToLineLogin(payload url.Values) (*http.Response, error) {
	requestUrl := "https://api.line.me/oauth2/v2.1/token"
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
