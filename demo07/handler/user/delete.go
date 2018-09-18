package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xingej-go/Apiserver-go/demo07/model"
	"xingej-go/Apiserver-go/demo07/pkg/errno"

	. "xingej-go/Apiserver-go/demo07/handler"
)

func Delete(c *gin.Context)  {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}

