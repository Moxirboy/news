package postgres

const (
	NewsCreate = `
		INSERT INTO news(
		title,content,author,category,publication_date)	
		VALUES($1,$2,$3,$4,$5)
		RETURNING id
		`

	NewsGetById = `	SELECT
		id,
		title,
		content,
		author,
		category,
		publication_date
	FROM news
WHERE id=CAST($1 AS INT) AND deletedAt IS NULL`

	NewsGetAll = `
SELECT
	id,
	title,
	content,
	author,
	category,
	publication_date
FROM news 
WHERE deletedAt IS NULL
OFFSET $1
LIMIT $2
`
	NewsCount = `
	SELECT COUNT(id) FROM news WHERE 1=1
`
	NewsUpdate = `
	Update news
SET title=$1,
content=$2,
author=$3,
category=$4,
updatedAt=CURRENT_TIMESTAMP
where id=CAST($5 AS INT)
	`
	NewsDelete = `
	UPDATE news
SET deletedAt=$2
WHERE id=CAST($1 AS INT)
`
)
