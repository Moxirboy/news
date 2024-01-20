package postgres

const(
	NewsCreate = `
		INSERT INTO news(
		title,content,author,category,publication_date)	
		VALUES($1,$2,$3,$4,$5)
		RETURNING id
		`
	
	)