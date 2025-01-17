package handler

import (
	"net/http"

	"github.com/Snake1-1eyes/todo-app"
	"github.com/gin-gonic/gin"
)

// @Summary      SignUp
// @Description  Create account
// @Tags         auth
// @ID           create-account
// @Accept       json
// @Produce      json
// @Param        input body todo.User true "Account Info"
// @Success      200  {integer}  integer 1
// @Failure      400,404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /auth/sign-up [post]

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary      SignIn
// @Description  login
// @Tags         auth
// @ID           login-account
// @Accept       json
// @Produce      json
// @Param        input body signInInput true "Credentials"
// @Success      200  {string}  string "token"
// @Failure      400,404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
