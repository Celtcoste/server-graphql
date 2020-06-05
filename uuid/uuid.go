package uuid

import (
	"github.com/Celtcoste/server-graphql/src/postgresql"
	"github.com/google/uuid"
	"log"
)

func GenerateUID(tableName string) *string {
	uid := uuid.New()
	var i = 0
	err := postgresql.WithTransaction(func(tx postgresql.Transaction) error {
		row := tx.QueryRow("SELECT COUNT(uid) FROM "+tableName+
			" WHERE uid = $1",
			uid)
		err := row.Scan(&i)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Println("sql : ", err.Error())
		return nil
	}
	if i == 1 {
		return GenerateUID(tableName)
	}
	var resUID string
	resUID = uid.String()
	return &resUID
}
