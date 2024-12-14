package api

type Usecase interface {
	FetchName() (string, error)
	SetName(msg string) error
}
