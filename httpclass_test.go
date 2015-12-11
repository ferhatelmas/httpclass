package httpclass

import (
	"net/http"
	"testing"
)

func checkCodes(codes []int, status StatusClass, shouldFail bool, t *testing.T) {
	for _, c := range codes {
		s, err := Get(c)
		if shouldFail {
			if err == nil {
				t.Errorf("Expected to fail but not for %d", c)
			}
		} else {
			if err != nil {
				t.Errorf("Shouldn't fail but failed for %d", c)
			} else if s != status {
				t.Errorf("Expected %d, got %d", status, s)
			}
		}
	}
}

func TestWrong(t *testing.T) {
	codes := []int{0, 99, 600, 10000}
	checkCodes(codes, 0, true, t)
}

func TestInformational(t *testing.T) {
	codes := []int{
		http.StatusContinue,
		http.StatusSwitchingProtocols,
	}
	checkCodes(codes, Informational, false, t)
}

func TestSuccess(t *testing.T) {
	codes := []int{
		http.StatusOK,
		http.StatusCreated,
	}
	checkCodes(codes, Success, false, t)
}

func TestRedirection(t *testing.T) {
	codes := []int{
		http.StatusMovedPermanently,
		http.StatusFound,
	}
	checkCodes(codes, Redirection, false, t)
}

func TestClientError(t *testing.T) {
	codes := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
	}
	checkCodes(codes, ClientError, false, t)
}

func TestServerError(t *testing.T) {
	codes := []int{
		http.StatusInternalServerError,
		http.StatusServiceUnavailable,
	}
	checkCodes(codes, ServerError, false, t)
}
