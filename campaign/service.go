package campaign

import (
	"bwastartupgochi/exception"
	"bwastartupgochi/helper"
	"errors"

	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	FindCampaigns(ctx context.Context, userId int) ([]Campaign, error)
	GetCampaignById(ctx context.Context, input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(ctx context.Context, input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(ctx context.Context, inputID GetCampaignDetailInput, inputparam CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(ctx context.Context, input CreateCampaignImageInput, filelocation string) (CampaignImage, error)
}

type service struct {
	repo     Repository
	db       *sql.DB
	Validate *validator.Validate
}

func NewService(repo Repository, db *sql.DB, Validate *validator.Validate) Service {
	return &service{repo, db, Validate}
}

func (s *service) FindCampaigns(ctx context.Context, userId int) ([]Campaign, error) {
	if userId == 0 {
		campaigns, err := s.repo.FindAll(ctx, s.db)
		exception.PanicIfNotFound(err, " error in finding campaigns")
		return campaigns, nil
	}
	campaigns, err := s.repo.FindByUserId(ctx, s.db, userId)
	exception.PanicIfNotFound(err, " error in finding campaigns by user id")
	return campaigns, nil

}

func (s *service) GetCampaignById(ctx context.Context, input GetCampaignDetailInput) (Campaign, error) {
	cmapign, err := s.repo.FindById(ctx, s.db, input.ID)
	exception.PanicIfNotFound(err, " errror in finding campaign bu campaign id")
	return cmapign, nil
}

func (s *service) CreateCampaign(ctx context.Context, input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{
		UserID:           input.User.ID,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		Slug:             fmt.Sprintf("%d-%s", input.User.ID, input.Name),
		// CreatedAt:        time.Time{},
		// UpdatedAt:        time.Time{},
		// CampaignImages:   []CampaignImage{},
		// User:             user.User{},
	}
	tx, err := s.db.Begin()
	helper.PanicIfError(err, " error in create tx increate campaign service")
	defer helper.CommitOrRollback(tx)
	newCampaign, err := s.repo.SaveCampaign(ctx, tx, campaign)
	helper.PanicIfError(err, " error in create in create campaign service")
	return newCampaign, nil
}

func (s *service) UpdateCampaign(ctx context.Context, inputID GetCampaignDetailInput, inputparam CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{
		ID:               inputID.ID,
		UserID:           inputparam.User.ID,
		Name:             inputparam.Name,
		ShortDescription: inputparam.ShortDescription,
		Description:      inputparam.Description,
		Perks:            inputparam.Perks,
		GoalAmount:       inputparam.GoalAmount,
		Slug:             fmt.Sprintf("%d-%s", inputparam.User.ID, inputparam.Name),
	}
	tx, err := s.db.Begin()
	helper.PanicIfError(err, " error in create tx in update campaign service")
	defer helper.CommitOrRollback(tx)
	cc, err := s.repo.FindById(ctx, s.db, inputID.ID)
	helper.PanicIfError(err, " error in finding campaign bu=y id in campaign service")
	if cc.UserID != inputparam.User.ID {
		exception.PanicIfNotOwner(errors.New("errror not owner"), "not owbner error  save campaign image service")
	}
	updatedCampaign, err := s.repo.UpdateCampaign(ctx, tx, campaign)
	helper.PanicIfError(err, " error in create in update campaign service")
	return updatedCampaign, nil

}

func (s *service) SaveCampaignImage(ctx context.Context, input CreateCampaignImageInput, filelocation string) (CampaignImage, error) {
	campaignimage := CampaignImage{
		CampaignID: input.CampaignID,
		FileName:   filelocation,
		IsPrimary:  0,
	}
	if input.IsPrimary {
		campaignimage.IsPrimary = 1
	}

	tx, err := s.db.Begin()
	helper.PanicIfError(err, " error in create tx in save  campaign image service")
	defer helper.CommitOrRollback(tx)

	cc, err := s.repo.FindById(ctx, s.db, input.CampaignID)
	helper.PanicIfError(err, " error in finding campaign bu=y id in campaign service")
	if cc.UserID != input.User.ID {
		exception.PanicIfNotOwner(errors.New("errror not owner"), "not owbner error save campaign image service")
	}
	newCampaignImage, err := s.repo.SaveImage(ctx, tx, campaignimage)
	helper.PanicIfError(err, " erorr in save campaign image")
	return newCampaignImage, nil

}
