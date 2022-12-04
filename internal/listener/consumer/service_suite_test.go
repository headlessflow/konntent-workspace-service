package consumer_test

//import (
//	"context"
//	"testing"
//	"konntent-workspace-service/internal/listener/consumer"
//	mcm "konntent-workspace-service/pkg/dummyclient/mocks"
//
//	"github.com/golang/mock/gomock"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//	"github.com/sirupsen/logrus"
//	"github.com/sirupsen/logrus/hooks/test"
//)
//
//func TestConsumerService(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "ConsumerService Suite")
//}
//
//var (
//	ctx                 context.Context
//	logger              *logrus.Logger
//	mockCtrl            *gomock.Controller
//	mobilisimClientMock *mcm.MockClient
//	service             consumer.Service
//)
//
//var _ = BeforeEach(func() {
//	ctx = context.Background()
//	logger, _ = test.NewNullLogger()
//	mockCtrl = gomock.NewController(GinkgoT())
//	mobilisimClientMock = mcm.NewMockClient(mockCtrl)
//	service = consumer.NewMobilisimConsumerService(logger, mobilisimClientMock)
//})
//
//var _ = AfterEach(func() {
//	mockCtrl.Finish()
//})
