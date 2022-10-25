package services

import (
	"altafashion_be/config"
	"altafashion_be/feature/users/domain"
	"altafashion_be/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{
		ID: 1, Fullname: "Lukman", Email: "lukman@gmail.com", Password: "lukman"}

	t.Run("Sukses Insert User", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{Fullname: "Lukman", Email: "lukman@gmail.com", Password: "lukman"}
		res, err := srv.Register(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Encrypt Error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		input := domain.Core{Fullname: "Lukman", Email: "lukman@gmail.com", Password: "lukman"}
		res, err := srv.Register(input)
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DATABASE_ERROR)
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicated Data", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{Fullname: "Lukman", Email: "lukman@gmail.com", Password: "lukman"}
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DUPLICATED_DATA)
		assert.Equal(t, domain.Core{}, res)
		assert.Empty(t, res, "karena object gagal dibuat")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{ID: 1, Fullname: "andi", Email: "andi@gmail.com", Password: "$2a$10$VZZa1fAxMBt852zMdVKPn.3jsikMBzf/9fBTwHweI5Hi2v/DeJLcC"}

	t.Run("Sukses Login", func(t *testing.T) {
		repo.On("GetUser", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{Email: "andi@gmail.com", Password: "123"}
		res, token, err := srv.Login(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotNil(t, token)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data empty", func(t *testing.T) {
		repo.On("GetUser", mock.Anything).Return(domain.Core{}, errors.New("email or password empty")).Once()
		srv := New(repo)
		input := domain.Core{}
		res, token, err := srv.Login(input)
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
		assert.EqualError(t, err, "email or password empty")
		repo.AssertExpectations(t)
	})

	// t.Run("No Data", func(t *testing.T) {
	// 	repo.On("Login", mock.Anything).Return(domain.Core{}, errors.New(config.NO_DATA)).Once()
	// 	srv := New(repo)
	// 	res, err := srv.Login(domain.Core{})
	// 	assert.NotNil(t, res)
	// 	assert.NotNil(t, err)
	// 	assert.EqualError(t, errors.New("no data in database"), config.NO_DATA)
	// 	repo.AssertExpectations(t)
	// })
}
