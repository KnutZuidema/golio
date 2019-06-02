package model

// DataDragonResponse represents a response from the Data Dragon service
type DataDragonResponse struct {
	Type    string
	Format  string
	Version string
	Data    interface{}
}
