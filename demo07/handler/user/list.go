package user

import (
	"xingej-go/Apiserver-go/demo07/pkg/errno"
	"github.com/gin-gonic/gin"
	"xingej-go/Apiserver-go/demo07/service"
	. "xingej-go/Apiserver-go/demo07/handler"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}

