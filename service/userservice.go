package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetUserList
// @Tags GetUserData
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserData [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"code":    0,
		"message": "獲取成功",
		"data":    data,
	})
}

func FindUserByName(name string) {}

// CreateUser
// @Tags CreateUSer
// @param name query string true "name"
// @param password query string true "password"
// @param repassword query string true "repassword"
// @param phone query string true "phone"
// @param email query string true "email"
// @param age query string true "age"
// @param sex query string true "sex"  Enums(male, female, none)
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	userAge, _ := strconv.Atoi(c.Query("age"))
	user.Age = userAge
	user.Sex = c.Query("sex")
	defaultTime := time.Now()
	user.LoginTime = defaultTime
	user.HeartbeatTime = defaultTime
	user.LoginOutTime = defaultTime
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	data := models.FindUserName(user.Name)

	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用戶已註冊",
		})
		return
	}
	if password != repassword {

		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密碼不一致",
		})
		return
	}

	user.Password = utils.MakePassword(password, salt)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "sign up failed",
		})
	} else {
		models.CreateUser(user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Success sign up",
			"data":    user,
		})
	}

}

// DeleteUser
// @Tags DeleteUser
// @param id query string true "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "Success delete",
		"data":    user,
	})

}

// UpdateUser
// @Tags UpdateUser
// @param id formData string true "id"
// @param name formData string true "name"
// @param password formData string true "password"
// @param phone formData string true "phone"
// @param email formData string true "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [patch]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "edit failed",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Success edit",
			"data":    user,
		})

	}

}

// FindUserByNameAndPwd
// @Tags FindUserByNameAndPwd
// @param name query string true "name"
// @param password query string true "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/find-user-by-name-pwd [post]
func FindUserByNameAndPwd(c *gin.Context) {

	c.String(200, "ds")
	return
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用戶不存在",
			"data":    data,
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "登入失敗",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "登入成功",
		"data":    data,
	})
}

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.SubscribeToRedis(c, utils.PublishKey)

	if err != nil {
		fmt.Println("err", err)
	}

	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println("err", err)
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
