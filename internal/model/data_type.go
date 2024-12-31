package model

// DataType - type of data
type DataType = int

// Constanst of data type
const (
	// CredentialsData - data type of credentials
	CredentialsData DataType = iota
	// BinariesData - data type of binaries
	BinariesData
	// NotesData - data type of notes
	NotesData
	// BankCardData - data type of vank cards
	BankCardData
)
