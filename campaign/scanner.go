package campaign

import "database/sql"

func CampaignScanner(r *sql.Row, c *Campaign) error {

	err := r.Scan(
		&c.ID,
		&c.Name,
		&c.ShortDescription,
		&c.Description,
		&c.Perks,
		&c.BackerCount,
		&c.GoalAmount,
		&c.CurrentAmount,
		&c.Slug,
		&c.UserID,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	return err
}

func CampaignImageScanner(r *sql.Row, c *CampaignImage) error {
	err := r.Scan(
		&c.ID,
		&c.CampaignID,
		&c.FileName,
		&c.IsPrimary,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	return err

}

func CampaignScanners(r *sql.Rows, c *Campaign) error {

	err := r.Scan(
		&c.ID,
		&c.Name,
		&c.ShortDescription,
		&c.Description,
		&c.Perks,
		&c.BackerCount,
		&c.GoalAmount,
		&c.CurrentAmount,
		&c.Slug,
		&c.UserID,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	return err
}

func CampaignImageScanners(r *sql.Rows, c *CampaignImage) error {
	err := r.Scan(
		&c.ID,
		&c.FileName,
		&c.IsPrimary,
		&c.CampaignID,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	return err

}
