package cmd_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-cli/cmd"
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	fakedir "github.com/cloudfoundry/bosh-cli/director/fakes"
	fakeui "github.com/cloudfoundry/bosh-cli/ui/fakes"
)

var _ = Describe("RecreateCmd", func() {
	var (
		ui         *fakeui.FakeUI
		deployment *fakedir.FakeDeployment
		command    RecreateCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		deployment = &fakedir.FakeDeployment{}
		command = NewRecreateCmd(ui, deployment)
	})

	Describe("Run", func() {
		var (
			opts RecreateOpts
		)

		BeforeEach(func() {
			opts = RecreateOpts{
				Args: AllOrPoolOrInstanceSlugArgs{
					Slug: boshdir.NewAllOrPoolOrInstanceSlug("some-name", ""),
				},
			}
		})

		act := func() error { return command.Run(opts) }

		It("recreate deployment, pool or instances", func() {
			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(deployment.RecreateCallCount()).To(Equal(1))

			slug, sd, force := deployment.RecreateArgsForCall(0)
			Expect(slug).To(Equal(boshdir.NewAllOrPoolOrInstanceSlug("some-name", "")))
			Expect(sd).To(Equal(boshdir.SkipDrain{}))
			Expect(force).To(BeFalse())
		})

		It("recreate allowing to skip drain scripts", func() {
			opts.SkipDrain = boshdir.SkipDrain{All: true}

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(deployment.RecreateCallCount()).To(Equal(1))

			slug, sd, force := deployment.RecreateArgsForCall(0)
			Expect(slug).To(Equal(boshdir.NewAllOrPoolOrInstanceSlug("some-name", "")))
			Expect(sd).To(Equal(boshdir.SkipDrain{All: true}))
			Expect(force).To(BeFalse())
		})

		It("recreate forcefully", func() {
			opts.Force = true

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(deployment.RecreateCallCount()).To(Equal(1))

			slug, sd, force := deployment.RecreateArgsForCall(0)
			Expect(slug).To(Equal(boshdir.NewAllOrPoolOrInstanceSlug("some-name", "")))
			Expect(sd).To(Equal(boshdir.SkipDrain{}))
			Expect(force).To(BeTrue())
		})

		It("does not recreate if confirmation is rejected", func() {
			ui.AskedConfirmationErr = errors.New("stop")

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("stop"))

			Expect(deployment.RecreateCallCount()).To(Equal(0))
		})

		It("returns error if restart failed", func() {
			deployment.RecreateReturns(errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
