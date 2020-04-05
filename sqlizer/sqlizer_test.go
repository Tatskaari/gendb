package sqlizer_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type sqlizerSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(sqlizerSuite))
}
