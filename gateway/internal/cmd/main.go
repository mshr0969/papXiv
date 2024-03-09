package main

import (
	"context"
	"fmt"
	"gateway/internal/handler"
	"gateway/internal/repository"
	"gateway/internal/server"
	"gateway/internal/usecase"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func dsn() string {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?parseTime=true`,
		"papXiv",
		"passw0rd",
		"127.0.0.1",
		3306,
		"papXiv")
	dsn += `&loc=Local&time_zone=%27Asia%2FTokyo%27`
	return dsn
}

func main() {
	var err error
	if time.Local, err = time.LoadLocation("Asia/Tokyo"); err != nil {
		fmt.Printf("faild to load location: %s\n", err)
		return
	}

	db, err := sqlx.Open("mysql", dsn())
	if err != nil {
		fmt.Printf("faild to open sql: %s\n", err)
		return
	}

	pr := repository.Repositories{
		Paper: repository.NewPaperRepository(db),
	}

	u := usecase.Usecases{
		Health: usecase.NewHealthUsecase(),
		Paper:  usecase.NewPaperUsecase(pr.Paper),
	}

	h := handler.Handlers{
		Health: handler.NewHealthHandler(u.Health),
		Paper:  handler.NewPaperHandler(u.Paper),
	}

	svr := server.NewServer(h)
	ctx := context.Background()
	router, err := server.NewRouter(ctx)
	if err != nil {
		log.Fatalf("error in server.NewRouter: %v", err)
	}
	handler.HandlerFromMux(svr, router)

	if err = http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
