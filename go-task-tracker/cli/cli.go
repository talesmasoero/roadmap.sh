package cli

import (
	"log"
	"task-cli/db"
	"task-cli/repository"
	"task-cli/service"
)

func Run(args []string) error {
	// Init "database"
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Init repository
	repo := repository.New(db)
	svc := service.New(repo)
	svc.AddTask("Test Task")

	return nil
}
