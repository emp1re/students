--name: GetStudents :many
SELECT * FROM students
ORDER BY id DESC;
--name: GetStudent :one
SELECT * FROM students
WHERE id = $1 LIMIT 1;
--name: CreateStudent :one
INSERT INTO students (first_name, age, grade, teacher_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
--name: UpdateStudent :one
UPDATE students
SET name = $1, age = $2, grade = $3, teacher_id = $4
WHERE id = $5
RETURNING *;
--name: DeleteStudent :one
DELETE FROM students
WHERE id = $1
RETURNING *;



