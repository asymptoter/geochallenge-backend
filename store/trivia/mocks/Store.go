// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ctx "github.com/asymptoter/practice-backend/base/ctx"
import mock "github.com/stretchr/testify/mock"
import models "github.com/asymptoter/practice-backend/models"

import uuid "github.com/google/uuid"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Answer provides a mock function with given fields: context, userID, gameID, answer
func (_m *Store) Answer(context ctx.CTX, userID uuid.UUID, gameID uuid.UUID, answer string) (*models.Quiz, *models.GameResult, error) {
	ret := _m.Called(context, userID, gameID, answer)

	var r0 *models.Quiz
	if rf, ok := ret.Get(0).(func(ctx.CTX, uuid.UUID, uuid.UUID, string) *models.Quiz); ok {
		r0 = rf(context, userID, gameID, answer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quiz)
		}
	}

	var r1 *models.GameResult
	if rf, ok := ret.Get(1).(func(ctx.CTX, uuid.UUID, uuid.UUID, string) *models.GameResult); ok {
		r1 = rf(context, userID, gameID, answer)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.GameResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(ctx.CTX, uuid.UUID, uuid.UUID, string) error); ok {
		r2 = rf(context, userID, gameID, answer)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateGame provides a mock function with given fields: context, game
func (_m *Store) CreateGame(context ctx.CTX, game *models.Game) error {
	ret := _m.Called(context, game)

	var r0 error
	if rf, ok := ret.Get(0).(func(ctx.CTX, *models.Game) error); ok {
		r0 = rf(context, game)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateQuiz provides a mock function with given fields: context, quiz
func (_m *Store) CreateQuiz(context ctx.CTX, quiz *models.Quiz) error {
	ret := _m.Called(context, quiz)

	var r0 error
	if rf, ok := ret.Get(0).(func(ctx.CTX, *models.Quiz) error); ok {
		r0 = rf(context, quiz)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGames provides a mock function with given fields: context, userID, name
func (_m *Store) GetGames(context ctx.CTX, userID uuid.UUID, name string) ([]*models.Game, error) {
	ret := _m.Called(context, userID, name)

	var r0 []*models.Game
	if rf, ok := ret.Get(0).(func(ctx.CTX, uuid.UUID, string) []*models.Game); ok {
		r0 = rf(context, userID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Game)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx.CTX, uuid.UUID, string) error); ok {
		r1 = rf(context, userID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuizzes provides a mock function with given fields: context, userID, content, category
func (_m *Store) GetQuizzes(context ctx.CTX, userID uuid.UUID, content string, category string) ([]*models.Quiz, error) {
	ret := _m.Called(context, userID, content, category)

	var r0 []*models.Quiz
	if rf, ok := ret.Get(0).(func(ctx.CTX, uuid.UUID, string, string) []*models.Quiz); ok {
		r0 = rf(context, userID, content, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Quiz)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx.CTX, uuid.UUID, string, string) error); ok {
		r1 = rf(context, userID, content, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartGame provides a mock function with given fields: context, userID, gameID
func (_m *Store) StartGame(context ctx.CTX, userID uuid.UUID, gameID uuid.UUID) (*models.Game, *models.Quiz, error) {
	ret := _m.Called(context, userID, gameID)

	var r0 *models.Game
	if rf, ok := ret.Get(0).(func(ctx.CTX, uuid.UUID, uuid.UUID) *models.Game); ok {
		r0 = rf(context, userID, gameID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Game)
		}
	}

	var r1 *models.Quiz
	if rf, ok := ret.Get(1).(func(ctx.CTX, uuid.UUID, uuid.UUID) *models.Quiz); ok {
		r1 = rf(context, userID, gameID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.Quiz)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(ctx.CTX, uuid.UUID, uuid.UUID) error); ok {
		r2 = rf(context, userID, gameID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
