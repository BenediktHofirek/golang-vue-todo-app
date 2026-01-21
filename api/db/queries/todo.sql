-- name: Todo_CreateBlank :one
WITH existing_blank AS (
  SELECT
    id
  FROM
    todos
  WHERE
    todos.user_id = @user_id
    AND title IS NULL
    AND description IS NULL
    AND due_date IS NULL
    AND scheduled_date IS NULL
    AND completed = FALSE
    AND starred = FALSE
  LIMIT
    1
), inserted AS (
  INSERT INTO
    todos (user_id)
  SELECT
    @user_id
  WHERE
    NOT EXISTS (
      SELECT
        1
      FROM
        existing_blank
    ) RETURNING id
)
SELECT
  id
FROM
  existing_blank
UNION ALL
SELECT
  id
FROM
  inserted;

-- name: Todo_GetOneById :one
SELECT
  id,
  title,
  description,
  completed,
  due_date,
  starred,
  scheduled_date,
  created_at,
  updated_at
FROM todos
WHERE id = @id
AND user_id = @user_id;

-- name: Todo_GetMany :many
SELECT
  id,
  title,
  description,
  completed,
  due_date,
  starred,
  scheduled_date,
  created_at,
  updated_at
FROM todos
WHERE user_id = $1
ORDER BY id;

-- name: Todo_UpdateOne :one
UPDATE todos
SET title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description),
    completed = COALESCE(sqlc.narg('completed'), completed), 
    due_date = COALESCE(sqlc.narg('due_date'), due_date),
    starred = COALESCE(sqlc.narg('starred'), starred), 
    scheduled_date = COALESCE(sqlc.narg('scheduled_date'), scheduled_date)
WHERE id = @id
AND user_id = @user_id
RETURNING *;

-- name: Todo_DeleteOne :exec
DELETE FROM todos
WHERE id = @id
AND user_id = @user_id;
