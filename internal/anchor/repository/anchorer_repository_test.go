package repository

import (
	"encoding/json"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity/dto"
	"github.com/enchainte/enchainte-sdk-go/internal/config/mockconfig"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http/mockhttp"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAnchorRepository(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	cs := mockconfig.NewMockConfigurerService(crtl)
	hc := mockhttp.NewMockHttpClient(crtl)
	ar := NewAnchorRepository(hc, cs)

	t.Run("Given an existing anchor, should return that anchor", func(t *testing.T) {
		resp := dto.AnchorRetrieveResponse{
			AnchorId: 1,
			BlockRoots: []string{"block_root"},
			Networks: []entity.Network{},
			Root: "root",
			Status: "Success",
		}
		respByte, err := json.Marshal(resp)
		require.Nil(t, err)

		cs.EXPECT().GetApiBaseUrl().Return("i'm definitely a URL")

		hc.EXPECT().Get(gomock.Any(), gomock.Any()).Return(respByte, nil).Times(1)

		actual, err := ar.GetAnchor(1)
		assert.Nil(t, err)
		assert.IsType(t, entity.Anchor{}, actual)
		assert.Equal(t, 1, actual.ID())
		assert.Equal(t, []string{"block_root"}, actual.BlockRoots())
		assert.Equal(t, []entity.Network{}, actual.Networks())
		assert.Equal(t, "root", actual.Root())
		assert.Equal(t, "Success", actual.Status())
	})
}
