package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func initialize() (*firebase.App, error) {
	opt := option.WithCredentialsFile("firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
