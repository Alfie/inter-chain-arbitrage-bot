package exchange

import "github.com/sirupsen/logrus"

// Implementation for the Osmosis Blockchain and DEX.
// Uses https://github.com/osmosis-labs/osmosis for key management and blockchain interaction.
// For more infos see the developer documentation.

type Osmosis struct{}

func (osmosis Osmosis) Deposit() {
	logrus.Debug("Deposit not implemented!")
}

func (osmosis Osmosis) Withdraw() {
	logrus.Debug("Withdraw not implemented!")
}

func (osmosis Osmosis) Swap(in, out string, amount int, poolIds []int) {
	logrus.Debug("Swap not implemented!")
}

func (osmosis Osmosis) Balances() (tokens []Token) {
	logrus.Debug("Balances not implemented!")
	return
}

func (osmosis Osmosis) Pools() (pools []Pool) {
	logrus.Debug("Pools not implemented!")
	return
}
