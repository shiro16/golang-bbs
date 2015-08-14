package tests

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel/testing"
	"github.com/shiro16/golang-bbs/app/controllers"
	. "github.com/shiro16/golang-bbs/app/controllers/api/v1"
	"github.com/shiro16/golang-bbs/app/models"
	"runtime"
	"strconv"
	"strings"
)

var comment *models.Comment

type ApiV1CommentTest struct {
	testing.TestSuite
}

func (t *ApiV1CommentTest) Before() {
	comment = &models.Comment{Body: "test"}
	controllers.DB.Create(comment)
}

func (t *ApiV1CommentTest) After() {
	controllers.DB.Table("comments").Delete(nil)
}

func (t *ApiV1CommentTest) TestIndex() {
	t.Get("/api/v1/comments")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	var r Response
	json.Unmarshal(t.ResponseBody, &r)
	comments := r.Results.([]interface{})
	c := comments[0].(map[string]interface{})

	h := c["id"].(float64)
	a := uint64(h)

	_, f, l, _ := runtime.Caller(1)
	m := fmt.Sprintf("In %s (arount line %d)", f, l)
	t.Assertf(a == comment.ID, "%s: json data not matched", m)
}

func (t *ApiV1CommentTest) TestShow() {
	t.Get("/api/v1/comments/0")
	t.AssertNotFound()

	c := &models.Comment{Body: "test"}
	controllers.DB.Create(c)
	t.Get("/api/v1/comments/" + strconv.FormatUint(c.ID, 10))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	var r Response
	json.Unmarshal(t.ResponseBody, &r)
	comment := r.Results.(map[string]interface{})
	h := comment["id"].(float64)
	a := uint64(h)

	_, f, l, _ := runtime.Caller(1)
	m := fmt.Sprintf("In %s (arount line %d)", f, l)
	t.Assertf(a == c.ID, "%s: json data not matched", m)
}

func (t *ApiV1CommentTest) TestCreate() {
	t.Post("/api/v1/comments", "application/json; charset=utf-8", strings.NewReader("{\"body\":\"\"}"))
	t.AssertStatus(400)

	t.Post("/api/v1/comments", "application/json; charset=utf-8", strings.NewReader("{\"body\":\"test\"}"))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	var r Response
	json.Unmarshal(t.ResponseBody, &r)
	c := r.Results.(map[string]interface{})

	_, f, l, _ := runtime.Caller(1)
	m := fmt.Sprintf("In %s (arount line %d)", f, l)
	t.Assertf(c["body"] == "test", "%s: json data not matched", m)
}

func (t *ApiV1CommentTest) TestDelete() {
	t.Delete("/api/v1/comments/0")
	t.AssertNotFound()

	c := &models.Comment{Body: "test"}
	controllers.DB.Create(c)
	t.Delete("/api/v1/comments/" + strconv.FormatUint(c.ID, 10))
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	var r Response
	json.Unmarshal(t.ResponseBody, &r)

	_, f, l, _ := runtime.Caller(1)
	m := fmt.Sprintf("In %s (arount line %d)", f, l)
	t.Assertf(r.Results == "success", "%s: json data not matched", m)
}
