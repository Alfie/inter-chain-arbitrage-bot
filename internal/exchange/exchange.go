package exchange

// Interface that all exchanges, e.g. Osmosis, Gravity and Sifchain, need to implement.

type Exchange interface {
	Deposit()                                       // IBC deposit token
	Withdraw()                                      // IBC withdraw tokens
	Swap(in, out string, amount int, poolIds []int) // Swap tokens via the specified pool route
	Balances() []Token                              // Query and return the current DEX balances
	Pools() []Pool                                  // Query and return the current DEX pools
}
