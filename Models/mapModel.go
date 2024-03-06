package models

type Map struct {
	Map_id   uint      `pg:"map_id,pk"`
	Cust_id  uint      `pg:"fk:cust_id"`
	Customer *Customer `pg:"rel:has-one"`
	Ac_id    uint      `pg:"fk:ac_id"`
	Account  *Account  `pg:"rel:has-one"`
}
