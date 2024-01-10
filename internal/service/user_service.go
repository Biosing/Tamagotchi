package service

import (
	"Tamagotchi/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userService struct {
	userRepository model.UserRepository
}

func NewUserService(repo model.UserRepository) model.UserService {
	return &userService{
		userRepository: repo,
	}
}

// @Summary Получение списка пользователей
// @Description возвращает список всех пользователей
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /user [get]
func (r *userService) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	r.userRepository.GetUser(ctx)

	// Реализация логики получения пользователей
	c.JSON(http.StatusOK, gin.H{"message": "Список пользователей"})
}

// @Summary Создание нового пользователя
// @Description создает нового пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /user [post]
func (r *userService) CreateUser(c *gin.Context) {
	login := c.Param("login")
	pass := c.Param("password")
	fullName := c.Param("fullName")

	// Реализация логики создания пользователя
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь создан", "login": login, "pass": pass, "fullName": fullName})
}

// @Summary Получение информации о пользователе
// @Description возвращает информацию о пользователе по ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /user/{id} [get]
func (r *userService) GetUser(c *gin.Context) {
	id := c.Param("id") // Получение ID пользователя из URL
	// Реализация логики получения данных пользователя
	c.JSON(http.StatusOK, gin.H{"message": "Информация о пользователе", "id": id})
}

// @Summary Обновление пользователя
// @Description обновляет информацию о пользователе
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /user/{id} [put]
func (r *userService) UpdateUser(c *gin.Context) {
	id := c.Param("id") // Получение ID пользователя из URL
	// Реализация логики обновления данных пользователя
	c.JSON(http.StatusOK, gin.H{"message": "Данные пользователя обновлены", "id": id})
}

// @Summary Удаление пользователя
// @Description удаляет пользователя по ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /user/{id} [delete]
func (r *userService) DeleteUser(c *gin.Context) {
	id := c.Param("id") // Получение ID пользователя из URL
	// Реализация логики удаления пользователя
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь удален", "id": id})
}
