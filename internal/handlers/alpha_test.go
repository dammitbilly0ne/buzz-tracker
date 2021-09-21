package handlers

import (
	"context"
	"errors"
	"testing"

	"github.com/dammitbilly0ne/buzz-tracker/internal/entities"

	"github.com/dammitbilly0ne/buzz-tracker/protos"

	"github.com/dammitbilly0ne/buzz-tracker/internal/repositories/beers"
	"github.com/stretchr/testify/assert"
)

func setupFixture() *Alpha {
	return &Alpha{repo: &beers.MockRepo{}}
}
func Test_NewAlpha(t *testing.T) {
	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewAlpha(nil)
		assert.Nil(t, actual)
		assert.NotNil(t, err)

		expErr := errors.New("cfg is required")
		assert.Equal(t, expErr, err)
	})
	t.Run("it requires a Repo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{})

		assert.Nil(t, actual)
		assert.NotNil(t, err)

		expERR := errors.New("cfg.Repo is required")
		assert.Equal(t, expERR, err)
	})
	t.Run("it returns a valid alpha", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			Repo: &beers.MockRepo{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.repo)
	})
}

func Test_GetBeer(t *testing.T) {
	ctx := context.Background()
	t.Run("it returns an error if the request is nil", func(t *testing.T) {
		handler := setupFixture()
		actual, err := handler.GetBeer(ctx, nil)

		assert.Nil(t, actual)
		assert.NotNil(t, err)

	})
	t.Run("it returns beer", func(t *testing.T) {
		handler := setupFixture()
		m := handler.repo.(*beers.MockRepo)

		expBeer := &entities.Beer{
			ID:      "testID",
			Brewery: "spokane",
			Name:    "Duey DeNada",
			Type:    "Imperial IPA",
		}

		m.On("GetBeer", "testID").Return(expBeer, nil)

		actual, err := handler.GetBeer(ctx, &protos.GetBeerRequest{
			Id: "testID",
		})

		assert.NotNil(t, actual)
		assert.Equal(t, expBeer.Name, actual.Beer.Name)
		assert.Equal(t, expBeer.Type, actual.Beer.BeerType)
		assert.Equal(t, expBeer.Brewery, actual.Beer.Brewery)
		assert.Equal(t, expBeer.ID, actual.Beer.Id)
		assert.Nil(t, err)

	})
	t.Run("it returns an error when the repo errors", func(t *testing.T) {
		handler := setupFixture()
		m := handler.repo.(*beers.MockRepo)
		expErr := errors.New("No beer for you")
		m.On("GetBeer", "testID").Return(nil, expErr)

		actual, err := handler.GetBeer(ctx, &protos.GetBeerRequest{
			Id: "testID",
		})

		assert.NotNil(t, err)
		assert.Nil(t, actual)

		assert.Equal(t, expErr, err)

	})
}
