package controllers

import (
	"github.com/go-ozzo/ozzo-routing"
	responseForm "router-api/response"
	requestArticleForm "router-api/request/site"
	responseArticle "router-api/response/site"
	"router-api/services"
	"fmt"
	"router-api/app"
	"time"
)

type (
	siteService interface {
		Add(services.ArticleAdd) (error)
		Index() (services.Articles,error)
	}

	siteResource struct {
		service siteService
	}
)


func ServeSiteResource(rg *routing.RouteGroup, service siteService) {
	r := &siteResource{service}
	rg.Post("/index", r.index)
	rg.Post("/add", r.add)
}

func (r *siteResource) add(c *routing.Context) error {
	var request requestArticleForm.ArticleAddRequest

	if err := c.Read(&request); err != nil {
		return c.Write(responseForm.NewClientError(err.Error()))
	}

	if err := request.Validate(); err != nil {
		return c.Write(responseForm.NewClientError(err.Error()))
	}

	article:=request.LoadParams()
	err:=r.service.Add(article)
	if err!=nil{
		c.Write(responseForm.NewServerError())
	}
	return c.Write(responseForm.HttpSuccessError())
}

func (r *siteResource) index(c *routing.Context) error {
	var request requestArticleForm.ArticleIndexRequest

	if err := c.Read(&request); err != nil {
		return c.Write(responseForm.NewClientError(err.Error()))
	}

	if err := request.Validate(); err != nil {
		return c.Write(responseForm.NewClientError(err.Error()))
	}
	res,err:=r.service.Index()
	if err!=nil{
		c.Write(responseForm.NewServerError())
	}

	var contents responseArticle.Contents
	for _,v:=range res {
		createdTime := time.Unix(app.NullInt64(v.CreatedTime), 0).Format("2006-01-02 15:04:05")
		content :=responseArticle.Content{
			Id:app.NullInt64(v.Id),
			Title:app.NullString(v.Title),
			Content:app.NullString(v.Content),
			CreatedTime:createdTime,
		}
		contents = append(contents, content)
	}

	return c.Write(responseArticle.QueryResponse{
		BaseResponse : responseForm.BaseResponse{
			Code:    200,
			Message: "Success",
		},
		Data : contents,
	})
}


func (r *siteResource) testRedis(c *routing.Context) error {
	client,err1 := app.ConnectRedis()
	if err1!=nil{
		fmt.Println(err1)
	}

	err := client.Set("foo", "bar", 0).Err()
	if err != nil {
		fmt.Printf("try set key[foo] to value[bar] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	value, err := client.Get("foo").Result()
	if err != nil {
		fmt.Printf("try get key[foo] error[%s]\n", err.Error())
		errHandler(err)
	}

	fmt.Printf("key[foo]'s value is %s\n", value)

	defer client.Close()

	return c.Write("test redis")
}

func errHandler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}
