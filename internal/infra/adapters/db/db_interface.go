package db

type IDatabase interface {
	InstanceDB() any
	Close()
}
