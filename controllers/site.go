package controllers

import (
	"github.com/go-ozzo/ozzo-routing"
	"fmt"
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
	name :="aaa"
	fmt.Println(name)
	res := r.service.SiteService(name)
	fmt.Println(res)
	return	c.Write("test api")
}