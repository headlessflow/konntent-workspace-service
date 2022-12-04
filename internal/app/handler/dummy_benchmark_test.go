// //go:build integration && benchmark
// // +build integration,benchmark
package handler_test

//
//import (
//	"encoding/json"
//	"konntent-workspace-service/internal/app/dto/request"
//	"time"
//
//	"github.com/brianvoe/gofakeit/v6"
//	"github.com/gofiber/fiber/v2"
//	"github.com/golang/mock/gomock"
//	. "github.com/onsi/ginkgo/v2"
//	"github.com/onsi/gomega/gmeasure"
//)
//
//const (
//	sendOneToNSMS = "/v1/mobilisim/oneToN"
//	sendNToNSMS   = "/v1/mobilisim/nToN"
//)
//
//var _ = Describe("[Benchmark] MobilisimIntegration", func() {
//	Describe("POST /v1/mobilisim/oneToN", func() {
//		It("should give benchmark reports", Serial, Label("measurement"), func() {
//			experiment := gmeasure.NewExperiment("[Experiment]POST /v1/mobilisim/oneToN")
//			AddReportEntry(experiment.Name, experiment)
//
//			var internalReq = request.OneToN{
//				Message:     gofakeit.Word(),
//				MessageType: "turkish",
//				Sender:      gofakeit.Word(),
//			}
//
//			for i := 0; i < 1000; i++ {
//				internalReq.Numbers = append(internalReq.Numbers, gofakeit.Phone())
//			}
//
//			reqAsBytes, _ := json.Marshal(internalReq)
//
//			jwtInstance.EXPECT().IsValid(gomock.Any(), gomock.Any()).AnyTimes().Return(jwtModelBytes, true)
//			messagingClient.EXPECT().PublishOnQueue(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
//
//			experiment.Sample(func(idx int) {
//				//mockSuccessfulEvent()
//
//				req := prepareRequestWithToken(fiber.MethodPost, sendOneToNSMS, reqAsBytes)
//
//				experiment.MeasureDuration("Push OneToN message into queue", func() {
//					_, _ = sendTestRequest(req)
//				})
//			}, gmeasure.SamplingConfig{N: 10 /*NumParallel: runtime.GOMAXPROCS(runtime.NumCPU()),*/, Duration: time.Minute})
//		})
//	})
//
//	Describe("POST /v1/mobilisim/nToN", func() {
//		It("should give benchmark reports", Serial, Label("measurement"), func() {
//			experiment := gmeasure.NewExperiment("[Experiment]POST /v1/mobilisim/nToN")
//			AddReportEntry(experiment.Name, experiment)
//
//			var internalReq = request.NToN{
//				MessageType: "turkish",
//				Sender:      gofakeit.Word(),
//			}
//
//			for i := 0; i < 10000; i++ {
//				internalReq.Destinations = append(internalReq.Destinations, request.NToNPhone{
//					Phone:   gofakeit.Phone(),
//					Message: gofakeit.Word(),
//				})
//			}
//
//			reqAsBytes, _ := json.Marshal(internalReq)
//
//			experiment.Sample(func(idx int) {
//				//mockSuccessfulEvent()
//				jwtInstance.EXPECT().IsValid(gomock.Any(), gomock.Any()).Times(1).Return(jwtModelBytes, true)
//				messagingClient.EXPECT().PublishOnQueue(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
//
//				req := prepareRequestWithToken(fiber.MethodPost, sendNToNSMS, reqAsBytes)
//
//				experiment.MeasureDuration("Push NToN message into queue", func() {
//					_, _ = sendTestRequest(req)
//				})
//			}, gmeasure.SamplingConfig{N: 10 /*NumParallel: runtime.GOMAXPROCS(runtime.NumCPU()),*/, Duration: time.Second * 45})
//		})
//	})
//})
