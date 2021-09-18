package phoneNumber

import (
	"backend/src/api/model"
	"backend/src/entity"
)

type UseCase interface {
	CreatePhoneNumber(model.PhoneNumber) error
	FindAllPhoneNumbers(offset, limit, country, state string) ([]entity.PhoneNumber, error)
}
