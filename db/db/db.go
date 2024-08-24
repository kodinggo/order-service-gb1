package db

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConfigDB() *gorm.DB {

	// dbDriver := "mysql"
	// dbName := "test"
	// dbUser := "user"
	// dbPassword := "your password"
	// dbTcp := "@tcp(127.0.0.1:3306)/"
	// gormDb, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+dbTcp+dbName+"?charset=utf8&parseTime=True")
	// if err != nil {
	// 	fmt.Println("gorm Db connection ", err)
	// 	return nil, err
	// }

	// Load the database configuration from Viper
	host := viper.GetString("mysql.host")
	port := viper.GetString("port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	dbname := viper.GetString("mysql.database")

	//connection string with gorm
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	// // Open a connection to the database
	// gormdb, err := gorm.Open(dbdbDriver, user+":"+password+"@tcp(" + host + ":" + port + ")/"+dbname+"?charset=utf8&parseTime=True")
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// Open a connection to the database
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ping the database to ensure it's connected
	ping, err := gormdb.DB()
	if err != nil {
		log.Fatal("Failed to get database handle:", err)
	}
	if err := ping.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Successfully connected to the database.")

	return gormdb
}
