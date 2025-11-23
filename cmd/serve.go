package cmd

import (
	"dedicationWall/config"
	"dedicationWall/dedication"
	"dedicationWall/infra/db"
	"dedicationWall/repo"
	"dedicationWall/rest"
	"dedicationWall/user"

	dedHandler "dedicationWall/rest/handlers/dedication"
	usrHandler "dedicationWall/rest/handlers/user"
	middleware "dedicationWall/rest/middlewares"

	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Migration failed:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully migrated database")

	// Repositories
	userRepo := repo.NewUserRepo(dbCon)
	dedicationRepo := repo.NewDedicationRepo(dbCon)

	// Services
	userSvc := user.NewService(userRepo)
	dedicationSvc := dedication.NewService(dedicationRepo)

	// Middlewares
	middlewares := middleware.NewMiddlewares(cnf)

	// Handlers
	userHandler := usrHandler.NewHandler(middlewares, userSvc)
	dedicationHandler := dedHandler.NewHandler(middlewares, dedicationSvc)

	// Server
	server := rest.NewServer(
		cnf,
		userHandler,
		dedicationHandler,
	)

	fmt.Println("dedicationWall server running on :4000")
	server.Start()
}
