package firelete

func Run[TID ID]() {
	application := BuildApplication[TID]()

	application.Run()
}
