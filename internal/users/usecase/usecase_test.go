package usecase

import (
	"testing"
)

func TestEnrichUser(t *testing.T) {
	ret, err := enrichUser("John")
	if err != nil {
		t.Error(err)
	}
	t.Logf("ret %+v", ret)
}
