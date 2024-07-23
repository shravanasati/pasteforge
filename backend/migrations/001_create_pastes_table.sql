CREATE TABLE pastes (
	id VARCHAR(8) PRIMARY KEY,
	content TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	expires_at TIMESTAMP,
	visibility VARCHAR(10) NOT NULL DEFAULT 'public',
	language  VARCHAR(20) NOT NULL,
	password VARCHAR(255)
);

---- create above / drop below ----

DROP TABLE pastes;
