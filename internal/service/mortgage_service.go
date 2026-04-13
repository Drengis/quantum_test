package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"github.com/user/quantum-server/internal/domain"
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/repository"
)

type MortgageService interface {
	CreateCalculation(ctx context.Context, input dto.CreateMortgageRequest) (*domain.MortgageCalculation, error)
	ProcessCalculation(ctx context.Context, id int, input dto.CreateMortgageRequest) error
	GetCalculation(ctx context.Context, id int) (*dto.MortgageResponse, error)
}

type mortgageService struct {
	db          *sql.DB
	profileRepo repository.MortgageProfileRepository
	calcRepo    repository.MortgageCalculationRepository
	taskChan    chan<- dto.MortgageTask
}

func NewMortgageService(db *sql.DB, profileRepo repository.MortgageProfileRepository, calcRepo repository.MortgageCalculationRepository, taskChan chan<- dto.MortgageTask) MortgageService {
	return &mortgageService{
		db:          db,
		profileRepo: profileRepo,
		calcRepo:    calcRepo,
		taskChan:    taskChan,
	}
}

func (s *mortgageService) CreateCalculation(ctx context.Context, input dto.CreateMortgageRequest) (*domain.MortgageCalculation, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	matCap := decimal.Zero
	if input.MatCapitalAmount != nil {
		matCap = *input.MatCapitalAmount
	}

	profile := &domain.MortgageProfile{
		UserID:             input.UserID,
		PropertyPrice:      input.PropertyPrice,
		PropertyType:       input.PropertyType,
		DownPaymentAmount:  input.DownPaymentAmount,
		MatCapitalAmount:   matCap,
		MatCapitalIncluded: input.MatCapitalIncluded,
		MortgageTermYears:  input.MortgageTermYears,
		InterestRate:       input.InterestRate,
	}
	if err := s.profileRepo.Create(ctx, tx, profile); err != nil {
		return nil, err
	}

	calculation := &domain.MortgageCalculation{
		UserID:            input.UserID,
		MortgageProfileID: profile.ID,
		Status:            domain.StatusPending,
	}

	if err := s.calcRepo.Create(ctx, tx, calculation); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Отправляем в воркер через канал
	if s.taskChan != nil {
		s.taskChan <- dto.MortgageTask{
			CalcID:  calculation.ID,
			Request: input,
		}
	}

	return calculation, nil
}

func (s *mortgageService) ProcessCalculation(ctx context.Context, id int, input dto.CreateMortgageRequest) error {
	calc, err := s.calcRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if calc == nil {
		return fmt.Errorf("calculation %d not found", id)
	}

	calc.Status = domain.StatusProcessing
	s.calcRepo.Update(ctx, calc)

	calcResult := s.calculate(input)

	scheduleJSON, _ := json.Marshal(calcResult.Schedule)
	calc.Status = domain.StatusCompleted
	calc.MonthlyPayment = calcResult.MonthlyPayment
	calc.TotalPayment = calcResult.TotalPayment
	calc.TotalOverpaymentAmount = calcResult.TotalOverpayment
	calc.PossibleTaxDeduction = calcResult.TaxDeduction
	calc.SavingsDueMotherCapital = calcResult.SavingsMotherCapital
	calc.RecommendedIncome = calcResult.RecommendedIncome
	calc.PaymentSchedule = scheduleJSON

	return s.calcRepo.Update(ctx, calc)
}

func (s *mortgageService) GetCalculation(ctx context.Context, id int) (*dto.MortgageResponse, error) {
	calc, err := s.calcRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if calc == nil {
		return nil, nil
	}

	var schedule dto.MortgagePaymentSchedule
	json.Unmarshal(calc.PaymentSchedule, &schedule)

	return &dto.MortgageResponse{
		ID:                      fmt.Sprintf("%d", calc.ID),
		Status:                  string(calc.Status),
		ErrorMessage:            calc.ErrorMessage,
		MonthlyPayment:          calc.MonthlyPayment.StringFixed(2),
		TotalPayment:            calc.TotalPayment.StringFixed(2),
		TotalOverpaymentAmount:  calc.TotalOverpaymentAmount.StringFixed(2),
		PossibleTaxDeduction:    calc.PossibleTaxDeduction.StringFixed(2),
		SavingsDueMotherCapital: calc.SavingsDueMotherCapital.StringFixed(2),
		RecommendedIncome:       calc.RecommendedIncome.StringFixed(2),
		MortgagePaymentSchedule: schedule,
	}, nil
}

