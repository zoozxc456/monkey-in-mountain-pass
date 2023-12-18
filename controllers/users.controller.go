package controllers

import (
	"net/http"
	"net/url"

	extensions "monkey-in-mountain-pass/extensions"

	"github.com/gin-gonic/gin"
)

func RegisterByLineLogin(ctx *gin.Context) {
	params := url.Values{
		"response_type": {"code"},
		"client_id":     {extensions.GetEnvironment("CLIENT_ID_FOR_LINE_LOGIN")},
		"redirect_uri":  {extensions.GetEnvironment("REDIRECT_URI_FOR_LINE_LOGIN")},
		"state":         {"12345abcde"},
		"scope":         {"profile openid email"},
		"nonce":         {"09876xyz"},
	}
	ctx.Redirect(http.StatusPermanentRedirect, "https://access.line.me/oauth2/v2.1/authorize?"+params.Encode())
}
