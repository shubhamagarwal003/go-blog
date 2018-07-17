package helper

import (
	"github.com/shubhamagarwal003/blog/models"
)

func (store *DbStore) CreateBlog(blog *models.Blog) error {
	_, err := store.Db.Query("INSERT INTO blogs(title, content, uid) VALUES ($1,$2, $3)",
		blog.Title, blog.Content, blog.Uid)
	return err
}

func (store *DbStore) GetBlogs() ([]*models.Blog, error) {
	rows, err := store.Db.Query("SELECT blogs.id, blogs.title, blogs.content,blogs.uid, users.username " +
		"from blogs inner join users on users.id=blogs.uid")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	blogs := []*models.Blog{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		blog := &models.Blog{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Uid, &blog.Username); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (store *DbStore) GetBlog(id string) (*models.Blog, error) {
	row, err := store.Db.Query("SELECT blogs.id, blogs.title, blogs.content, blogs.uid, users.username "+
		"from blogs inner join users on blogs.uid=users.id where blogs.id=$1", id)
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer row.Close()
	blog := &models.Blog{}
	if row.Next() {
		err = row.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Uid, &blog.Username)
	}
	return blog, err
}
