package firelete

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type FirestoreDataDeleter[TID ID] struct {
	configuration fireworks.Configuration
}

func NewFirestoreDataDeleter[TID ID](configuration fireworks.Configuration) *FirestoreDataDeleter[TID] {
	this := new(FirestoreDataDeleter[TID])

	this.configuration = configuration

	return this
}

func (this FirestoreDataDeleter[TID]) Delete(parameters Parameters[TID]) {
	ctx := context.Background()

	client, _ := firestore.NewClient(ctx, this.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(this.configuration.CollectionName)

	document := collection.Doc(string(parameters.DocumentID))

	document.Delete(ctx)
}
