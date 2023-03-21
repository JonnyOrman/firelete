# firelete

Create Go microservices that receive Pub/Sub push messages containing a Firestore document ID and delete the document with the matching ID.

Very easy to use, create a working service by adding just one line of code to your `main` method.

Either delete documents with a `string` ID:
```
firelete.RunStringId()
```
Or delete documents with an `int` ID:
```
firelete.RunIntId()
```

## Examples

Try working examples of firelete [here](https://github.com/JonnyOrman/firelete-examples)

## Getting started

Create a new Go project
```
mkdir firelete-example
cd firelete-example
go mod init firelete/example
```

Get `firelete`
```
go get github.com/jonnyorman/firelete
```

Add a `main.go` file with the following
```
package main

import "github.com/jonnyorman/firelete"

func main() {
	firelete.RunStringId()
}
```

Add a `firelete-config.json` file with the following
```
{
    "projectID": "your-firebase-project",
    "collectionName": "FirestoreCollection"
}
```

Tidy and run with access to a Firebase project or emulator
```
    go mod tidy
    go run .
```

Submit a `POST` to the service with a Pub/Sub push body, where the `data` includes a `documentID` value matching the ID of a document in the `FirestoreCollection` Firestore collection. You will see the document get deleted.
```

## Environment configuration

The configuration can also be provided by the environment with the following keys:
- `projectID` - `PROJECT_ID`
- `collectionName` - `COLLECTION_NAME`

A combination of the `firelete-config.json` file and environment variables can be used. For example, the project ID could be provided as the `PROJECT_ID` environment variable, while the collection name is provided with the following configuration file:
```
{
    "collectionName": "FirestoreCollection"
}
```

If a configuration value is provided in both `firelete-config.json` and the environment, then the configuration file with take priority. For example, if the `PROJECT_ID` envronment variable has value "env-project-id" and the following `firelete-config.json` file is provided:
```
{
    "projectID": "config-project-id",
    "collectionName": "FirestoreCollection"
}
```
then the project ID will be "config-project-id".

## Running integration tests

To run integration tests in Docker against a local Firebase emulator, run the following:
- For documents with `int` IDs: `make test-int-id`
- For documents with `string` IDs: `make test-string-id`