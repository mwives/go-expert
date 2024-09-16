package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/mwives/go-expert/sqlc/internal/db"
)

type CourseDB struct {
	*db.Queries
	dbConn *sql.DB
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		Queries: db.New(dbConn),
		dbConn:  dbConn,
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTX(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queries := db.New(tx)
	err = fn(queries)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return fmt.Errorf("fn err: %v, rollback err: %v", err, errRollback)
		}
		return err
	}
	return tx.Commit()
}

func (c CourseDB) CreateCourseAndCategory(
	ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams,
) error {
	return c.callTX(ctx, func(queries *db.Queries) error {
		err := queries.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		return queries.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})
	})
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	courseDB := NewCourseDB(dbConn)

	// Create a new category and course
	err = courseDB.CreateCourseAndCategory(ctx, CategoryParams{
		ID:   uuid.New().String(),
		Name: "Backend",
		Description: sql.NullString{
			String: "Backend category description", Valid: true,
		},
	}, CourseParams{
		ID:   uuid.New().String(),
		Name: "Go",
		Description: sql.NullString{
			String: "Go course description", Valid: true,
		},
		Price: 100.00,
	})
	if err != nil {
		panic(err)
	}

	queries := db.New(dbConn)

	// List all courses
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("ID: %s, Name: %s, Description: %s, Price: %f, CategoryID: %s, CategoryName: %s\n",
			course.ID, course.Name, course.Description.String, course.Price, course.CategoryID, course.CategoryName)
	}
}
