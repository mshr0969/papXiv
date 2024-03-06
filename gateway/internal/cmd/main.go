package main

import (
	"context"
	"gateway/handler"
	"gateway/server"
	"gateway/usecase"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	// r := repository.Repositories{
	// 	Paper: nil,
	// }

	u := usecase.Usecases{
		Health: usecase.NewHealthUsecase(),
	}

	h := handler.Handlers{
		Health: handler.NewHealthHandler(u.Health),
	}

	svr := server.NewServer(h)
	router, err := server.NewRouter(ctx)
	if err != nil {
		log.Fatalf("error in server.NewRouter: %v", err)
	}
	handler.HandlerFromMux(svr, router)

	http.ListenAndServe(":8080", router)
}
