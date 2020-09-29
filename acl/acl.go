package acl

import (
	"context"
	"github.com/Celtcoste/server-graphql/graph"
	"github.com/Celtcoste/server-graphql/graph/model"
	"github.com/Celtcoste/server-graphql/src/postgresql"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func containsRole(roles []*model.Role, role string) bool {
	for _, a := range roles {
		if a.String() == strings.ToUpper(role) {
			return true
		}
	}
	return false
}

func containsApplication(applications []model.Application, app string) bool {
	for _, a := range applications {
		if a.String() == strings.ToUpper(app) {
			return true
		}
	}
	return false
}

func getCurrentUser(ctx context.Context, roles []*model.Role, applications []model.Application) bool {
	var ginContext = ctx.Value("GinContextKey").(*gin.Context)
	var userRole = ginContext.GetHeader("role")
	var userApp = ginContext.GetHeader("application")

	err := postgresql.WithTransaction(func(tx postgresql.Transaction) error {
		row := tx.QueryRow("SELECT name "+
			"FROM account_states WHERE id = $1",
			userRole)
		if row == nil {
			return nil
		}
		err := row.Scan(&userRole)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
		return false
	}
	print(userRole)
	print(userApp)
	if containsRole(roles, userRole) && containsApplication(applications, userApp) {
		return true
	}
	return false
}