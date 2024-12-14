package usecase

type Provider interface {
	SelectName() (string, error)
	AddName(string) error
}
