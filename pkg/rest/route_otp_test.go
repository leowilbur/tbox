package rest_test

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leowilbur/tbox/pkg/rest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/resty.v1"
)

func TestRoute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "pkg/rest")
}

var _ = Describe("TBOX API", func() {
	Context("OTP Test", func() {
		var (
			restServer *httptest.Server
			restClient *resty.Client
			db         *sql.DB
			router     *rest.API
			err        error
		)

		BeforeEach(func() {
			if router == nil {
				db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tbox?timeout=360s&multiStatements=true&parseTime=true")
				Expect(err).To(BeNil())
				router, err = rest.New(db)
				Expect(err).To(BeNil())
			}
			/*
				Start Server + Client
			*/
			restServer = httptest.NewServer(router)

			restClient = resty.New().SetHostURL(restServer.URL)
		})

		AfterEach(func() {
			restServer.Close()
		})

		It("Should generate and send otp success! ", func() {
			resp, err := restClient.NewRequest().
				SetContext(context.Background()).
				SetBody(map[string]interface{}{
					"phoneNumber": "091558493",
				}).
				Post("/users/otp/generate")

			Expect(err).To(BeNil())
			Expect(resp.IsSuccess()).To(BeTrue())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
	})
})
