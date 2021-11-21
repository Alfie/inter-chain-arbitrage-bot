package exchange

// Representation of a token
// Example: Token{Name: "OSMO", Denom: "uosmo", Amount: 10000}

type Token struct {
	Name   string
	Denom  string
	Amount int
}
