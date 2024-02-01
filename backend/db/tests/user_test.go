package db_test

import (
	"context"
	db "github/ludo62/bank_db/db/sqlc"
	"github/ludo62/bank_db/utils"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	arg := db.CreateUserParams{
		Email:          utils.RandomEmail(),
		HashedPassword: "secret",
	}
	user, err := testQuery.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, arg.Email, user.Email)
	assert.Equal(t, arg.HashedPassword, user.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)
}
