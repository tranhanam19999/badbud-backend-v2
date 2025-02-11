package main

import (
	"fmt"
	"log"

	dbutil "github.com/badbud-backend-v2/infra"
	"github.com/badbud-backend-v2/internal/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db := dbutil.InitDB()
	// repos := repo.NewRepository(db)

	fmt.Println("Running migrations...")
	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	db.AutoMigrate(&model.User{}, &model.Match{}, &model.MatchRequest{}, &model.MatchParticipant{}, &model.Court{})
	fmt.Println("Running migrations completed...")
}
