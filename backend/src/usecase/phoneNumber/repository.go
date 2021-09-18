package phoneNumber

import "backend/src/entity"

type Repository interface {
	Create(*entity.PhoneNumber) error
	FindAll(offset, limit, state string, countries []string) ([]entity.PhoneNumber, error)
}
