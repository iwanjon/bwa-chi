package helper

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx) {
	fmt.Println("this is tx")
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback, "eror rollback")
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit, "error commit")
	}
}
