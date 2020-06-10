package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

var PostgresConn *sql.DB

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type TxFn func(Transaction) error

func Setup(host string, port string, user string, password string, databasename string) error {
	/*if os.Getenv("APP_ENV") == "PROD" {
		host, port, _ = net.SplitHostPort(host)
	}*/

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, databasename)

	fmt.Printf("psqlInfo : host=%s port=%s user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, databasename)
	var err error

	PostgresConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("%T", PostgresConn)
	//defer PostgresConn.Close()
	fmt.Println("Successfully connected!")
	err = PostgresConn.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetRow(ctx *gin.Context, query string, args ...interface{}) *sql.Row {
	PostgresConn.Begin()
	row := PostgresConn.QueryRowContext(ctx, query, args...)
	return row
}

func WithTransaction(fn TxFn) (err error) {
	tx, err := PostgresConn.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
