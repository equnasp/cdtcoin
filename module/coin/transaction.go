package coin

/*func NewTransaction(from, to string, amount int, ) *Transaction {

}*/

// Transaction ...
type Transaction struct {
	ID  []byte
	Out []Output
	In  []Input
}

// Output ...
type Output struct {
	Amount    int
	PublicKey string
}

// Input ...D
type Input struct {
	ID        []byte
	Amount    int
	PublicKey string
}
