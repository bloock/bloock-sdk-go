package service

import (
	"github.com/enchainte/enchainte-sdk-go/config/mockconfig"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/mockanchor"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAnchorService(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	cs := mockconfig.NewMockConfigurerService(crtl)
	ar := mockanchor.NewMockAnchorerRepository(crtl)
	as := NewAnchorService(ar, cs)

	t.Run("Given an existing anchor, should return that anchor", func(t *testing.T) {
		bl := []string{"bloock_root"}
		var nt []entity.Network
		a := entity.NewAnchor(1, bl, nt, "root", "Success")

		ar.EXPECT().GetAnchor(gomock.Any()).Return(a, nil).Times(1)

		actual, err := as.GetAnchor(1)
		assert.Nil(t, err)
		assert.IsType(t, entity.Anchor{}, actual)
		assert.Equal(t, 1, actual.ID())
		assert.Equal(t, []string{"bloock_root"}, actual.BlockRoots())
		assert.Equal(t, nt, actual.Networks())
		assert.Equal(t, "root", actual.Root())
		assert.Equal(t, "Success", actual.Status())
	})
}