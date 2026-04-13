package models

import (
	"encoding/json"
	"time"
)

type MortgageCalculation struct {
	ID                      int             `json:"id" gorm:"primaryKey"`
	UserID                  string          `json:"user_id" gorm:"type:uuid;not null"`
	MortgageProfileID       int             `json:"mortgage_profile_id" gorm:"not null"`
	MonthlyPayment          float64         `json:"monthly_payment" gorm:"not null"`
	TotalPayment            float64         `json:"total_payment" gorm:"not null"`
	TotalOverpaymentAmount  float64         `json:"total_overpayment_amount" gorm:"not null"`
	PossibleTaxDeduction    *float64        `json:"possible_tax_deduction"`
	SavingsDueMotherCapital *float64        `json:"savings_due_mother_capital"`
	RecommendedIncome       float64         `json:"recommended_income" gorm:"not null"`
	PaymentSchedule         json.RawMessage `json:"payment_schedule" gorm:"type:jsonb;not null"`
	CreatedAt               time.Time       `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt               time.Time       `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
