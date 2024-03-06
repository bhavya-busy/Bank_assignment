package models

type Bank struct {
	Bank_Id   uint      `pg:"bank_id,pk"`
	Bank_name string    `pg:"bank_name"`
	Branches  []*Branch `pg:"rel:has-many"`
}
