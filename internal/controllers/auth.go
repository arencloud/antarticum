package controllers

import (
	"fmt"
	"github.com/arencloud/antarticum/internal/helpers"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterInput struct {
	ID       int    `json:"patientId" binding:"required"`
	Name     string `json:"fullName" binding:"required"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var users []RegisterInput

func GetUserByID(uid uint) (RegisterInput, error) {

	var u RegisterInput

	//if err := DB.First(&u, uid).Error; err != nil {
	//	return u, errors.New("user not found")
	//}

	return u, nil

}

func CurrentUser(c *gin.Context) {

	userId, err := helpers.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Register(c *gin.Context) {
	var newPatient RegisterInput

	if err := c.ShouldBindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newPatient)

	c.JSON(http.StatusOK, gin.H{"message": newPatient})
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginCheck(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func LoginCheck(email string, password string) (string, error) {
	var patientToFind RegisterInput
	for _, patient := range users {
		if fmt.Sprint(patient.Email) == email {
			patientToFind = patient
		}
	}

	err := VerifyPassword(patientToFind.Password, password)
	token, err := helpers.GenerateToken(patientToFind.ID)

	if patientToFind.Email != email || (err != nil && err == bcrypt.ErrMismatchedHashAndPassword) {
		return "", err
	}

	return token, nil
}

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}
