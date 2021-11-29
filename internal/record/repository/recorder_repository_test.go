package repository

import (
	"encoding/json"
	"github.com/bloock/bloock-sdk-go/internal/config/mockconfig"
	"github.com/bloock/bloock-sdk-go/internal/infrastructure/http/mockhttp"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendRecordsRepository(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	cs := mockconfig.NewMockConfigurerService(crtl)
	hc := mockhttp.NewMockHttpClient(crtl)
	rr := NewRecordRepository(hc, cs)

	t.Run("Given a valid record request, should send with success", func(t *testing.T) {
		resp := dto.RecordWriteResponse{
			Anchor:   80,
			Client:   "ce10c769-022b-405e-8e7c-3b52eeb2a4ea",
			Messages: []string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
			Status:   "Pending",
		}
		respByte, err := json.Marshal(resp)
		require.Nil(t, err)

		cs.EXPECT().GetApiBaseUrl().Return("api").Times(1)
		hc.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(respByte, nil).Times(1)

		actual, err := rr.SendRecords([]entity.RecordEntity{})
		assert.Nil(t, err)
		assert.IsType(t, dto.RecordWriteResponse{}, actual)
		assert.Equal(t, 80, actual.Anchor)
		assert.Equal(t, "ce10c769-022b-405e-8e7c-3b52eeb2a4ea", actual.Client)
		assert.Equal(t, []string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"}, actual.Messages)
		assert.Equal(t, "Pending", actual.Status)
	})

	t.Run("Given an empty record, should return empty fields", func(t *testing.T) {
		resp := dto.RecordWriteResponse{}
		respByte, err := json.Marshal(resp)
		require.Nil(t, err)

		cs.EXPECT().GetApiBaseUrl().Return("api").Times(1)
		hc.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(respByte, nil).Times(1)

		actual, err := rr.SendRecords([]entity.RecordEntity{})
		assert.Nil(t, err)
		assert.IsType(t, dto.RecordWriteResponse{}, actual)
		assert.Equal(t, 0, actual.Anchor)
		assert.Equal(t, "", actual.Client)
		var emptyString []string
		assert.Equal(t, emptyString, actual.Messages)
		assert.Equal(t, "", actual.Status)

	})
}

func TestFetchRecordsRepository(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	cs := mockconfig.NewMockConfigurerService(crtl)
	hc := mockhttp.NewMockHttpClient(crtl)
	rr := NewRecordRepository(hc, cs)

	t.Run("Given a valid record request, should get records with success", func(t *testing.T) {
		resp := []dto.RecordRetrieveResponse{{
			Anchor:  80,
			Client:  "ce10c769-022b-405e-8e7c-3b52eeb2a4ea",
			Message: "02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5",
			Status:  "Pending",
		},
		}
		respByte, err := json.Marshal(resp)
		require.Nil(t, err)

		cs.EXPECT().GetApiBaseUrl().Return("api").Times(1)
		hc.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(respByte, nil).Times(1)

		actual, err := rr.FetchRecords([]entity.RecordEntity{})
		assert.Nil(t, err)
		assert.IsType(t, dto.RecordRetrieveResponse{}, actual[0])
		assert.Equal(t, 80, actual[0].Anchor)
		assert.Equal(t, "ce10c769-022b-405e-8e7c-3b52eeb2a4ea", actual[0].Client)
		assert.Equal(t, "02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5", actual[0].Message)
		assert.Equal(t, "Pending", actual[0].Status)
	})

	t.Run("Given an empty record retrieve, should return empty fields", func(t *testing.T) {
		resp := make([]dto.RecordRetrieveResponse, 0)
		respByte, err := json.Marshal(resp)
		require.Nil(t, err)

		cs.EXPECT().GetApiBaseUrl().Return("api").Times(1)
		hc.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(respByte, nil).Times(1)

		actual, err := rr.FetchRecords([]entity.RecordEntity{})
		assert.Nil(t, err)
		assert.Equal(t, []dto.RecordRetrieveResponse{}, actual)
	})
}
