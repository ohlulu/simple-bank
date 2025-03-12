package db

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateRandomAccount() Account {
	return Account{
		Owner:    randomOwner(),
		Balance:  randomMoney(),
		Currency: randomCurrency(),
	}
}

func CreateAndInsertRandomAccount() Account {
	account := CreateRandomAccount()
	result, _ := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    account.Owner,
		Balance:  account.Balance,
		Currency: account.Currency,
	})
	return result
}

// RandomInt generates a random integer between min and max
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func randomOwner() string {
	return randomString(6)
}

// RandomMoney generates a random amount of money
func randomMoney() int64 {
	return randomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func randomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func randomEmail() string {
	return fmt.Sprintf("%s@email.com", randomString(6))
}