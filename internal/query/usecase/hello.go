package usecase

func (u *Usecase) FetchName() (string, error) {
	msg, err := u.p.SelectName()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetName(msg string) error {
	err := u.p.AddName(msg)
	if err != nil {
		return err
	}

	return nil
}
