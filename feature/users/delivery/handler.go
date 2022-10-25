package delivery

import (
	"altafashion_be/config"
	"altafashion_be/feature/users/domain"
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

type userHandler struct {
	srv domain.Service
}

func InitJWT(c *config.AppConfig) {
	key = c.JWSecret
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}

	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
	e.GET("/users", handler.MyProfile(), middleware.JWT([]byte(key)))
	e.GET("/users/:fullname", handler.ShowByEmail(), middleware.JWT([]byte(key)))
	e.PUT("/users", handler.UpdateProfile(), middleware.JWT([]byte(key)))
	e.DELETE("/users", handler.Deactivate(), middleware.JWT([]byte(key)))
}

func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponses("berhasil register", ToResponse(res, "reg", "")))
	}

}

func (us *userHandler) Login() echo.HandlerFunc {
	//autentikasi user login
	return func(c echo.Context) error {
		var resQry LoginFormat
		if err := c.Bind(&resQry); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses("cannot bind input"))
		}

		cnv := ToDomain(resQry)
		res, token, err := us.srv.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponses("berhasil login", ToResponse(res, "login", token)))
	}
}

func (us *userHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponses(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		var input UpdateUserFormat
		input.Fullname = c.FormValue("fullname")
		input.Email = c.FormValue("email")
		input.Password = c.FormValue("password")
		input.Location = c.FormValue("location")

		file, err := c.FormFile("user_picture")
		if err == nil {
			src, err := file.Open()
			if err != nil {
				log.Println(err)
			}
			defer src.Close()

			s3Config := &aws.Config{
				Region:      aws.String(os.Getenv("AWS_REGION")),
				Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_USER"), os.Getenv("AWS_KEY"), ""),
			}
			temp := time.Now().Format("02 Jan 06 15:04")
			input.Profile = "https://altafashion.s3.ap-southeast-1.amazonaws.com/myfiles/" + temp + strings.ReplaceAll(file.Filename, " ", "+")
			s3Session := session.New(s3Config)

			uploader := s3manager.NewUploader(s3Session)
			inputData := &s3manager.UploadInput{
				Bucket: aws.String("altafashion"),                     // bucket's name
				Key:    aws.String("myfiles/" + temp + file.Filename), // files destination location
				Body:   src,                                           // content of the file

			}
			_, _ = uploader.UploadWithContext(context.Background(), inputData)
		}
		cnv := ToDomain(input)

		//userId := us.srv.ExtractToken(c)

		res, err := us.srv.UpdateProfile(cnv, c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, ToResponse(res, "update", ""))

	}
}

func (us *userHandler) Deactivate() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponses(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		if _, err := us.srv.Deactivate(c); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessDeleteResponses("Success Delete Data"))
	}
}

func (us *userHandler) ShowByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authorized request atau tidak dgn token
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponses(err.Error()))
		} else {
			log.Println("Authorized request.")
		}
		paramName := c.Param("Email")
		log.Println("Email :", paramName)
		// ID, err := strconv.Atoi(c.Param("email"))
		res, err := us.srv.ShowByEmail(paramName)
		log.Println("Email", paramName)
		log.Println("res", res)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponses("sucses get userBy email", ToResponse(res, "get", "")))
	}
}

func (us *userHandler) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := us.srv.IsAuthorized(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, FailResponses(err.Error()))
		} else {
			log.Println("Authorized request.")
		}

		res, err := us.srv.MyProfile(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponses(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponses("sucses get userBy fullname", ToResponse(res, "get", "")))

	}
}
