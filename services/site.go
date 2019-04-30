package services

import (
	"router-api/app"
	"fmt"
	"database/sql"
)

type SiteServices struct{}

func NeSiteService() *SiteServices {
	return &SiteServices{}
}

type ArticleAdd struct {
	Title        string
	Content      string
	CreatedTime  int64
}

type  Articles []ArticleList
type ArticleList struct {
	Id           sql.NullInt64
	Title        sql.NullString
	Content      sql.NullString
	CreatedTime  sql.NullInt64
}

func (site *SiteServices) Add(articleAdd ArticleAdd) error{
	db, err := app.ConnectMysql()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("insert into `article` (`title`,`content`,`created_time`) values (?,?,?) ")
	_, sqlError := stmt.Exec(articleAdd.Title,articleAdd.Content,articleAdd.CreatedTime)
	if sqlError != nil {
		fmt.Println("123123123123123123123")
		fmt.Println(sqlError)
		return  sqlError
	}
	 stmt.Close()
	return  nil
}

func (site *SiteServices) Index() (Articles,error){
	db, err := app.ConnectMysql()
	if err != nil {
		return nil,err
	}
	defer db.Close()
	rows, err := db.Query("select id ,title ,content ,created_time  from article ")
	if err != nil {
		return nil,err
	}
	var articles Articles
	for rows.Next() {
		var article ArticleList
		if err := rows.Scan(&article.Id,&article.Title,&article.Content,&article.CreatedTime); err != nil {
			fmt.Println(err)
		}
		articles = append(articles, article)
	}
	return  articles,nil
}


