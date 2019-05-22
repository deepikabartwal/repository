package repository

//Storage interface for implementing different system for data storage...
type Storage interface {
	Save(args []string)
	ShowToDos()
	Delete(indexToBeDeleted int64)
}
