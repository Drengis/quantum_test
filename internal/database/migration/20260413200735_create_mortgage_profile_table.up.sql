CREATE TYPE property_type AS ENUM (
    'apartment_in_new_building',
    'apartment_in_secondary_building',
    'house',
    'house_with_land_plot',
    'land_plot',
    'other'
);

CREATE TABLE IF NOT EXISTS mortgage_profiles (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    property_price NUMERIC NOT NULL,
    property_type property_type NOT NULL,
    down_payment_amount NUMERIC NOT NULL,
    mat_capital_amount NUMERIC,
    mat_capital_included BOOLEAN NOT NULL,
    mortgage_term_years INTEGER NOT NULL,
    interest_rate NUMERIC NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_mortgage_profiles_user_id ON mortgage_profiles(user_id);
