package controller

import (
	"api/ent"
	"api/ent/horse"
	"api/ent/user"
	"api/internal/dto/request"
	"api/internal/dto/response"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	db *ent.Client
}

func NewUserController(db *ent.Client) *UserController {
	return &UserController{
		db: db,
	}
}

func (c UserController) GetUserByUsername(ctx echo.Context) error {

	var body request.UserDTO
	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Body error")
	}

	_user, err := c.db.User.Query().Where(user.Username(body.Username)).Only(context.TODO())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return ctx.JSON(http.StatusFound, response.UserDTO{
		Username: _user.Username,
		Email:    _user.Email,
	})
}

func (c UserController) CreateUser(ctx echo.Context) error {

	var body request.UserCreateDTO
	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Body error")
	}

	_user, err := c.db.User.Create().SetEmail(body.Email).SetUsername(body.Email).SetPassword(body.Password).SetVerified(true).Save(context.TODO())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Creating user error")
	}

	return ctx.JSON(http.StatusFound, response.UserDTO{
		Username: _user.Username,
		Email:    _user.Email,
	})
}

func (c UserController) GetHorsesByUserID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Body error")
	}

	_horses, err := c.db.Horse.Query().Where(horse.HasOwnerWith(user.ID(id))).All(context.TODO())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Horses not found")
	}

	var horsesDTO []response.HorseDTO
	for _, v := range _horses {
		horsesDTO = append(horsesDTO, response.HorseDTO{Name: v.Name})
	}

	return ctx.JSON(http.StatusFound, horsesDTO)
}
