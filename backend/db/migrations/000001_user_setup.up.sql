CREATE TABLE "users" (
  id BIGSERIAL PRIMARY KEY
  email varchar(256) UNIQUE NOT NULL
  hashed_password varchar(256) NOT NULL
  created_at timestampz NOT NULL DEFAULT now()
  updated_at timestampz NOT NULL DEFAULT now()
);
