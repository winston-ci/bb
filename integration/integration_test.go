package integration_test

import (
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Big Brother", func() {
	It("works at a high level", func() {
		config := filepath.Join("fixtures", "winston.yml")
		bbCmd := exec.Command(bbPath, config)

		output := NewBuffer()
		session, err := Start(bbCmd, output, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(output).Should(Say("digraph"))
		Eventually(session).Should(Exit(0))
	})
})
