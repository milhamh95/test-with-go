package racefail

import (
	"fmt"
	"os"
	env "test-with-go/envjson"
	"testing"
)

func TestMain(m *testing.M) {
	err := env.Setup()
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	fmt.Println("======key=======")
	fmt.Printf("%+v\n", os.Getenv("HOST"))
	fmt.Printf("%+v\n", os.Getenv("DB"))
	fmt.Printf("%+v\n", os.Getenv("PORT"))
	fmt.Printf("%+v\n", os.Getenv("USERNAME"))
	fmt.Printf("%+v\n", os.Getenv("PASSWORD"))
	fmt.Println("=============")
	os.Exit(1)
}

// func run(m *testing.M) int {
// 	const (
// 		dropDB          = `DROP DATABASE IF EXISTS test_user_store;`
// 		createDB        = `CREATE DATABASE test_user_store;`
// 		createUserTable = `CREATE TABLE users (
// 			id SERIAL PRIMARY KEY,
// 			name TEXT,
// 			email TEXT UNIQUE NOT NULL,
// 			balance INTEGER
// 		);`
// 	)

// 	// open connection to postgres
// 	psql, err := sql.Open("postgres", "host=localhost port=5432 user=cocowork sslmode=disable password=rahasia")
// 	if err != nil {
// 		panic(fmt.Errorf("sql.Open() err = %s", err))
// 	}
// 	defer psql.Close()

// 	// drop database if exist
// 	_, err = psql.Exec(dropDB)
// 	if err != nil {
// 		panic(fmt.Errorf("psql.Exec() err = %s", err))
// 	}

// 	// teardown - defer drop db at the end of function
// 	defer func() {
// 		_, err = psql.Exec(dropDB)
// 		if err != nil {
// 			panic(fmt.Errorf("psql.Exec() err = %s", err))
// 		}
// 	}()

// 	db, err := sql.Open("postgres", "host=localhost port=5432 user=cocowork")

// }
