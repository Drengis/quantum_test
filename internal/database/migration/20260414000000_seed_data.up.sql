
INSERT INTO users (id, tg_id, username, first_name, last_name, lang_code, is_active) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '123456789', 'testuser', 'Test', 'User', 'ru', true),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', '987654321', 'john_doe', 'John', 'Doe', 'en', true)
ON CONFLICT (tg_id) DO NOTHING;

INSERT INTO mortgage_profiles (user_id, property_price, property_type, down_payment_amount, mat_capital_amount, mat_capital_included, mortgage_term_years, interest_rate) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 5000000, 'apartment_in_new_building', 1000000, 0, false, 20, 8.5),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 8000000, 'house', 2000000, 200000, true, 15, 7.5),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 12000000, 'apartment_in_secondary_building', 3000000, 0, false, 25, 9.0);

INSERT INTO mortgage_calculations (user_id, mortgage_profile_id, monthly_payment, total_payment, total_overpayment_amount, possible_tax_deduction, savings_due_mother_capital, recommended_income, payment_schedule) VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 1, 35000.00, 8400000.00, 4400000.00, 520000.00, 0, 100000.00, '[{"month":1,"payment":35000,"principal":2500,"interest":32500}]'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 2, 55000.00, 9900000.00, 3900000.00, 390000.00, 200000, 160000.00, '[{"month":1,"payment":55000,"principal":5000,"interest":50000}]'),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 3, 80000.00, 24000000.00, 9000000.00, 900000.00, 0, 230000.00, '[{"month":1,"payment":80000,"principal":3000,"interest":77000}]');