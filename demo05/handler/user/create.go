package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xingej-go/Apiserver-go/demo05/pkg/errno"
	"github.com/lexkong/log"
	"fmt"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context)  {
	var r struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r) ; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		log.Error("test-info", err, lager.Data{"create info":"struct has not init"})
	}

	log.Info("----------------1----------------")

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}
	log.Info("----------------2----------------")

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	log.Info("----------------3----------------")
	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	log.Info("----------------4----------------")
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}

