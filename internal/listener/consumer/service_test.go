package consumer_test

//
//import (
//	"errors"
//	"konntent-workspace-service/pkg/constants"
//	"konntent-workspace-service/pkg/dummyclient/model"
//
//	"github.com/brianvoe/gofakeit/v6"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//)
//
//var (
//	mobilisimClientErr = errors.New("mobilisim client didn't respond successfully")
//)
//
//var _ = Describe("ConsumerService", func() {
//	var (
//		messageDestinationGenerator = func() model.MessageDestination {
//			return model.MessageDestination{
//				To: gofakeit.Phone(),
//			}
//		}
//		messaageGenerator = func() model.Message {
//			return model.Message{
//				From:         gofakeit.Word(),
//				Text:         gofakeit.LoremIpsumSentence(3),
//				CallbackData: gofakeit.Word(),
//				LanguageEncoding: gofakeit.RandomString([]string{
//					constants.MobilisimEnglishMessageDecoder,
//					constants.MobilisimUnicodeMessageDecoder,
//					constants.MobilisimTurkishMessageDecoder,
//				}),
//				ValidityPeriod: 2880,
//				Destinations: []model.MessageDestination{
//					messageDestinationGenerator(),
//				},
//			}
//		}
//	)
//
//	Describe("OneToN", func() {
//		var (
//			oneToNRequest model.RequestOneToN
//		)
//
//		BeforeEach(func() {
//			oneToNRequest = model.RequestOneToN{
//				Messages: []model.Message{
//					messaageGenerator(),
//					messaageGenerator(),
//					messaageGenerator(),
//				},
//			}
//		})
//
//		It("should return error one to n request when mobilisim client throws an error", func() {
//			mobilisimClientMock.EXPECT().OneToN(ctx, oneToNRequest).Times(1).Return(mobilisimClientErr)
//
//			err := service.OneToN(ctx, oneToNRequest)
//
//			Expect(err).To(Equal(mobilisimClientErr))
//		})
//
//		It("should return no error one to n request succeeds", func() {
//			mobilisimClientMock.EXPECT().OneToN(ctx, oneToNRequest).Times(1).Return(nil)
//
//			err := service.OneToN(ctx, oneToNRequest)
//
//			Expect(err).To(BeNil())
//		})
//	})
//
//	Describe("NToN", func() {
//		var (
//			nToNRequest *model.RequestNToN
//		)
//
//		BeforeEach(func() {
//			nToNRequest = &model.RequestNToN{
//				Messages: []model.Message{
//					messaageGenerator(),
//					messaageGenerator(),
//					messaageGenerator(),
//				},
//			}
//		})
//
//		It("should return error one to n request when mobilisim client throws an error", func() {
//			mobilisimClientMock.EXPECT().NToN(ctx, nToNRequest).Times(1).Return(mobilisimClientErr)
//
//			err := service.NToN(ctx, nToNRequest)
//
//			Expect(err).To(Equal(mobilisimClientErr))
//		})
//
//		It("should return no error one to n request succeeds", func() {
//			mobilisimClientMock.EXPECT().NToN(ctx, nToNRequest).Times(1).Return(nil)
//
//			err := service.NToN(ctx, nToNRequest)
//
//			Expect(err).To(BeNil())
//		})
//	})
//})
