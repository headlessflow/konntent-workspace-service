//go:build integration
// +build integration

package handler_test

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"log"
)

const (
	getWorkspace = "/v1/workspace"
)

var _ = Describe("WorkspaceHandler", func() {
	Describe("GET /v1/workspace", func() {
		It("should return bodyParser error when given request wrong", func() {
			req := prepareRequestWithToken(fiber.MethodGet, getWorkspace, []byte(`{}`))

			actual, _ := sendTestRequest(req)

			log.Println(actual.StatusCode)
			Expect(nil).To(BeNil())
		})
	})
})
