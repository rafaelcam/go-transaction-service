package domain

import "github.com/jmoiron/sqlx"

type Store interface {
	getAll() ([]Transaction, error)
}

//TransactionStore for persistence
type TransactionStore struct {
	db *sqlx.DB
}

//NewStore creates new TransactionStore for Banks
func NewStore(db *sqlx.DB) *TransactionStore {
	return &TransactionStore{db: db}
}

func (s *TransactionStore) getAll() ([]Transaction, error) {
	var transactions []Transaction
	if err := s.db.Select(&transactions, "SELECT * FROM transactions"); err != nil {
		return nil, err
	}
	return transactions, nil
}