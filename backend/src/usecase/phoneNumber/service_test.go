package phoneNumber

import (
	"backend/src/api/model"
	"backend/src/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (r *MockRepository) Create(*entity.PhoneNumber) error {
	args := r.Called()
	return args.Error(0)
}
func (r *MockRepository) FindAll(offset, limit, state string, countries []string) ([]entity.PhoneNumber, error) {
	args := r.Called()
	result := args.Get(0)
	return result.([]entity.PhoneNumber), args.Error(1)
}
func (r *MockRepository) FindPhoneFromCustomerNotInPhoneNumbers() ([]model.PhoneNumber, error) {
	args := r.Called()
	result := args.Get(0)
	return result.([]model.PhoneNumber), args.Error(1)
}

func TestFindAllPhoneNumbers(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		mockRepo := new(MockRepository)
		phoneNumber := entity.PhoneNumber{
			Id:      1,
			Country: "Morocco",
			State:   true,
			Code:    "+212",
			Number:  "(212) 698054317",
		}
		mockRepo.On("FindAll").Return([]entity.PhoneNumber{phoneNumber}, nil)
		mockRepo.On("FindPhoneFromCustomerNotInPhoneNumbers").Return([]model.PhoneNumber{}, nil)
		testService := LoadService(mockRepo, nil)
		result, _ := testService.FindAllPhoneNumbers("1", "0", "", "")
		mockRepo.AssertExpectations(t)
		assert.Equal(t, phoneNumber, result[0])
	})
	t.Run("test fail", func(t *testing.T) {
		mockRepo := new(MockRepository)
		e := errors.New("Mock intended error")
		mockRepo.On("FindAll").Return([]entity.PhoneNumber{}, e)
		mockRepo.On("FindPhoneFromCustomerNotInPhoneNumbers").Return([]model.PhoneNumber{}, nil)
		testService := LoadService(mockRepo, nil)
		result, err := testService.FindAllPhoneNumbers("1", "0", "", "")
		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
