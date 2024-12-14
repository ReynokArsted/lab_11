package api

type Usecase interface {
	FindUser(log string, pas string) error
	AddUser(log string, pas string) error
}
