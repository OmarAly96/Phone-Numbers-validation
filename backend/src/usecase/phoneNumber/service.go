package phoneNumber

import (
	"backend/src/api/model"
	"backend/src/entity"
	"backend/src/pkg"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type Service struct {
	Repository Repository
	Logger     *zerolog.Logger
}

func LoadService(repository Repository, logger *zerolog.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}

func (s *Service) FindAllPhoneNumbers(offset, limit, country, state string) ([]entity.PhoneNumber, error) {
	s.MirgratePhoneNumbers()
	var countries []string
	if country != "" {
		countries = strings.Split(country, ",")
	}

	phoneNumbers, err := s.Repository.FindAll(offset, limit, state, countries)
	if err != nil {
		return []entity.PhoneNumber{}, fmt.Errorf("can't get phone numbers %s", err)
	}
	return phoneNumbers, nil
}

func (s *Service) CreatePhoneNumber(n model.PhoneNumber) error {
	number, ok := n.(string)
	if ok {
		code := pkg.SegregateCode(number)

		if err := pkg.CodeExists(code); err != nil {
			return err
		}

		c := pkg.CodeCountryExpression(code)

		phoneNumber := &entity.PhoneNumber{
			Country: c.Country,
			Code:    "+" + code,
			Number:  number,
		}
		phoneNumber.ValidateState(c.Exp)
		if err := s.Repository.Create(phoneNumber); err != nil {
			s.Logger.Error().Msgf("can't create phone number: %v", number)
			return fmt.Errorf("can't create phone number: %s", err)
		}
		return nil
	}
	return entity.ErrInvalidInput
}

func (s *Service) MirgratePhoneNumbers() error {
	phoneNumbers, err := s.Repository.FindPhoneFromCustomerNotInPhoneNumbers()
	if err != nil {
		return fmt.Errorf("can't load numbers: %s", err)
	}
	for _, number := range phoneNumbers {
		err := s.CreatePhoneNumber(number)
		if err != nil {
			s.Logger.Error().Msgf("can't create phone number: %v", number)
		}
	}

	return nil
}
