
-- +migrate Up
CREATE TABLE IF NOT EXISTS comments (
  id BIGSERIAL,
  post_id BIGINT NOT NULL REFERENCES posts(id),
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

comment on table comments is 'コメント情報';
comment on column comments.content is 'コメント内容';

-- +migrate Down
DROP TABLE IF EXISTS comments;
