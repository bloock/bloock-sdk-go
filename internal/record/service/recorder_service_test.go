package service

import (
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity/dto"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity/exception"
	"github.com/enchainte/enchainte-sdk-go/internal/record/mockrecord"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSendRecordsService(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	rr := mockrecord.NewMockRecorderRepository(crtl)
	rs := NewRecordService(rr)

	t.Run("Given a valid record request, should send with success", func(t *testing.T) {
		resp := dto.RecordWriteResponse{
			Anchor: 80,
			Client: "ce10c769-022b-405e-8e7c-3b52eeb2a4ea",
			Messages: []string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
			Status: "Pending",
		}
		arr := make([]entity.RecordEntity, 0)
		var rec entity.RecordEntity
		rec = rec.FromHash("02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5")
		arr = append(arr, rec)

		rr.EXPECT().SendRecords(gomock.Any()).Return(resp, nil).Times(1)

		actual, err := rs.SendRecords(arr)
		assert.Nil(t, err)
		assert.IsType(t, reflect.Array, reflect.TypeOf(actual).Kind())
		assert.IsType(t, entity.RecordReceipt{}, actual[0])
		assert.Equal(t, 80, actual[0].Anchor)
		assert.Equal(t, "ce10c769-022b-405e-8e7c-3b52eeb2a4ea", actual[0].Client)
		assert.Equal(t, "02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5", actual[0].Record)
		assert.Equal(t, "Pending", actual[0].Status)
	})
	
	t.Run("Given an invalid record, should return an error", func(t *testing.T) {
		arr := make([]entity.RecordEntity, 0)
		var rec entity.RecordEntity
		rec = rec.FromHash("record")
		arr = append(arr, rec)

		actual, err := rs.SendRecords(arr)
		assert.NotNil(t, err)
		assert.Equal(t, []entity.RecordReceipt{}, actual)
		assert.Equal(t, exception.InvalidRecordException{}.Error(), err.Error())
	})
}

func TestGetRecordsService(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	rr := mockrecord.NewMockRecorderRepository(crtl)
	rs := NewRecordService(rr)

	t.Run("Given a valid record request, should get records with success", func(t *testing.T) {
		resp := []dto.RecordRetrieveResponse{{
			Anchor: 80,
			Client: "ce10c769-022b-405e-8e7c-3b52eeb2a4ea",
			Message: "02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5",
			Status: "Pending",
		},
		}
		arr := make([]entity.RecordEntity, 0)
		var rec entity.RecordEntity
		rec = rec.FromString("record")
		arr = append(arr, rec)

		rr.EXPECT().FetchRecords(gomock.Any()).Return(resp, nil).Times(1)

		actual, err := rs.GetRecords(arr)
		assert.Nil(t, err)
		assert.IsType(t, entity.RecordReceipt{}, actual[0])
		assert.Equal(t, 80, actual[0].Anchor)
		assert.Equal(t, "ce10c769-022b-405e-8e7c-3b52eeb2a4ea", actual[0].Client)
		assert.Equal(t, "02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5", actual[0].Record)
		assert.Equal(t, "Pending", actual[0].Status)
	})
}
