package webhooks

import "google.golang.org/grpc"

func NewClient() WebhooksClient {
	conn, err := grpc.Dial("webhook-service.general.svc.cluster.local:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return NewWebhooksClient(conn)
}
