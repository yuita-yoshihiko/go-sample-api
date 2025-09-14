DELETE FROM comments;
DELETE FROM posts;
DELETE FROM users;

SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('users', 'id'), 1, false);
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('posts', 'id'), 1, false);
SELECT pg_catalog.setval(pg_catalog.pg_get_serial_sequence('comments', 'id'), 1, false);

INSERT INTO users (
    id,
    name,
    email,
    created_at,
    updated_at
) VALUES (
    1,
    'テストユーザー1',
    'user@example.com',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
),(
    2,
    'テストユーザー2',
    'user2@example.com',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
);
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));

INSERT INTO posts (
    id,
    user_id,
    title,
    content,
    created_at,
    updated_at
) VALUES (
    1,
    1,
    'テスト投稿1',
    'テスト投稿1の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
), (
    2,
    1,
    'テスト投稿2',
    'テスト投稿2の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
), (
    3,
    2,
    'テスト投稿3',
    'テスト投稿3の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
), (
    4,
    2,
    'テスト投稿4',
    'テスト投稿4の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
);
SELECT setval('posts_id_seq', (SELECT MAX(id) FROM posts));

INSERT INTO comments (
    id,
    post_id,
    content,
    created_at,
    updated_at
) VALUES (
    1,
    1,
    'テストコメント1の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
), (
    2,
    1,
    'テストコメント2の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
), (
    3,
    2,
    'テストコメント3の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
), (
    4,
    3,
    'テストコメント4の内容',
    '2025-01-01T00:00:00Z',
    '2025-01-01T00:00:00Z'
);
SELECT setval('comments_id_seq', (SELECT MAX(id) FROM comments));
