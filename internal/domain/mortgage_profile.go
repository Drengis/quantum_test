package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type PropertyType string

const (
	PropertyTypeApartmentNew       PropertyType = "apartment_in_new_building"
	PropertyTypeApartmentSecondary PropertyType = "apartment_in_secondary_building"
	PropertyTypeHouse              PropertyType = "house"
	PropertyTypeHouseLand          PropertyType = "house_with_land_plot"
	PropertyTypeLand               PropertyType = "land_plot"
	PropertyTypeOther              PropertyType = "other"
)

type MortgageProfile struct {
	ID                 int             `json:"id" gorm:"primaryKey"`
	UserID             string          `json:"user_id" gorm:"type:uuid;not null"`
	PropertyPrice      decimal.Decimal `json:"property_price" gorm:"type:numeric;not null"`
	PropertyType       PropertyType    `json:"property_type" gorm:"type:property_type;not null"`
	DownPaymentAmount  decimal.Decimal `json:"down_payment_amount" gorm:"type:numeric;not null"`
	MatCapitalAmount   decimal.Decimal `json:"mat_capital_amount" gorm:"type:numeric"`
	MatCapitalIncluded bool            `json:"mat_capital_included" gorm:"not null"`
	MortgageTermYears  int             `json:"mortgage_term_years" gorm:"not null"`
	InterestRate       decimal.Decimal `json:"interest_rate" gorm:"type:numeric;not null"`
	CreatedAt          time.Time       `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt          time.Time       `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
