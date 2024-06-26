package db

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

func CreateContainer() (*dockertest.Resource, *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = 2 * time.Minute
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "latest",
		Env: []string{
			"MYSQL_USER=papXiv",
			"MYSQL_DATABASE=papXiv",
			"MYSQL_PASSWORD=passw0rd",
			"MYSQL_ROOT_PASSWORD=passw0rd",
			"TZ=Asia/Tokyo",
		},
	}

	resource, err := pool.RunWithOptions(runOptions,
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
	}

	return resource, pool
}

func ConnectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sqlx.DB {
	var db *sqlx.DB
	var err error
	if err := pool.Retry(func() error {
		time.Sleep(time.Second * 5)

		db, err = sqlx.Open("mysql", fmt.Sprintf("root:passw0rd@(localhost:%s)/papXiv?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return errors.Wrap(err, "Could not connect to database")
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("failed to connect DB: %s", err)
	}
	db.SetConnMaxLifetime(time.Second)
	return db
}

func CloseContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}
}

func MigrateDB(db *sqlx.DB) error {
	pwd, _ := os.Getwd()

	migrations := &migrate.FileMigrationSource{
		Dir: pwd + "../../../configs/migrations",
	}

	_, err := migrate.Exec(db.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "Could not migrate")
	}

	return nil
}
