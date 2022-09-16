package firelete

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type FirestoreDataDeleter struct {
	configuration fireworks.Configuration
}

func NewFirestoreDataDeleter(configuration fireworks.Configuration) *FirestoreDataDeleter {
	this := new(FirestoreDataDeleter)

	this.configuration = configuration

	return this
}

func (this FirestoreDataDeleter) Delete(parameters Parameters) {
	ctx := context.Background()

	client, _ := firestore.NewClient(ctx, this.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(this.configuration.CollectionName)

	document := collection.Doc(parameters.DocumentID)

	document.Delete(ctx)
}
