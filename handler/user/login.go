package user

import (
	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"
	"todolist_backend/pkg/token"
	"todolist_backend/service"
	"todolist_backend/util"

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
// @Router /login [post]
func Login(c *gin.Context) {
	request := LoginRequest{}
	if err := c.BindJSON(&request); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 根据 code 从 QQ 获取 openID 和 sessionKey
	openID, sessionKey, err := service.CodeToSession(request.Code)
	if err != nil {
		handler.SendResponse(c, err, "登录失败")
		return
	}

	// 检查用户是否存在
	// 若不存在，则注册，新建用户
	var isNew = false
	user, err := model.GetUserByQQ(openID)
	if err == gorm.ErrRecordNotFound {
		// 创建新用户
		user, err = service.CreateNewAccount(openID)
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
		Expired:      util.GetTokenExpiredTime(),
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
