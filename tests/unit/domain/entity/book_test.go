//go:build unit
// +build unit

package entity_test

import (
	"github.com/andreis3/foodtosave-case/internal/domain/entity"
	"github.com/andreis3/foodtosave-case/internal/domain/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INTERNAL :: DOMAIN :: ENTITY :: BOOK", func() {
	Describe("#Validate", func() {
		Context("when the all fields is empty", func() {
			It("should return error", func() {
				book := &entity.Book{
					ID:     "",
					Title:  "",
					Gender: "",
				}
				err := book.Validate()

				Expect(err).NotTo(BeNil())
				Expect(err.ReturnErrors()).To(HaveLen(2))
				Expect(err.ReturnErrors()).To(ContainElement("title: is required"))
				Expect(err.ReturnErrors()).To(ContainElement("gender: is required"))
			})
			It("should return an author valid ", func() {
				book := &entity.Book{
					ID:     "1",
					Title:  "title",
					Gender: "gender",
				}

				err := book.Validate()

				Expect(err).To(Equal(&errors.NotificationErrors{}))
				Expect(err.ReturnErrors()).To(HaveLen(0))
			})
		})
	})
})
