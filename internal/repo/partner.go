package repo

import (
	"database/sql"

	"github.com/bi6o/aroundhome-challenge/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RepoInterface interface {
	GetMatchingPartners(ctx *gin.Context, floorMaterial string, addressLong, addressLat float64) ([]model.Partner, error)
	Get(ctx *gin.Context, id uuid.UUID) (*model.Partner, error)
}

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetMatchingPartners(ctx *gin.Context, floorMaterial string, addressLong, addressLat float64) ([]model.Partner, error) {
	var partners []model.Partner

	query := `
		SELECT *
		FROM partners
		WHERE flooring_materials @> ARRAY[$1]::text[]
		AND earth_distance(
			ll_to_earth(address_lat, address_long),
			ll_to_earth($2, $3)
		) <= operating_radius * 1000
		ORDER BY rating DESC, earth_distance(
			ll_to_earth(address_lat, address_long),
			ll_to_earth($2, $3)
		) ASC
		`

	err := r.db.Select(&partners, query, floorMaterial, addressLat, addressLong)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}

	return partners, nil
}

func (r *Repo) Get(ctx *gin.Context, id uuid.UUID) (*model.Partner, error) {
	partner := &model.Partner{}
	query := `
		SELECT *
		FROM partners
		WHERE id = $1
		`
	err := r.db.Get(partner, query, id)
	if err != nil {
		return nil, err
	}

	return partner, nil
}
