package e2e

import (
	"net/http"
	"testing"

	"go-clean-app/config"
	"go-clean-app/test/helpers"

	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, &UserSuite{})
}

func (s *UserSuite) SetupTest() {}

func (s *UserSuite) TearDownTest() {}

func TestUser(t *testing.T) {
	helpers.Initialize("../../.env.test")
	apiConfig := config.GetAPIConfig()
	hostURL := apiConfig.HostURL
	version := apiConfig.ApiVersion
	url := hostURL + "/" + version
	e := httpexpect.Default(t, url)
	t.Run("すべてのユーザー一覧を取得できる", func(t *testing.T) {
		res := e.GET("/users").
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()
		res.Value("users").IsArray().Array().Length().IsEqual(3)
	})
}
