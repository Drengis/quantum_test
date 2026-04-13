package models

import "time"

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
	ID                 int          `json:"id" gorm:"primaryKey"`
	UserID             string       `json:"user_id" gorm:"type:uuid;not null"`
	PropertyPrice      float64      `json:"property_price" gorm:"not null"`
	PropertyType       PropertyType `json:"property_type" gorm:"type:property_type;not null"`
	DownPaymentAmount  float64      `json:"down_payment_amount" gorm:"not null"`
	MatCapitalAmount   *float64     `json:"mat_capital_amount"`
	MatCapitalIncluded bool         `json:"mat_capital_included" gorm:"not null"`
	MortgageTermYears  int          `json:"mortgage_term_years" gorm:"not null"`
	InterestRate       float64      `json:"interest_rate" gorm:"not null"`
	CreatedAt          time.Time    `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt          time.Time    `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
