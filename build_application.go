package firelete

import (
	"github.com/jonnyorman/fireworks"
)

func BuildApplication[TID ID]() *fireworks.Application {
	configuration := fireworks.GenerateConfiguration("firelete-config")

	pubSubBodyDeserialiser := fireworks.JsonDataDeserialiser[fireworks.PubSubBody]{}

	ioutilReader := fireworks.IoutilReader{}

	pubSubBodyReader := fireworks.NewGinPubSubBodyReader(
		ioutilReader,
		pubSubBodyDeserialiser)

	dataDeserialiser := fireworks.JsonDataDeserialiser[Parameters[TID]]{}

	dataReader := fireworks.NewHttpRequestBodyDataReader[Parameters[TID]](
		pubSubBodyReader,
		dataDeserialiser)

	dataDeleter := NewFirestoreDataDeleter[TID](configuration)

	requestHandler := NewPubSubPushRequestHandler[TID](
		dataReader,
		dataDeleter,
	)

	routerBuilder := fireworks.NewGinRouterBuilder()

	routerBuilder.AddPost("/", requestHandler.Handle)

	router := routerBuilder.Build()

	application := fireworks.NewApplication(router)

	return application
}
