-- name: Todo_CreateOne :one
INSERT INTO todos (
  user_id,
  title,
  description,
  completed,
  due_date
)
VALUES (@user_id, @title, @description, @completed, @due_date)
RETURNING *;

-- name: Todo_GetOneById :one
SELECT
  id,
  user_id,
  title,
  description,
  completed,
  due_date,
  created_at,
  updated_at
FROM todos
WHERE id = @id
AND user_id = @user_id;

-- name: Todo_GetManyByUserId :many
SELECT
  id,
  user_id,
  title,
  description,
  completed,
  due_date,
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
    due_date = COALESCE(sqlc.narg('due_date'), due_date)
WHERE id = @id
AND user_id = @user_id
RETURNING *;

-- name: Todo_DeleteOne :exec
DELETE FROM todos
WHERE id = @id
AND user_id = @user_id;
