package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	user := CreateRandomAccount()

	result, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user.Owner,
		Balance:  user.Balance,
		Currency: user.Currency,
	})

	assert.NoError(t, err)
	assert.Equal(t, result.Owner, user.Owner)
	assert.Equal(t, result.Balance, user.Balance)
	assert.Equal(t, result.Currency, user.Currency)
}

func TestGetAccount(t *testing.T) {
	account := insertRandomAccount()

	result, err := testQueries.GetAccount(context.Background(), account.ID)

	assert.NoError(t, err)
	assert.Equal(t, result.ID, account.ID)
}

func TestUpdateAccount(t *testing.T) {
	account := insertRandomAccount()

	result, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:      account.ID,
		Balance: 9527,
	})

	assert.NoError(t, err)
	assert.Equal(t, result.Balance, int64(9527))
}

func TestDeleteAccount(t *testing.T) {
	account := insertRandomAccount()

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	assert.NoError(t, err)

	result, err := testQueries.GetAccount(context.Background(), account.ID)
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		insertRandomAccount()
	}

	result, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  5,
		Offset: 0,
	})

	assert.NoError(t, err)
	assert.Equal(t, len(result), 5)
}

func insertRandomAccount() Account {
	user := CreateRandomAccount()

	account, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user.Owner,
		Balance:  user.Balance,
		Currency: user.Currency,
	})

	if err != nil {
		fmt.Println(err)
	}

	return account
}