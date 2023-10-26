package repository

const (
	DELETE_BY_ID = "DELETE FROM users WHERE id = $1"
	ADD_USER     = "INSERT INTO users (name, surname, patronymic, age, gender, nationalize) VALUES (:name, :surname, :patronymic, :age, :gender, :nationalize) RETURNING *"
	UPDATE_BY_ID = `UPDATE users SET name = COALESCE(NULLIF(:name, ''), name), surname = COALESCE(NULLIF(:surname, ''), surname), patronymic = COALESCE(NULLIF(:patronymic, ''), patronymic), age = COALESCE(NULLIF(:age, 0), age), gender = COALESCE(NULLIF(:gender, ''), gender), nationalize = COALESCE(NULLIF(:nationalize, ''), nationalize)
WHERE id = :id
RETURNING *
`
)
