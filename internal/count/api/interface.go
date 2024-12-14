package api

type Usecase interface {
	FetchCount() (string, error)
	SetCount(n int) error
}
