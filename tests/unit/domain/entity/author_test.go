//go:build unit
// +build unit

package entity_test

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/notification"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITY :: AUTHOR", func() {
	Describe("#Validate", func() {
		Context("when the all fields is empty", func() {
			It("should return error", func() {
				author := &entity.Author{
					ID:          "",
					Name:        "",
					Nationality: "",
				}
				err := author.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.ReturnErrors()).To(HaveLen(2))
				Expect(err.ReturnErrors()).To(ContainElement("name: is required"))
				Expect(err.ReturnErrors()).To(ContainElement("nationality: is required"))
			})
			It("should return an author valid ", func() {
				author := &entity.Author{
					ID:          "1",
					Name:        "author",
					Nationality: "BR",
				}

				err := author.Validate()

				Expect(err).To(Equal(&notification.Error{}))
				Expect(err.ReturnErrors()).To(HaveLen(0))
			})
		})
	})
})
