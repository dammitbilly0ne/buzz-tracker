package breweries

import (
	"github.com/dammitbilly0ne/buzz-tracker/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateBrewery(id string) (*entities.Brewery, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Brewery), nil
}
