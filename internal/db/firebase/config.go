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

// func cloudStorageCustomBucket(app *firebase.App) {
// 	client, err := app.Storage(context.Background())
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// [START cloud_storage_custom_bucket_golang]
// 	bucket, err := client.Bucket("loap-img-storage")
// 	// [END cloud_storage_custom_bucket_golang]
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Created bucket handle: %v\n", bucket)
// }
