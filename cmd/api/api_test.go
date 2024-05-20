package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServeHTTP(t *testing.T) {
	t.Run("Server starts successfully", func(t *testing.T) {
		go func() {
			// We expect this call not to block, hence running it in a goroutine.
			err := ServeHTTP(":9999") // Use a test-specific port
			assert.Nil(t, err, "Expected no error when starting server")
		}()

		// Allow some time for the server to start
		time.Sleep(time.Second)

		// Make a test request to verify the server is listening
		_, err := http.Get("http://localhost:9999")
		assert.Nil(t, err, "Expected to connect to the server successfully")
	})
}
