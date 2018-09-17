package user

import (
	"github.com/gin-gonic/gin"
	"xingej-go/Apiserver-go/demo06/pkg/errno"
	"github.com/lexkong/log"
	"fmt"
	"xingej-go/Apiserver-go/demo06/handler"
)



func Create(c *gin.Context)  {
	var r struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r) ; err != nil {
		handler.SendResponse(c,errno.ErrBind, nil)
		return
	}

	//Param() 读取的是/:username
	param2 := c.Param("username")
	log.Infof("URL username: %s", param2)

	//读取 // GET /path?id=1234&name=Manu&value=
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)
	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	log.Info("----------------1----------------")

	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}
	log.Info("----------------2----------------")

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	log.Info("----------------3----------------")
	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
	}
	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	handler.SendResponse(c, nil, rsp)
}

