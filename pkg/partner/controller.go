package partner

import (
	"github.com/bi6o/aroundhome-challenge/internal/repo"

	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.Logger
	repo   repo.RepoInterface
}

func NewController(repo repo.RepoInterface, l *zap.Logger) *Controller {
	return &Controller{
		repo:   repo,
		logger: l,
	}
}
