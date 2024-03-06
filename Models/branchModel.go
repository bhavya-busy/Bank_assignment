package models

type Branch struct {
	BranchId       uint        `pg:"branch_id,pk"`
	Bank           *Bank       `pg:"rel:has-one"`
	BankId         uint        `pg:"fk:bank_id"`
	Branch_Address string      `pg:"branch_address"`
	Branch_Name    string      `pg:"branch_name"`
	IFSC_Code      string      `pg:"ifsc_code"`
	Account        []*Account  `pg:"rel:has-many"`
	Customer       []*Customer `pg:"rel:has-many"`
}
