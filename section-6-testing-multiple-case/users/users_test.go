package users

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	exitCode := run(m)
	os.Exit(exitCode)
}

func run(m *testing.M) int {
	const (
		createDB        = `CREATE DATABASE test_user_store`
		dropDB          = `DROP DATABASE IF EXISTS test_user_store`
		createUserTable = `CREATE TABLE users(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);`
	)

	// connect to postgres
	psql, err := sql.Open("postgres", "host=localhost port=5432 user=cocowork password=rahasia sslmode=disable")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err = %s", err))
	}
	fmt.Println("Success connect to postgres")
	defer psql.Close()

	// drop database if exist
	_, err = psql.Exec(dropDB)
	if err != nil {
		panic(fmt.Errorf("Failed to drop db err = %s", err))
	}
	fmt.Println("Success drop db")

	// create database if exists
	_, err = psql.Exec(createDB)
	if err != nil {
		panic(fmt.Errorf("Failed to create db err = %s", err))
	}
	fmt.Println("Success create db")

	// defer delete db after test is completed
	defer func() {
		_, err = psql.Exec(dropDB)
		if err != nil {
			panic(fmt.Errorf("Failed to drop db err = %s", err))
		}
		fmt.Println("Success to drop db")
	}()

	// connect to spesific db in postgres
	db, err := sql.Open("postgres", "host=localhost port=5432 user=cocowork password=rahasia sslmode=disable dbname=test_user_store")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err = %s", err))
	}
	fmt.Println("Success connect to test_user_store db")
	defer db.Close()

	// create user table
	_, err = db.Exec(createUserTable)
	if err != nil {
		panic(fmt.Errorf("db.Exec() err = %s", err))
	}
	fmt.Println("Success create user table")

	return m.Run()
}

func TestUserStore(t *testing.T) {
	// TestUserStore_Find - query to create and drop db, and create user table for testing
	db, err := sql.Open("postgres", "host=localhost port=5432 user=cocowork password=rahasia sslmode=disable dbname=test_user_store")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err = %s", err))
	}
	fmt.Println("Success connect to test_user_store db")
	defer db.Close()

	// create user store object
	us := &UserStore{
		sql: db,
	}
	t.Run("Find", testUserStore_Find(us))

}

func testUserStore_Find(us *UserStore) func(t *testing.T) {
	return func(t *testing.T) {
		jon := &User{
			Name:  "Jon Calhoun",
			Email: "jon@calhoun.io",
		}

		// insert data to database
		err := us.Create(jon)
		if err != nil {
			t.Errorf("us.Create() err = %s", err)
		}
		t.Logf("Success insert new data to user table")

		// defer delete data
		defer func() {
			err := us.Delete(jon.ID)
			if err != nil {
				t.Errorf("us.Delete() err = %s", err)
			}
			t.Logf("Sucess delete data for id : %d", jon.ID)
		}()

		// table driven test cases
		tests := []struct {
			name    string
			id      int
			want    *User
			wantErr error
		}{
			{"Found", jon.ID, jon, nil},
			{"Not Found", -1, nil, ErrorNotFound},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				got, err := us.Find(tc.id)
				if err != tc.wantErr {
					t.Errorf("us.Find() err = %s", err)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("us.Find() = %+v, want %+v", got, tc.want)
				}
			})
		}
	}
}
