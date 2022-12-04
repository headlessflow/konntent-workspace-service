// //go:build integration && !benchmark
// // +build integration,!benchmark
package handler_test

//
//import (
//	"net/http"
//
//	"github.com/gofiber/fiber/v2"
//	"github.com/golang/mock/gomock"
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//)
//
//const (
//	sendOneToNSMS = "/v1/mobilisim/oneToN"
//	sendNToNSMS   = "/v1/mobilisim/nToN"
//)
//
//var _ = Describe("MobilisimIntegration", func() {
//	Describe("POST /v1/mobilisim/oneToN-nToN Unauthorized", func() {
//		It("should return send unauthorized status code when you don't have correct signed jwt", func() {
//			jwtInstance.EXPECT().IsValid(gomock.Any(), gomock.Any()).Times(1).Return(nil, false)
//
//			req := prepareRequestWithToken(fiber.MethodPost, sendOneToNSMS, []byte(`{
//				"message": "12345",
//				"message_type": "turkish",
//				"sender": "VATANSMS",
//				"numbers": [
//					"905442666417"
//				]
//			}`))
//
//			actual, _ := sendTestRequest(req)
//
//			Expect(actual.StatusCode).To(Equal(http.StatusUnauthorized))
//		})
//	})
//
//	Describe("POST /v1/mobilisim/oneToN", func() {
//		//var (
//		//	url string
//		//)
//		//
//		//BeforeEach(func() {
//		//	url = "%s/sms/1/text/advanced"
//		//})
//		//
//		//prepareExternalRequest := func(body interface{}) httpclient.Request {
//		//	return httpclient.Request{
//		//		URL:    fmt.Sprintf(url, mobilisimURL),
//		//		Method: fiber.MethodPost,
//		//		Body:   body,
//		//		Headers: map[string]string{
//		//			constants.HeaderKeyContentType:   "application/json",
//		//			constants.HeaderKeyAuthorization: "Bearer " + mobilisimBearer,
//		//		},
//		//	}
//		//}
//
//		BeforeEach(func() {
//			jwtInstance.EXPECT().IsValid(gomock.Any(), gomock.Any()).Times(1).Return(jwtModelBytes, true)
//		})
//
//		It("should return send one to n message error when request is not parsed", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendOneToNSMS, []byte(`{
//				"message": 12345,
//				"message_type": "turkish",
//				"sender": "VATANSMS",
//				"numbers": [
//					"905442666417"
//				]
//			}`))
//
//			actual, _ := sendTestRequest(req)
//
//			Expect(actual.StatusCode).To(Equal(http.StatusBadRequest))
//		})
//
//		It("should return send one to n message error when request is not validated", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendOneToNSMS, []byte(`{
//				"message_type": "turkish",
//				"sender": "VATANSMS",
//				"numbers": [
//					"905442666417"
//				]
//			}`))
//
//			actual, _ := sendTestRequest(req)
//
//			Expect(actual.StatusCode).To(Equal(http.StatusUnprocessableEntity))
//		})
//
//		It("should return send one to n error when mobilisim event throws an error", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendOneToNSMS, []byte(`{
//				"message": "my message",
//				"message_type": "turkish",
//				"sender": "VATANSMS",
//				"numbers": [
//					"905442666417"
//				]
//			}`))
//
//			mockFailedEvent()
//
//			actual, _ := sendTestRequest(req)
//			Expect(actual.StatusCode).To(Equal(http.StatusInternalServerError))
//		})
//
//		It("should return send one to n response", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendOneToNSMS, []byte(`{
//				"message": "my message",
//				"message_type": "turkish",
//				"sender": "VATANSMS",
//				"numbers": [
//					"905442666417"
//				]
//			}`))
//			expected := `{
//			  "data": {
//				"status": "success",
//				"description": "SMS gönderiminiz başarıyla başlatıldı.",
//				"numberCount": 1,
//				"quantity": 1
//			  }
//			}`
//
//			mockSuccessfulEvent()
//
//			actual, body := sendTestRequest(req)
//			Expect(actual.StatusCode).To(Equal(http.StatusOK))
//			Expect(body).To(MatchJSON(expected))
//		})
//	})
//
//	Describe("POST /v1/mobilisim/nToN", func() {
//		BeforeEach(func() {
//			jwtInstance.EXPECT().IsValid(gomock.Any(), gomock.Any()).Times(1).Return(jwtModelBytes, true)
//		})
//
//		It("should return send n to n message error when request is not parsed", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendNToNSMS, []byte(`{
//			  "message_type": "turkish",
//			  "sender": "VATANSMS",
//			  "destinations": [
//				{
//				  "phone": 905552554411,
//				  "message": "example message"
//				}
//			  ]
//			}`))
//
//			actual, _ := sendTestRequest(req)
//
//			Expect(actual.StatusCode).To(Equal(http.StatusBadRequest))
//		})
//
//		It("should return send n to n message error when request is not validated", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendNToNSMS, []byte(`{
//			  "message_type": "turkish",
//			  "sender": "VATANSMS",
//			  "destinations": [
//				{
//				  "message": "example message"
//				}
//			  ]
//			}`))
//
//			actual, _ := sendTestRequest(req)
//
//			Expect(actual.StatusCode).To(Equal(http.StatusUnprocessableEntity))
//		})
//
//		It("should return send n to n error when mobilisim event throws an error", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendNToNSMS, []byte(`{
//			  "message_type": "turkish",
//			  "sender": "VATANSMS",
//			  "destinations": [
//				{
//				  "phone": "905552221133",
//				  "message": "example message"
//				}
//			  ]
//			}`))
//
//			mockFailedEvent()
//
//			actual, _ := sendTestRequest(req)
//			Expect(actual.StatusCode).To(Equal(http.StatusInternalServerError))
//		})
//
//		It("should return send n to n response", func() {
//			req := prepareRequestWithToken(fiber.MethodPost, sendNToNSMS, []byte(`{
//			  "message_type": "turkish",
//			  "sender": "VATANSMS",
//			  "destinations": [
//				{
//				  "phone": "905552221133",
//				  "message": "example message"
//				}
//			  ]
//			}`))
//			expected := `{
//			  "data": {
//				"status": "success",
//				"description": "SMS gönderiminiz başarıyla başlatıldı.",
//				"numberCount": 1,
//				"quantity": 1
//			  }
//			}`
//
//			mockSuccessfulEvent()
//
//			actual, body := sendTestRequest(req)
//			Expect(actual.StatusCode).To(Equal(http.StatusOK))
//			Expect(body).To(MatchJSON(expected))
//		})
//	})
//})
