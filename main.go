package main

import (
	"log"

	dbutil "github.com/badbud-backend-v2/infra"
	bdhttpauth "github.com/badbud-backend-v2/internal/https/app/httpauth"
	bdhttpmatch "github.com/badbud-backend-v2/internal/https/app/httpmatch"
	bdhttpuser "github.com/badbud-backend-v2/internal/https/app/httpuser"
	"github.com/badbud-backend-v2/internal/https/middlewares"
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
	e.Use(middlewares.BDContextMiddleware)
	// Define services
	jwtSvc := service.NewJWT()
	usrSvc := service.NewUser(repos)
	authSvc := service.NewAuth(repos, *jwtSvc)
	matchSvc := service.NewMatch(repos)

	// Define routes
	publicRouter := e.Group("/public")
	{
		v1Router := publicRouter.Group("/v1")
		{
			matchV1Router := v1Router.Group("/match")
			authV1Router := v1Router.Group("/auth")

			bdhttpuser.RegisterHttpUser(usrSvc, v1Router)
			bdhttpmatch.RegisterHttpMatch(matchSvc, matchV1Router)
			bdhttpauth.RegisterHttpAuth(authSvc, authV1Router)
		}
	}

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
