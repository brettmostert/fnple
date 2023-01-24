CREATE TABLE transaction (
  id serial PRIMARY KEY,
  coalation_id UUID NOT NULL UNIQUE,
  description VARCHAR(50) NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_time TIMESTAMP NOT NULL,
  modified_time TIMESTAMP NOT NULL
);

---- create above / drop below ----
DROP TABLE transaction;
