package site

import (
	"errors"
	"strings"
	"router-api/app"
)

type ArticleIndexRequest struct {
	Token   string `json:"token"  form:"token"`
}

//请求验证
func (m ArticleIndexRequest) Validate() error {
	if !m.validateToken() {
		return errors.New("token error")
	}
	return nil
}

//验证token
func (m ArticleIndexRequest) validateToken() bool {
	Token := app.Config.Token
	return strings.EqualFold(Token, m.Token)
}


