package suitetest

import (
	"test-with-go/section-18-interface/suite"
	"testing"
)

func UserStore(t *testing.T, us suite.UserStore) {
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

}
