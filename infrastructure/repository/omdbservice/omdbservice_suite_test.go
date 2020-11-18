package omdbservice_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOmdbservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Omdbservice Suite")
}
