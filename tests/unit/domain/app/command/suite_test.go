//go:build unit
// +build unit

package entity_test

import "testing"

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_SuiteEntityContext(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()
	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true
	RegisterFailHandler(Fail)
	RunSpecs(t, "Command Suite Test Context", suiteConfig, reporterConfig)
}
