package postgres

const (
	createComment=`
	insert into comment(blog_id,body) values($1,$2)
`
	getAllComment=`
SELECT
*
FROM comment
WHERE id=$1
OFFSET $2
LIMIT $3
`
	countComment=`
SELECT COUNT(id) FROM comment WHERE 1=1
`
	getOneComment=`
SELECT
*
FROM comment
WHERE id=$1
`
	updateComment=`
Update comment
SET body=$1
where id=CAST($2 AS INT)
`
	deleteComment=`
DELETE FROM comment
WHERE id = CAST($1 AS INT);

`
)