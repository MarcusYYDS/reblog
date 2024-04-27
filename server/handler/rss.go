package handler

import (
	"reblog/internal/query"
	"reblog/internal/rss"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		获取Rss
//	@Description	获取包含所有文章的RSS
//	@Tags			Rss
//	@Produce		xml
//	@Success		200	"RSS Feed"
//	@Failure		500	{object}	common.Resp	"服务器错误"
//	@Router			/rss [get]
func Rss(router fiber.Router) {
	router.Get("/rss", func(c fiber.Ctx) error {
		a := query.Article

		articles, err := a.Find()

		if err != nil {
			return common.RespServerError(c, err)
		}

		rssString, err := rss.GenerateRSS(articles)

		if err != nil {
			return common.RespServerError(c, err)
		}

		return c.XML(rssString)
	})
}