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
	var start = time.Now().Unix()
	var nextTry = start + int64(waitDefault)
	var timeout = start + int64(limit)

	for true {
		anchor, err := a.anchorRepository.GetAnchor(anchorId)
		if err != nil {
			return entity.Anchor{}, err
		}
		if anchor.Status() == "Success" {
			return anchor, nil
		}
		currentTime := time.Now().Unix()

		if currentTime > timeout {
			return entity.Anchor{}, anchorException.NewWaitAnchorTimeoutException()
		}
		time.Sleep(time.Millisecond * 1000)

		for currentTime < nextTry && currentTime < timeout {
			time.Sleep(time.Millisecond * 200)
			currentTime = time.Now().Unix()
		}
		if currentTime > timeout {
			return entity.Anchor{}, anchorException.NewWaitAnchorTimeoutException()
		}
		nextTry += int64(attempts*waitFactor + waitDefault)
		attempts += 1

		if currentTime > timeout {
			return entity.Anchor{}, anchorException.NewWaitAnchorTimeoutException()
		}
	}

	return entity.Anchor{}, nil
}
