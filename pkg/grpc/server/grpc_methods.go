package grpc

import (
	"context"
	"fmt"
	"smartcard/internal/app/model/mongocontrol"
	service "smartcard/internal/app/model/service"

	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"
)

func (server *GRPCServer) RegisterCardData(ctx context.Context, req *api.RegistrateRequest) (*api.RegistrateResponse, error) {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Logrus.Panic("Recovered panic ", rec)
		}
	}()

	var status, errText string
	var id string

	regCardData, err := getDecodingCardData(req)
	if err != nil {
		log.Logrus.Errorf("Error Unmarshal = %v", err)
		errText = fmt.Sprintf("Error Unmarshal card data : %v", err)
		status = string(service.FAILED)
	}

	if errText == "" {
		id, err = service.AddOne(ctx, regCardData)
		if err != nil {
			log.Logrus.Errorf("Error inserting data : %v", err)
			errText = fmt.Sprintf("Error inserting card data : %v", err)
			status = string(service.FAILED)
		}
	}
	response := &api.RegistrateResponse{
		Id:        id,
		Status:    service.ConvertStatus(status),
		ErrorText: errText,
	}
	return response, nil
}

func (server *GRPCServer) GetCardData(ctx context.Context, req *api.GetDataRequest) (*api.GetDataResponse, error) {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Logrus.Panic("Recovered panic ", rec)
		}
	}()

	var card *mongocontrol.CardData
	var byteCard []byte
	var status, errText string

	idCard, err := getIdCard(req)
	if err != nil {
		errText = fmt.Sprintf("Error get ObjectID : %v", err)
		status = string(service.FAILED)
	}
	if errText == "" {
		card, err = service.GetOne(ctx, idCard)
		if err != nil {
			log.Logrus.Errorf("Error get data : %v", err)
			errText = fmt.Sprintf("Error get card data : %v", err)
			status = string(service.FAILED)
		}
	}
	if errText == "" {
		byteCard, err = getEncodingCardData(card)
		if err != nil {
			log.Logrus.Errorf("Error Marshal = %v", err)
			errText = fmt.Sprintf("Error Marshal bytes : %v", err)
			status = string(service.FAILED)
		}
	}
	response := &api.GetDataResponse{
		Data:      byteCard,
		Status:    service.ConvertStatus(status),
		ErrorText: errText,
	}
	return response, nil
}
