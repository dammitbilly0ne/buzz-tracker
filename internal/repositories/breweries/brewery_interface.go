package breweries

import "github.com/dammitbilly0ne/buzz-tracker/internal/entities"

type Repository interface {
	CreateBrewery(id string) (*entities.Brewery, error)
}