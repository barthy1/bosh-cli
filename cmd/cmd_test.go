package cmd_test

import (
	"errors"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-cli/cmd"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	fakeui "github.com/cloudfoundry/bosh-cli/ui/fakes"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"
)

var _ = Describe("Cmd", func() {
	var (
		ui     *fakeui.FakeUI
		confUI *boshui.ConfUI
		fs     *fakesys.FakeFileSystem
		cmd    Cmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		logger := boshlog.NewLogger(boshlog.LevelNone)
		confUI = boshui.NewWrappingConfUI(ui, logger)

		fs = fakesys.NewFakeFileSystem()

		deps := NewBasicDeps(confUI, logger)
		deps.FS = fs

		cmd = NewCmd(BoshOpts{}, nil, deps)
	})

	Describe("Execute", func() {
		It("succeeds executing at least one command", func() {
			cmd.Opts = &BuildManifestOpts{}

			err := cmd.Execute()
			Expect(err).ToNot(HaveOccurred())

			Expect(ui.Blocks).To(Equal([]string{"null\n"}))
		})

		It("prints message if specified", func() {
			cmd.Opts = &MessageOpts{Message: "output"}

			err := cmd.Execute()
			Expect(err).ToNot(HaveOccurred())

			Expect(ui.Blocks).To(Equal([]string{"output"}))
		})

		It("allows to enable json output", func() {
			cmd.BoshOpts = BoshOpts{JSONOpt: true}
			cmd.Opts = &BuildManifestOpts{}

			err := cmd.Execute()
			Expect(err).ToNot(HaveOccurred())

			confUI.Flush()

			Expect(ui.Blocks[0]).To(ContainSubstring(`Blocks": [`))
		})

		Describe("color", func() {
			executeCmdAndPrintTable := func() {
				err := cmd.Execute()
				Expect(err).ToNot(HaveOccurred())

				// Tables have emboldened header values
				confUI.PrintTable(boshtbl.Table{Header: []string{"State"}})
			}

			It("has color in the output enabled by default", func() {
				cmd.BoshOpts = BoshOpts{}
				cmd.Opts = &BuildManifestOpts{}

				executeCmdAndPrintTable()

				// Expect that header values are bold
				Expect(ui.Tables[0].HeaderVals[0].(boshtbl.ValueFmt).Func).ToNot(BeNil())
			})

			It("allows to disable color in the output", func() {
				cmd.BoshOpts = BoshOpts{NoColorOpt: true}
				cmd.Opts = &BuildManifestOpts{}

				executeCmdAndPrintTable()

				// Expect that header values are empty because they were not emboldened
				Expect(ui.Tables[0].HeaderVals).To(BeEmpty())
			})
		})

		It("returns error if changing tmp root fails", func() {
			fs.ChangeTempRootErr = errors.New("fake-err")

			err := cmd.Execute()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("fake-err"))
		})

		It("returns error for unknown commands", func() {
			err := cmd.Execute()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Unhandled command: <nil>"))
		})
	})
})