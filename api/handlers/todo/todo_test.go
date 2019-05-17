package todo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"bitbucket.org/aj5110/tpo-lp4/api/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
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

func (ts *TodoTestSuite) TestList() {
	// mock todo repo
	ts.todoRepoMock.On("List").Return([]entities.Todo{
		{Id: bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), Description: "description 1"},
		{Id: bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}), Description: "description 2"},
	}, nil)

	for _, tc := range []struct {
		resCode int
		resBody string
	}{
		{http.StatusOK, `[{"_id":"000000000000000000000000","description":"description 1"},{"_id":"000000000000000000000001","description":"description 2"}]`},
	} {
		req, err := http.NewRequest(http.MethodGet, ts.server.URL+"/todo/", nil)
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

		ts.JSONEq(tc.resBody, string(resBody))
	}

	ts.todoRepoMock.AssertCalled(ts.T(), "List")
}

func (ts *TodoTestSuite) TestAdd() {
	t1 := &entities.Todo{
		Description: "description 1",
	}

	// mock todo repo
	ts.todoRepoMock.On("Add", t1).Return(&entities.Todo{
		Id:          bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		Description: "description 1",
	}, nil)

	for _, tc := range []struct {
		reqBody io.Reader
		resCode int
		resBody string
	}{
		{strings.NewReader(`{"description":"description 1"}`), http.StatusOK, `{"_id":"000000000000000000000001","description":"description 1"}`},
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

		ts.JSONEq(tc.resBody, string(resBody))
	}

	ts.todoRepoMock.AssertCalled(ts.T(), "Add", t1)
}

func (ts *TodoTestSuite) TestEdit() {
	var (
		oid1 = bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
		t1   = &entities.Todo{
			Description: "description 1",
		}
	)

	// mock todo repo
	ts.todoRepoMock.On("Edit", oid1, t1).Return(&entities.Todo{Id: oid1, Description: t1.Description}, nil)

	for _, tc := range []struct {
		reqBody io.Reader
		resCode int
		resBody string
	}{
		{strings.NewReader(`{"description":"description 1"}`), http.StatusOK, `{"_id":"000000000000000000000001","description":"description 1"}`},
	} {
		req, err := http.NewRequest(http.MethodPut, ts.server.URL+"/todo/000000000000000000000001", tc.reqBody)
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

		ts.JSONEq(tc.resBody, string(resBody))
	}

	ts.todoRepoMock.AssertCalled(ts.T(), "Edit", oid1, t1)
}

func (ts *TodoTestSuite) TestDelete() {
	var (
		oid1 = bson.ObjectId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	)

	// mock todo repo
	ts.todoRepoMock.On("Remove", oid1).Return(nil)

	for _, tc := range []struct {
		reqId   string
		resCode int
	}{
		{`000000000000000000000001`, http.StatusOK},
	} {
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/todo/%s", ts.server.URL, tc.reqId), nil)
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

		ts.Empty(resBody)
	}

	ts.todoRepoMock.AssertCalled(ts.T(), "Remove", oid1)
}

func (ts *TodoTestSuite) TeardownTest() {
	ts.server.Close()
}

func TestTodoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoTestSuite))
}
