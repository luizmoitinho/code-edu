package domain_test

import (
	"testing"

	"github.com/service-user/domain"
	"github.com/stretchr/testify/require"
)

//TestNewUser ... user create test
func TestNewUser(t *testing.T) {

	//success
	_, err := domain.NewUser("Test", "email@test.com", "123456")
	require.Nil(t, err)

	//invalid email
	_, err = domain.NewUser("Test", "emailtest.com", "123456")
	require.Error(t, err)

	//invalid password
	_, err = domain.NewUser("Test", "emailtest.com", "")
	require.Error(t, err)

}
