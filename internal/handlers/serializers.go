package handlers

import (
	"github.com/dammitbilly0ne/buzz-tracker/internal/entities"
	"github.com/dammitbilly0ne/buzz-tracker/protos"
)

func beerToProto(beer *entities.Beer) *protos.Beer {
	if beer == nil {
		return nil
	}

	return &protos.Beer{
		Id:       beer.ID,
		Name:     beer.Name,
		Brewery:  beer.Brewery,
		BeerType: beer.Type,
	}
}
