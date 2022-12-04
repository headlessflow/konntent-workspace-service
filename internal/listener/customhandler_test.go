//go:build integration
// +build integration

package listener_test

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"sync"
//	"konntent-workspace-service/internal/listener"
//	consumerservice "konntent-workspace-service/internal/listener/consumer"
//	"konntent-workspace-service/pkg/constants"
//	"konntent-workspace-service/pkg/event"
//	"konntent-workspace-service/pkg/event/schema"
//	"konntent-workspace-service/pkg/eventmanager"
//	"konntent-workspace-service/pkg/rabbit"
//	"konntent-workspace-service/pkg/rabbit/utils"
//
//	"github.com/brianvoe/gofakeit/v6"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//	"github.com/sirupsen/logrus/hooks/test"
//	"github.com/streadway/amqp"
//)
//
//type x struct {
//	consumerManager rabbit.ConsumerManager
//	maxRetries      int
//}
//
//func (z *x) ConsumeClaim(ctx context.Context, queue <-chan amqp.Delivery) error {
//	errCh := make(chan error)
//	for {
//		select {
//		case delivery := <-queue:
//			deathCount := utils.GetXDeathCount(delivery.Headers)
//			if deathCount >= z.maxRetries-1 {
//				delivery.Headers[acknowledgeStateKey] = acknowledgeStateAck
//				return nil
//			}
//
//			go func() {
//				errCh <- z.consumerManager.Process(ctx, delivery)
//			}()
//			select {
//			case err := <-errCh:
//				if err != nil {
//					delivery.Headers[acknowledgeStateKey] = acknowledgeStateReject
//					return err
//				}
//			}
//
//		default:
//			return nil
//		}
//	}
//}
//func (z *x) Close() {}
//func (z *x) Status() chan bool {
//	ch := make(chan bool)
//	return ch
//}
//
//var (
//	unexpectedEventTypeErr = schema.MobilisimUnexpectedEventType
//
//	acknowledgeStateKey    = "acknowledge-state"
//	acknowledgeStateReject = "reject"
//	acknowledgeStateAck    = "ack"
//)
//
//var _ = Describe("OneToN ConsumerTest", func() {
//	Describe("[Dispatch]Event OneToN", func() {
//		It("should return unexpected type error when one to n event dispatches by unknown delivery", func() {
//			var (
//				wg                  = new(sync.WaitGroup)
//				logger, _           = test.NewNullLogger()
//				consumerService     = consumerservice.NewMobilisimConsumerService(logger, mobilisimClientMock)
//				eventHandlerFactory = listener.NewEventHandlerFactory(consumerService)
//				eventManager        = eventmanager.NewEventManager(eventHandlerFactory, event.NewEventCreator())
//				customHandler       = listener.NewCustomHandler(logger, eventManager)
//				consumerManager     = rabbit.NewConsumerManager(logger, customHandler)
//				handler             = &x{consumerManager: consumerManager, maxRetries: 3}
//				consumerInstance    = rabbit.NewConsumerInstance(logger, rabbitClientMock, handler)
//			)
//
//			var (
//				unexpectedEventType = ""
//				errCh               = make(chan error)
//				deliveries          = make(chan amqp.Delivery, 1)
//			)
//
//			deliveries <- amqp.Delivery{
//				Headers: amqp.Table{},
//			}
//
//			rabbitClientMock.EXPECT().Consume(ctx, handler).AnyTimes().
//				DoAndReturn(func(ctx context.Context, cg rabbit.ConsumerGroupHandler) error {
//					for {
//						errCh <- cg.ConsumeClaim(ctx, deliveries)
//						select {
//						case err := <-errCh:
//							if err != nil {
//								return err
//							}
//						}
//						break
//					}
//					return nil
//				})
//
//			wg.Add(1)
//			go consumerInstance.Consume(ctx)
//
//			wg.Done()
//			wg.Wait()
//
//			actualErr := <-errCh
//			expectedErr := fmt.Errorf("event type: %s - err: %v", unexpectedEventType, unexpectedEventTypeErr)
//
//			Expect(actualErr).To(Equal(expectedErr))
//		})
//
//		It("should call Reject function and writes the information into header map", func() {
//			var (
//				logger, _           = test.NewNullLogger()
//				consumerService     = consumerservice.NewMobilisimConsumerService(logger, mobilisimClientMock)
//				eventHandlerFactory = listener.NewEventHandlerFactory(consumerService)
//				eventManager        = eventmanager.NewEventManager(eventHandlerFactory, event.NewEventCreator())
//				customHandler       = listener.NewCustomHandler(logger, eventManager)
//				consumerManager     = rabbit.NewConsumerManager(logger, customHandler)
//				wg                  = new(sync.WaitGroup)
//				handler             = &x{consumerManager: consumerManager, maxRetries: 3}
//				consumerInstance    = rabbit.NewConsumerInstance(logger, rabbitClientMock, handler)
//			)
//
//			var (
//				errCh      = make(chan error)
//				deliveries = make(chan amqp.Delivery, 1)
//			)
//
//			delivery := amqp.Delivery{
//				Acknowledger: &amqp.Channel{},
//				Headers: amqp.Table{
//					utils.XDeathKeyName: []interface{}{
//						amqp.Table{
//							utils.XDeathCountKeyName: int64(1),
//						},
//					},
//				},
//			}
//			deliveries <- delivery
//
//			rabbitClientMock.EXPECT().Consume(ctx, handler).AnyTimes().
//				DoAndReturn(func(ctx context.Context, cg rabbit.ConsumerGroupHandler) error {
//					for {
//						errCh <- cg.ConsumeClaim(ctx, deliveries)
//						break
//					}
//					return nil
//				})
//
//			wg.Add(1)
//			go consumerInstance.Consume(ctx)
//
//			wg.Done()
//			wg.Wait()
//
//			<-errCh
//
//			Expect(delivery.Headers[acknowledgeStateKey]).To(Equal(acknowledgeStateReject))
//		})
//
//		It("should call Ack function when consumer tries to proceed same message for 3 times and writes the information into header map", func() {
//			var (
//				logger, _           = test.NewNullLogger()
//				consumerService     = consumerservice.NewMobilisimConsumerService(logger, mobilisimClientMock)
//				eventHandlerFactory = listener.NewEventHandlerFactory(consumerService)
//				eventManager        = eventmanager.NewEventManager(eventHandlerFactory, event.NewEventCreator())
//				customHandler       = listener.NewCustomHandler(logger, eventManager)
//				consumerManager     = rabbit.NewConsumerManager(logger, customHandler)
//				wg                  = new(sync.WaitGroup)
//				handler             = &x{consumerManager: consumerManager, maxRetries: 3}
//				consumerInstance    = rabbit.NewConsumerInstance(logger, rabbitClientMock, handler)
//			)
//
//			var (
//				errCh      = make(chan error)
//				deliveries = make(chan amqp.Delivery, 1)
//			)
//
//			event := event.OneToNEvent{
//				EventType: schema.MobilisimOneToNEventType,
//				EventData: event.OneToNEventData{
//					Message: gofakeit.Word(),
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
//					},
//				},
//				EventOwner: event.Owner{
//					UserID: gofakeit.IntRange(1000, 9999),
//					Credit: gofakeit.IntRange(100, 150),
//				},
//			}
//			eventBytes, _ := json.Marshal(event)
//
//			delivery := amqp.Delivery{
//				Acknowledger: &amqp.Channel{},
//				Headers: amqp.Table{
//					utils.XDeathKeyName: []interface{}{
//						amqp.Table{
//							utils.XDeathCountKeyName: int64(3),
//						},
//					},
//				},
//				Body: eventBytes,
//			}
//			deliveries <- delivery
//
//			rabbitClientMock.EXPECT().Consume(ctx, handler).AnyTimes().
//				DoAndReturn(func(ctx context.Context, cg rabbit.ConsumerGroupHandler) error {
//					for {
//						errCh <- cg.ConsumeClaim(ctx, deliveries)
//						break
//					}
//					return nil
//				})
//
//			wg.Add(1)
//			go consumerInstance.Consume(ctx)
//
//			wg.Done()
//			wg.Wait()
//
//			err := <-errCh
//
//			Expect(err).To(BeNil())
//			Expect(delivery.Headers[acknowledgeStateKey]).To(Equal(acknowledgeStateAck))
//		})
//	})
//})
