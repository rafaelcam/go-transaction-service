package domain

import "time"

// Transaction domain model
type Transaction struct {
	ID          string    `json:"id" db:"id"`
	CustomerId  string    `json:"customer_id" db:"customer_id"`
	Description string    `json:"description" db:"description"`
	Amount      int64     `json:"amount" db:"amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

//Service instance to manage Transactions
type Service struct {
	s Store
}

//NewService creates new Transaction Service
func NewService(store Store) *Service {
	return &Service{s: store}
}

func (svc Service) GetTransactions() ([]Transaction, error) {
	return svc.s.getAll()
}