package domain_test

import (
	"testing"

	"github.com/service-user/domain"
	"github.com/stretchr/testify/require"
)

//TestNewUser ... user create test
func TestNewUser(t *testing.T) {
	_, err := domain.NewUser("Test", "email@test.com", "123456")
	require.Nil(t, err)

}
