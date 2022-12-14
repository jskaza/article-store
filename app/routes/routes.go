package routes

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jskaza/article-store/app/controllers"
	"github.com/jskaza/article-store/app/models"
	"github.com/jskaza/article-store/app/utils"
)

func SetupRoutes(views, css, js, favicon string) {
	router := gin.Default()
	router.LoadHTMLGlob(views)
	// engine := html.New("./app/ui/views", ".html")

	// app := fiber.New(fiber.Config{
	// 	Views: engine,
	// })

	router.Static("./css", css)
	router.Static("./js", js)
	router.StaticFile("./favicon.ico", favicon)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("category/:category", func(c *gin.Context) {
		papers, _ := controllers.GetPapers(strings.ToLower(c.Param("category")))
		c.HTML(http.StatusOK, "list.html", papers)
	})

	router.GET("category/:category/paper/:id", func(c *gin.Context) {
		paper, _ := controllers.GetPaper(c.Param("id"))
		c.HTML(http.StatusOK, "paper.html", paper)
	})

	router.GET("about/mission", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mission.html", gin.H{})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("uploaded-paper")
		extension := filepath.Ext(file.Filename)
		uuid := uuid.New().String()
		c.SaveUploadedFile(file, uuid+extension)
		utils.ParsePaper(uuid, extension)
		os.Remove(uuid + extension)
		content, _ := ioutil.ReadFile(uuid + ".html")
		category := strings.ToLower(c.PostForm("category"))
		paper := models.Paper{
			// Author:   controllers.GetName(), // hard code for now,
			Title:    c.PostForm("title"),
			Author:   c.PostForm("author"),
			Category: category,
			Abstract: c.PostForm("abstract"),
			Content:  string(content)}
		os.Remove(uuid + ".html")
		controllers.InsertPaper(paper)

		// need middleware to redirect
		// id, _ := controllers.InsertPaper(paper)
		// c.Redirect(http.StatusMovedPermanently, "/category/"+category+"/paper/"+id)
	})

	router.Run(":8080")

}
