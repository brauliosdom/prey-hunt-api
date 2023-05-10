package tests

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_PreyIsFaster(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPut, "/v1/prey", `{
		"speed": 10
	}`)
	r.ServeHTTP(rr, req)

	req, rr = createRequestTest(http.MethodPut, "/v1/shark", `{
		"x_position": 10,
  		"y_position": 10,
		"speed": 5
	}`)
	r.ServeHTTP(rr, req)

	req, rr = createRequestTest(http.MethodPost, "/v1/simulate", ``)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code) // 500
}

func Test_PreyIsToFar(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPut, "/v1/prey", `{
		"speed": 10
	}`)
	r.ServeHTTP(rr, req)

	req, rr = createRequestTest(http.MethodPut, "/v1/shark", `{
		"x_position": 10,
  		"y_position": 10,
		"speed": 10.1
	}`)
	r.ServeHTTP(rr, req)

	req, rr = createRequestTest(http.MethodPost, "/v1/simulate", ``)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code) // 500
}

func Test_PreyIsHunted(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPut, "/v1/prey", `{
		"speed": 10
	}`)
	r.ServeHTTP(rr, req)

	req, rr = createRequestTest(http.MethodPut, "/v1/shark", `{
		"x_position": 10,
  		"y_position": 10,
		"speed": 10.589
	}`)
	r.ServeHTTP(rr, req)

	req, rr = createRequestTest(http.MethodPost, "/v1/simulate", ``)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code) // 200
}

func Test_PreyBadData(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPut, "/v1/prey", `{
		"speed": "10"
	}`)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code) // 400
}

func Test_SharkBadData(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPut, "/v1/shark", `{
		"speed": "10"
	}`)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code) // 400
}
