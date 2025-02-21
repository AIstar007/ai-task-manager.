package auth

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var secretKey = []byte("secret")

func GenerateToken(userID int) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString(secretKey)
}

func Login(c *gin.Context) {
    token, _ := GenerateToken(1) // Example user ID
    c.JSON(200, gin.H{"token": token})
}
