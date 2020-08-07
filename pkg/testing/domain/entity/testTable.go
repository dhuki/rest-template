package entity

// by default gorm convert camelcase (TestTable) become snakecase and lowercase (new_price)
// and name of type struct become plurar (test_tables)

// pacakge for defined mapping database
type TestTable struct {
	ID          string `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
