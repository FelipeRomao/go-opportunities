package handler

import (
	"net/http"

	"github.com/FelipeRomao/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "get-opening", opening)
}
