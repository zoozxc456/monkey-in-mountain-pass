package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	linelogin "monkey-in-mountain-pass/services/line-login"
)

type LineLoginCallbackDto struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func HandleLineLoginCallback(ctx *gin.Context) {
	var callbackDto LineLoginCallbackDto
	if err := ctx.ShouldBindQuery(&callbackDto); err != nil {
		ctx.JSON(http.StatusBadRequest, "Queries Error")
		return
	}

	accessToken, err := linelogin.GetAccessToken(callbackDto.Code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Get AccessToken Error")
		return
	}
	verifyIdTokenDto, err := linelogin.VerifyIdToken(accessToken.IdToken)
	fmt.Println(verifyIdTokenDto, err)
}
