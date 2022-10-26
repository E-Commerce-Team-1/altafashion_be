package delivery

import (
	"altafashion_be/feature/products/domain"
	"altafashion_be/utils/jwt"
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := productHandler{
		srv: srv,
	}

	/* Routing endpoints: */
	e.POST("/products", handler.AddProduct())
}

func (ph *productHandler) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddProductFormat
		input.Name = c.FormValue("name")
		input.Description = c.FormValue("description")
		input.Category = c.FormValue("category")
		cnvQty, _ := strconv.Atoi(c.FormValue("qty"))
		input.Qty = uint(cnvQty)
		cnvPrice, _ := strconv.Atoi(c.FormValue("price"))
		input.Price = cnvPrice
		input.UserID = jwt.ExtractTokenProd(c)
		file, err := c.FormFile("image")
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
			input.Image = "https://altafashion.s3.ap-southeast-1.amazonaws.com/myfiles/" + temp + strings.ReplaceAll(file.Filename, " ", "+")
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
		res, err := ph.srv.AddProduct(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Berhasil add product", ToResponse(res, "add")))
	}
}

func (ph *productHandler) EditProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input EditProductFormat
		input.Name = c.FormValue("name")
		input.Description = c.FormValue("description")
		input.Category = c.FormValue("category")
		cnvQty, _ := strconv.Atoi(c.FormValue("qty"))
		input.Qty = uint(cnvQty)
		cnvPrice, _ := strconv.Atoi(c.FormValue("price"))
		input.Price = cnvPrice
		file, err := c.FormFile("image")
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
			input.Image = "https://altafashion.s3.ap-southeast-1.amazonaws.com/myfiles/" + temp + strings.ReplaceAll(file.Filename, " ", "+")
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

		productID := jwt.ExtractTokenProd(c)
		res, err := ph.srv.EditProduct(cnv, productID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, ToResponse(res, "edit"))
	}
}

func (ph *productHandler) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			return c.JSON(http.StatusInternalServerError, "cant convert id")
		}

		err := ph.srv.Destroy(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, SuccessDeleteResponse("success delete data."))
	}
}
