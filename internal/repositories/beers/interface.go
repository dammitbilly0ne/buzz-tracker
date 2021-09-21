package beers

import "github.com/dammitbilly0ne/buzz-tracker/internal/entities"

type Repository interface {
	GetBeer(id string) (*entities.Beer, error)
}
