//go:build integration
// +build integration

package handler_test

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	di "konntent-workspace-service"
	"konntent-workspace-service/pkg/middlewarepkg"
	nrcmock "konntent-workspace-service/pkg/nrclient/mocks"
	pgimock "konntent-workspace-service/pkg/pg/mocks"
	"konntent-workspace-service/pkg/utils"
	"konntent-workspace-service/pkg/validation"
	"testing"
)

//
//import (
//	"bytes"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io"
//	di "konntent-workspace-service"
//	appctx "konntent-workspace-service/internal/app"
//	"konntent-workspace-service/pkg/claimer"
//	jwtmock "konntent-workspace-service/pkg/claimer/mocks"
//	"konntent-workspace-service/pkg/constants"
//	"konntent-workspace-service/pkg/middlewarepkg"
//	"konntent-workspace-service/pkg/nrclient"
//	rabbitmock "konntent-workspace-service/pkg/rabbit/mocks"
//	"konntent-workspace-service/pkg/utils"
//	"konntent-workspace-service/pkg/validation"
//
//	"github.com/gofiber/fiber/v2"
//	"github.com/golang/mock/gomock"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"konntent-workspace-service/pkg/dummyclient"
//
//	"github.com/sirupsen/logrus/hooks/test"
//)
//
var (
	server *fiber.App
)

var (
	mockCtrl *gomock.Controller
	pgMock   *pgimock.MockInstance
	nrMock   *nrcmock.MockNewRelicInstance

	errEvent = errors.New("something went wrong")
)

//const (
//	mobilisimURL = "http://mobilisim.com/"
//
//	contentTypeHeaderKey = "Content-Type"
//	jsonHeaderValue      = "application/json"
//
//	authorizationHeaderKey       = "X-Authorization"
//	authorizationMobilisimApiKey = "X-Mobilisim-Key"
//	bearer                       = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNTI0LCJjcmVkaXQiOjExMH0.aVSN0EShTyVnzqFAiI5Vf1XPixZGnrkvPXxPbNupSLo"
//	mobilisimBearer              = "1234-5678-9123-0012"
//)

func TestHandlerIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Integration Suite")
}

//
//var _ = BeforeEach(func() {
//	jwtInstance.EXPECT().GetModel(gomock.Any()).AnyTimes().Return(&jwtModel)
//})
//
//var _ = BeforeSuite(func() {
//	loadMockDependencies()
//
//	initServer()
//
//	go func() {
//		if serveErr := server.Listen(fmt.Sprintf(":%s", constants.TestPort)); serveErr != nil {
//			Expect(serveErr).NotTo(HaveOccurred())
//		}
//	}()
//})
//
var _ = AfterSuite(func() {
	defer func() {
		_ = server.Shutdown()
	}()
})

func loadMockDependencies() {
	mockCtrl = gomock.NewController(GinkgoT())
	pgMock = pgimock.NewMockInstance(mockCtrl)
}

func initServer() {
	var logger = zap.L()

	server = fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	validator := validation.InitValidator()

	server.Use(func(c *fiber.Ctx) error {
		c.Locals(utils.Claimer, s.jwtInstance)
		return c.Next()
	})
	server.Use(func(c *fiber.Ctx) error {
		c.Locals(utils.Validator, validator)
		return c.Next()
	})
	server.Use(middlewarepkg.PutHeaders)

	route := di.InitAll(
		logger,
		pgMock,
		nrInstance,
	)

	route.SetupRoutes(&appctx.RouteCtx{
		App: server,
	})
}

//
//func prepareRequestWithToken(method, url string, data []byte) *http.Request {
//	var body io.Reader
//	if data != nil {
//		body = bytes.NewBuffer(data)
//	}
//
//	req := httptest.NewRequest(method, url, body)
//	req.Header.Add(contentTypeHeaderKey, jsonHeaderValue)
//	req.Header.Add(authorizationHeaderKey, bearer)
//	req.Header.Add(authorizationMobilisimApiKey, mobilisimBearer)
//
//	return req
//}
//
//func sendTestRequest(req *http.Request) (*http.Response, []byte) {
//	resp, _ := server.Test(req)
//	defer resp.Body.Close()
//
//	body, _ := io.ReadAll(resp.Body)
//	return resp, body
//}
//
////func mockFailedRequest(_req httpclient.Request) {
////	httpClientMock.EXPECT().
////		HandleRequest(ctx, _req).Times(1).Return(nil, errRequestTimeout)
////	httpClientMock.EXPECT().GetJSONHeaders().Times(1).Return(map[string]string{})
////}
//
////func mockHTTPSuccessfulRequest(_req httpclient.Request, _res *httpclient.Response) {
////	httpClientMock.EXPECT().HandleRequest(ctx, _req).Times(1).Return(_res, nil)
////	httpClientMock.EXPECT().GetJSONHeaders().Times(1).Return(map[string]string{})
////	httpClientMock.EXPECT().IsSuccessStatusCode(_res).Times(1).Return(true)
////}
//
//func mockFailedEvent() {
//	messagingClient.EXPECT().PublishOnQueue(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(errEvent)
//}
//
//func mockSuccessfulEvent() {
//	messagingClient.EXPECT().PublishOnQueue(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
//}
