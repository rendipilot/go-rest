package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"pilotkode/belajar-golang-restfull-api/app"
	"pilotkode/belajar-golang-restfull-api/controller"
	"pilotkode/belajar-golang-restfull-api/exception"
	"pilotkode/belajar-golang-restfull-api/helper"
	"pilotkode/belajar-golang-restfull-api/middleware"
	"pilotkode/belajar-golang-restfull-api/repository"
	"pilotkode/belajar-golang-restfull-api/service"
)

func main() {

	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	// server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
