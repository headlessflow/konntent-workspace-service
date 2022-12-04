package dummy_test

//
//import (
//	"context"
//	"konntent-workspace-service/internal/app/mobilisim"
//	rcm "konntent-workspace-service/pkg/rabbit/mocks"
//	"testing"
//
//	"github.com/golang/mock/gomock"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//	"github.com/sirupsen/logrus"
//	"github.com/sirupsen/logrus/hooks/test"
//)
//
//func TestMobilisimService(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "MobilisimService Suite")
//}
//
//var (
//	ctx              context.Context
//	logger           *logrus.Logger
//	mockCtrl         *gomock.Controller
//	rabbitClientMock *rcm.MockClient
//	service          mobilisim.Service
//)
//
//var _ = BeforeEach(func() {
//	ctx = context.Background()
//	logger, _ = test.NewNullLogger()
//	mockCtrl = gomock.NewController(GinkgoT())
//	rabbitClientMock = rcm.NewMockClient(mockCtrl)
//	service = mobilisim.NewMobilisimService(logger, rabbitClientMock)
//})
//
//var _ = AfterEach(func() {
//	mockCtrl.Finish()
//})
