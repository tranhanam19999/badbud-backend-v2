package main

import (
	"log"

	dbutil "github.com/badbud-backend-v2/infra"
	bdhttpmatch "github.com/badbud-backend-v2/internal/https/app/httpmatch"
	bdhttpuser "github.com/badbud-backend-v2/internal/https/app/httpuser"
	"github.com/badbud-backend-v2/internal/repo"
	"github.com/badbud-backend-v2/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// load
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db := dbutil.InitDB()
	repos := repo.NewRepository(db)

	e := echo.New()

	// Define services
	usrSvc := service.NewUser(repos)
	matchSvc := service.NewMatch(repos)

	// Define routes
	publicRouter := e.Group("/public")
	{
		v1Router := publicRouter.Group("/v1")
		{
			bdhttpuser.RegisterHttpUser(usrSvc, v1Router)
			bdhttpmatch.RegisterHttpMatch(matchSvc, v1Router)
		}
	}

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
