package director_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-cli/director"
)

var _ = Describe("NewConfigFromURL", func() {
	It("sets host and port (25555) if scheme is specified", func() {
		config, err := NewConfigFromURL("https://host")
		Expect(err).ToNot(HaveOccurred())
		Expect(config).To(Equal(Config{Host: "host", Port: 25555}))
	})

	It("sets host and port (25555) if scheme is not specified", func() {
		config, err := NewConfigFromURL("host")
		Expect(err).ToNot(HaveOccurred())
		Expect(config).To(Equal(Config{Host: "host", Port: 25555}))
	})

	It("extracts port if scheme is specified", func() {
		config, err := NewConfigFromURL("https://host:4443")
		Expect(err).ToNot(HaveOccurred())
		Expect(config).To(Equal(Config{Host: "host", Port: 4443}))
	})

	It("extracts port if scheme is not specified", func() {
		config, err := NewConfigFromURL("host:4443")
		Expect(err).ToNot(HaveOccurred())
		Expect(config).To(Equal(Config{Host: "host", Port: 4443}))
	})

	It("returns error if url is empty", func() {
		_, err := NewConfigFromURL("")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Expected non-empty Director URL"))
	})

	It("returns error if host is not specified", func() {
		_, err := NewConfigFromURL("https://:25555")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Expected to extract host from"))
	})

	It("returns error if parsing url fails", func() {
		_, err := NewConfigFromURL(":/")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Parsing Director URL"))
	})

	It("returns error if port cannot be extracted", func() {
		_, err := NewConfigFromURL("https://host::")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Extracting host/port from URL"))
	})

	It("returns error if port is empty", func() {
		_, err := NewConfigFromURL("host:")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Extracting port from URL"))
	})

	It("returns error if port cannot be parsed as int", func() {
		_, err := NewConfigFromURL("https://host:abc")
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Extracting port from URL"))
	})
})

var _ = Describe("FactoryConfig", func() {
	Describe("Validate", func() {
		It("returns without error for basic config", func() {
			err := Config{Host: "host", Port: 1}.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if host is empty", func() {
			err := Config{}.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Missing 'Host'"))
		})

		It("returns error if host is empty", func() {
			err := Config{Host: "host"}.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Missing 'Port'"))
		})

		It("returns error if cannot parse PEM formatted block", func() {
			err := Config{
				Host:   "host",
				Port:   1,
				CACert: "-",
			}.Validate()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing CA certificate: Missing PEM block"))
		})

		It("returns error if PEM formatted block is not a certificate", func() {
			err := Config{
				Host: "host",
				Port: 1,
				CACert: `-----BEGIN PRIVATE KEY-----
MIIDXzCCAkegAwIBAgIJAPerMgLAne5vMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
-----END PRIVATE KEY-----`,
			}.Validate()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Parsing CA certificate: Not a certificate"))
		})

		It("returns error if parsing certificate fails", func() {
			err := Config{
				Host: "host",
				Port: 1,
				CACert: `-----BEGIN CERTIFICATE-----
MIIDXzCCAkegAwIBAgIJAPerMgLAne5vMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
-----END CERTIFICATE-----`,
			}.Validate()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Parsing CA certificate: asn1: syntax error:"))
		})
	})

	Describe("CACertPool", func() {
		It("returns without error for basic config", func() {
			caCert := `-----BEGIN CERTIFICATE-----
MIIDXzCCAkegAwIBAgIJAPerMgLAne5vMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwIBcNMTYwMTE2MDY0NTA0WhgPMjI4OTEwMzAwNjQ1MDRa
MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJ
bnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQCtSo3KPjnVPzodb6+mNwbCdcpzVop8OmfwJ3ynQtyBEzGaKsAn4tlz
/wfQQrKFHgxqVpqcoxAlWPNMs5+iO2Jst3Gz2+oLcaDyz/EWorw0iF5q1F6+WYHp
EijY20MzaWYMyu4UhhlbJCkSGZSjujh5SFOAXQwWYJXsqjyxA9KaTD6OdH5Kpger
B9D4zogX0We00eouyvvz/sAeDbTshk9sJRGWHNFJr+TjVx2D01alU49liAL94yF6
1eEOEbE50OAhv9RNsRh6O58idaHg30bbMf1yAzcgBvh8CzIHH0BPofoF2pRfztoY
uudZ0ftJjTz4fA2h/7GOVzxemrTjx88vAgMBAAGjUDBOMB0GA1UdDgQWBBQjz5Q2
YW2kBTb4XLqKFZMSBLpi6zAfBgNVHSMEGDAWgBQjz5Q2YW2kBTb4XLqKFZMSBLpi
6zAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQA/s94M/mSGELHJWIb1
oE0IKHWajBd3Pc8+O1TZRE+ke3q+rZRfcxd2dAjq6zQHJUs2+fs0B3DyT9Wtyyoq
UrRdsgprOdf2Cuw8bMIsCQOvqWKhhdlLTnCi2xaGJawGsIkheuD1n+Il9gRQ2WGy
lACxVngPwjNYxjOE+CUnSZCuAmAfQYzqto3bNPqkgEwb7ueODeOiyhR8SKsH7ySW
QAOSxgrLBblGLWcDF9fjMeYaUnI34pHviCKeVxfgsxDR+Jg11F78sPdYLOF6ipBe
/5qTYucsY20B2EKtlscD0mSYBRwbVrSQt2RYbTCwaibxWUC13VV+YEk0NAv9Mm04
6sKO
-----END CERTIFICATE-----`

			certPool, err := Config{CACert: caCert}.CACertPool()
			Expect(err).ToNot(HaveOccurred())
			Expect(certPool.Subjects()[0]).To(ContainSubstring("Internet Widgits Pty Ltd"))
		})
	})
})
