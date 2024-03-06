package models

type Account struct {
	Ac_no       uint           `pg:"ac_no"`
	Ac_id       uint           `pg:"ac_id,pk"`
	Ac_type     string         `pg:"ac_type"`
	Balance     uint           `pg:"balance"`
	BranchId    uint           `pg:"fk:branch_id"`
	Transaction []*Transaction `pg:"rel:has-many"`
	Branch      *Branch        `pg:"rel:has-one"`
	Map         []*Map         `pg:"rel:has-many"`
}
