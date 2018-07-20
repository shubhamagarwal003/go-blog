package helper

import (
	"github.com/shubhamagarwal003/go-blog/models"
)

func (store *DbStore) CreateTag(tag *models.Tag) error {
	_, err := store.Db.Query("INSERT INTO tags(value, blogId) VALUES ($1,$2)",
		tag.Value, tag.BlogId)
	return err
}

func (store *DbStore) GetBlogsForTag(tagValue string) ([]*models.Blog, error) {
	rows, err := store.Db.Query("SELECT blogs.id, blogs.title, blogs.content,blogs.uid, users.username "+
		"from tags inner join blogs on tags.blogid=blogs.id inner join users on blogs.uid=users.id where tags.value=$1",
		tagValue)
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
