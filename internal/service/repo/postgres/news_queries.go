package postgres

const (
	NewsCreate = `
		INSERT INTO news(
		title,content,author,category,publication_date)	
		VALUES($1,$2,$3,$4,$5)
		RETURNING id
		`
	NewsGetById = `
	SELECT
	title,
	content,
	author,
	category,
	publication_date
	FROM news WHERE id=$1
    WHERE deletedAt NULL
	`
	NewsGetAll = `
SELECT
id,
title,
content,
author,
category,
publication_date
FROM news LIMIT 10 OFFSET $1
WHERE deletedAt NULL
`
	NewsUpdate = `
	Update news
SET title=$1,
content=$2,
author=$3,
category=$4,
updatedAt=CURRENT_TIMESTAMP
where id=$5
	`
	NewsDelete = `
	UPDATE news
SET DeletedAt=CURRENT_TIMESTAMP
WHERE id=$1
`
	)
