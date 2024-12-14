package usecase

func (u *Usecase) FetchCount() (string, error) {
	msg, err := u.p.SelectCount()
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (u *Usecase) SetCount(n int) error {
	err := u.p.UpdateCount(n)
	if err != nil {
		return err
	}

	return nil
}
