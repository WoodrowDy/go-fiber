package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var DBConn *sql.DB

func InitDB() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v ", err)
	}

	HOST := os.Getenv("TEST_DB_CONFIG__HOST")
	PORT := os.Getenv("TEST_DB_CONFIG__PORT")
	USER := os.Getenv("TEST_DB_CONFIG__USERNAME")
	PASS := os.Getenv("TEST_DB_CONFIG__PASSWORD")
	DBNAME := os.Getenv("TEST_DB_CONFIG__DBNAME")
	MAX_IDLE_CONNS := os.Getenv("TEST_DB_CONFIG__MAX_IDLE_CONNS")
	MAX_OPEN_CONNS := os.Getenv("TEST_DB_CONFIG__MAX_OPEN_CONNS")

	DSN := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME
	fmt.Println("DSN:", DSN)

	DBConn, err = sql.Open("mysql", DSN)

	_MAX_IDLE_CONNS, _ := strconv.Atoi(MAX_IDLE_CONNS)
	DBConn.SetMaxIdleConns(_MAX_IDLE_CONNS)

	_MAX_OPEN_CONNS, _ := strconv.Atoi(MAX_OPEN_CONNS)
	DBConn.SetMaxOpenConns(_MAX_OPEN_CONNS)

	DBConn.SetConnMaxLifetime(time.Hour)

}

// TODO: GORM 사용 및 모듈화
func GetAdminList() []map[string]interface{} {

	result_slice := []map[string]interface{}{}

	rows, err := DBConn.Query("SELECT ADMIN_NO, LOGIN_ID, PASSWD, NICK, EMAIL FROM TB_ADMIN")
	if err != nil {
		fmt.Println(err)
	}
	if rows != nil {
		defer rows.Close()
	}
	for rows.Next() {
		var admin_no int
		var login_id string
		var passwd string
		var nick string
		var email string
		result_map := map[string]interface{}{}
		err := rows.Scan(&admin_no, &login_id, &passwd, &nick, &email)
		if err != nil {
			fmt.Println(err)
		}

		result_map["admin_no"] = admin_no
		result_map["login_id"] = login_id
		result_map["passwd"] = passwd
		result_map["nick"] = nick
		result_map["email"] = email

		result_slice = append(result_slice, result_map)

	}

	return result_slice
}
