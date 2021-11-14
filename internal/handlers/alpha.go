package handlers

import (
	"context"
	"errors"

	"github.com/dammitbilly0ne/buzz-tracker/internal/repositories/beers"
	"github.com/dammitbilly0ne/buzz-tracker/protos"
)

type Alpha struct {
	repo beers.Repository
}
type AlphaConfig struct {
	Repo beers.Repository
}

func NewAlpha(cfg *AlphaConfig) (*Alpha, error) {
	if cfg == nil {
		return nil, errors.New("cfg is required")
	}

	if cfg.Repo == nil {
		return nil, errors.New("cfg.Repo is required")
	}

	return &Alpha{
		repo: cfg.Repo,
	}, nil
}

func (a *Alpha) GetBeer(ctx context.Context, req *protos.GetBeerRequest) (*protos.GetBeerResponse, error) {
	if req == nil {
		return nil, errors.New("req is required")
	}

	beer, err := a.repo.GetBeer(req.Id)
	if err != nil {
		return nil, err
	}
	return &protos.GetBeerResponse{
		Beer: beerToProto(beer),
	}, nil
}
