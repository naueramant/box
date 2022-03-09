package routes

import (
	"box/internal/authentication"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PostLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PostLoginResponse struct {
	Token        string                `json:"token"`
	RefreshToken string                `json:"refresh_token"`
	User         PostLoginResponseUser `json:"user"`
}

type PostLoginResponseUser struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Container) PostLogin(ctx echo.Context, l *logrus.Entry) error {
	var req PostLoginRequest
	if err := ctx.Bind(&req); err != nil {
		l.WithError(err).Debug("Failed to bind PostLoginRequest")

		return echo.ErrBadRequest
	}

	l = l.WithField("username", req.Username)

	if err := ctx.Validate(&req); err != nil {
		l.WithError(err).Debug("Validation failed for PostLoginRequest")

		return echo.ErrBadRequest
	}

	token, user, err := c.Authentication.Login(
		ctx.Request().Context(),
		req.Username,
		req.Password,
	)
	if err != nil {
		if errors.Is(err, authentication.ErrFailedLoginAttempt) {
			l.WithError(err).Warn("Failed login attempt")
		} else {
			l.WithError(err).Error("Failed to login due to internal error")
		}

		return echo.ErrUnauthorized
	}

	l.Info("Successful login")

	return ctx.JSON(http.StatusOK, PostLoginResponse{
		Token:        token.JWT,
		RefreshToken: token.RefreshToken,
		User: PostLoginResponseUser{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

func (c *Container) PostRefresh(ctx echo.Context, l *logrus.Entry) error {
	return nil
}
