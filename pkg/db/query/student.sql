-- name: CreateStudent :one
INSERT INTO students (
   student_id,first_name, last_name, age, email, gender, favourite_color, student_address, created_at, updated_at, deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;
-- name: CreateAddress :one
INSERT INTO address
    (address_id, street, city, planet, phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetStudent :one
SELECT * FROM students
WHERE id = $1 AND deleted = false;

-- name: ListStudents :many
SELECT * FROM students
WHERE deleted = false
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateStudent :exec
UPDATE students
SET first_name = $2, last_name = $3, age = $4 , email = $5
WHERE id = $1 AND deleted = false;

-- name: DeleteStudent :exec
UPDATE students
SET deleted = true
WHERE id = $1 AND deleted = false;

