package site

import (
	"errors"
	"github.com/go-ozzo/ozzo-validation"
	"strings"
	"router-api/app"
	"router-api/services"
	"time"
)

type ArticleAddRequest struct {
	Title   string `json:"title"  form:"title"`
	Content string `json:"content"  form:"content"`
	Token   string `json:"token"  form:"token"`
}

//请求验证
func (m ArticleAddRequest) Validate() error {
	if err := m.ValidateParams(); err != nil {
		return err
	}

	if !m.validateToken() {
		return errors.New("token error")
	}
	return nil
}

func (m ArticleAddRequest) ValidateParams() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required),
		validation.Field(&m.Content, validation.Required),
		validation.Field(&m.Token, validation.Required),
	)
}

//验证token
func (m ArticleAddRequest) validateToken() bool {
	Token := app.Config.Token
	return strings.EqualFold(Token, m.Token)
}

//赋值
func (m ArticleAddRequest) LoadParams()(article services.ArticleAdd)  {
	article.Title=m.Title
	article.Content=m.Content
	article.CreatedTime=time.Now().Unix()
	return
}

