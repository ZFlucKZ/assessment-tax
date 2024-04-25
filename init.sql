-- Creation of product table
CREATE TYPE deduction_type AS ENUM ('Personal', 'KReceipt', 'Donation');

CREATE TABLE IF NOT EXISTS deduction (
	id SERIAL PRIMARY KEY,
	deduction_type deduction_type NOT NULL,
	amount DECIMAL(10, 1) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO deduction (deduction_type, amount) VALUES
('Personal', 60000.0),
('Donation', 100000.0),
('KReceipt', 50000.0);
