package dummy_test

//
//import (
//	"errors"
//	"konntent-workspace-service/pkg/constants"
//	"konntent-workspace-service/pkg/event"
//	"konntent-workspace-service/pkg/event/schema"
//	"konntent-workspace-service/pkg/utils"
//	"time"
//
//	"github.com/brianvoe/gofakeit/v6"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//)
//
//var (
//	resourceStatusErr     = "error"
//	resourceStatusSuccess = "success"
//	queueErr              = errors.New("some queue not found")
//	insufficientCreditErr = errors.New("insufficient credit")
//)
//
//var _ = Describe("MobilisimService", func() {
//	Describe("OneToNEvent", func() {
//		var (
//			oneToNEvent *event.OneToNEvent
//
//			now time.Time
//			ttl int64
//		)
//
//		BeforeEach(func() {
//			oneToNEvent = &event.OneToNEvent{
//				EventType: schema.MobilisimOneToNEventType,
//				EventData: event.OneToNEventData{
//					Message: gofakeit.LoremIpsumWord(),
//					MessageType: gofakeit.RandomString([]string{
//						constants.MobilisimEnglishMessageDecoder,
//						constants.MobilisimUnicodeMessageDecoder,
//						constants.MobilisimTurkishMessageDecoder,
//					}),
//					Sender: gofakeit.Word(),
//					Numbers: []string{
//						gofakeit.Phone(),
//						gofakeit.Phone(),
//						gofakeit.Phone(),
//						gofakeit.Phone(),
//						gofakeit.Phone(),
//						gofakeit.Phone(),
//					},
//				},
//				EventOwner: event.Owner{
//					UserID: gofakeit.IntRange(1000, 9999),
//					Credit: gofakeit.IntRange(100, 150),
//				},
//			}
//
//			now, _ = time.Parse(constants.ParseLayout, time.Now().Format(constants.ParseLayout))
//			ttl = utils.PreCalculateTimeDiff(now, oneToNEvent.EventData.SendTime)
//		})
//
//		It("should return error when user sms credit insufficient", func() {
//			actual, err := service.OneToNEvent(ctx, &event.OneToNEvent{
//				EventType: oneToNEvent.EventType,
//				EventData: oneToNEvent.EventData,
//				EventOwner: event.Owner{
//					UserID: gofakeit.IntRange(1000, 9999),
//					Credit: gofakeit.IntRange(1, 3),
//				},
//			})
//
//			Expect(err).To(Equal(insufficientCreditErr))
//			Expect(actual).To(BeNil())
//		})
//
//		It("should return error when rabbit client cannot make all declarations", func() {
//			chunk, _ := utils.Chunk(oneToNEvent.EventData.Numbers, utils.DefaultChunkSize)
//
//			for _, messages := range chunk {
//				rabbitClientMock.
//					EXPECT().
//					PublishOnQueue(
//						oneToNEvent.ToPrepareQueue(messages),
//						oneToNEvent.EventType,
//						ttl,
//					).
//					Times(1).
//					Return(queueErr)
//			}
//
//			actual, err := service.OneToNEvent(ctx, oneToNEvent)
//
//			Expect(err).To(Equal(queueErr))
//			Expect(actual.Status).To(Equal(resourceStatusErr))
//		})
//
//		It("should return resource successfully", func() {
//			chunk, _ := utils.Chunk(oneToNEvent.EventData.Numbers, utils.DefaultChunkSize)
//
//			for _, messages := range chunk {
//				rabbitClientMock.
//					EXPECT().
//					PublishOnQueue(
//						oneToNEvent.ToPrepareQueue(messages),
//						oneToNEvent.EventType,
//						ttl,
//					).
//					Times(1).
//					Return(nil)
//			}
//
//			actual, err := service.OneToNEvent(ctx, oneToNEvent)
//
//			Expect(err).To(BeNil())
//			Expect(actual.Status).To(Equal(resourceStatusSuccess))
//		})
//	})
//
//	Describe("NToNEvent", func() {
//		var (
//			nToNEvent    *event.NToNEvent
//			destinations []event.Destination
//
//			now time.Time
//			ttl int64
//		)
//
//		BeforeEach(func() {
//			destinations = append(destinations, event.Destination{
//				Phone:   gofakeit.Phone(),
//				Message: gofakeit.Word(),
//			})
//
//			nToNEvent = &event.NToNEvent{
//				EventType: schema.MobilisimNToNEventType,
//				EventData: event.NToNEventData{
//					MessageType: gofakeit.RandomString([]string{
//						constants.MobilisimEnglishMessageDecoder,
//						constants.MobilisimUnicodeMessageDecoder,
//						constants.MobilisimTurkishMessageDecoder,
//					}),
//					Sender: gofakeit.Word(),
//					Messages: []string{
//						destinations[0].Message,
//					},
//					Phones: destinations,
//				},
//				EventOwner: event.Owner{
//					UserID: gofakeit.IntRange(1000, 9999),
//					Credit: gofakeit.IntRange(100, 150),
//				},
//			}
//
//			now, _ = time.Parse(constants.ParseLayout, time.Now().Format(constants.ParseLayout))
//			ttl = utils.PreCalculateTimeDiff(now, nToNEvent.EventData.SendTime)
//		})
//
//		It("should return error when user sms credit insufficient", func() {
//			actual, err := service.NToNEvent(ctx, &event.NToNEvent{
//				EventType: nToNEvent.EventType,
//				EventData: nToNEvent.EventData,
//				EventOwner: event.Owner{
//					UserID: gofakeit.IntRange(1000, 9999),
//					Credit: 0,
//				},
//			})
//
//			Expect(err).To(Equal(insufficientCreditErr))
//			Expect(actual).To(BeNil())
//		})
//
//		It("should return error when rabbit client cannot make all declarations", func() {
//			chunk, _ := utils.Chunk(nToNEvent.EventData.Phones, utils.DefaultChunkSize)
//
//			for _, messages := range chunk {
//				rabbitClientMock.
//					EXPECT().
//					PublishOnQueue(
//						nToNEvent.ToPrepareQueue(messages),
//						nToNEvent.EventType,
//						ttl,
//					).
//					AnyTimes().
//					Return(queueErr)
//			}
//
//			actual, err := service.NToNEvent(ctx, nToNEvent)
//
//			Expect(err).To(Equal(queueErr))
//			Expect(actual.Status).To(Equal(resourceStatusErr))
//		})
//
//		It("should return resource successfully", func() {
//			chunk, _ := utils.Chunk(nToNEvent.EventData.Phones, utils.DefaultChunkSize)
//
//			for _, messages := range chunk {
//				rabbitClientMock.
//					EXPECT().
//					PublishOnQueue(
//						nToNEvent.ToPrepareQueue(messages),
//						nToNEvent.EventType,
//						ttl,
//					).
//					Times(1).
//					Return(nil)
//			}
//
//			actual, err := service.NToNEvent(ctx, nToNEvent)
//
//			Expect(err).To(BeNil())
//			Expect(actual.Status).To(Equal(resourceStatusSuccess))
//		})
//	})
//})
