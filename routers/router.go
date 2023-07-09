package routers

import (
	"github.com/gin-gonic/gin"
	"lcrm2/controllers"
	"lcrm2/middlewares"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")

	{
		authGroup := apiGroup.Group("/auth/")
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
		//authGroup.Use(middlewares.JwtAuthMiddleware())
		authGroup.GET("/user", controllers.GetUser)
		authGroup.PUT("/user/:id", controllers.UpdateUserById)
		authGroup.POST("/logout", controllers.Logout)
	}

	{
		protected := apiGroup.Group("/admin")
		protected.Use(middlewares.JwtAuthMiddleware())
		//protected.Use(middlewares.JwtRoleMiddleware(1))
		protected.GET("/user", controllers.CurrentUser)
	}

	{
		teacher := apiGroup.Group("/teacher")
		teacher.Use(middlewares.JwtAuthMiddleware())
		teacher.Use(middlewares.JwtRoleMiddleware([]uint{2, 3}))
		teacher.GET("/users", controllers.GetUsers)

		//protected.GET("/project", controllers.GetUsers)
		teacher.GET("/project", controllers.GetAllProject)
		teacher.POST("/project", controllers.CreateProject)

		teacher.GET("/project/:id", controllers.GetProjectById)
		teacher.PUT("/project/:id", controllers.UpdateProjectById)
		teacher.DELETE("/project/:id", controllers.DeleteProjectById)

		teacher.POST("/project/member", controllers.AddMember)
		teacher.DELETE("/project/member", controllers.RemoveMember)

	}

	project := apiGroup.Group("/project")
	project.Use(middlewares.JwtRoleMiddleware([]uint{1, 2}))

	// Роуты для тасков
	boardGroup := project.Group("/board")
	{
		boardGroup.POST("/", controllers.CreateBoard)
		boardGroup.GET("/:id", controllers.GetBoardById)
		boardGroup.POST("/task/", controllers.CreateTask)
		boardGroup.DELETE("/task/:id", controllers.DeleteTaskById)
		boardGroup.PATCH("/task/:id", controllers.UpdateTaskById)

		boardGroup.GET("/users", controllers.GetUsers)

	}

	// Роуты для комментариев
	commentGroup := project.Group("/comments")

	{
		commentGroup.POST("/", controllers.CreateComment)
		commentGroup.GET("/:id", controllers.GetCommentByID)
		commentGroup.GET("/:id/replies", controllers.GetRepliesByCommentID)
	}

	// Роуты для ответов
	replyGroup := project.Group("/replies")
	{
		replyGroup.POST("/", controllers.CreateReply)
		replyGroup.GET("/:id", controllers.GetReplyByID)
	}

	{
		gitApi := apiGroup.Group("/git")
		gitApi.Use(middlewares.JwtAuthMiddleware())
		gitApi.GET("/test", controllers.Test)
		gitApi.GET("/history", controllers.GetCommitHistory)

	}
	return router
}
