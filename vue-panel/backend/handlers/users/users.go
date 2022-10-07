package users

import (
	"errors"
	"net/http"

	database "../../db"
	lggr "../../logger"
	utils "../../utils"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

func GetUsersData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		return
	}

	userDataFull := login.GetUsersData()
	var usersData []login.UsersData
	for _, user := range userDataFull {
		userData := login.UsersData{
			Id:         user.Id,
			Username:   user.Username,
			Access:     user.Access,
			LastOnline: user.LastOnline,
		}
		usersData = append(usersData, userData)
	}
	ctx.JSON(http.StatusOK, usersData)
}

type UserForm struct {
	Cookie string `json:"cookie"`
	Data   struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"form"`
}

func Create(ctx *gin.Context) {
	var userForm UserForm
	ctx.BindJSON(&userForm)

	if !login.RootAuthCheck(userForm.Cookie) {
		return
	}

	_, err := db.Exec("INSERT INTO `users` (`Username`, `PasswordHash`, `Access`, `Token`, `LastOnline`) VALUES (?, ?, ?, ?, ?)",
		userForm.Data.Username,
		utils.HashAndSalt(userForm.Data.Password),
		0,
		utils.GenRandStr(32),
		0,
	)

	if err != nil {
		logger.Println("users.Create: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Edit(ctx *gin.Context) {
	var userForm UserForm
	ctx.BindJSON(&userForm)

	if !login.RootAuthCheck(userForm.Cookie) {
		return
	}

	var err error
	if userForm.Data.Username != "" {
		_, err = db.Exec("UPDATE `users` SET `Username` = ? WHERE `Id` = ?",
			userForm.Data.Username,
			userForm.Data.Id,
		)
	}

	if userForm.Data.Password != "" {
		_, err = db.Exec("UPDATE `users` SET `PasswordHash` = ? WHERE `Id` = ?",
			utils.HashAndSalt(userForm.Data.Password),
			userForm.Data.Id,
		)
	}

	if err != nil {
		logger.Println("users.Edit: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Delete(ctx *gin.Context) {
	var userForm UserForm
	ctx.BindJSON(&userForm)

	if !login.RootAuthCheck(userForm.Cookie) {
		return
	}

	var access bool
	err := db.Get(&access, "SELECT `Access` FROM `users` WHERE `Username` = ?", userForm.Data.Username)
	if err != nil {
		logger.Println("users.UserDelete: ", err)
	}

	if !access {
		_, err = db.Exec("DELETE FROM `users` WHERE `Username` = ?", userForm.Data.Username)
	} else {
		err = errors.New("you cannot delete root user")
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}
