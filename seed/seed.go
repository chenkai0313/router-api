package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)


var Config config

type config struct {
	DbConfigs []dbConfig `mapstructure:"db"`
}

type dbConfig struct {
	Name string `mapstructure:"name"`
	Dsn  string `mapstructure:"dsn"`
}

func main() {
	sqlContent, err := ioutil.ReadFile("config/sql")
	if err != nil {
		panic(fmt.Errorf("can't read sql file"))
	}

	viper.AddConfigPath("./config")
	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("Fatal error on set value to dbConfig: %s \n", err))
	}

	c := make(chan string)
	for _, dbs := range Config.DbConfigs {
		go execsql(dbs, sqlContent, c)
	}

	for range Config.DbConfigs {
		x := <-c
		fmt.Println(x)
	}
}

func execsql(Dbconfig dbConfig, sqlContent []byte, c chan string) {
	db, err := sql.Open("mysql", Dbconfig.Dsn)
	if err != nil {
		str := fmt.Sprintf("can't open db [ %s ]: %s \n", Dbconfig.Name, err)
		c <- str
		return
	}
	defer db.Close()

	content := string(sqlContent)
	sqls := strings.Split(content, ";")

	for _, sql := range sqls {
		if strings.TrimSpace(sql) == "" {
			continue
		}
		if _, err := db.Exec(sql); err != nil {
			str := fmt.Sprintf("[%s] error : sql = %s, error= %s \n", Dbconfig.Name, sql, err)
			c <- str
			return
		}
	}

	str := fmt.Sprintf("[%s] succss", Dbconfig.Name)
	c <- str
}
