package tc

import (
	"log"
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/go-tstr/tstr"
	"github.com/go-tstr/tstr/dep/compose"
)

func TestMain(m *testing.M) {
	os.Setenv("POSTGRES_PORT", strconv.Itoa(mustFreePort()))
	tstr.RunMain(m, tstr.WithDeps(
		compose.New(
			compose.WithFile("docker-compose.yaml"),
			compose.WithOsEnv(),
		),
	))
}

func TestValidate(t *testing.T) {
	t.Log("Test environment is working correctly")
}

func mustFreePort() int {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	tcpAddr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		log.Fatal(err)
	}

	return tcpAddr.Port
}