type internalCalcResult struct {
	MonthlyPayment       decimal.Decimal
	TotalPayment         decimal.Decimal
	TotalOverpayment     decimal.Decimal
	TaxDeduction         decimal.Decimal
	SavingsMotherCapital decimal.Decimal
	RecommendedIncome    decimal.Decimal
	Schedule             dto.MortgagePaymentSchedule
}

func (s *mortgageService) calculate(input dto.CreateMortgageRequest) internalCalcResult {
	// 1. Сумма кредита
	matCap := decimal.Zero
	if input.MatCapitalAmount != nil && input.MatCapitalIncluded {
		matCap = *input.MatCapitalAmount
	}
	loanAmount := input.PropertyPrice.Sub(input.DownPaymentAmount).Sub(matCap)

	// 2. Количество месяцев
	totalMonths := int64(input.MortgageTermYears * 12)

	// 3. Месячная ставка
	monthlyRate := input.InterestRate.Div(decimal.NewFromInt(12)).Div(decimal.NewFromInt(100))

	// 4. Ежемесячный платеж
	monthlyPayment := decimal.Zero
	if monthlyRate.GreaterThan(decimal.Zero) {
		onePlusI := decimal.NewFromInt(1).Add(monthlyRate)
		pow := onePlusI.Pow(decimal.NewFromInt(totalMonths))

		num := loanAmount.Mul(monthlyRate).Mul(pow)
		den := pow.Sub(decimal.NewFromInt(1))
		monthlyPayment = num.Div(den)
	} else {
		monthlyPayment = loanAmount.Div(decimal.NewFromInt(totalMonths))
	}

	// 5. Сумма выплат
	totalPayment := monthlyPayment.Mul(decimal.NewFromInt(totalMonths))

	// 6. Переплата
	overpayment := totalPayment.Sub(loanAmount)

	// 7. Налоговый вычет
	// Покупка: min(Price, 2,000,000) * 0.13
	taxPurchase := decimal.Min(input.PropertyPrice, decimal.NewFromInt(2000000)).Mul(decimal.NewFromFloat(0.13))
	// Проценты: min(Overpayment, 3,000,000) * 0.13
	taxInterests := decimal.Min(overpayment, decimal.NewFromInt(3000000)).Mul(decimal.NewFromFloat(0.13))
	taxDeduction := taxPurchase.Add(taxInterests)

	// 8. Рекомендованный доход
	recommendedIncome := monthlyPayment.Div(decimal.NewFromFloat(0.4))

	// 9. График
	schedule := make(dto.MortgagePaymentSchedule)
	balance := loanAmount
	startTime := time.Now()

	for m := int64(1); m <= totalMonths; m++ {
		currDate := startTime.AddDate(0, int(m), 0)
		yearKey := currDate.Format("2006")
		monthKey := currDate.Format("January")

		if _, ok := schedule[yearKey]; !ok {
			schedule[yearKey] = make(map[string]dto.MortgagePayment)
		}

		interestMonth := balance.Mul(monthlyRate)
		principalMonth := monthlyPayment.Sub(interestMonth)
		balance = balance.Sub(principalMonth)
		if balance.IsNegative() {
			balance = decimal.Zero
		}

		schedule[yearKey][monthKey] = dto.MortgagePayment{
			TotalPayment:                monthlyPayment.Round(2),
			RepaymentOfMortgageBody:     principalMonth.Round(2),
			RepaymentOfMortgageInterest: interestMonth.Round(2),
			MortgageBalance:             balance.Round(2),
		}
	}

	return internalCalcResult{
		MonthlyPayment:       monthlyPayment.Round(2),
		TotalPayment:         totalPayment.Round(2),
		TotalOverpayment:     overpayment.Round(2),
		TaxDeduction:         taxDeduction.Round(2),
		SavingsMotherCapital: matCap,
		RecommendedIncome:    recommendedIncome.Round(2),
		Schedule:             schedule,
	}
}
