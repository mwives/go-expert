package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/mwives/go-expert/sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// Create a new category
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:   uuid.New().String(),
		Name: "Backend",
		Description: sql.NullString{
			String: "Backend category description", Valid: true,
		},
	})
	if err != nil {
		panic(err)
	}

	// List all categories
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	// Update a category
	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:   categories[0].ID,
	// 	Name: "Backend Updated",
	// 	Description: sql.NullString{
	// 		String: "Backend category description updated",
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// Get a category
	// category, err := queries.GetCategory(ctx, categories[0].ID)
	// if err != nil {
	// 	panic(err)
	// }
	// println(category.ID, category.Name, category.Description.String)

	// Delete a category
	// err = queries.DeleteCategory(ctx, categories[0].ID)
	// if err != nil {
	// 	panic(err)
	// }
}
