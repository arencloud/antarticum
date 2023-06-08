package controllers

import (
	"fmt"
	"github.com/arencloud/antarticum/internal/helpers"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Patient struct {
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

var patients []Patient

func GetUserByID(uid uint) (Patient, error) {

	var u Patient

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
	var newPatient Patient

	if err := c.ShouldBindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patients = append(patients, newPatient)

	c.JSON(http.StatusOK, gin.H{"message": "Validated and registered"})
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
	var patientToFind Patient
	for _, patient := range patients {
		if fmt.Sprint(patient.Email) == email {
			patientToFind = patient
		}
	}

	err := VerifyPassword(patientToFind.Password, password)
	token, err := helpers.GenerateToken(patientToFind.ID)

	//compare the user from the request, with the one we defined:
	if patientToFind.Email != email || (err != nil && err == bcrypt.ErrMismatchedHashAndPassword) {
		return "", err
	}

	return token, nil
}

func GetPatients(c *gin.Context) {
	c.JSON(200, patients)
}
