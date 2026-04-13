package service

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/user/quantum-server/internal/dto"
)

func TestMortgageService_Calculate(t *testing.T) {
	s := &mortgageService{}

	t.Run("Standard Mortgage", func(t *testing.T) {
		input := dto.CreateMortgageRequest{
			PropertyPrice:     decimal.NewFromInt(10000000),
			DownPaymentAmount: decimal.NewFromInt(2000000),
			MortgageTermYears: 20,
			InterestRate:      decimal.NewFromInt(10),
		}
		result := s.calculate(input)

		if !result.MonthlyPayment.Equal(decimal.NewFromFloat(77201.73)) {
			t.Errorf("Expected monthly payment 77201.73, got %s", result.MonthlyPayment.String())
		}

		if result.TotalPayment.LessThan(decimal.NewFromInt(18000000)) {
			t.Errorf("Total payment seems too low: %s", result.TotalPayment.String())
		}
	})

	t.Run("Zero Interest Rate", func(t *testing.T) {
		input := dto.CreateMortgageRequest{
			PropertyPrice:     decimal.NewFromInt(1200000),
			DownPaymentAmount: decimal.NewFromInt(0),
			MortgageTermYears: 10,
			InterestRate:      decimal.NewFromInt(0),
		}
		result := s.calculate(input)

		expectedMonthly := decimal.NewFromInt(10000)
		if !result.MonthlyPayment.Equal(expectedMonthly) {
			t.Errorf("Expected monthly payment %s, got %s", expectedMonthly.String(), result.MonthlyPayment.String())
		}
	})

	t.Run("Maternity Capital Included", func(t *testing.T) {
		matCap := decimal.NewFromInt(600000)
		input := dto.CreateMortgageRequest{
			PropertyPrice:      decimal.NewFromInt(5000000),
			DownPaymentAmount:  decimal.NewFromInt(1000000),
			MatCapitalAmount:   &matCap,
			MatCapitalIncluded: true,
			MortgageTermYears:  10,
			InterestRate:       decimal.NewFromInt(10),
		}
		result := s.calculate(input)

		if !result.MonthlyPayment.Equal(decimal.NewFromFloat(44931.25)) {
			t.Errorf("Expected monthly payment 44931.25, got %s", result.MonthlyPayment.String())
		}
	})

	t.Run("Maternity Capital Not Included", func(t *testing.T) {
		matCap := decimal.NewFromInt(600000)
		inputIncluded := dto.CreateMortgageRequest{
			PropertyPrice:      decimal.NewFromInt(5000000),
			DownPaymentAmount:  decimal.NewFromInt(1000000),
			MatCapitalAmount:   &matCap,
			MatCapitalIncluded: true,
			MortgageTermYears:  10,
			InterestRate:       decimal.NewFromInt(10),
		}
		inputExcluded := inputIncluded
		inputExcluded.MatCapitalIncluded = false

		resultIncluded := s.calculate(inputIncluded)
		resultExcluded := s.calculate(inputExcluded)

		if !resultExcluded.MonthlyPayment.GreaterThan(resultIncluded.MonthlyPayment) {
			t.Errorf("Payment without matcap (%s) should be higher than with matcap (%s)",
				resultExcluded.MonthlyPayment.String(), resultIncluded.MonthlyPayment.String())
		}
	})

	t.Run("Tax Deduction Cap at 2m and 3m", func(t *testing.T) {
		input := dto.CreateMortgageRequest{
			PropertyPrice:     decimal.NewFromInt(20000000),
			DownPaymentAmount: decimal.NewFromInt(5000000),
			MortgageTermYears: 25,
			InterestRate:      decimal.NewFromInt(12),
		}
		result := s.calculate(input)

		maxTaxDeduction := decimal.NewFromInt(650000)
		if result.TaxDeduction.GreaterThan(maxTaxDeduction) {
			t.Errorf("Tax deduction should not exceed 650000, got %s", result.TaxDeduction.String())
		}

		expectedPurchasePart := decimal.NewFromInt(2000000).Mul(decimal.NewFromFloat(0.13))
		if result.TaxDeduction.LessThan(expectedPurchasePart) {
			t.Errorf("Tax deduction should include at least purchase part %s, got %s",
				expectedPurchasePart.String(), result.TaxDeduction.String())
		}
	})

	t.Run("Recommended Income Is 40% Of Monthly Payment", func(t *testing.T) {
		input := dto.CreateMortgageRequest{
			PropertyPrice:     decimal.NewFromInt(3000000),
			DownPaymentAmount: decimal.NewFromInt(1000000),
			MortgageTermYears: 10,
			InterestRate:      decimal.NewFromInt(10),
		}
		result := s.calculate(input)

		expectedIncome := result.MonthlyPayment.Mul(decimal.NewFromFloat(2.5)).Round(2)
		diff := result.RecommendedIncome.Sub(expectedIncome).Abs()
		if diff.GreaterThan(decimal.NewFromFloat(0.01)) {
			t.Errorf("Recommended income %s differs too much from expected %s (diff: %s)",
				result.RecommendedIncome.String(), expectedIncome.String(), diff.String())
		}
	})

	t.Run("Payment Schedule Has Correct Length", func(t *testing.T) {
		years := 2
		input := dto.CreateMortgageRequest{
			PropertyPrice:     decimal.NewFromInt(1000000),
			DownPaymentAmount: decimal.NewFromInt(0),
			MortgageTermYears: years,
			InterestRate:      decimal.NewFromInt(10),
		}
		result := s.calculate(input)

		totalMonths := 0
		for _, months := range result.Schedule {
			totalMonths += len(months)
		}
		if totalMonths != years*12 {
			t.Errorf("Expected %d months in schedule, got %d", years*12, totalMonths)
		}
	})

	t.Run("Schedule Balance Decreases Over Time", func(t *testing.T) {
		input := dto.CreateMortgageRequest{
			PropertyPrice:     decimal.NewFromInt(2000000),
			DownPaymentAmount: decimal.NewFromInt(0),
			MortgageTermYears: 1,
			InterestRate:      decimal.NewFromInt(10),
		}
		result := s.calculate(input)

		startTime := time.Now()
		firstMonth := startTime.AddDate(0, 1, 0).Format("January")
		lastMonth := startTime.AddDate(0, 12, 0).Format("January")
		year := startTime.AddDate(0, 1, 0).Format("2006")

		firstBalance := result.Schedule[year][firstMonth].MortgageBalance
		lastBalance := result.Schedule[year][lastMonth].MortgageBalance

		if !firstBalance.GreaterThan(lastBalance) {
			t.Errorf("Balance should decrease: first month %s, last month %s",
				firstBalance.String(), lastBalance.String())
		}

		if lastBalance.GreaterThan(decimal.NewFromFloat(1)) {
			t.Errorf("Final balance should be 0 or near 0, got %s", lastBalance.String())
		}
	})

	t.Run("Total Overpayment = TotalPayment - LoanAmount", func(t *testing.T) {
		input := dto.CreateMortgageRequest{
			PropertyPrice:     decimal.NewFromInt(5000000),
			DownPaymentAmount: decimal.NewFromInt(1000000),
			MortgageTermYears: 15,
			InterestRate:      decimal.NewFromInt(9),
		}
		result := s.calculate(input)

		loanAmount := input.PropertyPrice.Sub(input.DownPaymentAmount)
		expectedOverpayment := result.TotalPayment.Sub(loanAmount).Round(2)
		if !result.TotalOverpayment.Equal(expectedOverpayment) {
			t.Errorf("Overpayment should equal TotalPayment - LoanAmount: expected %s, got %s",
				expectedOverpayment.String(), result.TotalOverpayment.String())
		}
	})
}
