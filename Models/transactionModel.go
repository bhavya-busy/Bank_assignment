package models

type Transaction struct {
	Transc_id    uint     `pg:"trans_id,pk"`
	Ac_id        uint     `pg:"fk:ac_id"`
	Type         string   `pg:"type"`
	Tranc_Medium string   `pg:"tranc_medium"`
	Tranc_Amt    uint     `pg:"tranc_amt"`
	Account      *Account `pg:"rel:has-one"`
}
