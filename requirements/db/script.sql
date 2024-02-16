
CREATE TABLE client (
  id INT PRIMARY KEY NOT NULL,
  name VARCHAR(70) NOT NULL,
  debit_balance BIGINT DEFAULT 0,
  credit_limit BIGINT NOT NULL,
  createAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transation (
  id UUID PRIMARY KEY NOT NULL,
  value BIGINT NOT NULL,
  type CHAR(1) CHECK (type IN ('c', 'd')) NOT NULL,
  description VARCHAR(10),
  client_id INT NOT NULL,
  createAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (client_id) REFERENCES client(id)
);

DO $$
BEGIN
  INSERT INTO client (id, name, credit_limit)
  VALUES
    (1, 'o barato sai caro', 1000 * 100),
    (2, 'zan corp ltda', 800 * 100),
    (3, 'les cruders', 10000 * 100),
    (4, 'padaria joia de cocaia', 100000 * 100),
    (5, 'kid mais', 5000 * 100);
END; $$



