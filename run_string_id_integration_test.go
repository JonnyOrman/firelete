//go:build integrationstringid
// +build integrationstringid

package firelete

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
)

type PubSubPushBody struct {
	Message Message
}

type Message struct {
	Data string
}

var project = os.Getenv("PROJECT")
var firebaseEmulatorHost = os.Getenv("FIRESTORE_EMULATOR_HOST")
var appUrl = os.Getenv("APP_URL")
var collectionName = os.Getenv("COLLECTION_NAME")

func TestMain(m *testing.M) {
	ctx := context.Background()

	client, _ := firestore.NewClient(ctx, project)

	defer client.Close()

	collection := client.Collection(collectionName)

	data := make(map[string]interface{})
	data["prop1"] = "def"
	data["prop2"] = 456

	collection.Doc("abc").Set(ctx, data)

	m.Run()

	documentsUrl := fmt.Sprintf("http://%s/emulator/v1/projects/%s/databases/(default)/documents", firebaseEmulatorHost, project)
	req, _ := http.NewRequest(http.MethodDelete, documentsUrl, nil)
	deleteClient := &http.Client{}
	deleteClient.Do(req)
}

func TestDocumentWithStringIdIsDeleted(t *testing.T) {
	data := make(map[string]interface{})
	data["documentID"] = "abc"

	dataJson, _ := json.Marshal(data)

	dataBase64 := base64.StdEncoding.EncodeToString([]byte(dataJson))

	body := PubSubPushBody{
		Message: Message{
			Data: dataBase64,
		},
	}

	bodyJson, _ := json.Marshal(body)

	bodyBuffer := bytes.NewBuffer(bodyJson)

	resp, _ := http.Post(appUrl, "application/json", bodyBuffer)

	defer resp.Body.Close()

	var snapshots []*firestore.DocumentSnapshot

	ctx := context.Background()

	client, _ := firestore.NewClient(ctx, project)

	iter := client.Collection(collectionName).Documents(ctx)

	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		snapshots = append(snapshots, doc)
	}

	assert.Equal(t, 0, len(snapshots))
}
