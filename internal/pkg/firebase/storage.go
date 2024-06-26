package firebase

import (
	"firebase.google.com/go/v4/storage"
	"fmt"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/helper"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"time"
)

type FirebaseStorage struct {
	log    *logrus.Logger
	client *storage.Client
}

func NewFirebaseStorage(client *storage.Client, log *logrus.Logger) FirebaseStorage {
	return FirebaseStorage{
		log:    log,
		client: client,
	}
}

func (f FirebaseStorage) UploadFile(file []byte, fileName string) (string, error) {
	bucket, err := f.client.DefaultBucket()
	if err != nil {
		f.log.Fatalf("error when get bucket %v", err)
		return "", err
	}

	bucketName := fmt.Sprintf("%s.appspot.com", os.Getenv("FIREBASE_BUCKET_NAME"))

	id := fmt.Sprintf("at-%s-%s", helper.GenerateUID(12), time.Now().UnixNano())

	obj := bucket.Object(fileName)
	w := obj.NewWriter(context.Background())
	defer w.Close()
	w.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id}

	if _, err := w.Write(file); err != nil {
		f.log.Fatalf("error when uploading file %v", err)
		return "", err
	}

	link := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", bucketName, fileName, id)
	return link, nil
}
