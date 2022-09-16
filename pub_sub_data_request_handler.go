package firelete

import (
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

type PubSubPushRequestHandler struct {
	dataReader  fireworks.DataReader[Parameters]
	dataDeleter DataDeleter
}

func NewPubSubPushRequestHandler(
	dataReader fireworks.DataReader[Parameters],
	dataDeleter DataDeleter,
) *PubSubPushRequestHandler {
	this := new(PubSubPushRequestHandler)

	this.dataReader = dataReader
	this.dataDeleter = dataDeleter

	return this
}

func (this PubSubPushRequestHandler) Handle(ginContext *gin.Context) {
	data := this.dataReader.Read(ginContext)

	this.dataDeleter.Delete(data)
}
