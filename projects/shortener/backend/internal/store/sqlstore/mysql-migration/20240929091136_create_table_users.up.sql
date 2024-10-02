CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  email varchar(100) NOT NULL,
  password varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE INDEX idx_email ON users (email);
