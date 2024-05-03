package repository

import (
	"context"
	"database/sql"
	"errors"
	"pilotkode/belajar-golang-restfull-api/helper"
	"pilotkode/belajar-golang-restfull-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me

	SQL := "insert into category(name) values (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	category.Id = int(id)

	return category
}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me

	SQL := "update category set name = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)

	return category
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	//TODO implement me
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfErr(err)
}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	//TODO implement me

	SQL := "select id,name from category where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfErr(err)

	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	//TODO implement me

	SQL := "select id,name from category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}

		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfErr(err)
		categories = append(categories, category)
	}

	return categories
}
