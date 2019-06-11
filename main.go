package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {

	// engine := gin.Default()
	// engine.Any("/", WebRoot)
	// engine.Run(":9205")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "hello %s", name)
	// })

	// router.GET("/someGet", middleware1, middleware2, handler)
	// router.Run(":9025")

	// router.GET("/welcome", func(c *gin.Context) {

	// 	firstname := c.DefaultQuery("firstname", "xuchen")
	// 	lastname := c.Query("lastname")
	// 	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	// })

	// router.POST("/form_post", func(c *gin.Context) {
	// 	message := c.PostForm("message")
	// 	nick := c.DefaultPostForm("nick", "anonymous")

	// 	c.JSON(200, gin.H{
	// 		"status":  "posted",
	// 		"message": message,
	// 		"nick":    nick,
	// 	})
	// })

	// router.Run(":9025")

	// 设置文件上传大小 router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 处理单一的文件上传
	// router.POST("/upload", func(c *gin.Context) {
	// 	// 拿到这个文件
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)
	// 	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// })

	// // 处理多个文件的上传
	// router.POST("/uploads", func(c *gin.Context) {
	// 	form, _ := c.MultipartForm()
	// 	// 拿到集合
	// 	files := form.File["upload[]"]
	// 	for _, file := range files {
	// 		log.Println(file.Filename)
	// 	}
	// 	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	// })
	// router.Run(":9025")

	// router.POST("/post", func(c *gin.Context) {
	// 	d, err := c.GetRawData()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	log.Println(string(d))
	// 	c.String(200, "ok")
	// })

	// router.Run(":9025")

	// router.Any("/testing", startPage)
	// router.Run(":9025")

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON; err == nil {
			if json.User == "xuchen" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
	})

	// POST 到这个路由一个 Form 表单 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 验证数据并绑定
		if err := c.ShouldBind(&form); err == nil {
			if form.User == "xuchen" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
	})

	// gin.H 本质是 map[string]interface{}
	router.GET("/someJSON", func(c *gin.Context) {
		// 会输出头格式为 application/json; charset=UTF-8 的 json 字符串
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/moreJSON", func(c *gin.Context) {
		// 直接使用结构体定义
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 会输出  {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(c *gin.Context) {
		// 会输出头格式为 text/xml; charset=UTF-8 的 xml 字符串
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/someYAML", func(c *gin.Context) {
		// 会输出头格式为 text/yaml; charset=UTF-8 的 yaml 字符串
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.Run()
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "success")
}

func handler(c *gin.Context) {
	log.Println("exec handler")
}

func middleware1(c *gin.Context) {
	log.Println("exec middleware1")

	c.Next()
}

func middleware2(c *gin.Context) {
	log.Println("arrive at middleware2")
	// 执行该中间件之前，先跳到流程的下一个方法
	c.Next()
	// 流程中的其他逻辑已经执行完了
	log.Println("exec middleware2")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
