package db_test

import (
	"context"
	db "github/ludo62/bank_db/db/sqlc"
	"github/ludo62/bank_db/utils"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func createRandomUser(t *testing.T) db.User {
	hashedPassword, err := utils.GenerateHashPassword(utils.RandomString(6))

	if err != nil {
		log.Fatal("Impossible de générer un mot de passe hashé:", err)
	}

	arg := db.CreateUserParams{
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}
	user, err := testQuery.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, arg.Email, user.Email)
	assert.Equal(t, arg.HashedPassword, user.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)

	return user
}

func TestCreateUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQuery.CreateUser(context.Background(), db.CreateUserParams{
		Email:          user1.Email,
		HashedPassword: user1.HashedPassword,
	})
	assert.Error(t, err)
	assert.Empty(t, user2)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	newPassword, err := utils.GenerateHashPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Impossible de générer un mot de passe hashé:", err)
	}

	arg := db.UpdateUserPasswordParams{
		HashedPassword: newPassword,
		ID:             user1.ID,
		UpdatedAt:      time.Now(),
	}
	newUser, err := testQuery.UpdateUserPassword(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, arg.HashedPassword, newUser.HashedPassword)
	assert.Equal(t, user1.Email, newUser.Email)
	assert.WithinDuration(t, user1.UpdatedAt, time.Now(), 2*time.Second)
}
