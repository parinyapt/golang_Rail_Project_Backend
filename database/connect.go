package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// dsn := "root:123456789@tcp(localhost:3306)/jpa_erp_v0?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := dsnMySQLGenerate(dsnMySQLGenerateParameter{
		Host:      os.Getenv("DATABASE_MYSQL_HOST"),
		DBName:    os.Getenv("DATABASE_MYSQL_DBNAME"),
		Username:  os.Getenv("DATABASE_MYSQL_USERNAME"),
		Password:  os.Getenv("DATABASE_MYSQL_PASSWORD"),
		Charset:   "utf8mb4",
		ParseTime: "True",
		loc:       "Local",
	})
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[Error]->Failed to connect database : %s", err)
	}

	// database.AutoMigrate(&models.{})

	DB = database

}

// MySQL
type dsnMySQLGenerateParameter struct {
	Host      string
	DBName    string
	Username  string
	Password  string
	Charset   string
	ParseTime string
	loc       string
}

func dsnMySQLGenerate(param dsnMySQLGenerateParameter) string {
	return param.Username + `:` + param.Password + `@tcp(` + param.Host + `)/` + param.DBName + `?charset=` + param.Charset + `&parseTime=` + param.ParseTime + `&loc=` + param.loc
}
