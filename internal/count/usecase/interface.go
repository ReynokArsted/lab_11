package usecase

type Provider interface {
	SelectCount() (string, error)
	UpdateCount(int) error
}
