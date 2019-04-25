package controllers

import (
	"github.com/go-ozzo/ozzo-routing"
	"fmt"
	"router-api/app"
)

type (
	siteService interface {
		SiteService(name string)(error)
	}

	siteResource struct {
		service siteService
	}
)


func ServeSiteResource(rg *routing.RouteGroup, service siteService) {
	r := &siteResource{service}
	rg.Post("/index", r.index)
}

func (r *siteResource) index(c *routing.Context) error {

	client := app.ConnectRedis()
	err := client.Set("foo", "bar", 0).Err()
	if err != nil {
		fmt.Printf("try set key[foo] to value[bar] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	value, err := client.Get("foo").Result()
	if err != nil {
		fmt.Printf("try get key[foo] error[%s]\n", err.Error())
		err_handler(err)
	}

	fmt.Printf("key[foo]'s value is %s\n", value)

	defer client.Close()

	name :="aaa"
	fmt.Println(name)
	res := r.service.SiteService(name)
	fmt.Println(res)
	return	c.Write("test api")
}

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}