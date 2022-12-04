CREATE TABLE entry (
  id serial PRIMARY KEY,
  transaction_id int NOT NULL,
  account_id int NOT NULL,
  amount int NOT NULL,
  description VARCHAR(50) NOT NULL,
  type VARCHAR(20) NOT NULL,
  CONSTRAINT fk_transaction_id FOREIGN KEY (transaction_id) REFERENCES transaction (id),
  CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES account (id)

);

---- create above / drop below ----
DROP TABLE entry;
