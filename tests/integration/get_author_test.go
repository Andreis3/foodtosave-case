//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("INTEGRATION :: TEST :: GET :: EXISTS :: AUTHOR AND BOOKS", func() {
	Describe("#POST", func() {
		Context("When send request GET to route /v1/author", func() {
			When("should return a author and all books", func() {
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

					req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v1/author/%s", response["data"].(map[string]any)["id"]), nil)

					res, err = client.Do(req)

					response = map[string]any{}

					json.NewDecoder(res.Body).Decode(&response)

					Expect(err).To(BeNil())
					Expect(res.StatusCode).To(Equal(http.StatusOK))
					Expect(response["request_id"]).NotTo(BeNil())
					Expect(response["status_code"].(float64)).To(Equal(float64(http.StatusOK)))
					Expect(response["data"].(map[string]any)["name"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["nationality"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["books"].([]any)[0].(map[string]any)["title"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["books"].([]any)[0].(map[string]any)["gender"]).To(Equal("test 1"))

				})
			})

			When("should return error when not exists author", func() {
				var client = &http.Client{}
				BeforeEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				AfterEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				It("Should return status 400", func() {
					req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v1/author/%s", "9e952ad4-bc44-4e44-bfdf-fd6c6c371d72"), nil)

					res, err := client.Do(req)

					response := map[string]any{}

					json.NewDecoder(res.Body).Decode(&response)

					Expect(err).To(BeNil())

					Expect(res.StatusCode).To(Equal(http.StatusNotFound))
					Expect(response["request_id"]).NotTo(BeNil())
					Expect(response["status_code"].(float64)).To(Equal(float64(http.StatusNotFound)))
					Expect(response["error_message"].([]any)).To(ContainElements("Author not found"))
				})
			})
		})
	})
})
