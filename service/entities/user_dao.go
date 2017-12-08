package entities

import (
	"crypto/md5"
	"fmt"
	"io"
)

type userDAO DataAccessObject

func (dao *userDAO) Insert(u *User) error {
	hash := md5.New()
	io.WriteString(hash, u.Username)
	digest := fmt.Sprintf("%x", hash.Sum(nil))
	sqlStmt := `
        INSERT INTO users(key, username, password, email, phone) VALUES(
            '` + digest + `',
            '` + u.Username + `',
            '` + u.Password + `',
            '` + u.Email + `',
            '` + u.Phone + `'
        );
    `
	result, err := db.Exec(sqlStmt)
	panicIfErr(err)
	id, err := result.LastInsertId()
	panicIfErr(err)
	u.ID = int(id)
	u.Key = digest
	return nil
}

func (dao *userDAO) FindAll() ([]User, error) {
	sqlStmt := `SELECT * FROM users"`

	rows, err := dao.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := make([]User, 0, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(
			&u.ID, &u.Key, &u.Username, &u.Password, &u.Email, &u.Phone)
		if err != nil {
			return nil, err
		}
		userList = append(userList, u)
	}

	return userList, nil
}

func (dao *userDAO) FindBy(col string, val string) (User, error) {
	sqlStmt := `SELECT * FROM users WHERE ` + col + ` = '` + val + `';`

	rows, err := dao.Query(sqlStmt)
	panicIfErr(err)
	u := User{}
	if rows.Next() {
		err = rows.Scan(
			&u.ID, &u.Key, &u.Username, &u.Password, &u.Email, &u.Phone)
	}
	panicIfErr(err)

	return u, nil
}

func (dao *userDAO) DeleteByKey(key string) error {
	sqlStmt := `DELETE FROM users WHERE key = '` + key + `';`
	_, err := dao.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
