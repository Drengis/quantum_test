CREATE TABLE IF NOT EXISTS mortgage_calculations (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    mortgage_profile_id INTEGER NOT NULL,
    monthly_payment NUMERIC NOT NULL,
    total_payment NUMERIC NOT NULL,
    total_overpayment_amount NUMERIC NOT NULL,
    possible_tax_deduction NUMERIC,
    savings_due_mother_capital NUMERIC,
    recommended_income NUMERIC NOT NULL,
    payment_schedule JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (mortgage_profile_id) REFERENCES mortgage_profiles(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_mortgage_calculations_user_id ON mortgage_calculations(user_id);
CREATE INDEX IF NOT EXISTS idx_mortgage_calculations_profile_id ON mortgage_calculations(mortgage_profile_id);
