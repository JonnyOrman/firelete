package firelete

import (
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

type PubSubPushRequestHandler[TID ID] struct {
	dataReader  fireworks.DataReader[Parameters[TID]]
	dataDeleter DataDeleter[TID]
}

func NewPubSubPushRequestHandler[TID ID](
	dataReader fireworks.DataReader[Parameters[TID]],
	dataDeleter DataDeleter[TID],
) *PubSubPushRequestHandler[TID] {
	this := new(PubSubPushRequestHandler[TID])

	this.dataReader = dataReader
	this.dataDeleter = dataDeleter

	return this
}

func (this PubSubPushRequestHandler[TID]) Handle(ginContext *gin.Context) {
	data := this.dataReader.Read(ginContext)

	this.dataDeleter.Delete(data)
}
