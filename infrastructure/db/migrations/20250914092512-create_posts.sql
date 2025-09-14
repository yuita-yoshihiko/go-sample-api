
-- +migrate Up
CREATE TABLE IF NOT EXISTS posts (
  id BIGSERIAL,
  user_id BIGINT NOT NULL REFERENCES users(id),
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

comment on table posts is '投稿情報';
comment on column posts.title is 'タイトル';
comment on column posts.content is '投稿内容';

-- +migrate Down
DROP TABLE IF EXISTS posts;
