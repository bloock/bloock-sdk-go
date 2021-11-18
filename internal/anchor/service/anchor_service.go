package service

import (
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	exception2 "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity/exception"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/repository"
	"github.com/enchainte/enchainte-sdk-go/internal/shared/entity/exception"
	"reflect"
	"time"
)

type AnchorService struct {
	anchorRepository repository.AnchorerRepository
	configService service.ConfigurerService
}

func NewAnchorService(ar repository.AnchorerRepository, conf service.ConfigurerService) AnchorService {
	return AnchorService{
		anchorRepository: ar,
		configService: conf,
	}
}

func(a AnchorService) GetAnchor(anchorId int) (entity.Anchor, error) {
	id := reflect.TypeOf(anchorId).Kind()
	if id != reflect.Int {
		return entity.Anchor{}, exception.NewInvalidArgumentException()
	}

	anchor, err := a.anchorRepository.GetAnchor(anchorId)
	if err != nil {
		return entity.Anchor{}, err
	}

	if anchor.Status() != "Success" {
		return entity.Anchor{}, exception2.NewAnchorNotFoundException()
	}

	return anchor, nil
}

func(a AnchorService) WaitAnchor(anchorId int, limit int) (entity.Anchor, error) {
	id := reflect.TypeOf(anchorId).Kind()
	timeLimit := reflect.TypeOf(limit).Kind()
	if id != reflect.Int || timeLimit != reflect.Int {
		return entity.Anchor{}, exception.NewInvalidArgumentException()
	}

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
			return entity.Anchor{}, exception2.NewWaitAnchorTimeoutException()
		}
		time.Sleep(time.Millisecond * 1000)

		for currentTime < nextTry && currentTime < timeout {
			time.Sleep(time.Millisecond * 200)
			currentTime = time.Now().Unix()
		}
		if currentTime > timeout {
			return entity.Anchor{}, exception2.NewWaitAnchorTimeoutException()
		}
		nextTry += int64(attempts * waitFactor + waitDefault)
		attempts += 1

		if currentTime > timeout {
			return entity.Anchor{}, exception2.NewWaitAnchorTimeoutException()
		}
	}

	return entity.Anchor{}, nil
}
