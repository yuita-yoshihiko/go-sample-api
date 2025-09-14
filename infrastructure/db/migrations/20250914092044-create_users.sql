
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

comment on table users is 'ユーザー情報';
comment on column users.name is 'ユーザー名';
comment on column users.email is 'メールアドレス';

-- +migrate Down
DROP TABLE IF EXISTS users;
