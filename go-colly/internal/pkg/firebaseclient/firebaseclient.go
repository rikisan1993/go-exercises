package firebaseclient

import (
	"os"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

// CreateApp create firebase App
func CreateApp() (*firebase.App, error) {
	sapath := os.Getenv("SERVICE_ACCOUNT_PATH")
	opt := option.WithCredentialsFile(sapath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// return nil, fmt.Errorf("error initializing app: %v", err)
		return nil, app
	}

	return app, nil
}