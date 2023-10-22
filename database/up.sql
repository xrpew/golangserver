DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id varchar(30) PRIMARY KEY NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW()
);
