package services

import "fmt"

type SiteServices struct{}

func NeSiteService() *SiteServices {
	return &SiteServices{}
}

func (site *SiteServices) SiteService(name string)(error){
	fmt.Println("siteService123")
	return  nil
}
