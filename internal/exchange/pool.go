package exchange

// Representation of a pool.
// Example: pool := exchange.Pool{Id: 3, Tokens: []exchange.Token{
//	exchange.Token{Name: "OSMO", Denom: "uosmo", Amount: 3},
// 	exchange.Token{Name: "ION", Denom: "uion", Amount: 2}}
// }

type Pool struct {
	Id     int
	Tokens []Token
}
