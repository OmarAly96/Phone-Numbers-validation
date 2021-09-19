package dbRepo

import (
	"backend/src/api/model"
	"backend/src/entity"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
)

type SqlLiteDB struct {
	db *sqlx.DB
}

func NewSqlLiteDB(db *sqlx.DB) *SqlLiteDB {
	return &SqlLiteDB{
		db: db,
	}
}

func (s *SqlLiteDB) FindAll(offset, limit, state string, countries []string) ([]entity.PhoneNumber, error) {
	baseQuery := "SELECT id,country,state,code,number FROM phone_numbers WHERE 1=1"
	var arguments []string

	if len(countries) > 0 {
		baseQuery = baseQuery + " AND country IN(?"
		for i := 1; i < len(countries); i++ {
			baseQuery = baseQuery + ",?"
		}
		baseQuery = baseQuery + ")"
		arguments = append(arguments, countries...)
	}
	if state != "" {
		baseQuery = baseQuery + " AND state = ?"
		arguments = append(arguments, state)
	}

	arguments = append(arguments, limit, offset)
	baseQuery = baseQuery + " LIMIT ? OFFSET ?"

	phoneNumbers := []entity.PhoneNumber{}

	args := make([]interface{}, len(arguments))
	for i, v := range arguments {
		args[i] = v
	}
	tx := s.db.MustBegin()
	err := s.db.Select(&phoneNumbers, baseQuery, args...)

	if err != nil {
		tx.Rollback()
		return []entity.PhoneNumber{}, fmt.Errorf("can't select phone numbers from database %s", err)
	}
	tx.Commit()
	return phoneNumbers, nil
}

func (s *SqlLiteDB) FindPhoneFromCustomerNotInPhoneNumbers() ([]model.PhoneNumber, error) {
	baseQuery := "SELECT phone FROM customer WHERE phone NOT IN (SELECT number FROM phone_numbers)"
	phones := []model.PhoneNumber{}
	tx := s.db.MustBegin()
	err := s.db.Select(&phones, baseQuery)
	if err != nil {
		tx.Rollback()
		return []model.PhoneNumber{}, fmt.Errorf("can't select phones from database %s", err)
	}
	tx.Commit()
	return phones, nil
}

func (s *SqlLiteDB) Create(p *entity.PhoneNumber) error {
	var state int
	if p.State {
		state = 1
	}
	tx := s.db.MustBegin()
	tx.MustExec("INSERT INTO phone_numbers (country, state, code, number) VALUES ($1, $2, $3, $4)", p.Country, state, p.Code, p.Number)
	// err := tx.MustExec("INSERT INTO phone_numbers (country, state, code, number) VALUES ($1, $2, $3, $4)", p.Country, state, p.Code, p.Number)
	// if err != nil {
	// 	tx.Rollback()
	// 	return fmt.Errorf("can't insert phone number in database: %s", err)
	// }
	tx.Commit()
	return nil
}
