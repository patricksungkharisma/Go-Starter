package app

const (
	queryGetExampleData = `
		SELECT *
		FROM example_table
		WHERE
			id = $1
	`
)
