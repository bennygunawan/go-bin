package service

func (u *UserServices) GetUser() error {
	return u.repo.GetUser()
}
func (u *UserServices) SetUser() error {
	return u.repo.SetUser()
}
func (u *UserServices) UpsertUser() error {
	return u.repo.UpsertUser()
}
func (u *UserServices) DeleteUser() error {
	return u.repo.DeleteUser()
}
