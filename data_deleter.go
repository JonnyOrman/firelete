package firelete

type DataDeleter[TID ID] interface {
	Delete(parameters Parameters[TID])
}
