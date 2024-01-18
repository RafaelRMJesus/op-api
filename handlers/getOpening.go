package handler

import (
	"fmt"
	"net/http"

	"github.com/RafaelRMJesus/op-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, 400, errParamIsReq("id", "queryParam").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}
	sendSuccess(ctx, "get-opening", opening)
}
