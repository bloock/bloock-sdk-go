package service

import (
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity/exception"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/mockanchor"
	configEntity "github.com/enchainte/enchainte-sdk-go/internal/config/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/config/mockconfig"
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

func TestWaitAnchorService(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	cs := mockconfig.NewMockConfigurerService(crtl)
	ar := mockanchor.NewMockAnchorerRepository(crtl)
	as := NewAnchorService(ar, cs)

	counter := 0
	maxCount := 3
	var getAnchorSideEffect func(args ...interface{}) (entity.Anchor, error)
	getAnchorSideEffect = func(args ...interface{}) (entity.Anchor, error) {
		if counter < maxCount {
			counter += 1
			return entity.Anchor{}, nil
		}
		return entity.NewAnchor(1, []string{"block_root"}, []entity.Network{}, "root", "Success"), nil
	}

	t.Run("Given one try, should wait and return the anchor", func(t *testing.T) {
		counter = 0
		maxCount = 0

		conf := configEntity.NewConfiguration("api", 0, 1)
		cs.EXPECT().GetConfiguration().Return(conf).Times(2)

		ar.EXPECT().GetAnchor(gomock.Any()).DoAndReturn(getAnchorSideEffect).Times(maxCount + 1)

		actual, err := as.WaitAnchor(1, 5000)

		assert.Nil(t, err)
		assert.IsType(t, entity.Anchor{}, actual)
		assert.Equal(t, 1, actual.ID())
		assert.Equal(t, []string{"block_root"}, actual.BlockRoots())
		assert.Equal(t, []entity.Network{}, actual.Networks())
		assert.Equal(t, "root", actual.Root())
		assert.Equal(t, "Success", actual.Status())
	})

	t.Run("Given three tries, should wait and return the anchor", func(t *testing.T) {
		counter = 0
		maxCount = 3

		conf := configEntity.NewConfiguration("api", 0, 1)
		cs.EXPECT().GetConfiguration().Return(conf).Times(2)

		ar.EXPECT().GetAnchor(gomock.Any()).DoAndReturn(getAnchorSideEffect).Times(maxCount + 1)

		actual, err := as.WaitAnchor(1, 5000)

		assert.Nil(t, err)
		assert.IsType(t, entity.Anchor{}, actual)
		assert.Equal(t, 1, actual.ID())
		assert.Equal(t, []string{"block_root"}, actual.BlockRoots())
		assert.Equal(t, []entity.Network{}, actual.Networks())
		assert.Equal(t, "root", actual.Root())
		assert.Equal(t, "Success", actual.Status())
	})

	t.Run("Given no time to wait the anchor, should return timeout error", func(t *testing.T) {
		counter = 0
		maxCount = 3

		conf := configEntity.NewConfiguration("api", 0, 10)
		cs.EXPECT().GetConfiguration().Return(conf).Times(2)

		ar.EXPECT().GetAnchor(gomock.Any()).DoAndReturn(getAnchorSideEffect)

		_, err := as.WaitAnchor(1, 1)

		assert.Equal(t, exception.NewWaitAnchorTimeoutException().Error(), err.Error())
	})

}
