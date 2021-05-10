package pscan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestOpenedPort(t *testing.T) {
	s, _ := net.Listen("tcp", ":9001")

	t.Run("opened port", func(t *testing.T) {
		opened := OpenedPort("localhost", 9001)

		assert.True(t, opened)
	})

	t.Run("closed port", func(t *testing.T) {
		opened := OpenedPort("localhost", 9000)

		assert.False(t, opened)
	})

	s.Close()
}

func ExampleOpenedPort() {
	opened := OpenedPort("jandersonmartins.com.br", 80)
	fmt.Println(opened)
	//Output: true
}
