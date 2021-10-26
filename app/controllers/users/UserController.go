package users

import (
	"fmt"
	"gin-react-auth/app/domain/users"
	"gin-react-auth/app/services"
	"gin-react-auth/utils/config"
	"gin-react-auth/utils/errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

func Register(c *gin.Context) {
	config.LoadEnv()
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequestError("invalid json")
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequestError("invalid json")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString(JwtSecretKey)
	fmt.Println(token, err)
	if err != nil {
		err := errors.InternalServerError("login failed")
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "127.0.0.1", false, true)

	c.JSON(http.StatusOK, result)
}

func Home(c *gin.Context) {
	config.LoadEnv()

	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.InternalServerError("could not retrieve cookie")
		c.JSON(getErr.Status, getErr)
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if err != nil {
		restErr := errors.InternalServerError("error parsing cookie")
		c.JSON(restErr.Status, restErr)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.BadRequestError("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)

}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "logout successfully!",
	})

}
