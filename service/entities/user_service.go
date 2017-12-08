package entities

// A UserServiceProvider provides a series of
// operations to Entity User
type UserServiceProvider struct{}

// UserService is an instance of UserServiceProvider
var UserService = UserServiceProvider{}

// Insert is an atomic operation that inserts a new
// user to the database
func (*UserServiceProvider) Insert(u *User) error {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	dao := userDAO{tx}
	err = dao.Insert(u)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	return nil
}

// FindAll returns all users in the database
func (*UserServiceProvider) FindAll() ([]User, error) {
	dao := userDAO{db}
	users, err := dao.FindAll()
	panicIfErr(err)
	return users, nil
}

// FindBy finds a user by a specified column value
// If result contains multiple rows, only returns the first one
func (*UserServiceProvider) FindBy(col string, val string) (User, error) {
	dao := userDAO{db}
	user, err := dao.FindBy(col, val)
	panicIfErr(err)
	return user, nil
}

// DeleteByKey removes a user specified by a key
func (*UserServiceProvider) DeleteByKey(key string) error {
	dao := userDAO{db}
	err := dao.DeleteByKey(key)
	if err != nil {
		return err
	}
	return nil
}
