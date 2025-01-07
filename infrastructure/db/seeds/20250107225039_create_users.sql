-- +goose Up
-- +goose StatementBegin
INSERT INTO users
  (id, name, created_at, updated_at)
  VALUES
  (
    1,
    'test-1',
    NOW(),
    NOW()
  ),
  (
    2,
    'test-2',
    NOW(),
    NOW()
  ),
  (
    3,
    'test-3',
    NOW(),
    NOW()
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
-- +goose StatementEnd
