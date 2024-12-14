package usecase

type Provider interface {
	SelectUser(log string, pas string) error
	InsertUser(log string, pas string) error
}
