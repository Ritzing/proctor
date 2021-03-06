package describe

import (
	"github.com/gojektech/proctor/daemon"
	"github.com/gojektech/proctor/io"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
)

type ScheduleCreateCmdTestSuite struct {
	suite.Suite
	mockPrinter        *io.MockPrinter
	mockProctorDClient *daemon.MockClient
	testScheduleDescribeCmd   *cobra.Command
}

func (s *ScheduleCreateCmdTestSuite) SetupTest() {
	s.mockPrinter = &io.MockPrinter{}
	s.mockProctorDClient = &daemon.MockClient{}
	s.testScheduleDescribeCmd = NewCmd(s.mockPrinter, s.mockProctorDClient)
}

func (s *ScheduleCreateCmdTestSuite) TestScheduleCreateCmdHelp() {
	assert.Equal(s.T(), "Describe scheduled job", s.testScheduleDescribeCmd.Short)
	assert.Equal(s.T(), "This command helps to describe scheduled job", s.testScheduleDescribeCmd.Long)
	assert.Equal(s.T(), "proctor schedule describe D958FCCC-F2B3-49D1-B83A-4E70A2A775A0", s.testScheduleDescribeCmd.Example)
}

func TestScheduleCreateCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleCreateCmdTestSuite))
}
