package Controllers

import (
	"fmt"
	"gin-restful-api-mysql/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []Models.User
	err := Models.GetAllUsers(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSONP(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.Create(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSONP(http.StatusOK, user)
	}
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserById(&user, id)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSONP(http.StatusOK, user)
	}
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserById(&user, id)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = Models.DeleteUser(&user, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSONP(http.StatusOK, user)
	}
}

func GetById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserById(&user, id)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSONP(http.StatusOK, user)
	}
}
