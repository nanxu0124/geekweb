package web

import (
	"geekweb/geek"
	"geekweb/internal/domain"
	"geekweb/internal/service"
	"github.com/dlclark/regexp2"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type UserHandler struct {
	svc              service.UserService
	emailRegexExp    *regexp2.Regexp
	passwordRegexExp *regexp2.Regexp

	cmd redis.Cmdable
}

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	bizLogin             = "login"
)

func NewUserHandler(svc service.UserService) *UserHandler {

	return &UserHandler{
		svc:              svc,
		emailRegexExp:    regexp2.MustCompile(emailRegexPattern, regexp2.None),
		passwordRegexExp: regexp2.MustCompile(passwordRegexPattern, regexp2.None),
	}
}

func (u *UserHandler) RegisterRoutes(g *geek.HTTPServer) {

	g.Post("/users/signup", u.SignUp)

}

func (u *UserHandler) SignUp(ctx *geek.Context) {
	type SignReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.RespJSON(http.StatusOK, "系统错误")
		return
	}

	ok, err := u.emailRegexExp.MatchString(req.Email)
	if err != nil {
		ctx.RespJSON(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.RespJSON(http.StatusOK, "邮箱格式不正确")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.RespJSON(http.StatusOK, "两次输入的密码不相同")
		return
	}
	ok, err = u.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		ctx.RespJSON(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.RespJSON(http.StatusOK, "密码必须包含数字、特殊字符，并且长度不能小于8位")
		return
	}
	err = u.svc.SignUp(ctx.Req.Context(), domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicateEmail {
		ctx.RespJSON(http.StatusOK, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.RespJSON(http.StatusOK, "系统异常")
		return
	}
	ctx.RespJSON(http.StatusOK, "注册成功")
}
