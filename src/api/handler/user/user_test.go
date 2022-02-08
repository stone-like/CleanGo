package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stonelike/CleanGo/src/api/presenter"
	"github.com/stonelike/CleanGo/src/domain/entity"
	infra "github.com/stonelike/CleanGo/src/infra/user"
	usecase "github.com/stonelike/CleanGo/src/usecase/user"
	mock_user "github.com/stonelike/CleanGo/src/usecase/user/mock"
	"github.com/stretchr/testify/require"
)

func Test_createUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock_user.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("createUser").GetPathTemplate()
	require.NoError(t, err)
	require.Equal(t, "/user", path)

	user, _ := entity.NewUser("test")
	m.EXPECT().
		CreateUser("test").
		Return(user, nil)

	h := createUser(m)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
       "name": "test"
    }`)
	//適当なurlでもルーティングしてしまう、おそらくここではurlの正しさは、require.Equal(t, "/user", path)で検証している、ts.URLとすれば今回httptest.NewServer(h)のhはcreateUserのみなので、絶対にcreateUserが起動する
	resp, err := http.Post(ts.URL+"/user", "application/json", strings.NewReader(payload))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.User
	json.NewDecoder(resp.Body).Decode(&u)
	require.Equal(t, "test", u.Name)
}

func Test_getUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock_user.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, m)
	path, err := r.GetRoute("getUser").GetPathTemplate()
	require.Nil(t, err)
	require.Equal(t, "/user/{id}", path)
	user, _ := entity.NewUser("test")
	m.EXPECT().
		FindById(user.GetId()).
		Return(user, nil)

	handler := getUser(m)
	r.Handle("user/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/user/" + user.GetId())
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	var u *presenter.User
	json.NewDecoder(res.Body).Decode(&u)
	require.NotNil(t, u)
	require.Equal(t, user.GetId(), u.Id)
}

func Test_getUserError(t *testing.T) {
	repo := infra.NewInmem()
	u := usecase.NewService(repo)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, u)

	handler := getUser(u)
	r.Handle("user/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()

	payload := fmt.Sprintf(`{
		"name": "test"
	 }`)
	resp, err := http.Post(ts.URL+"/user", "application/json", strings.NewReader(payload))
	require.Nil(t, err)
	defer resp.Body.Close()

	res, err := http.Get(ts.URL + "/user/" + "dummy")
	require.Nil(t, err)
	defer res.Body.Close()

}

func Test_createUserError(t *testing.T) {
	repo := infra.NewInmem()
	u := usecase.NewService(repo)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, u)

	handler := getUser(u)
	r.Handle("user/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()

	payload := fmt.Sprintf(`{
		"name": "testtttttttt"
	 }`)
	resp, err := http.Post(ts.URL+"/user", "application/json", strings.NewReader(payload))
	require.Nil(t, err)

	defer resp.Body.Close()

}
