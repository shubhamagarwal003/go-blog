package helper

import (
	"fmt"
	"github.com/shubhamagarwal003/go-blog/models"
)

func (store *DbStore) CreateUser(user *models.User) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	// THe first underscore means that we don't care about what's returned from
	// this insert query. We just want to know if it was inserted correctly,
	// and the error will be populated if it wasn't

	// Make Username as Unique
	_, err := store.Db.Query("INSERT INTO users(username, password) VALUES ($1,$2)", user.Username, user.Password)
	return err
}

func (store *DbStore) CheckUser(user *models.User) (bool, int) {
	// Query the database for all birds, and return the result to the
	// `rows` object
	var userId int
	row, err := store.Db.Query("SELECT id from users where username=$1 and password=$2",
		user.Username, user.Password)
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		fmt.Println(err)
		return false, 0
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&userId)
		if err != nil {
			fmt.Println(err)
			return false, 0
		}

		if userId >= 1 {
			return true, userId
		}
	}
	return false, 0
}

func (store *DbStore) SaveToken(sessionToken *models.SessionToken) error {
	_, err := store.Db.Query("INSERT INTO sessiontokens(token, uid) VALUES ($1,$2)", sessionToken.Token,
		sessionToken.Uid)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (store *DbStore) DeleteToken(token string) error {
	_, err := store.Db.Query("DELETE * from sessionTokens where token=$1 and uid=$2", token)
	return err
}

func (store *DbStore) GetUser(token string) *models.User {
	user := &models.User{}
	row, err := store.Db.Query("SELECT users.id, users.username from users inner join "+
		"sessionTokens on users.id = sessionTokens.uid where sessionTokens.token=$1", token)
	if err == nil {
		if row.Next() {
			err := row.Scan(&user.Id, &user.Username)
			if err != nil {
				return nil
			}
			return user
		}
	}
	return nil
}
