package controllers

import "github.com/gin-gonic/gin"

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func hashPassword(password string) string {

}

func verifyPassword(userPassword string, providePassword string) (bool, string) {

}
