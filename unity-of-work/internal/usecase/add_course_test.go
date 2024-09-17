package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mwives/go-expert/unity-of-work/internal/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE IF EXISTS `courses`;")
	dbt.Exec("DROP TABLE IF EXISTS `categories`;")

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

	input := InputUseCase{
		CategoryName:     "Category 1", // ID->1
		CourseName:       "Course 1",
		CourseCategoryID: 2,
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCase(
		repository.NewCourseRepository(dbt),
		repository.NewCategoryRepository(dbt),
	)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
