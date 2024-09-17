package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mwives/go-expert/unity-of-work/internal/db"
	"github.com/mwives/go-expert/unity-of-work/internal/repository"
	"github.com/mwives/go-expert/unity-of-work/pkg/uow"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE IF EXISTS courses;")
	dbt.Exec("DROP TABLE IF EXISTS categories;")

	dbt.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL
		);`)

	dbt.Exec(`
		CREATE TABLE IF NOT EXISTS courses (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			category_id INT NOT NULL,
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);`)

	ctx := context.Background()
	uow := uow.NewUow(ctx, dbt)

	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUseCase{
		CategoryName:     "Category 1", // ID = 1
		CourseName:       "Course 1",
		CourseCategoryID: 1, // Should fail if ID 2 is passed
	}

	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
