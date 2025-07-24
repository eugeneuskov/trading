package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"trading/app/models"
)

const (
	signalValueColumn = "value"
)

type signalRepository struct {
	db *sqlx.DB
}

func newSignalRepository(db *sqlx.DB) *signalRepository {
	return &signalRepository{db}
}

func (s *signalRepository) Insert(symbol, value string) (*models.Signal, error) {
	createdAt := time.Now()

	_, err := s.db.Exec(fmt.Sprintf(`
			insert into %s
			(%s, %s, %s)
		`,
		signalTableName,
		symbolColumn,
		signalValueColumn,
		createdAtColumn,
	),
		symbol,
		value,
		createdAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.Signal{
		Symbol:    symbol,
		Value:     value,
		CreatedAt: createdAt,
	}, nil
}

func (s *signalRepository) GetBySymbol(symbol string) (*models.Signal, error) {
	var signal models.Signal

	err := s.db.Select(
		&signal,
		fmt.Sprintf(`
			select * from %s
			where %s = $1
		`,
			signalTableName,
			symbolColumn,
		),
		symbol,
	)
	if err != nil {
		return nil, err
	}

	return &signal, nil
}

func (s *signalRepository) Update(signal *models.Signal, value string) (*models.Signal, error) {
	createdAt := time.Now()

	_, err := s.db.Exec(
		fmt.Sprintf(`
			update %s
			set %s = $1, %s = $2
			where %s = $3
		`,
			signalTableName,
			signalValueColumn,
			createdAtColumn,
			symbolColumn,
		),
		value,
		createdAt,
		signal.Symbol,
	)
	if err != nil {
		return nil, err
	}

	return &models.Signal{
		Symbol:    signal.Symbol,
		Value:     value,
		CreatedAt: createdAt,
	}, nil
}
