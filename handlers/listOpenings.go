package handler

import (
	"github.com/RafaelRMJesus/op-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, 500, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
