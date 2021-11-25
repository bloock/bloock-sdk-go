package service

import (
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	anchorException "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity/exception"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/repository"
	"github.com/enchainte/enchainte-sdk-go/internal/config/service"
	"time"
)

type AnchorService struct {
	anchorRepository repository.AnchorerRepository
	configService    service.ConfigurerService
}

func NewAnchorService(ar repository.AnchorerRepository, conf service.ConfigurerService) AnchorService {
	return AnchorService{
		anchorRepository: ar,
		configService:    conf,
	}
}

func (a AnchorService) GetAnchor(anchorId int) (entity.Anchor, error) {
	anchor, err := a.anchorRepository.GetAnchor(anchorId)
	if err != nil {
		return entity.Anchor{}, err
	}

	if anchor.Status() != "Success" {
		return entity.Anchor{}, anchorException.NewAnchorNotFoundException()
	}

	return anchor, nil
}

func (a AnchorService) WaitAnchor(anchorId int, limit int) (entity.Anchor, error) {
	waitDefault := a.configService.GetConfiguration().WaitMessageIntervalDefault
	waitFactor := a.configService.GetConfiguration().WaitMessageIntervalFactor

	var attempts = 0
	var start = time.Now()
	var nextTry = start.Add(time.Millisecond * time.Duration(waitDefault))
	var timeout = start.Add(time.Millisecond * time.Duration(limit))

	for true {
		anchor, err := a.anchorRepository.GetAnchor(anchorId)
		if err != nil {
			return entity.Anchor{}, err
		}
		if anchor.Status() == "Success" {
			return anchor, nil
		}
		currentTime := time.Now()

		if currentTime.After(timeout) {
			return entity.Anchor{}, anchorException.NewWaitAnchorTimeoutException()
		}
		time.Sleep(time.Millisecond * 1000)

		for currentTime.Before(nextTry) && currentTime.Before(timeout) {
			time.Sleep(time.Millisecond * 200)
			currentTime = time.Now()
		}
		if currentTime.After(timeout) {
			return entity.Anchor{}, anchorException.NewWaitAnchorTimeoutException()
		}
		nextTry.Add(time.Millisecond * time.Duration(attempts*waitFactor+waitDefault))
		attempts += 1

		if currentTime.After(timeout) {
			return entity.Anchor{}, anchorException.NewWaitAnchorTimeoutException()
		}
	}

	return entity.Anchor{}, nil
}
