package api

import (
	"crawlab-lite/forms"
	"crawlab-lite/models"
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
)

func TestUserAPI(t *testing.T) {
	Convey("Test User API", t, func() {
		app := InitTestApp()
		users, err := models.GetUserList()
		So(err, ShouldBeNil)
		user := users[0]

		Convey("Test right user", func() {
			w := httptest.NewRecorder()
			form := forms.UserForm{
				Username: user.Username,
				Password: user.Password,
			}
			req, err := PostJson("/api/login", form)
			So(err, ShouldBeNil)
			app.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, 200)

			resp := GetResponse(w.Body)
			So(resp.Code, ShouldEqual, 200)
			So(resp.Message, ShouldEqual, "success")
			So(resp.Data, ShouldNotBeEmpty)
		})

		Convey("Test wrong user", func() {
			w := httptest.NewRecorder()
			values := map[string]string{"username": "abcdefg", "password": "000000"}
			req, err := PostJson("/api/login", values)
			So(err, ShouldBeNil)
			app.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, 401)

			resp := GetResponse(w.Body)
			So(resp.Code, ShouldEqual, 401)
			So(resp.Message, ShouldEqual, "not authorized")
			So(resp.Data, ShouldBeEmpty)
		})
	})
}