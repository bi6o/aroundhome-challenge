package partner

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router			/partners/:id [get]
// @Tags			Partners
// @Summary		Get Partner
// @Description	This endpoint retrieves a partner using its UUID.
// @Success		200	{object}	model.Partner
// @Failure		401	{error}		mwerror.MWErrorConnector
// @Failure		403	{error}		mwerror.MWErrorConnector
// @Failure		404	{error}		mwerror.MWErrorConnector
// @Failure		500	{error}		mwerror.MWErrorConnector
func (c *Controller) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	partnerUUID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to parse id from the request URL: %v", err)})
		return
	}

	partner, err := c.repo.Get(ctx, partnerUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "could not find partner with provided id"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, partner)
}
