package main

import (
	"net/http"
	"log"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	"router-api/controllers"
	"router-api/services"
	"router-api/app"
	"fmt"
)


func main() {
	configErr := app.LoadConfig("./config") //加载配置文件
	if configErr != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", configErr))
	}

	http.Handle("/", buildRouter())
	err := http.ListenAndServe(":"+app.Config.ServerPort+"", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func buildRouter() *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort()
		return c.Write("OK " )
	})
	router.Use(
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
	)

	odan := router.Group("/site")
	siteService:= services.NeSiteService()
	controllers.ServeSiteResource(odan,siteService)

	return router

}