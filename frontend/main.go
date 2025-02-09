package main

import (
	"net/http"

	"Platform/backend/config"
	"Platform/backend/models"
	"Platform/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	// Run migrations
	config.DB.AutoMigrate(&models.User{}, &models.Business{}, &models.Appointment{})
	println("✅ Database migrations completed!")

	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLFiles("frontend/templates/register.html",
		"frontend/templates/login.html",
		"frontend/templates/sto-register.html",
		"frontend/templates/appointment.html",
		"frontend/templates/userInfo.html",
		"frontend/templates/business.html",
		"frontend/templates/business_profile.html",
		"frontend/templates/business_appointment.html",
		"frontend/templates/index.html",
	)

	router.LoadHTMLGlob("templates/*.html")
	// Serve static files
	router.Static("/static", "./frontend/static")

	routes.SetupRoutes(router)

	// Define routes
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.GET("/sto", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sto-register.html", gin.H{})
	})

	router.GET("/appointment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "appointment.html", gin.H{})
	})

	router.GET("/userInfo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "userInfo.html", gin.H{})
	})

	router.GET("/business", func(c *gin.Context) {
		c.HTML(http.StatusOK, "business.html", gin.H{})
	})

	router.GET("/business_profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "business_profile.html", gin.H{})
	})

	router.GET("/business_appointment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "business_appointment.html", gin.H{})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Start server
	router.Run(":8080")
}
