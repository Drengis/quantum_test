package domain

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type CalculationStatus string

const (
	StatusPending    CalculationStatus = "pending"
	StatusProcessing CalculationStatus = "processing"
	StatusCompleted  CalculationStatus = "completed"
	StatusFailed     CalculationStatus = "failed"
)

type MortgageCalculation struct {
	ID                      int               `json:"id" gorm:"primaryKey"`
	UserID                  string            `json:"user_id" gorm:"type:uuid;not null"`
	MortgageProfileID       int               `json:"mortgage_profile_id" gorm:"not null"`
	Status                  CalculationStatus `json:"status" gorm:"not null;default:pending"`
	ErrorMessage            *string           `json:"error_message"`
	MonthlyPayment          decimal.Decimal   `json:"monthly_payment" gorm:"type:numeric;not null"`
	TotalPayment            decimal.Decimal   `json:"total_payment" gorm:"type:numeric;not null"`
	TotalOverpaymentAmount  decimal.Decimal   `json:"total_overpayment_amount" gorm:"type:numeric;not null"`
	PossibleTaxDeduction    decimal.Decimal   `json:"possible_tax_deduction" gorm:"type:numeric"`
	SavingsDueMotherCapital decimal.Decimal   `json:"savings_due_mother_capital" gorm:"type:numeric"`
	RecommendedIncome       decimal.Decimal   `json:"recommended_income" gorm:"type:numeric;not null"`
	PaymentSchedule         json.RawMessage   `json:"payment_schedule" gorm:"type:jsonb;not null"`
	CreatedAt               time.Time         `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt               time.Time         `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
