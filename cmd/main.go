package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	ctx := context.Background()

	//sqldb, err := sql.Open(sqliteshim.ShimName, ":memory:")
	sqldb, err := sql.Open(sqliteshim.ShimName, "./mydatabase.db")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	// Migrate the schema
	_, err = db.NewCreateTable().Model(&User{}).IfNotExists().Exec(ctx)
	if err != nil {
		panic(err)
	}

	// Insert a new user
	user := &User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		panic(err)
	}

	// Query the user
	user = new(User)
	err = db.NewSelect().Model(user).Where("name = ?", "John Doe").Scan(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
}

//package main
//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/uptrace/bun"
//	"github.com/uptrace/bun/driver/sqliteshim"
//
//	"github.com/uptrace/bun"
//	"github.com/uptrace/bun/dialect/sqlitedialect"
//)
//
//func main() {
//	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
//	if err != nil {
//		panic(err)
//	}
//	db := bun.NewDB(sqldb, sqlitedialect.New())
//
//	res, err := db.ExecContext(ctx, "SELECT 1")
//	fmt.Println(res)
//	var num int
//	err2 := db.QueryRowContext(ctx, "SELECT 1").Scan(&num)
//	fmt.Println(err2)
//
//}
