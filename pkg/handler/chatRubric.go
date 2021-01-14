package handler

import (
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/alex-dev-master/chat-golang/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllChatRubricsResponse struct {
	Result []model.ChatRubricUser `json:"result"`
}

func (h *Handler) createChatRubric(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input model.ChatRubric
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.UserId = userId
	id, err := h.services.ChatRubric.Create(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getChatRubrics(c *gin.Context) {
	rubrics, err := h.services.ChatRubric.GetAll()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllChatRubricsResponse{
		Result: rubrics,
	})
}
