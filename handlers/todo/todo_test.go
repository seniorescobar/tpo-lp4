package todo

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bitbucket.org/aj5110/tpo-lp4/entities"
	"bitbucket.org/aj5110/tpo-lp4/repositories"
	"bitbucket.org/aj5110/tpo-lp4/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type TodoTestSuite struct {
	suite.Suite

	server *httptest.Server

	todoRepoMock *repositories.TodoRepoMock
}

func (ts *TodoTestSuite) SetupSuite() {
	r := mux.NewRouter()

	ts.todoRepoMock = repositories.NewTodoRepoMock()
	SetTodoHandler(r, services.NewTodoService(ts.todoRepoMock))

	ts.server = httptest.NewServer(r)
}

func (ts *TodoTestSuite) TestAdd() {
	desc1 := "description 1"
	t1 := &entities.Todo{
		Description: &desc1,
	}

	// mock todo repo
	ts.todoRepoMock.On("Add", t1).Return(nil)

	for _, tc := range []struct {
		reqBody io.Reader
		resCode int
		resBody []byte
	}{
		{strings.NewReader(`{"description":"description 1"}`), http.StatusOK, []byte("")},
	} {
		req, err := http.NewRequest(http.MethodPost, ts.server.URL+"/todo/", tc.reqBody)
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		ts.Equal(tc.resCode, res.StatusCode)

		resBody, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		ts.Equal(tc.resBody, resBody)
	}

	ts.todoRepoMock.AssertCalled(ts.T(), "Add", t1)
}

func (ts *TodoTestSuite) TeardownTest() {
	ts.server.Close()
}

func TestTodoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoTestSuite))
}
