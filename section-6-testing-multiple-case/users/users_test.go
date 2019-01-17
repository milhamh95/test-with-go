package users

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestUserStore_Find(t *testing.T) {
	// TestUserStore_Find - query to create and drop db, and create user table for testing
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
		t.Fatalf("sql.Open() err = %s", err)
	}
	t.Logf("Success connect to postgres")
	defer psql.Close()

	// drop database if exist
	_, err = psql.Exec(dropDB)
	if err != nil {
		t.Fatalf("Failed to drop db err = %s", err)
	}
	t.Logf("Success drop db")

	// create database if exists
	_, err = psql.Exec(createDB)
	if err != nil {
		t.Fatalf("Failed to create db err = %s", err)
	}
	t.Logf("Success create db")

	// defer delete db after test is completed
	defer func() {
		_, err = psql.Exec(dropDB)
		if err != nil {
			t.Errorf("Failed to drop db err = %s", err)
		}
		t.Logf("Success to drop db")
	}()

	// connect to spesific db in postgres
	db, err := sql.Open("postgres", "host=localhost port=5432 user=cocowork password=rahasia sslmode=disable dbname=test_user_store")
	if err != nil {
		t.Fatalf("sql.Open() err = %s", err)
	}
	t.Logf("Success connect to test_user_store db")
	defer db.Close()

	// create user table
	_, err = db.Exec(createUserTable)
	if err != nil {
		t.Fatalf("db.Exec() err = %s", err)
	}
	t.Logf("Success create user table")

	// create user store object
	us := &UserStore{
		sql: db,
	}

	jon := &User{
		Name:  "Jon Calhoun",
		Email: "jon@calhoun.io",
	}

	// insert data to database
	err = us.Create(jon)
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
