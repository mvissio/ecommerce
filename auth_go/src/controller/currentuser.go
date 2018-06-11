package controller

import (
	"auth/token"
	"auth/tools/rest"
	"auth/user"

	"github.com/gin-gonic/gin"
)

// CurrentUser is the controller to get the current logged in user
func CurrentUser(c *gin.Context) {
	payload, err := token.ValidateToken(c)

	if err != nil {
		rest.HandleError(c, err)
		return
	}

	user, err := user.CurrentUser(payload.UserID)

	if err != nil {
		rest.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"id":    user.ID(),
		"name":  user.Name,
		"roles": user.Roles,
		"login": user.Login,
	})
}