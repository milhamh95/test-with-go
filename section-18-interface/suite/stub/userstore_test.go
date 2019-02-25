package stub_test

import (
	"test-with-go/section-18-interface/suite"
	"test-with-go/section-18-interface/suite/stub"
	"test-with-go/section-18-interface/suite/suitetest"
	"testing"
)

var _ suite.UserStore = &stub.UserStore{}

func TestUserStore(t *testing.T) {
	us := &stub.UserStore{}
	suitetest.UserStore(t, us, nil, nil)
}

func TestUserStore_withStruct(t *testing.T) {
	us := &stub.UserStore{}
	tests := suitetest.UserStoreSuite{
		UserStore: us,
	}
	tests.All(t)
}
