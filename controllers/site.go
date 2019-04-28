package controllers

import (
	"github.com/go-ozzo/ozzo-routing"
	"fmt"
	"router-api/app"
	"database/sql"
	responseForm "router-api/response"
	requestArticleForm "router-api/request/article"
)

type (
	siteService interface {
		SiteService(name string) (error)
	}

	siteResource struct {
		service siteService
	}
)
type Articles []Article

type Article struct {
	id           sql.NullInt64
	title        sql.NullString
	content      sql.NullString
	createdTime sql.NullInt64
}

func ServeSiteResource(rg *routing.RouteGroup, service siteService) {
	r := &siteResource{service}
	rg.Post("/index", r.index)
	rg.Post("/add", r.add)
	rg.Post("/update", r.index)
	rg.Post("/delte", r.index)
}

func (r *siteResource) add(c *routing.Context) error {
	var request requestArticleForm.ArticleAddRequest
	if err := c.Read(&request); err != nil {
		return c.Write(responseForm.NewClientError())
	}
	fmt.Println(request)

	return c.Write(responseForm.HttpSuccessError())
}


func (r *siteResource) index(c *routing.Context) error {
	db, err := app.ConnectMysql()
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("select id ,title ,content ,created_time  from article ")
	if err != nil {
		fmt.Println(err)
	}
	var articles Articles
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.id,&article.title,&article.content,&article.createdTime); err != nil {
			fmt.Println(err)
		}
		articles = append(articles, article)
	}
	fmt.Println(articles)

	//fmt.Println(app.NullString(article.title))

	defer db.Close()


	//
	//client,err1 := app.ConnectRedis()
	//if err1!=nil{
	//	fmt.Println(err1)
	//}
	//
	//err := client.Set("foo", "bar", 0).Err()
	//if err != nil {
	//	fmt.Printf("try set key[foo] to value[bar] error[%s]\n",
	//		err.Error())
	//	err_handler(err)
	//}
	//
	//value, err := client.Get("foo").Result()
	//if err != nil {
	//	fmt.Printf("try get key[foo] error[%s]\n", err.Error())
	//	err_handler(err)
	//}
	//
	//fmt.Printf("key[foo]'s value is %s\n", value)
	//
	//defer client.Close()

	name := "aaa"
	fmt.Println(name)
	res := r.service.SiteService(name)
	fmt.Println(res)
	return c.Write("test api")
}

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}
