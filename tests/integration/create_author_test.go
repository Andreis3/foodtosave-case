//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_IntegrationCreateGroup(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Integration Create Group Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("INTEGRATION :: TEST :: CREATE :: NEW :: AUTHOR AND BOOKS", func() {
	Describe("#POST", func() {
		Context("When send request to route /v1/author", func() {
			When("When send request with valid data", func() {
				var client = &http.Client{}
				BeforeEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				AfterEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				It("Should return status 201", func() {
					body, _ := json.Marshal(map[string]interface{}{
						"name":        "test 1",
						"nationality": "test 1",
						"books": []map[string]interface{}{
							{
								"title":  "test 1",
								"gender": "test 1",
							},
						},
					})
					payload := bytes.NewBuffer(body)

					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/v1/author", payload)
					defer req.Body.Close()
					req.Header.Set("Content-Type", "application/json")

					res, err := client.Do(req)

					response := map[string]any{}

					json.NewDecoder(res.Body).Decode(&response)

					Expect(err).To(BeNil())
					Expect(res.StatusCode).To(Equal(http.StatusCreated))
					Expect(response["request_id"]).NotTo(BeNil())
					Expect(response["status_code"].(float64)).To(Equal(float64(http.StatusCreated)))
					Expect(response["data"].(map[string]any)["name"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["nationality"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["books"].([]any)[0].(map[string]any)["title"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["books"].([]any)[0].(map[string]any)["gender"]).To(Equal("test 1"))

				})
			})

			When("When send invalid data should return error", func() {
				var client = &http.Client{}
				BeforeEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				AfterEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				It("Should return status 400", func() {
					body, _ := json.Marshal(map[string]interface{}{
						"name":        "",
						"nationality": "test 1",
						"books": []map[string]interface{}{
							{
								"title":  "test 1",
								"gender": "",
							},
						},
					})

					payload := bytes.NewBuffer(body)

					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/v1/author", payload)
					req.Header.Set("Content-Type", "application/json")

					res, err := client.Do(req)

					response := map[string]any{}

					json.NewDecoder(res.Body).Decode(&response)

					Expect(err).To(BeNil())

					Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(response["request_id"]).NotTo(BeNil())
					Expect(response["status_code"].(float64)).To(Equal(float64(http.StatusBadRequest)))
					Expect(response["error_message"].([]any)).To(ContainElements("books[0].gender: is required", "name: is required"))
				})
			})
		})
	})
})
