package proof_test
//
//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestProofNewProof(t *testing.T) {
//	tests := []struct {
//		leaves []string
//		nodes  []string
//		depth  string
//		bitmap string
//		err    error
//	}{
//		{
//			[]string{
//				"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//			},
//			[]string{
//				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
//				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
//				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
//				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//				"517e320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
//			},
//			"020304050501",
//			"f4",
//			nil,
//		},
//		{
//			[]string{
//				"xxaae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//			},
//			[]string{
//				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
//				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
//				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
//				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//				"517e320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
//			},
//			"020304050501",
//			"f4",
//			ErrInvalidLeave,
//		},
//		{
//			[]string{
//				"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97",
//			},
//			[]string{
//				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
//				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
//				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
//				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//				"517e320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
//			},
//			"020304050501",
//			"f4",
//			ErrInvalidLeave,
//		},
//		{
//			[]string{
//				"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//			},
//			[]string{
//				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
//				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
//				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
//				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//				"xyze320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
//			},
//			"020304050501",
//			"f4",
//			ErrInvalidNode,
//		},
//		{
//			[]string{
//				"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//			},
//			[]string{
//				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
//				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
//				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
//				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//				"517e320992fb35553575750153992d6360268d04a1e4d9",
//			},
//			"020304050501",
//			"f4",
//			ErrInvalidNode,
//		},
//		{
//			[]string{
//				"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//			},
//			[]string{
//				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
//				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
//				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
//				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//				"517e320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
//			},
//			"xyz304050501",
//			"f4",
//			ErrInvalidDepth,
//		},
//		{
//			[]string{
//				"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//			},
//			[]string{
//				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
//				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
//				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
//				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
//				"517e320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
//			},
//			"020304050501",
//			"z4",
//			ErrInvalidBitmap,
//		},
//	}
//	for _, test := range tests {
//		_, err := New(test.leaves, test.nodes, test.depth, test.bitmap)
//		assert.Equal(t, err, test.err)
//	}
//}
