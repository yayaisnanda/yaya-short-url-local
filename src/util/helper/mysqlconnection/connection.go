package db

import (
	"fmt"
	"os"
	"time"

	// "log"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var DB *gorm.DB

func init() {
	godotenv.Load()
}

func DBInit() (*gorm.DB, error) {

	mysqlCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_MYSQL_USERNAME"),
		os.Getenv("DB_MYSQL_PASSWORD"),
		os.Getenv("DB_MYSQL_HOST"),
		os.Getenv("DB_MYSQL_PORT"),
		os.Getenv("DB_MYSQL_DATABASE"),
	)

	var err error
	DB, err = gorm.Open("mysql", mysqlCon)

	if err != nil {
		fmt.Println(fmt.Sprintf("Failed connected to database %s", mysqlCon))
		return DB, err
	}

	fmt.Println(fmt.Sprintf("Successfully connected to database %s", mysqlCon))
	DB.DB().SetConnMaxLifetime(5 * time.Minute)
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(200)
	DB.LogMode(true)

	log, err := zap.NewProduction()
	DB.SetLogger(CustomLogger(log))
	// DB.SetLogger(log.New(os.Stdout, "\r\n", 0))

	fmt.Println("Connection is created")
	return DB, err
}

func GetConnection() *gorm.DB {
	if DB == nil {
		fmt.Println("No Active Connection Found")
		DB, _ = DBInit()
	}
	return DB
}

func CustomLogger(zap *zap.Logger) *Logger {
	return &Logger{
		zap: zap,
	}
}

type Logger struct {
	zap *zap.Logger
}

func (l *Logger) Print(values ...interface{}) {
	var additionalString = ""
	for _, item := range values {
		if _, ok := item.(string); ok {
			additionalString = additionalString + fmt.Sprintf("\n%v", item)
		}
		if err, ok := item.(*mysql.MySQLError); ok {
			err.Message = err.Message + additionalString
		}
	}
}
