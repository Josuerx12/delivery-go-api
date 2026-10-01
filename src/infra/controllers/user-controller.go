package controllers

import (
	"delivery-go-api/src/domain/entities"
	usecases "delivery-go-api/src/domain/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(database *gorm.DB) *UserController {
	return &UserController{
		db: database,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var payload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload inválido",
			"error":   err.Error(),
		})
		return
	}
	
	var count int64

	if err := uc.db.Model(&entities.User{}).Where("email = ?", payload.Email).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao consultar usuário"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Usuário já existe com o email fornecido",
		})
		return
	}

	user, err := usecases.CreateUser(payload.Name, payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error ao tentar criar novo usuário",
			"error":   err.Error(),
		})
		return
	}

	if err := uc.db.Create(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error ao tentar salvar usuário no banco de dados",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    user,
	})
}