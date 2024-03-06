package models

type Customer struct {
	Cust_id      int     `pg:"cust_id,pk"`
	Cust_name    string  `pg:"cust_name"`
	Cust_address string  `pg:"cust_address"`
	Branch       *Branch `pg:"rel:has-one"`
	BranchId     uint    `pg:"fk:branch_id"`
	Map          []*Map  `pg:"rel:has-many"`
}
