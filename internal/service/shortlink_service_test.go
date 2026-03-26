package service

import (
	"errors"
	"testing"
	"time"

	"urlshortener/internal/model"
	"urlshortener/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepo for the ShortlinkRepository
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(shortlink *model.Shortlink) error {
	args := m.Called(shortlink)
	return args.Error(0)
}

func (m *MockRepo) FindByID(id string) (*model.Shortlink, error) {
	args := m.Called(id)
	var msl *model.Shortlink
	if rf, ok := args.Get(0).(func(string) *model.Shortlink); ok {
		msl = rf(id)
	} else {
		if args.Get(0) != nil {
			msl = args.Get(0).(*model.Shortlink)
		}
	}
	return msl, args.Error(1)
}

func (m *MockRepo) FindByOriginalURL(originalURL string) (*model.Shortlink, error) {
	args := m.Called(originalURL)
	var msl *model.Shortlink
	if rf, ok := args.Get(0).(func(string) *model.Shortlink); ok {
		msl = rf(originalURL)
	} else {
		if args.Get(0) != nil {
			msl = args.Get(0).(*model.Shortlink)
		}
	}
	return msl, args.Error(1)
}

func TestCreateShortlink_InvalidURL(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewShortlinkService(mockRepo)

	_, err := svc.CreateShortlink("not-a-valid-url")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidURL, err)
}

func TestCreateShortlink_ExistingURL(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewShortlinkService(mockRepo)

	expectedModel := &model.Shortlink{
		ID:          "abcd123",
		OriginalURL: "https://example.com",
		CreatedAt:   time.Now(),
	}

	mockRepo.On("FindByOriginalURL", "https://example.com").Return(expectedModel, nil)

	shortlink, err := svc.CreateShortlink("https://example.com")
	assert.NoError(t, err)
	assert.Equal(t, expectedModel, shortlink)
	mockRepo.AssertExpectations(t)
}

func TestCreateShortlink_NewURL(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewShortlinkService(mockRepo)

	mockRepo.On("FindByOriginalURL", "https://example.com/new").Return(nil, repository.ErrNotFound)

	// First FindByID returns not found so code is available
	mockRepo.On("FindByID", mock.AnythingOfType("string")).Return(nil, repository.ErrNotFound)
	mockRepo.On("Create", mock.AnythingOfType("*model.Shortlink")).Return(nil)

	shortlink, err := svc.CreateShortlink("https://example.com/new")
	assert.NoError(t, err)
	assert.NotNil(t, shortlink)
	assert.Equal(t, "https://example.com/new", shortlink.OriginalURL)
	assert.Len(t, shortlink.ID, 7)
	mockRepo.AssertExpectations(t)
}

func TestGetShortlink_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewShortlinkService(mockRepo)

	expectedModel := &model.Shortlink{
		ID:          "abcd123",
		OriginalURL: "https://example.com",
		CreatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", "abcd123").Return(expectedModel, nil)

	shortlink, err := svc.GetShortlink("abcd123")
	assert.NoError(t, err)
	assert.Equal(t, expectedModel, shortlink)
	mockRepo.AssertExpectations(t)
}

func TestGetShortlink_ErrorNotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewShortlinkService(mockRepo)

	mockRepo.On("FindByID", "not-found").Return(nil, repository.ErrNotFound)

	shortlink, err := svc.GetShortlink("not-found")
	assert.Error(t, err)
	assert.Nil(t, shortlink)
	assert.True(t, errors.Is(err, repository.ErrNotFound))
	mockRepo.AssertExpectations(t)
}
