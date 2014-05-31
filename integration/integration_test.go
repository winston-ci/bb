package integration_test

import (
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Big Brother", func() {
	It("works at a high level", func() {
		config := filepath.Join("fixtures", "winston.yml")
		bbCmd := exec.Command(bbPath, config)

		session, err := gexec.Start(bbCmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))
	})
})
