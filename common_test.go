package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	// set up Gin to test mode before running tests
	gin.SetMode(gin.TestMode)

	// run the other tests
	os.Exit(m.Run())
}

// Helper func to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// Helper func to process a request and test response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// create a response recorder
	w := httptest.NewRecorder()

	// create the service and process the request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}

}

// store the main lists into temp one for testing
func saveLists() {
	tmpArticleList = articleList
}

// restore main list from temp one
func restoreLists() {
	articleList = tmpArticleList
}
