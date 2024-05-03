package service

import (
	"context"
	"pilotkode/belajar-golang-restfull-api/model/web"
)

type CategoryService interface {
	// disarankan membuat mode untuk request dan response untuk menghindari menampilkan model
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
