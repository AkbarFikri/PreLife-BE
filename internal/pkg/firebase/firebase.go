package firebase

import (
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"os"
	"strings"
)

type FirebaseClient struct {
	app *firebase.App
	log *logrus.Logger
}

func Init(log *logrus.Logger) (*FirebaseClient, error) {
	firebaseServiceCredential := map[string]interface{}{
		"type":                        os.Getenv("FIREBASE_TYPE"),
		"project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
		"private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		"private_key":                 strings.Replace(os.Getenv("FIREBASE_PRIVATE_KEY"), "/\\n/gm", "\n", -1),
		"client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
		"auth_uri":                    os.Getenv("FIREBASE_AUTH_URI"),
		"token_uri":                   os.Getenv("FIREBASE_TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_x509_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("FIREBASE_CLIENT_x509_CERT_URL"),
		"universe_domain":             os.Getenv("FIREBASE_UNIVERSE_DOMAIN"),
	}

	credential, err := json.Marshal(firebaseServiceCredential)
	if err != nil {
		log.Warnf("error marshalling firebase credential: %v", err)
		return nil, err
	}

	opt := option.WithCredentialsJSON(credential)
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}
	return &FirebaseClient{app: firebaseApp, log: log}, nil
}

func (f *FirebaseClient) Auth() *auth.Client {
	client, err := f.app.Auth(context.Background())
	if err != nil {
		f.log.Fatal(err)
	}
	return client
}
