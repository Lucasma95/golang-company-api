package database

import (
	"context"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func save(ctx context.Context, client *gorm.DB, entity interface{}, fullSaveAsociations ...bool) error {
	if len(fullSaveAsociations) > 0 {
		if fullSaveAsociations[0] {
			client = client.Session(&gorm.Session{FullSaveAssociations: true})
		}
	}

	if err := client.WithContext(ctx).Save(entity).Error; err != nil {
		log.Println("error saving entity")
		return err
	}
	return nil
}

func create(ctx context.Context, client *gorm.DB, entity interface{}, columnsToOmit ...string) error {
	if err := client.WithContext(ctx).Omit(columnsToOmit...).Create(entity).Error; err != nil {
		if validateIfEntityIsDuplicated(err) {
			log.Println("duplicated entity.")
			return err
		} else {
			log.Println("error on save entity")
			return err
		}
	}
	return nil
}

func validateIfEntityIsDuplicated(err error) bool {
	return strings.Contains(err.(*pgconn.PgError).Message, "duplicate key value violates unique constraint")
}

func delete(ctx context.Context, client *gorm.DB, entity interface{}) error {
	if err := client.WithContext(ctx).Delete(entity).Error; err != nil {
		log.Println("error deleting entity")
		return err
	}
	return nil
}

func getByID[T any](ctx context.Context, client *gorm.DB, id any, preload ...string) (*T, error) {
	return getEntityByParams[T](ctx, client, map[string]any{"id": id}, preload...)
}

func getEntityByParams[T any](ctx context.Context, client *gorm.DB, params map[string]any, preload ...string) (*T, error) {
	for _, v := range preload {
		client = client.Preload(v)
	}

	entity := new(T)
	if err := client.WithContext(ctx).Where(params).First(&entity).Error; err != nil {
		log.Println("error getting entity by params")
		return nil, err
	}

	return entity, nil
}

func getEntitiesByParams[T any](ctx context.Context, client *gorm.DB, params map[string]any, preload ...string) ([]T, error) {
	for _, v := range preload {
		client = client.Preload(v)
	}

	entities := []T{}
	if err := client.WithContext(ctx).Where(params).Find(&entities).Error; err != nil {

		log.Println("error getting entities by params")
		return nil, err
	}

	if len(entities) == 0 {
		return nil, nil
	}

	return entities, nil
}
