package mysql_test

import (
	"testing"

	. "gopkg.in/check.v1"
	//"github.com/screwyprof/gosandbox/services/sandbox"
	"github.com/screwyprof/gosandbox/test"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type SandboxMysqlStorageSuite struct {
	test.IntegrationTestSuite
}

var _ = Suite(&SandboxMysqlStorageSuite{})

func (s *SandboxMysqlStorageSuite) TestLoadUserNameById(c *C) {

	sandboxService := s.SandboxService
	c.Assert(sandboxService.LoadUserNameById(1), Equals, "admin")
}
