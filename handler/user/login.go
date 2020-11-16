package user

import (
	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"
	"todolist_backend/pkg/token"
	"todolist_backend/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type LoginRequest struct {
	Code string `json:"code"` // 小程序 auth code
}

type LoginResponse struct {
	Token string `json:"token"`
	IsNew bool   `json:"is_new"` // 是否是新用户
}

// @Summary 登录
// @Tags user
// @Param data body user.LoginRequest true "body data"
// @Success 200 {object} user.LoginResponse
// @Router /user/login [post]
func Login(c *gin.Context) {
	request := LoginRequest{}
	if err := c.BindJSON(&request); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 根据 code 从 QQ 获取 openID 和 sessionKey
	openID, sessionKey, err := service.CodeToSession(request.Code)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	// 检查用户是否存在
	// 若不存在，则注册，新建用户
	var isNew = false
	user, err := model.GetUserByQQ(openID)
	if err == gorm.ErrRecordNotFound {
		// 注册新用户
		user, err = Register(request.Code, openID, sessionKey)
		if err != nil {
			handler.SendError(c, errno.ErrDatabase, nil, err.Error())
			return
		}
		isNew = true
	} else if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	// 生成 token
	token, err := token.GenerateToken(&token.TokenPayload{
		ID:           user.ID,
		QQOpenID:     openID,
		QQSessionKey: sessionKey,
	})
	if err != nil {
		handler.SendError(c, errno.ErrToken, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, LoginResponse{
		Token: token,
		IsNew: isNew,
	})
}

func Register(code, openID, sessionKey string) (*model.UserModel, error) {
	// 从 QQ 获取用户信息，昵称

	user := &model.UserModel{
		Username: "",
		Nickname: "",
		QQ:       openID,
	}
	if err := user.Create(); err != nil {
		return nil, err
	}

	return nil, nil
}
