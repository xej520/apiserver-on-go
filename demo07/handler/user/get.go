package user

import (
	"github.com/gin-gonic/gin"
	"xingej-go/Apiserver-go/demo07/model"
	. "xingej-go/Apiserver-go/demo07/handler"
	"xingej-go/Apiserver-go/demo07/pkg/errno"
)

func Get(c *gin.Context)  {
	username := c.Param("username")

	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}


