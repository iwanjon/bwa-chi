package transaction

import "database/sql"

func TransactionScanner(r *sql.Row, t *Transaction) error {
	err := r.Scan(
		&t.ID,
		&t.Amount,
		&t.Status,
		&t.Code,
		&t.PaymentURL,
		&t.CampaignID,
		&t.UserID,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	return err
}

func TransactionScanners(r *sql.Rows, t *Transaction) error {
	err := r.Scan(
		&t.ID,
		&t.Amount,
		&t.Status,
		&t.Code,
		&t.PaymentURL,
		&t.CampaignID,
		&t.UserID,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	return err
}
