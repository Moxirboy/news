
package postgres

const (
	BlogCreate = `
INSERT INTO blog(
title,content,author,publication_date)	
VALUES($1,$2,$3,$4)
RETURNING id
`

	BlogGetById = `	SELECT
id,
title,
content,
author,
publication_date
FROM blog
WHERE id=CAST($1 AS INT) AND deletedAt IS NULL`

	BlogGetAll = `
SELECT
id,
title,
content,
author,
publication_date
FROM blog 
WHERE deletedAt IS NULL
OFFSET $1
LIMIT $2
`
	BlogCount = `
SELECT COUNT(id) FROM blog WHERE 1=1
`
	BlogUpdate = `
Update blog
SET title=$1,
content=$2,
author=$3,
updatedAt=CURRENT_TIMESTAMP
where id=CAST($4 AS INT)
`
	BlogDelete = `
UPDATE blog
SET deletedAt=$2
WHERE id=CAST($1 AS INT)
`
)
