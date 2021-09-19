package phoneNumber

import (
	"backend/src/api/model"
	"backend/src/entity"
)

type Repository interface {
	Create(*entity.PhoneNumber) error
	FindAll(offset, limit, state string, countries []string) ([]entity.PhoneNumber, error)
	FindPhoneFromCustomerNotInPhoneNumbers() ([]model.PhoneNumber, error)
}
