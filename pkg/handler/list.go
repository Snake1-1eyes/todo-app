package handler

import (
	"net/http"
	"strconv"

	"github.com/Snake1-1eyes/todo-app"
	"github.com/gin-gonic/gin"
)

// @Summary      Create todo list
// @Security     ApiKeyAuth
// @Description  Create todo list
// @Tags         lists
// @ID           create-list
// @Accept       json
// @Produce      json
// @Param        input body todo.TodoList true "List Info"
// @Success      200  {integer}  integer 1
// @Failure      400,404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary      Get all todo lists
// @Security     ApiKeyAuth
// @Description  get all todo lists
// @Tags         lists
// @ID           get-all-lists
// @Accept       json
// @Produce      json
// @Success      200  {object}  getAllListsResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary      Get List By Id
// @Security     ApiKeyAuth
// @Description  get list by id
// @Tags         lists
// @ID           get-list-by-id
// @Accept       json
// @Produce      json
// @Success      200  {object}  todo.TodoList
// @Failure      400,404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary      Update list by id
// @Security     ApiKeyAuth
// @Description  Update list by id
// @Tags         lists
// @ID           update-list-by-id
// @Accept       json
// @Produce      json
// @Param        input body todo.UpdateListInput true "List Info"
// @Success      200  {object}  statusResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/lists/:id [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

// @Summary      Delete list by id
// @Security     ApiKeyAuth
// @Description  delete list by id
// @Tags         lists
// @ID           del-list-by-id
// @Accept       json
// @Produce      json
// @Success      200  {object}  statusResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/lists/:id [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
