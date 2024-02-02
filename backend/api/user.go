package api

import (
	"context"
	"database/sql"
	db "github/ludo62/bank_db/db/sqlc"

	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	server *Server
}

func (u User) router(server *Server) {
	u.server = server

	serverGroup := server.router.Group("/users", AuthenticatedMiddleware())
	serverGroup.GET("", u.listUsers)
	serverGroup.GET("/me", u.getLoggedInUser)
}

type UserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) listUsers(c *gin.Context) {
	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  10,
	}

	users, err := u.server.queries.ListUsers(context.Background(), arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUsers := []UserResponse{}

	for _, v := range users {
		n := UserResponse{}.toUserResponse(&v)
		newUsers = append(newUsers, *n)
	}

	c.JSON(http.StatusOK, newUsers)
}

func (u *User) getLoggedInUser(c *gin.Context) {
	value, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Vous n'êtes pas autorisé à accéder à cette ressource"})
		return
	}

	userId, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "A renconter une erreur"})
		return
	}

	user, err := u.server.queries.GetUserByID(context.Background(), userId)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Vous n'êtes pas autorisé à accéder à cette ressource"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{}.toUserResponse(&user))
}

type UserResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u UserResponse) toUserResponse(user *db.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}
