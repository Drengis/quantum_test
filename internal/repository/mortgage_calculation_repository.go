package repository

import (
	"context"
	"database/sql"

	"github.com/user/quantum-server/internal/domain"
)

type MortgageCalculationRepository interface {
	Create(ctx context.Context, tx *sql.Tx, calc *domain.MortgageCalculation) error
	Update(ctx context.Context, calc *domain.MortgageCalculation) error
	FindByID(ctx context.Context, id int) (*domain.MortgageCalculation, error)
}

type mortgageCalculationRepository struct {
	db *sql.DB
}

func NewMortgageCalculationRepository(db *sql.DB) MortgageCalculationRepository {
	return &mortgageCalculationRepository{db: db}
}

func (r *mortgageCalculationRepository) Create(ctx context.Context, tx *sql.Tx, calc *domain.MortgageCalculation) error {
	query := `INSERT INTO mortgage_calculations (user_id, mortgage_profile_id, status, error_message, monthly_payment, total_payment, total_overpayment_amount, possible_tax_deduction, savings_due_mother_capital, recommended_income, payment_schedule) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, created_at, updated_at`

	return tx.QueryRowContext(ctx, query,
		calc.UserID, calc.MortgageProfileID, calc.Status, calc.ErrorMessage,
		calc.MonthlyPayment, calc.TotalPayment,
		calc.TotalOverpaymentAmount, calc.PossibleTaxDeduction, calc.SavingsDueMotherCapital, calc.RecommendedIncome, calc.PaymentSchedule).
		Scan(&calc.ID, &calc.CreatedAt, &calc.UpdatedAt)
}

func (r *mortgageCalculationRepository) Update(ctx context.Context, calc *domain.MortgageCalculation) error {
	query := `UPDATE mortgage_calculations 
	SET status = $1, error_message = $2, monthly_payment = $3, total_payment = $4, total_overpayment_amount = $5, 
	    possible_tax_deduction = $6, savings_due_mother_capital = $7, recommended_income = $8, payment_schedule = $9, updated_at = NOW() 
	WHERE id = $10`

	_, err := r.db.ExecContext(ctx, query,
		calc.Status, calc.ErrorMessage, calc.MonthlyPayment, calc.TotalPayment,
		calc.TotalOverpaymentAmount, calc.PossibleTaxDeduction, calc.SavingsDueMotherCapital, calc.RecommendedIncome, calc.PaymentSchedule, calc.ID)
	return err
}

func (r *mortgageCalculationRepository) FindByID(ctx context.Context, id int) (*domain.MortgageCalculation, error) {
	query := `SELECT id, user_id, mortgage_profile_id, status, error_message, monthly_payment, total_payment, total_overpayment_amount, possible_tax_deduction, savings_due_mother_capital, recommended_income, payment_schedule, created_at, updated_at 
	FROM mortgage_calculations WHERE id = $1`

	c := &domain.MortgageCalculation{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.UserID, &c.MortgageProfileID, &c.Status, &c.ErrorMessage, &c.MonthlyPayment, &c.TotalPayment,
		&c.TotalOverpaymentAmount, &c.PossibleTaxDeduction, &c.SavingsDueMotherCapital, &c.RecommendedIncome, &c.PaymentSchedule, &c.CreatedAt, &c.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return c, err
}
