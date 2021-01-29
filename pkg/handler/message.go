package handler

import (
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/alex-dev-master/chat-golang/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type getAllMessagesWithUserResponse struct {
	Result []model.MessageWithUser `json:"result"`
}

func (h *Handler) createMessage(c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.Message
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserId = userId
	rubric := c.Param("rubric")
	input.ChatRubric, _ = strconv.Atoi(rubric)
	now := time.Now()
	input.Created =  now
	id, err := h.services.Message.Create(input)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getMessages(c *gin.Context)  {
	rubric := c.Param("rubric")
	rubricInt, _ := strconv.Atoi(rubric)
	messages, err := h.services.Message.GetAllOfRubric(rubricInt)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllMessagesWithUserResponse{
		Result: messages,
	})

}