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
	query := "SELECT id,country,state,code,number FROM phone_numbers WHERE 1=1"
	var arguments []string

	if len(countries) > 0 {
		query = query + " AND country IN(?"
		for i := 1; i < len(countries); i++ {
			query = query + ",?"
		}
		query = query + ")"
		arguments = append(arguments, countries...)
	}
	if state != "" {
		query = query + " AND state = ?"
		arguments = append(arguments, state)
	}

	query = query + " ORDER BY code LIMIT ? OFFSET ?"
	arguments = append(arguments, limit, offset)

	args := make([]interface{}, len(arguments))
	for i, v := range arguments {
		args[i] = v
	}

	phoneNumbers := []entity.PhoneNumber{}
	tx := s.db.MustBegin()
	err := s.db.Select(&phoneNumbers, query, args...)

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
	tx := s.db.MustBegin()
	sqlResult := tx.MustExec("INSERT INTO phone_numbers (country, state, code, number) VALUES ($1, $2, $3, $4)", p.Country, p.State, p.Code, p.Number)
	RowsAffected, err := sqlResult.RowsAffected()
	if RowsAffected == 0 || err != nil {
		tx.Rollback()
		return fmt.Errorf("can't insert phone number in database: %s", err)
	}
	tx.Commit()
	return nil
}
