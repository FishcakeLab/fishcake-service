package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"

	"github.com/ethereum/go-ethereum/log"
)

type NotificationRequest struct {
	Token string
	Title string
	Body  string
	Data  map[string]string
}

type FireBaseClient struct {
	client *messaging.Client
}

func NewFirebase(firebasePath string) (*FireBaseClient, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(firebasePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Error("error initializing app", "error", err)
		return nil, err
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Error("error getting Messaging client", "error", err)
		return nil, err
	}
	return &FireBaseClient{client: client}, nil
}

func (fbc *FireBaseClient) SendNotification(ntr NotificationRequest) error {
	message := &messaging.Message{
		Token: ntr.Token,
		Notification: &messaging.Notification{
			Title: ntr.Title,
			Body:  ntr.Body,
		},
		Data: ntr.Data,
	}
	_, err := fbc.client.Send(context.Background(), message)
	if err != nil {
		log.Error("send message fail", "err", err)
		return err
	}
	return nil
}
