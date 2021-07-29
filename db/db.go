package db

import (
	"create_sp/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() error {
	var err error
	// dsn := "root:asd@asd@@tcp(127.0.0.1:3306)/nozom_ecommerce"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	fmt.Println(connectionString)
	// DBConn, err = gorm.Open("mysql", "sqlserver://mcs:123@41.38.87.59:1433?database=stock_main")
	DBConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to external database")
		panic(err)
	}
	fmt.Println("Connection Opened to External Database")
	return nil

}
