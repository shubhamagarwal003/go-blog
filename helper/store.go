package helper

// The sql go library is needed to interact with the database
import (
	"database/sql"
	"github.com/shubhamagarwal003/go-blog/models"
)

// Our store will have two methods, to add a new bird,
// and to get all existing birds
// Each method returns an error, in case something goes wrong
type Store interface {
	CreateUser(user *models.User) error
	CheckUser(user *models.User) (bool, int)
	SaveToken(sessionToken *models.SessionToken) error
	DeleteToken(token string) error
	GetUser(token string) *models.User
	CreateBlog(blog *models.Blog) (int, error)
	GetBlogs() ([]*models.Blog, error)
	GetBlog(id string) (*models.Blog, error)
	CreateTag(tag *models.Tag) error
	BeginTransaction() (*sql.Tx, error)
	GetBlogsForTag(tagValue string) ([]*models.Blog, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type DbStore struct {
	Db *sql.DB
}

func (store *DbStore) BeginTransaction() (*sql.Tx, error) {
	tx, err := store.Db.Begin()
	return tx, err
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var Str Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	Str = s
}
