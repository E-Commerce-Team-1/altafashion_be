package services

import (
	"altafashion_be/config"
	"altafashion_be/feature/products/domain"
	mocks "altafashion_be/mocks/products"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddProduct(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add Product", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Image: "img.jpg", Name: "hoodie", Description: "size XL", Category: "top wear", Qty: 1, Price: 100000}, nil).Once()

		srv := New(repo)
		input := domain.Core{ID: uint(1), Image: "img.jpg", Name: "hoodie", Description: "size XL", Category: "top wear", Qty: 1, Price: 100000}
		res, err := srv.AddProduct(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Add Products", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		res, err := srv.AddProduct(domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDestroy(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete Product", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Destroy(1)
		assert.Nil(t, err)
		assert.Equal(t, nil, nil)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Delete Product", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		err := srv.Destroy(1)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New(config.DATABASE_ERROR), "error, database cant process data")
		repo.AssertExpectations(t)
	})
}
