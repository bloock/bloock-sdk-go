package mock
//
//import (
//	"errors"
//	"github.com/enchainte/enchainte-sdk-go/internal/message"
//)
//
//type ServiceMock struct {
//	receipts message.Receipts
//	err error
//}
//
//
//func NewServiceMock() ServiceMock {
//	return ServiceMock{}
//}
//
//func (s ServiceMock) Write(hash []byte) error {
//	if s.err != nil {
//		return s.err
//	}
//	return nil
//}
//
//func (s ServiceMock) Search(hash [][]byte) (*message.Receipts, error) {
//	if s.err != nil {
//		return nil, s.err
//	}
//	return &s.receipts, nil
//}
//
//func (s ServiceMock) Wait(hashes [][]byte) (*message.Receipts, error) {
//	if s.err != nil {
//		return nil, s.err
//	}
//	return &s.receipts, nil
//}
//
//func (s ServiceMock) Receipts(receipts... message.Receipts) {
//	var rs message.Receipts
//	for _, receipt := range receipts {
//		receipts = append(receipts, receipt)
//	}
//	s.receipts = rs
//}
//
//func (s ServiceMock) Error(msg string) {
//	s.err = errors.New(msg)
//}
