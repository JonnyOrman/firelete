package firelete

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type DataReaderMock[T any] struct {
	mock.Mock
}

func (this DataReaderMock[T]) Read(ginContext *gin.Context) T {
	args := this.Called(ginContext)
	return args.Get(0).(T)
}

type DataDeleterMock[TID ID] struct {
	mock.Mock
}

func (this DataDeleterMock[TID]) Delete(parameters Parameters[TID]) {
	_ = this.Called(parameters)
}

func TestHandle(t *testing.T) {
	ginContext := gin.Context{}

	documentID := "abc"

	data := Parameters[string]{
		DocumentID: documentID,
	}

	dataReader := new(DataReaderMock[Parameters[string]])
	dataReader.On("Read", &ginContext).Return(data)

	dataDeleter := new(DataDeleterMock[string])
	dataDeleter.On("Delete", data).Return()

	sut := PubSubPushRequestHandler[string]{dataReader, dataDeleter}

	sut.Handle(&ginContext)
}
