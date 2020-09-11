package user_test

import (
	"testing"

	"github.com/ardanlabs/service/business/tests"
)

func TestUser(t *testing.T) {
	_, teardown := tests.NewUnit(t)
	t.Cleanup(teardown)
}
