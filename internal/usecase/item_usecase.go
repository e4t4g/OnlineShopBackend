package usecase

import (
	"OnlineShopBackend/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// CreateItem call database method and returns id of created item or error
func (storage *Storage) CreateItem(ctx context.Context, item *models.Item) (uuid.UUID, error) {
	storage.logger.Debug("Enter in usecase CreateItem()")
	id, err := storage.itemStore.CreateItem(ctx, item)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error on create item: %w", err)
	}
	return id, nil
}

// UpdateItem call database method to update item and returns error or nil
func (storage *Storage) UpdateItem(ctx context.Context, item *models.Item) error {
	storage.logger.Debug("Enter in usecase UpdateItem()")
	return storage.itemStore.UpdateItem(ctx, item)
}

// GetItem call database and returns *models.Item with given id or returns error
func (storage *Storage) GetItem(ctx context.Context, id uuid.UUID) (*models.Item, error) {
	storage.logger.Debug("Enter in usecase GetItem()")
	item, err := storage.itemStore.GetItem(ctx, id)
	if err != nil {
		return &models.Item{}, fmt.Errorf("error on get item: %w", err)
	}
	return item, nil
}

// ItemsList call database method and returns chan with all models.Item or error
func (storage *Storage) ItemsList(ctx context.Context, number int) (chan models.Item, error) {
	storage.logger.Debug("Enter in usecase ItemsList()")
	itemIncomingChan, err := storage.itemStore.ItemsList(ctx, number)
	if err != nil {
		return nil, err
	}
	itemOutChan := make(chan models.Item, 100)
	go func() {
		defer close(itemOutChan)
		for {
			select {
			case <-ctx.Done():
				return
			case item, ok := <-itemIncomingChan:
				if !ok {
					return
				}
				itemOutChan <- item
			}
		}
	}()
	return itemOutChan, nil
}

// SearchLine call database method and returns chan with all models.Item with given params or error
func (storage *Storage) SearchLine(ctx context.Context, param string, number int) (chan models.Item, error) {
	storage.logger.Debug("Enter in usecase SearchLine()")
	itemIncomingChan, err := storage.itemStore.SearchLine(ctx, param, number)
	if err != nil {
		return nil, err
	}
	itemOutChan := make(chan models.Item, 100)
	go func() {
		defer close(itemOutChan)
		for {
			select {
			case <-ctx.Done():
				return
			case item, ok := <-itemIncomingChan:
				if !ok {
					return
				}
				itemOutChan <- item
			}
		}
	}()
	return itemOutChan, nil
}
