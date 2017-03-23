package table_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/bosh-cli/ui/table"
)

var _ = Describe("Headers", func() {

	Describe("KeyifyHeader", func() {
		It("should convert alphanumeric to lowercase ", func() {
			keyifyHeader := table.KeyifyHeader("Header1")
			Expect(keyifyHeader).To(Equal("header1"))
		})

		Context("given a header that only contains non-alphanumeric and alphanumeric", func() {
			It("should non-alphanumeric to underscore", func() {
				keyifyHeader := table.KeyifyHeader("FOO!@AND#$BAR")
				Expect(keyifyHeader).To(Equal("foo_and_bar"))
			})
		})

		Context("given a header that only contains non-alphanumeric", func() {
			It("should convert to underscore", func() {
				keyifyHeader := table.KeyifyHeader("!@#$")
				Expect(keyifyHeader).To(Equal("_"))
			})

			It("should convert empty header to underscore", func() {
				keyifyHeader := table.KeyifyHeader("")
				Expect(keyifyHeader).To(Equal("_"))
			})
		})

	})
})
