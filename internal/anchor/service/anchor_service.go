package service

import (
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	exception2 "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity/exception"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/repository"
	"github.com/enchainte/enchainte-sdk-go/internal/shared/entity/exception"
	"log"
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

	var attempts = 0
	var start = time.Now().Unix()
	log.Printf("First time: %v\n", start)
	var nextTry = start + int64(a.configService.GetConfiguration().WaitMessageIntervalDefault)
	log.Printf("First nextTry: %v\n", nextTry)
	var timeout = start + int64(limit)
	log.Printf("First timeout: %v\n", timeout)

	for true {
		anchor, err := a.anchorRepository.GetAnchor(anchorId)
		log.Printf("Anchor waitting: %+v\n", anchor)
		if err != nil {
			return entity.Anchor{}, err
		}
		if anchor.Status() == "Success" {
			return anchor, nil
		}
		currentTime := time.Now().Unix()
		log.Printf("Second time: %v\n", currentTime)

		if currentTime > timeout {
			return entity.Anchor{}, exception2.NewWaitAnchorTimeoutException()
		}
		log.Println("Waiting sleep 1000...")
		time.Sleep(time.Millisecond * 1000)

		for currentTime < nextTry && currentTime < timeout {
			time.Sleep(time.Millisecond * 200)
			currentTime = time.Now().Unix()
			log.Println("I'm in loop")
		}
		if currentTime > timeout {
			return entity.Anchor{}, exception2.NewWaitAnchorTimeoutException()
		}
		nextTry += int64(attempts * a.configService.GetConfiguration().WaitMessageIntervalFactor +
			a.configService.GetConfiguration().WaitMessageIntervalDefault)
		attempts += 1
		log.Printf("Next Try: %v, Attempts: %v", nextTry, attempts)

		if currentTime > timeout {
			return entity.Anchor{}, exception2.NewWaitAnchorTimeoutException()
		}
	}

	/*var start = time.Now().Unix()
	log.Printf("First time: %v\n", start)
	var nextTry = (start + int64(a.configService.GetConfiguration().WaitMessageIntervalFactor))/1000
	log.Printf("First nextTry: %v\n", nextTry)
	var timeout = (start + int64(limit))/1000
	log.Printf("First timeout: %v\n", timeout)
	old_fib := 0
	fib := 1

	for true {
		anchor, err := a.anchorRepository.GetAnchor(anchorId)
		log.Printf("Anchor waitting: %+v\n", anchor)
		if err != nil {
			return entity.Anchor{}, err
		}
		if anchor.Status() == "Success" {
			return anchor, nil
		}
		currentTime := time.Now().Unix()
		log.Printf("Second time: %v\n", currentTime)

		for currentTime < nextTry && currentTime < timeout {
			time.Sleep(time.Millisecond * 200)
			currentTime = time.Now().Unix()
			log.Println("I'm in loop")
		}
		if currentTime > timeout {
			return entity.Anchor{}, exception2.NewWaitAnchorTimeoutException()
		}

		nextTry += int64(old_fib + fib)
		aux := old_fib
		old_fib = fib
		fib = old_fib + aux
	}*/






	return entity.Anchor{}, nil
}
