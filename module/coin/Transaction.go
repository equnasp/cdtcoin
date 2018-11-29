package coin

/*func NewTransaction(from, to string, amount int, ) *Transaction {

}*/

// Transaction ...
type Transaction struct {
	Id  []byte
	Out []Output
	In  []Input
}

// Output ...
type Output struct {
	Amount    int
	PublicKey string
}

// Input ...
type Input struct {
	Id        []byte
	Amount    int
	PublicKey string
}