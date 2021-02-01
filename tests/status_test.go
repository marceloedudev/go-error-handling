package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-error-handling/dtos"
	"go-error-handling/routes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func AddRouters() *gin.Engine {
	return routes.AddRoutes()
}

func main() {
	app := AddRouters()
	app.Run()
}

func TestStatus(t *testing.T) {

	testRouter := AddRouters()

	t.Run("should be able to return the data", func(t *testing.T) {

		var formLogin = &dtos.StatusJSON{
			Name:  "Marcelo",
			Email: "example@gmail.com",
		}

		data, _ := json.Marshal(formLogin)

		req, err := http.NewRequest("POST", "/status", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		assert.Equal(t, resp.Code, http.StatusOK)

	})

	t.Run("should be able to show the 'causes' field", func(t *testing.T) {

		var formLogin = &dtos.StatusJSON{
			Name:  "Marcelo",
			Email: "example",
		}

		data, _ := json.Marshal(formLogin)

		req, err := http.NewRequest("POST", "/status", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var res = &struct {
			Causes []string
		}{}

		json.Unmarshal(body, &res)

		assert.NotEqual(t, res.Causes, nil)

	})

	t.Run("should not be able to with just the name field", func(t *testing.T) {

		var formLogin = &dtos.StatusJSON{
			Name: "Marcelo",
		}

		data, _ := json.Marshal(formLogin)

		req, err := http.NewRequest("POST", "/status", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		assert.Equal(t, resp.Code, http.StatusInternalServerError)

	})

}

func TestPanic(t *testing.T) {

	testRouter := AddRouters()

	t.Run("should be able to have 'test' return", func(t *testing.T) {

		req, err := http.NewRequest("POST", "/panic1", nil)
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var res = &struct {
			Message string
		}{}

		json.Unmarshal(body, &res)

		assert.Equal(t, res.Message, "test")

	})

	t.Run("You should be able to an error", func(t *testing.T) {

		req, err := http.NewRequest("POST", "/panic2", nil)
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)

		assert.Equal(t, resp.Code, http.StatusInternalServerError)

	})

}
