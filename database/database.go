package database

import (
	"context"
	"example/gin-api-server/ent"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *ent.Client
	once   sync.Once
)

func Connect() {
	once.Do(func() {
		var err error
		client, err = ent.Open("mysql", serverInfo())
		if err != nil {
			log.Fatalf("failed opening connection to mysql: %v", err)
		}

		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	})
}

func GetClient() *ent.Client {
	return client
}

func serverInfo() string {
	user := "root"
	pass := ""
	host := "localhost"
	port := "3306"
	databaseName := "gin_test"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, pass, host, port, databaseName)
}
