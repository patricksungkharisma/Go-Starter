package app

const (
	QueryGetExampleData = `
		SELECT *
		FROM example_table
		WHERE
			id = $1
	`
)
