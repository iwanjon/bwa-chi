package payment

import (
	"bwastartupgochi/helper"
	"bwastartupgochi/user"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type service struct {
}

type Service interface {
	GetPaymentUrl(transaction Transaction, user user.User) (string, error)
}

func NewPaymentService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(transaction Transaction, user user.User) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ServerKey := os.Getenv("ServerKey")
	// ServerKey := "kingkong"
	Environment := midtrans.Sandbox

	//Initiate client for Midtrans Snap
	var ss snap.Client
	ss.New(ServerKey, Environment)

	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
			// Phone: "081234567890",
			// BillAddr: custAddress,
			// ShipAddr: custAddress,
		},
	}
	// custAddress := &midtrans.CustomerAddress{
	// 	FName:       "John",
	// 	LName:       "Doe",
	// 	Phone:       "081234567890",
	// 	Address:     "Baker Street 97th",
	// 	City:        "Jakarta",
	// 	Postcode:    "16000",
	// 	CountryCode: "IDN",
	// }
	// snapRequest := &snap.Request{
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  "MID-GO-ID-eeefrwb-v14555",
	// 		GrossAmt: 200000,
	// 	},
	// 	CreditCard: &snap.CreditCardDetails{
	// 		Secure: true,
	// 	},
	// 	CustomerDetail: &midtrans.CustomerDetails{
	// 		FName:    "John",
	// 		LName:    "Doe",
	// 		Email:    "john@doe.com",
	// 		Phone:    "081234567890",
	// 		BillAddr: custAddress,
	// 		ShipAddr: custAddress,
	// 	},
	// 	EnabledPayments: snap.AllSnapPaymentType,
	// 	Items: &[]midtrans.ItemDetails{
	// 		{
	// 			ID:    "ITEM1",
	// 			Price: 200000,
	// 			Qty:   1,
	// 			Name:  "Someitem",
	// 		},
	// 	},
	// }

	// tokenstring, err := ss.CreateTransactionToken(snapRequest)
	// fmt.Println(tokenstring, "mkmk", err)
	// helper.PanicIfError(err, " erro in create token payment service")

	urlString, errr := ss.CreateTransactionUrl(snapRequest)
	if errr != nil {
		helper.PanicIfError(errors.New("error in create transaction url in payemnet service"), " error in create transaction url in payemnet service")
	}
	// helper.PanicIfError(errr.GetRawError(), " error in create transaction url in payemnet service")

	fmt.Println(urlString, "    ffffffffffffff           ")
	return urlString, nil
}

// midclient := midtrans.NewClient()
// midclient.ServerKey = "SB-Mid-server-rfoEAZy_b7QibK3ZZTM4-Jsp"
// midclient.ClientKey = "SB-Mid-client-S1AvvSdTk2QBf7D9"
// midclient.APIEnvType = midtrans.Sandbox

// var snapGateway midtrans.SnapGateway
// snapGateway := midtrans.SnapGateway{
// 	Client: midclient,
// }

// snapReq := &midtrans.SnapReq{
// 	TransactionDetails: midtrans.TransactionDetails{
// 		OrderID:  strconv.Itoa(transaction.ID),
// 		GrossAmt: int64(transaction.Amount),
// 	},
// 	CustomerDetail: &midtrans.CustDetail{
// 		FName: user.Name,
// 		Email: user.Email,
// 	},
// }
// log.Println("GetToken:")
// snapTokenResp, err := snapGateway.GetToken(snapReq)
// if err != nil {
// 	fmt.Println(snapTokenResp, "madang", err)
// 	return "error getting token", err
// }
// return snapTokenResp.RedirectURL, nil
