package partner

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	floorMaterialWood   = "wood"
	floorMaterialCarpet = "carpet"
	floorMaterialTile   = "tile"
)

type MatcherRequest struct {
	FloorMaterial string  `json:"floor_material"`
	AddressLong   float64 `json:"address_long"`
	AddressLat    float64 `json:"address_lat"`
	FloorArea     float64 `json:"floor_area"`
	PhoneNumber   string  `json:"phone_number"`
}

// @Router			/partners/match [post]
// @Tags			Partners
// @Summary		Match Partners
// @Description	This endpoint is used to show customers the available partners within their radius
// @Param			request	body		MatcherRequest	true	"The match request"
// @Success		200		{object}	[]model.Partner
// @Failure		401		{error}		mwerror.MWErrorConnector
// @Failure		403		{error}		mwerror.MWErrorConnector
// @Failure		404		{error}		mwerror.MWErrorConnector
// @Failure		500		{error}		mwerror.MWErrorConnector
func (c *Controller) Match(ctx *gin.Context) {
	var req MatcherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("error binding json", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.FloorMaterial != floorMaterialWood && req.FloorMaterial != floorMaterialCarpet && req.FloorMaterial != floorMaterialTile {
		c.logger.Error("invalid floor material", zap.String("floor_material", req.FloorMaterial))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid floor material"})
		return
	}

	matchingPartners, err := c.repo.GetMatchingPartners(ctx, req.FloorMaterial, req.AddressLong, req.AddressLat)
	if err != nil {
		c.logger.Error("failed to get matches from the db", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get matches from the db: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, matchingPartners)
}
