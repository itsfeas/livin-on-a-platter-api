package firebase

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type FireDB struct {
	*db.Client
}

var fireDB FireDB

func (db *FireDB) Connect() error {
	home, err := os.Getwd()
	credFile := os.Getenv("DB_CREDENTIALS")
	firebaseClientUrl := os.Getenv("FB_RTDB_URL")
	if err != nil {
		return err
	}

	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: firebaseClientUrl,
	}

	// Fetch the service account key JSON file contents
	fmt.Println("gdb - finding firebase json @: " + home + credFile)
	opt := option.WithCredentialsFile(home + credFile)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database client: %v", err)
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("restricted_access/secret_document")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		return fmt.Errorf("error reading from database: %v", err)
	}
	db.Client = client
	// fmt.Println(data)
	return nil
}

func GetDB() *FireDB {
	return &fireDB
}
