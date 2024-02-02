package campaign

import (
	"bwastartupgochi/exception"
	"bwastartupgochi/helper"
	"bwastartupgochi/user"
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	FindAll(ctx context.Context, tx *sql.DB) ([]Campaign, error)
	FindByUserId(ctx context.Context, tx *sql.DB, userId int) ([]Campaign, error)
	FindById(ctx context.Context, tx *sql.DB, campaignId int) (Campaign, error)
	SaveCampaign(ctx context.Context, tx *sql.Tx, campaign Campaign) (Campaign, error)
	UpdateCampaign(ctx context.Context, tx *sql.Tx, campaign Campaign) (Campaign, error)
	SaveImage(ctx context.Context, tx *sql.Tx, campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(ctx context.Context, tx *sql.Tx, campaignId int) (bool, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll(ctx context.Context, tx *sql.DB) ([]Campaign, error) {
	var campaign Campaign
	var campaigns []Campaign
	var u user.User
	var campaignImage CampaignImage

	// sqlSelect := "update campaigns set name = 'akakak' where id = $1;"
	sqlSelect := "select * from campaigns"
	stete, err := tx.PrepareContext(ctx, sqlSelect)
	helper.PanicIfError(err, " error in select repo campaign find all")
	defer stete.Close()

	results, err := stete.QueryContext(ctx)
	helper.PanicIfError(err, " error in select statement")

	sqlSelectUser := "select * from users where id = $1"
	// sqlSelectUser := "update campaigns set name = 'akakak' where id = $1"
	stete2, err := tx.PrepareContext(ctx, sqlSelectUser)
	helper.PanicIfError(err, " errror in create statement 2 campaign repo find all")
	defer stete2.Close()

	sql_campaign_images := "select * from campaign_images where campaign_id = $1"
	stet3, err := tx.PrepareContext(ctx, sql_campaign_images)
	helper.PanicIfError(err, " error in create statement 3 repsiotry campaign find all")
	defer stet3.Close()

	for results.Next() {
		var campaignImages []CampaignImage
		err := CampaignScanners(results, &campaign)
		helper.PanicIfError(err, "error in scanner campaign repository find all")

		fmt.Println(results, "ffffffffffffffffffffffffffffff", campaign)
		userResult := stete2.QueryRowContext(ctx, campaign.UserID)
		err = user.UserScanner(userResult, &u)
		helper.PanicIfError(err, " erro ins canner user in campaing repository find all")

		allImages, err := stet3.QueryContext(ctx, campaign.ID)
		helper.PanicIfError(err, " errro in select campaign images by campaign id")

		for allImages.Next() {
			err = CampaignImageScanners(allImages, &campaignImage)
			helper.PanicIfError(err, " errro in scan campaign images repository campaign find all")
			campaignImages = append(campaignImages, campaignImage)
		}
		campaign.CampaignImages = campaignImages
		campaign.User = u
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}

func (r *repository) FindByUserId(ctx context.Context, tx *sql.DB, userId int) ([]Campaign, error) {
	var campaigns []Campaign
	var campaign Campaign
	var u user.User
	var campaignImage CampaignImage

	sql_script := " select * from campaigns where user_id = $1"
	stete, err := tx.PrepareContext(ctx, sql_script)
	helper.PanicIfError(err, " error in crate statement 1 find by user id respository campaign")
	defer stete.Close()

	rows, err := stete.QueryContext(ctx, userId)
	helper.PanicIfError(err, " error in exceute find by user i drepository campaign")

	sql_2 := "  select * from users where id =$1"
	ste2, err := tx.PrepareContext(ctx, sql_2)
	helper.PanicIfError(err, " error in cerate stet 2 ")
	defer ste2.Close()

	sql3 := "select * from campaign_images where campaign_id = $1"
	ste3, err := tx.PrepareContext(ctx, sql3)
	helper.PanicIfError(err, " error incrteate ste3")
	defer ste3.Close()

	for rows.Next() {
		var campaignImages []CampaignImage
		err = CampaignScanners(rows, &campaign)
		helper.PanicIfError(err, "error in scaning campaign repository campaign find by id")

		row := ste2.QueryRowContext(ctx, campaign.UserID)
		err := user.UserScanner(row, &u)
		helper.PanicIfError(err, " error in scan user find by ser id")

		roimage, err := ste3.QueryContext(ctx, campaign.ID)
		helper.PanicIfError(err, " error in exectue sql fetch campaign images find by user id repository")
		for roimage.Next() {
			err = CampaignImageScanners(roimage, &campaignImage)
			helper.PanicIfError(err, " error inscan campaign images repository ")
			campaignImages = append(campaignImages, campaignImage)
		}
		campaign.CampaignImages = campaignImages
		campaign.User = u
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}

func (r *repository) FindById(ctx context.Context, tx *sql.DB, campaignId int) (Campaign, error) {
	// var campaigns []Campaign
	var campaign Campaign
	var u user.User
	var campaignImage CampaignImage
	var campaignImages []CampaignImage

	sql_script := " select * from campaigns where id = $1"
	stete, err := tx.PrepareContext(ctx, sql_script)
	helper.PanicIfError(err, " error in crate statement 1 find by id respository campaign")
	defer stete.Close()

	rows := stete.QueryRowContext(ctx, campaignId)
	// helper.PanicIfError(err, " error in exceute find by i drepository campaign")

	sql_2 := "  select * from users where id =$1"
	ste2, err := tx.PrepareContext(ctx, sql_2)
	helper.PanicIfError(err, " error in cerate stet 2 find by id ")
	defer ste2.Close()

	sql3 := "select * from campaign_images where campaign_id = $1"
	ste3, err := tx.PrepareContext(ctx, sql3)
	helper.PanicIfError(err, " error incrteate ste3 find by id")
	defer ste3.Close()

	err = CampaignScanner(rows, &campaign)
	helper.PanicIfError(err, " error in scan campaign find by id")

	row := ste2.QueryRowContext(ctx, campaign.UserID)
	err = user.UserScanner(row, &u)
	helper.PanicIfError(err, " error in scan user find by id")

	roimage, err := ste3.QueryContext(ctx, campaign.ID)
	helper.PanicIfError(err, " error in exectue sql fetch campaign images find by  id repository")
	for roimage.Next() {
		err = CampaignImageScanners(roimage, &campaignImage)
		helper.PanicIfError(err, " error inscan campaign images repository ")
		campaignImages = append(campaignImages, campaignImage)
	}
	campaign.CampaignImages = campaignImages
	campaign.User = u

	return campaign, nil
}

func (r *repository) SaveCampaign(ctx context.Context, tx *sql.Tx, campaign Campaign) (Campaign, error) {

	sql_query := "insert into campaigns( user_id, name, short_description, description, perks, backer_count, goal_amount, current_amount,slug) values($1, $2,$3,$4,$5,$6,$7,$8,$9) returning id"
	// var campaign Campaign
	statement_inser_campaign, err := tx.PrepareContext(ctx, sql_query)
	helper.PanicIfError(err, " error in create stetement save campaign")
	defer statement_inser_campaign.Close()

	row := statement_inser_campaign.QueryRowContext(ctx, campaign.UserID, campaign.Name, campaign.ShortDescription, campaign.Description, campaign.Perks, campaign.BackerCount, campaign.GoalAmount, campaign.CurrentAmount, campaign.Slug)
	fmt.Println(campaign.ID, "fggggg")
	err = row.Scan(&campaign.ID)
	fmt.Println(campaign.ID, "fggggg")
	helper.PanicIfError(err, "error in scan insert id")
	return campaign, nil

}

func (r *repository) UpdateCampaign(ctx context.Context, tx *sql.Tx, campaign Campaign) (Campaign, error) {

	sql_statement := "update campaigns set user_id =$1 , name = $2 , short_description= $3 , description= $4 , perks= $5 , backer_count= $6 , goal_amount= $7 , current_amount= $8 ,slug= $9 where id = $10 returning id"
	staement_update, err := tx.PrepareContext(ctx, sql_statement)
	helper.PanicIfError(err, "error create stetement update campaign")
	defer staement_update.Close()

	row := staement_update.QueryRowContext(ctx, campaign.UserID, campaign.Name, campaign.ShortDescription, campaign.Description, campaign.Perks, campaign.BackerCount, campaign.GoalAmount, campaign.CurrentAmount, campaign.Slug, campaign.ID)
	helper.PanicIfError(err, "error in exceute update campaign query")

	err = row.Scan(&campaign.ID)
	helper.PanicIfError(err, "error in scan update campaign query")

	return campaign, nil

}

func (r *repository) SaveImage(ctx context.Context, tx *sql.Tx, campaignImage CampaignImage) (CampaignImage, error) {

	sql_statement := "insert into campaign_images(campaign_id, file_name, is_primary) values($1,$2,$3) returning id"
	stetement, err := tx.PrepareContext(ctx, sql_statement)
	helper.PanicIfError(err, "error in create statement save images")
	defer stetement.Close()

	row := stetement.QueryRowContext(ctx, campaignImage.CampaignID, campaignImage.FileName, campaignImage.IsPrimary)
	helper.PanicIfError(err, " error in execute save image campaign")
	// id, err := result.LastInsertId()
	fmt.Println(campaignImage.ID, "fggggg")
	err = row.Scan(&campaignImage.ID)
	fmt.Println(campaignImage.ID, "fggggg")
	helper.PanicIfError(err, "error in getting last id campain images repository")
	// campaignImage.ID = int(id)

	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(ctx context.Context, tx *sql.Tx, campaignId int) (bool, error) {

	var campaign_image CampaignImage

	sql_get_images := " select * from campaign_images where campaign_id = $1"
	statement_get_campaign_images, err := tx.PrepareContext(ctx, sql_get_images)
	helper.PanicIfError(err, "  error in prepare statement  get campaign images")
	defer func() {
		statement_get_campaign_images.Close()
		fmt.Println("madang dap")
	}()

	result_campaign_images := statement_get_campaign_images.QueryRowContext(ctx, campaignId)

	err = CampaignImageScanner(result_campaign_images, &campaign_image)

	exception.PanicIfNotFound(err, "error not found data")
	fmt.Println(campaign_image, "alll")
	helper.PanicIfError(err, "error in execute query get campagin imagesby campaign id")

	sql_update_image_isaprimary := "update campaign_images set is_primary = 0 where campaign_id = $1 returning id, is_primary ;"

	statement_update_images, err := tx.PrepareContext(ctx, sql_update_image_isaprimary)
	helper.PanicIfError(err, " error create update isprymary")
	defer statement_update_images.Close()

	var id int
	var rr int
	result_update_images, err := statement_update_images.QueryContext(ctx, campaignId)
	helper.PanicIfError(err, "error in execute update images")
	fmt.Println(result_update_images, "xxxxx")
	for result_update_images.Next() {
		result_update_images.Scan(&id, &rr)
		fmt.Println(id, rr, "ffff")
	}

	return true, nil

}
