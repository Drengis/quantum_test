package dto

import (
	"github.com/shopspring/decimal"
	"github.com/user/quantum-server/internal/domain"
)

type CreateMortgageRequest struct {
	UserID             string              `json:"user_id"`
	PropertyPrice      decimal.Decimal     `json:"propertyPrice"`
	PropertyType       domain.PropertyType `json:"propertyType"`
	DownPaymentAmount  decimal.Decimal     `json:"downPaymentAmount"`
	MatCapitalAmount   *decimal.Decimal    `json:"matCapitalAmount"`
	MatCapitalIncluded bool                `json:"matCapitalIncluded"`
	MortgageTermYears  int                 `json:"mortgageTermYears"`
	InterestRate       decimal.Decimal     `json:"interestRate"`
}

type MortgagePayment struct {
	TotalPayment                decimal.Decimal `json:"totalPayment"`
	RepaymentOfMortgageBody     decimal.Decimal `json:"repaymentOfMortgageBody"`
	RepaymentOfMortgageInterest decimal.Decimal `json:"repaymentOfMortgageInterest"`
	MortgageBalance             decimal.Decimal `json:"mortgageBalance"`
}

type MortgagePaymentSchedule map[string]map[string]MortgagePayment

type MortgageTask struct {
	CalcID  int
	Request CreateMortgageRequest
}

type MortgageResponse struct {
	ID                      string                  `json:"id"`
	Status                  string                  `json:"status"`
	ErrorMessage            *string                 `json:"errorMessage,omitempty"`
	MonthlyPayment          string                  `json:"monthlyPayment"`
	TotalPayment            string                  `json:"totalPayment"`
	TotalOverpaymentAmount  string                  `json:"totalOverpaymentAmount"`
	PossibleTaxDeduction    string                  `json:"possibleTaxDeduction"`
	SavingsDueMotherCapital string                  `json:"savingsDueMotherCapital"`
	RecommendedIncome       string                  `json:"recommendedIncome"`
	MortgagePaymentSchedule MortgagePaymentSchedule `json:"mortgagePaymentSchedule"`
}
