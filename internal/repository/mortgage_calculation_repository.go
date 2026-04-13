package repository

import (
	"context"
	"database/sql"

	"github.com/user/quantum-server/internal/models"
)

type MortgageCalculationRepository interface {
	Create(ctx context.Context, tx *sql.Tx, calc *models.MortgageCalculation) error
	FindByID(ctx context.Context, id int) (*models.MortgageCalculation, error)
}

type mortgageCalculationRepository struct {
	db *sql.DB
}

func NewMortgageCalculationRepository(db *sql.DB) MortgageCalculationRepository {
	return &mortgageCalculationRepository{db: db}
}

func (r *mortgageCalculationRepository) Create(ctx context.Context, tx *sql.Tx, calc *models.MortgageCalculation) error {
	query := `INSERT INTO mortgage_calculations (user_id, mortgage_profile_id, monthly_payment, total_payment, total_overpayment_amount, possible_tax_deduction, savings_due_mother_capital, recommended_income, payment_schedule) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at, updated_at`

	return tx.QueryRowContext(ctx, query,
		calc.UserID, calc.MortgageProfileID, calc.MonthlyPayment, calc.TotalPayment,
		calc.TotalOverpaymentAmount, calc.PossibleTaxDeduction, calc.SavingsDueMotherCapital, calc.RecommendedIncome, calc.PaymentSchedule).
		Scan(&calc.ID, &calc.CreatedAt, &calc.UpdatedAt)
}

func (r *mortgageCalculationRepository) FindByID(ctx context.Context, id int) (*models.MortgageCalculation, error) {
	query := `SELECT id, user_id, mortgage_profile_id, monthly_payment, total_payment, total_overpayment_amount, possible_tax_deduction, savings_due_mother_capital, recommended_income, payment_schedule, created_at, updated_at 
	FROM mortgage_calculations WHERE id = $1`

	c := &models.MortgageCalculation{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.UserID, &c.MortgageProfileID, &c.MonthlyPayment, &c.TotalPayment,
		&c.TotalOverpaymentAmount, &c.PossibleTaxDeduction, &c.SavingsDueMotherCapital, &c.RecommendedIncome, &c.PaymentSchedule, &c.CreatedAt, &c.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return c, err
}
