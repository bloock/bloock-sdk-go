package service

import (
	entity2 "github.com/enchainte/enchainte-sdk-go/config/entity"
	"github.com/enchainte/enchainte-sdk-go/config/mockconfig"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/mockanchor"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
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

	/*counter := 0
	maxCount := 0
	var getAnchorSideEffect func() (entity.Anchor, error)
	getAnchorSideEffect = func() (entity.Anchor, error) {
		log.Printf("Counter: %d, MaxCount: %v", counter, maxCount)
		if counter < maxCount {
			counter += 1
			return entity.Anchor{}, exception.NewHttpRequestException("Anchor not ready yet")
		}
		return entity.NewAnchor(1, []string{"block_root"}, []entity.Network{}, "root", "Success"), nil
	}*/

	t.Run("Given one try, should return the anchor", func(t *testing.T) {
		counter := 0
		maxCount := 3
		var getAnchorSideEffect func() (entity.Anchor, error)
		getAnchorSideEffect = func() (entity.Anchor, error) {
			log.Printf("Counter: %d, MaxCount: %v", counter, maxCount)
			if counter < maxCount {
				counter += 1
				return entity.Anchor{}, nil
			}
			return entity.NewAnchor(1, []string{"block_root"}, []entity.Network{}, "root", "Success"), nil
		}
		conf := entity2.NewConfiguration("api", 0, 1)
		cs.EXPECT().GetConfiguration().Return(conf)

		ar.EXPECT().GetAnchor(gomock.Any()).Return(getAnchorSideEffect())

		actual, err := as.WaitAnchor(1, 5000)

		assert.Nil(t, err)
		assert.IsType(t, entity.Anchor{}, actual)
		assert.Equal(t, 1, actual.ID())
		assert.Equal(t, []string{"block_root"}, actual.BlockRoots())
		assert.Equal(t, []entity.Network{}, actual.Networks())
		assert.Equal(t, "root", actual.Root())
		assert.Equal(t, "Success", actual.Status())
	})

}
