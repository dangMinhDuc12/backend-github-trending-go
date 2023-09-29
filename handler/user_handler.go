package handler

import (
	"example/backend-github-trending/model"
	"example/backend-github-trending/model/req"
	"example/backend-github-trending/repository"
	"example/backend-github-trending/security"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)


type UserHandler struct {
	UserRepo repository.UserRepo
}


func (u *UserHandler) HandleSignup(c echo.Context) error {
	req := req.ReqSignUp{}

	//----- Start Bind user request to req variable -----//
	if err := c.Bind(&req); err != nil {
		log.Error(err)

		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

		//----- End Bind user request to req variable -----//


	//----- Start validate request -----//
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())

		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	//----- End validate request -----//


	//----- Start format data for insert db -----//
	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()

	if err != nil {
		log.Error(err)

		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message: err.Error(),
			Data: nil,
		})
	}

	user := model.User {
		UserId: userId.String(),
		FullName: req.FullName,
		Email: req.Email,
		Password: hash,
		Role: role,
	}

	//----- End format data for insert db -----//


	//----- Start Handle Insert db -----//



	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)


	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message: err.Error(),
			Data: nil,
		})
	}


	//----- End Handle Insert db -----//


	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "success",
		Data: user,
	})
}

func (u *UserHandler) HandleGetListUser(c echo.Context) error {
	//----- Start query db -----//
	users, _ := u.UserRepo.GetListUser(c.Request().Context())

	//----- End query db -----//


	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "success",
		Data: users,
	})

}

func (u *UserHandler) HandleSignin(c echo.Context) error {
	req := req.ReqSignIn{}

	//----- Start Bind user request to req variable -----//
	if err := c.Bind(&req); err != nil {
		log.Error(err)

		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	//----- End Bind user request to req variable -----//

	//----- Start validate request -----//
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())

		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	//----- End validate request -----//


	//----- Start check user exist -----//

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message: err.Error(),
			Data: nil,
		})
	}

	//----- End check user exist -----//

	//----- Start check password -----//

	isSamePassword := security.ComparePasswords(user.Password, []byte(req.Password))

	if !isSamePassword {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message: "Mật khẩu không đúng, đăng nhập thất bại",
			Data: nil,
		})
	}


	//----- End check password -----//


		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message: "Đăng nhập thành công",
			Data: user,
		})

}