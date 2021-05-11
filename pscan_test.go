package pscan

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net"
	"strings"
	"testing"
)

func TestPscan(t *testing.T) {
	s, _ := net.Listen("tcp", ":8081")

	t.Run("write all opened ports", func(t *testing.T) {
		buff := &bytes.Buffer{}

		Pscan("localhost", 5, buff)

		assert.True(t, strings.Contains(buff.String(), "8081\n"))
	})

	s.Close()
}
