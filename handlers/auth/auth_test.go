package auth

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bitbucket.org/aj5110/tpo-lp4/repositories"
	"bitbucket.org/aj5110/tpo-lp4/services"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite

	server *httptest.Server

	userRepoMock *repositories.UserRepoMock
}

func (ts *AuthTestSuite) SetupSuite() {
	mux := http.NewServeMux()

	ts.userRepoMock = repositories.NewUserRepoMock()
	SetAuthHandler(mux, services.NewAuthService(ts.userRepoMock))

	ts.server = httptest.NewServer(mux)
}

func (ts *AuthTestSuite) TestRegister() {
	// mock user repo
	ts.userRepoMock.On("Register", "john@doe.com", "password123").Return(nil)

	for _, tc := range []struct {
		reqMethod string
		reqBody   io.Reader
		resCode   int
		resBody   []byte
	}{
		{http.MethodPost, strings.NewReader(`{"email":"john@doe.com","password":"password123"}`), http.StatusOK, []byte("")},
		{http.MethodGet, strings.NewReader(`{"email":"a@b.c","password":"p"}`), http.StatusMethodNotAllowed, []byte("")},
	} {
		req, err := http.NewRequest(tc.reqMethod, ts.server.URL+"/register/", tc.reqBody)
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

	ts.userRepoMock.AssertCalled(ts.T(), "Register", "john@doe.com", "password123")
	ts.userRepoMock.AssertNotCalled(ts.T(), "Register", "a@b.c", "p")
}

func (ts *AuthTestSuite) TeardownTest() {
	ts.server.Close()
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
