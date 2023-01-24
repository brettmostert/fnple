CREATE TABLE account (
  id serial PRIMARY KEY,
  coalation_id UUID NOT NULL UNIQUE,
  balance int NOT NULL DEFAULT 0,
  description VARCHAR(50) NOT NULL,
  created_time TIMESTAMP NOT NULL,
  modified_time TIMESTAMP NOT NULL
);

---- create above / drop below ----
DROP TABLE account;
