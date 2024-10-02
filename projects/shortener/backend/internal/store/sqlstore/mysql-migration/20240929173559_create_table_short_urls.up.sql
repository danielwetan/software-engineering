CREATE TABLE short_urls (
  shortcode varchar(10) NOT NULL,
  user_id int NOT NULL DEFAULT 1,
  target varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (shortcode)
);

CREATE INDEX user_id ON short_urls (user_id);
CREATE INDEX idx_shortcode_user_id ON short_urls (shortcode, user_id);

-- guest user seed
INSERT INTO users(email, password) 
VALUES ("guest@shortener.com", "-");