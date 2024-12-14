package usecase

func (u *Usecase) FindUser(log string, pas string) error {
	err := u.p.SelectUser(log, pas)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) AddUser(log string, pas string) error {
	err := u.p.InsertUser(log, pas)
	if err != nil {
		return err
	}

	return nil
}
