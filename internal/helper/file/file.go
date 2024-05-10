package file

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

type ResizeOption struct {
	Width    int
	Height   int
	IsResize bool
}

func Save(e echo.Context, key string, opt ResizeOption) (string, error) {
	f, err := e.FormFile(key)
	if err != nil {
		log.Error().Err(err).Msg("failed to get form file")
		return "", pkgerr.ErrWithStackTrace(err)
	}

	src, err := f.Open()
	defer src.Close()
	if err != nil {
		log.Error().Err(err).Msg("failed to open file")
		return "", pkgerr.ErrWithStackTrace(err)
	}

	img, ext, err := image.Decode(src)
	if err != nil {
		log.Error().Err(err).Msg("failed to decode image")
		return "", pkgerr.ErrWithStackTrace(customerror.ErrorImageNotSupported)
	}

	if opt.IsResize {
		img = transform.Resize(img, opt.Width, opt.Height, transform.Linear)
	}

	var enc imgio.Encoder
	switch ext {
	case "png":
		enc = imgio.PNGEncoder()
	default:
		enc = imgio.JPEGEncoder(90)
	}

	newFilePath := fmt.Sprintf("%s/%s.%s", _const.PrefixUploadPath, uuid.NewString(), ext)
	if err := imgio.Save(newFilePath, img, enc); err != nil {
		log.Error().Err(err).Msg("failed to save resized image")
		return "", pkgerr.ErrWithStackTrace(err)
	}

	return newFilePath, nil
}
