package transaction

import "database/sql"

func TransactionScanner(r *sql.Row, t *Transaction) error {
	err := r.Scan(
		&t.ID,
		&t.CampaignID,
		&t.UserID,
		&t.Amount,
		&t.Status,
		&t.Code,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.PaymentURL,
	)
	return err
}

func TransactionScanners(r *sql.Rows, t *Transaction) error {
	err := r.Scan(
		&t.ID,
		&t.CampaignID,
		&t.UserID,
		&t.Amount,
		&t.Status,
		&t.Code,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.PaymentURL,
	)
	return err
}
