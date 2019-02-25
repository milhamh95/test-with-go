package suitetest

import (
	"test-with-go/section-18-interface/suite"
	"testing"
)

type UserStoreSuite struct {
	suite.UserStore
	BeforeEach func()
	AfterEach  func()
}

func (uss *UserStoreSuite) All(t *testing.T) {
	UserStore(t, uss.UserStore, uss.BeforeEach, uss.AfterEach)
}

func UserStore(t *testing.T, us suite.UserStore, beforeEach, afterEach func()) {
	_, err := us.ByID(123)
	if err != suite.ErrNotFound {
		t.Errorf("ByID(123) err = nil; want ErrNotFound")
	}

	t.Run("create", func(t *testing.T) {
		user := &suite.User{
			Email: "jon@calhoun.io",
		}

		err = us.Create(user)

		if err != nil {
			t.Errorf("Create() err = %s; want nil", err)
		}

		if user.ID <= 0 {
			t.Errorf("Create() user.ID = %d; want a positive value", user.ID)
		}
	})

	// t.Run("ByID", func(t *testing.T) {
	// 	if beforeEach != nil {
	// 		beforeEach()
	// 	}
	// 	user := &suite.User{
	// 		Email: "jon@calhoun.io",
	// 	}

	// 	us.Create(user)
	// 	// teardown
	// 	defer func() {
	// 		us.Delete(user)
	// 		if afterEach != nil {
	// 			afterEach()
	// 		}
	// 	}()

	// 	got, err := us.ByID(user.ID)
	// 	if err != nil {
	// 		t.Errorf("ByID() err = %s; want nil", err)
	// 	}
	// 	if got != user {
	// 		t.Errorf("ByID() = %v; want %v", got, user)
	// 	}

	// 	us.Delete(user)
	// })

}
