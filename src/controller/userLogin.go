package controller

import (
	"douyin/src/database"
	"douyin/src/repository"
	"douyin/src/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	username := c.Query("username")
	//查看数据库有无用户
	var user database.User
	if users, exist := repository.UserQueryByName(username); !exist {
		//如果用户不存在，返回登录失败响应
		c.JSON(http.StatusOK, service.UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	} else {
		user = users[0]
	}
	var password string
	switch user.Encryption {
	case "pbkdf2":
		password = service.Encryption_PBKDF2(c.Query("password"), username, user.Iter)
	case "sha256":
		password = service.Encryption(c.Query("password"))
	default:
		password = service.Encryption(c.Query("password"))
	}

	if password != user.Password {
		//密码错误
		c.JSON(http.StatusOK, service.UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Password error"},
		})
		return
	}

	//获取新的token，每次登陆都会更新token
	tokenString, err := service.GetToken(username, user.Id)
	//token生成失败
	if err != nil {
		//返回登录失败响应
		c.JSON(http.StatusOK, service.UserLoginResponse{
			Response: service.Response{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("Token release error:%v", err),
			}})
		return
	}

	//返回登录成功响应
	c.JSON(http.StatusOK, service.UserLoginResponse{
		Response: service.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    tokenString,
	})

}
