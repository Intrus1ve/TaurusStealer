package login

import (
	"net/http"

	database "../../db"
	lggr "../../logger"
	utils "../../utils"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

func SetToken(user string) string {
	token := utils.GenRandStr(32)
	_, err := db.Exec("UPDATE `users` SET `Token` = ?, `LastOnline` = ? WHERE `Username` = ?", token, utils.GetTimeStamp(), user)
	if err != nil {
		logger.Println("login.SetToken: ", err)
	}

	return token
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersData struct {
	Id           int    `db:"Id" json:"id"`
	Username     string `db:"Username" json:"username"`
	PasswordHash string `db:"PasswordHash" json:"password_hash"`
	Access       int    `db:"Access" json:"access"`
	Token        string `db:"Token" json:"token"`
	LastOnline   int64  `db:"LastOnline" json:"last_online"`
}

func GetUsersData() (usersData []UsersData) {
	err := db.Select(&usersData, "SELECT * FROM `users` ORDER BY `Id` DESC")
	if err != nil {
		logger.Println("login.GetUsersData: ", err)
	}

	return usersData
}

type CookieData struct {
	User   string `json:"user"`
	Cookie string `json:"cookie"`
	IsRoot bool   `json:"is_root"`
}

func IsUserRoot(username string) (isRoot bool) {
	err := db.Get(&isRoot, "SELECT `Access` FROM `users` WHERE `Username` = ?", username)
	if err != nil {
		logger.Println("login.IsUserRoot: ", err)
	}
	return
}

func Login(ctx *gin.Context) {
	var data LoginData
	ctx.BindJSON(&data)

	usersData := GetUsersData()
	for _, user := range usersData {
		if user.Username == data.Username {
			if utils.HashCompare(user.PasswordHash, data.Password) {
				token := SetToken(data.Username)
				cookieData := CookieData{
					User:   user.Username,
					Cookie: utils.HashAndSalt(token),
					IsRoot: IsUserRoot(user.Username),
				}

				ctx.JSON(http.StatusOK, cookieData)
				return
			}
		}
	}

	logger.Printf("login.Login: unsuccessful login attempt. Username: %s, password: %s, IP: %s", data.Username, data.Password, ctx.ClientIP())
}

func IsUserValid(cookie string) (username string) {
	usersData := GetUsersData()
	for _, user := range usersData {
		if utils.HashCompare(cookie, user.Token) {
			username = user.Username
			return username
		}
	}

	return ""
}

func CtxIsUserValid(ctx *gin.Context) string {
	var cookie CookieData
	ctx.BindJSON(&cookie)
	return IsUserValid(cookie.Cookie)
}

func CtxAuthCheck(ctx *gin.Context) bool {
	var cookie CookieData
	ctx.BindJSON(&cookie)
	username := IsUserValid(cookie.Cookie)
	if username == "" {
		ctx.Status(http.StatusUnauthorized)
		return false
	}
	return true
}

func AuthCheck(cookie string) bool {
	username := IsUserValid(cookie)
	if username == "" {
		return false
	}
	return true
}

func CtxRootAuthCheck(ctx *gin.Context) bool {
	var cookie CookieData
	ctx.BindJSON(&cookie)
	username := IsUserValid(cookie.Cookie)
	if username == "" {
		ctx.Status(http.StatusUnauthorized)
		return false
	}
	return IsUserRoot(username)
}

func RootAuthCheck(cookie string) bool {
	username := IsUserValid(cookie)
	if username == "" {
		return false
	}
	return IsUserRoot(username)
}

func GetRootUser() (rootUser string) {
	err := db.Get(&rootUser, "SELECT `Username` FROM `users` WHERE `Access` = '1'")
	if err != nil {
		logger.Println("login.GetRootUser: ", err)
	}
	return
}
