package firelete

import (
	"github.com/jonnyorman/fireworks"
)

func BuildApplication() *fireworks.Application {
	configuration := fireworks.GenerateConfiguration("firelete-config")

	pubSubBodyDeserialiser := fireworks.JsonDataDeserialiser[fireworks.PubSubBody]{}

	ioutilReader := fireworks.IoutilReader{}

	pubSubBodyReader := fireworks.NewGinPubSubBodyReader(
		ioutilReader,
		pubSubBodyDeserialiser)

	dataDeserialiser := fireworks.JsonDataDeserialiser[Parameters]{}

	dataReader := fireworks.NewHttpRequestBodyDataReader[Parameters](
		pubSubBodyReader,
		dataDeserialiser)

	dataDeleter := NewFirestoreDataDeleter(configuration)

	requestHandler := NewPubSubPushRequestHandler(
		dataReader,
		dataDeleter,
	)

	routerBuilder := fireworks.NewGinRouterBuilder()

	routerBuilder.AddPost("/", requestHandler.Handle)

	router := routerBuilder.Build()

	application := fireworks.NewApplication(router)

	return application
}
