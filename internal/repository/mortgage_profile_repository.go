package repository

import (
	"context"
	"database/sql"

	"github.com/user/quantum-server/internal/domain"
)

type MortgageProfileRepository interface {
	Create(ctx context.Context, tx *sql.Tx, profile *domain.MortgageProfile) error
	FindByID(ctx context.Context, id int) (*domain.MortgageProfile, error)
}

type mortgageProfileRepository struct {
	db *sql.DB
}

func NewMortgageProfileRepository(db *sql.DB) MortgageProfileRepository {
	return &mortgageProfileRepository{db: db}
}

func (r *mortgageProfileRepository) Create(ctx context.Context, tx *sql.Tx, profile *domain.MortgageProfile) error {
	query := `INSERT INTO mortgage_profiles (user_id, property_price, property_type, down_payment_amount, mat_capital_amount, mat_capital_included, mortgage_term_years, interest_rate) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`

	idRow := tx.QueryRowContext(ctx, query,
		profile.UserID, profile.PropertyPrice, profile.PropertyType, profile.DownPaymentAmount,
		profile.MatCapitalAmount, profile.MatCapitalIncluded, profile.MortgageTermYears, profile.InterestRate)

	return idRow.Scan(&profile.ID, &profile.CreatedAt, &profile.UpdatedAt)
}

func (r *mortgageProfileRepository) FindByID(ctx context.Context, id int) (*domain.MortgageProfile, error) {
	query := `SELECT id, user_id, property_price, property_type, down_payment_amount, mat_capital_amount, mat_capital_included, mortgage_term_years, interest_rate, created_at, updated_at 
	FROM mortgage_profiles WHERE id = $1`

	p := &domain.MortgageProfile{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.UserID, &p.PropertyPrice, &p.PropertyType, &p.DownPaymentAmount,
		&p.MatCapitalAmount, &p.MatCapitalIncluded, &p.MortgageTermYears, &p.InterestRate, &p.CreatedAt, &p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}
