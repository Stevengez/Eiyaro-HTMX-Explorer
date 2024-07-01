package models

type TxInput struct {
	Type              string
	Asset_id          string
	Asset_definition  any
	Amount            int64
	Arbitrary         string
	Input_id          string
	Witness_arguments any
	Sign_data         string
}

type TxOutput struct {
	Type             string
	Id               string
	Position         int64
	Asset_id         string
	Asset_definition any
	Amount           int64
	Control_program  string
	Address          string
}

type BlockTx struct {
	Id          string
	Version     int64
	Size        int64
	Time_range  int64
	Inputs      []TxInput
	Outputs     []TxOutput
	Status_fail bool
	Mux_id      string
}

type Block struct {
	Hash                    string
	Size                    int64
	Version                 int64
	Height                  int64
	Previous_block_hash     string
	Timestamp               int64
	Nonce                   int64
	Bits                    int64
	Difficulty              string
	Transaction_merkle_root string
	Transaction_status_hash string
	Transactions            []BlockTx
}
