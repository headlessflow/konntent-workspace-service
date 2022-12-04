package orchestration_test

//
//import (
//	"errors"
//	"testing"
//	"konntent-workspace-service/internal/app/dto/request"
//	"konntent-workspace-service/pkg/claimer"
//	"konntent-workspace-service/pkg/constants"
//	"konntent-workspace-service/pkg/event"
//
//	"github.com/golang/mock/gomock"
//
//	"github.com/brianvoe/gofakeit/v6"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//)
//
//var (
//	serviceErr = errors.New("something went wrong while adding sms into queue")
//)
//
//var _ = Describe("MobilisimOrchestrator", func() {
//	Describe("OneToN", func() {
//		var (
//			oneToNRequest request.OneToN
//		)
//
//		BeforeEach(func() {
//			oneToNRequest = request.OneToN{
//				Message: gofakeit.LoremIpsumSentence(5),
//				MessageType: gofakeit.RandomString([]string{
//					constants.MobilisimEnglishMessageDecoder,
//					constants.MobilisimUnicodeMessageDecoder,
//					constants.MobilisimTurkishMessageDecoder,
//				}),
//				Sender: gofakeit.Word(),
//				Numbers: []string{
//					gofakeit.Phone(),
//					gofakeit.Phone(),
//					gofakeit.Phone(),
//				},
//			}
//
//			claimerMock.EXPECT().GetModel(ctx).AnyTimes().Return(owner)
//		})
//
//		It("should return error when send one to n event send throws an error", func() {
//			mobilisimEventServiceMock.
//				EXPECT().
//				OneToNEvent(ctx, oneToNRequest.ToEvent(owner, headerMap)).
//				Times(1).
//				Return(nil, serviceErr)
//
//			actual, err := orchestrator.OneToN(ctx, oneToNRequest)
//
//			Expect(err).To(Equal(serviceErr))
//			Expect(actual).To(BeNil())
//		})
//
//		It("should return one to n resource successfully", func() {
//			expectedResponse := &event.ResourceOneToNEvent{
//				Status:      constants.MobilisimSuccessStatus,
//				Description: constants.MobilisimSuccessDescription,
//				NumberCount: 3,
//				SMSQuantity: 3,
//			}
//
//			mobilisimEventServiceMock.
//				EXPECT().
//				OneToNEvent(ctx, oneToNRequest.ToEvent(owner, headerMap)).
//				Times(1).
//				Return(expectedResponse, nil)
//
//			actual, err := orchestrator.OneToN(ctx, oneToNRequest)
//
//			Expect(err).To(BeNil())
//			Expect(actual).To(Equal(expectedResponse))
//		})
//	})
//
//	Describe("NToN", func() {
//		var (
//			nToNRequest request.NToN
//		)
//
//		BeforeEach(func() {
//			destinationnGenerator := func() request.NToNPhone {
//				return request.NToNPhone{
//					Phone:   gofakeit.Phone(),
//					Message: gofakeit.LoremIpsumSentence(3),
//				}
//			}
//			nToNRequest = request.NToN{
//				MessageType: gofakeit.RandomString([]string{
//					constants.MobilisimEnglishMessageDecoder,
//					constants.MobilisimUnicodeMessageDecoder,
//					constants.MobilisimTurkishMessageDecoder,
//				}),
//				Sender: gofakeit.Word(),
//				Destinations: []request.NToNPhone{
//					destinationnGenerator(),
//					destinationnGenerator(),
//					destinationnGenerator(),
//				},
//			}
//
//			claimerMock.EXPECT().GetModel(ctx).AnyTimes().Return(owner)
//		})
//
//		It("should return error when send n to n event send throws an error", func() {
//			mobilisimEventServiceMock.
//				EXPECT().
//				NToNEvent(ctx, nToNRequest.ToEvent(owner, headerMap)).
//				Times(1).
//				Return(nil, serviceErr)
//
//			actual, err := orchestrator.NToN(ctx, &nToNRequest)
//
//			Expect(err).To(Equal(serviceErr))
//			Expect(actual).To(BeNil())
//		})
//
//		It("should return n to n resource successfully", func() {
//			expectedResponse := &event.ResourceNToNEvent{
//				Status:      constants.MobilisimSuccessStatus,
//				Description: constants.MobilisimSuccessDescription,
//				NumberCount: 3,
//				SMSQuantity: 3,
//			}
//
//			mobilisimEventServiceMock.
//				EXPECT().
//				NToNEvent(ctx, nToNRequest.ToEvent(owner, headerMap)).
//				Times(1).
//				Return(expectedResponse, nil)
//
//			actual, err := orchestrator.NToN(ctx, &nToNRequest)
//
//			Expect(err).To(BeNil())
//			Expect(actual).To(Equal(expectedResponse))
//		})
//
//		It("should return n to n resource successfully via WithRoutine function", func() {
//			const size = 10000
//			nToNRequest.Destinations = nil
//			expectedResponse := &event.ResourceNToNEvent{
//				Status:      constants.MobilisimSuccessStatus,
//				Description: constants.MobilisimSuccessDescription,
//				NumberCount: size,
//				SMSQuantity: size,
//			}
//
//			for i := 0; i < size; i++ {
//				nToNRequest.Destinations = append(nToNRequest.Destinations, request.NToNPhone{
//					Phone:   gofakeit.Phone(),
//					Message: gofakeit.Word(),
//				})
//			}
//
//			mobilisimEventServiceMock.
//				EXPECT().
//				NToNEvent(ctx, gomock.Any()).
//				Times(1).
//				Return(expectedResponse, nil)
//
//			actual, err := orchestrator.NToN(ctx, &nToNRequest)
//
//			Expect(err).To(BeNil())
//			Expect(actual).To(Equal(expectedResponse)) // inordinary equal ??
//		})
//	})
//})
//
//func BenchmarkA(b *testing.B) {
//	var nToNRequest = &request.NToN{
//		MessageType: gofakeit.RandomString([]string{
//			constants.MobilisimEnglishMessageDecoder,
//			constants.MobilisimUnicodeMessageDecoder,
//			constants.MobilisimTurkishMessageDecoder,
//		}),
//		Sender: gofakeit.Word(),
//	}
//	const size = 500000
//
//	for i := 0; i < size; i++ {
//		nToNRequest.Destinations = append(nToNRequest.Destinations, request.NToNPhone{
//			Phone:   "6136459948",
//			Message: "gofakeit.Word()",
//		})
//	}
//
//	b.SetParallelism(8)
//	b.ReportMetric(float64(b.N), "[SAMPLING]")
//	b.ReportAllocs()
//	b.ResetTimer()
//	b.StartTimer()
//
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			var (
//				msg     = make([]string, 0, size)
//				dest    = make([]event.Destination, 0, size)
//				routine = &request.Routine{
//					Msg:  msg,
//					Dest: dest,
//				}
//			)
//			nToNRequest.WithRoutine(routine)
//		}
//	})
//	b.StopTimer()
//}
//
//func BenchmarkB(b *testing.B) {
//	var nToNRequest = &request.NToN{
//		MessageType: gofakeit.RandomString([]string{
//			constants.MobilisimEnglishMessageDecoder,
//			constants.MobilisimUnicodeMessageDecoder,
//			constants.MobilisimTurkishMessageDecoder,
//		}),
//		Sender: gofakeit.Word(),
//	}
//	owner = &claimer.Model{
//		UserID: gofakeit.IntRange(1000, 9999),
//		Credit: gofakeit.IntRange(100, 150),
//	}
//	const size = 3
//
//	for i := 0; i < size; i++ {
//		nToNRequest.Destinations = append(nToNRequest.Destinations, request.NToNPhone{
//			Phone:   gofakeit.Phone(),
//			Message: gofakeit.Word(),
//		})
//	}
//
//	//b.SetParallelism(16)
//	//b.ReportMetric(float64(b.N), "[SAMPLING]")
//	//b.ReportAllocs()
//	//b.ResetTimer()
//	//b.StartTimer()
//	//
//	//b.RunParallel(func(pb *testing.PB) {
//	//	for pb.Next() {
//	//		nToNRequest.ToEvent(owner)
//	//	}
//	//})
//	//b.StopTimer()
//
//	nToNRequest.ToEvent(owner, headerMap)
//}
