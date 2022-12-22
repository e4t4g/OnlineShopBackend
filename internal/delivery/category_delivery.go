/*
 * Backend for Online Shop
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package delivery

import (
	"OnlineShopBackend/internal/handlers"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// CreateCategory - create a new category
func (delivery *Delivery) CreateCategory(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery CreateCategory()")
	ctx := c.Request.Context()
	var deliveryCategory handlers.Category
	if err := c.ShouldBindJSON(&deliveryCategory); err != nil {
		delivery.logger.Sugar().Errorf("error on bind json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delivery.logger.Sugar().Debugf("Binded struct: %v", deliveryCategory)
	if deliveryCategory.Name == "" && deliveryCategory.Description == "" {
		delivery.logger.Debug("Empty category")
		c.JSON(http.StatusBadRequest, "empty category is not correct")
		return
	}

	id, err := delivery.categoryHandlers.CreateCategory(ctx, deliveryCategory)
	if err != nil {
		delivery.logger.Sugar().Errorf("error on handlers.CreateCategory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": id.String()})
}

// UpdateCategory updating category
func (delivery *Delivery) UpdateCategory(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery UpdateCategory()")
	ctx := c.Request.Context()
	id := c.Param("categoryID")
	delivery.logger.Debug(fmt.Sprintf("Category id from request is %v", id))
	if id == "" {
		delivery.logger.Sugar().Error("empty id in request")
		c.JSON(http.StatusBadRequest, "empty id")
	}
	var deliveryCategory handlers.Category
	if err := c.ShouldBindJSON(&deliveryCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deliveryCategory.Id = id
	err := delivery.categoryHandlers.UpdateCategory(ctx, deliveryCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// UploadCategoryImage - upload an image
func (delivery *Delivery) UploadCategoryImage(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery UploadCategoryImage()")
	ctx := c.Request.Context()
	id := c.Param("categoryID")
	if id == "" {
		delivery.logger.Sugar().Error("error empty category id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty category id"})
		return
	}
	var name string
	contentType := c.ContentType()

	if contentType == "image/jpeg" {
		name = carbon.Now().ToShortDateTimeString() + ".jpeg"
	} else if contentType == "image/png" {
		name = carbon.Now().ToShortDateTimeString() + ".png"
	} else {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{})
		return
	}

	file, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{})
		return
	}

	delivery.logger.Info("Read id", zap.String("id", id))
	delivery.logger.Info("File len=", zap.Int32("len", int32(len(file))))
	path, err := delivery.filestorage.PutCategoryImage(id, name, file)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{})
		return
	}

	category, err := delivery.categoryHandlers.GetCategory(ctx, id)
	if err != nil {
		delivery.logger.Sugar().Errorf("error on get category: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	category.Image = path

	err = delivery.categoryHandlers.UpdateCategory(ctx, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "upload image success"})
}

// DeleteCategoryImage delete category image
func (delivery *Delivery) DeleteCategoryImage(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery DeleteCategoryImage()")
	var imageOptions ImageOptions
	err := c.Bind(&imageOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delivery.logger.Debug(fmt.Sprintf("image options is %v", imageOptions))

	if imageOptions.Id == "" || imageOptions.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("empty category id or file name")})
		return
	}
	err = delivery.filestorage.DeleteCategoryImage(imageOptions.Id, imageOptions.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx := c.Request.Context()
	category, err := delivery.categoryHandlers.GetCategory(ctx, imageOptions.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if strings.Contains(category.Image, imageOptions.Name) {
		category.Image = ""
	}
	err = delivery.categoryHandlers.UpdateCategory(ctx, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "delete image success"})
}

// GetCategory - get a specific category by id
func (delivery *Delivery) GetCategory(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery GetCategory()")
	id := c.Param("categoryID")
	delivery.logger.Debug(fmt.Sprintf("Category id from request is %v", id))
	if id == "" {
		delivery.logger.Sugar().Error("empty id from request")
		c.JSON(http.StatusBadRequest, "empty id is not correct")
		return
	}
	ctx := c.Request.Context()
	category, err := delivery.categoryHandlers.GetCategory(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// GetCategoryList - get a list of categories
func (delivery *Delivery) GetCategoryList(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery GetCategoryList()")
	list, err := delivery.categoryHandlers.GetCategoryList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, list)
}

// DeleteCategory deleted category by id
func (delivery *Delivery) DeleteCategory(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery DeleteCategory()")
	ctx := c.Request.Context()
	stringId := c.Param("categoryID")
	if stringId == "" {
		delivery.logger.Sugar().Error("recieved empty category id")
		c.JSON(http.StatusBadRequest, "empty id is not correct")
	}
	id, err := uuid.Parse(stringId)
	if err != nil {
		delivery.logger.Error(fmt.Sprintf("error on parsing uuid: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deletedCategory, err := delivery.categoryHandlers.GetCategory(ctx, stringId)
	if err != nil {
		delivery.logger.Debug(fmt.Sprintf("error on get deleted category: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	delivery.logger.Debug(fmt.Sprintf("deletedCategory: %v", deletedCategory))

	err = delivery.categoryHandlers.DeleteCategory(ctx, id)
	if err != nil {
		delivery.logger.Debug(fmt.Sprintf("error on delete category: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if deletedCategory.Image != "" {
		err = delivery.filestorage.DeleteCategoryImageById(stringId)
		if err != nil {
			delivery.logger.Error(fmt.Sprintf("error on delete category images: %v", err))
		}
	}

	quantity, err := delivery.itemHandlers.ItemsQuantity(ctx)
	if err != nil {
		delivery.logger.Error(fmt.Sprintf("error on get items quantity: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items, err := delivery.itemHandlers.GetItemsByCategory(ctx, deletedCategory.Name, 0, quantity)
	if err != nil {
		delivery.logger.Error(fmt.Sprintf("error on get items by category: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	noCategory, err := delivery.categoryHandlers.GetCategoryByName(ctx, "NoCategory")
	if err != nil {
		delivery.logger.Error(fmt.Sprintf("error on get category by name: %v", err))
		noCategory := handlers.Category{
			Name:        "NoCategory",
			Description: "Category for items from deleting categories",
		}
		noCategoryId, err := delivery.categoryHandlers.CreateCategory(ctx, noCategory)
		if err != nil {
			delivery.logger.Error(fmt.Sprintf("error on create no category: %v", err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		noCategory.Id = noCategoryId.String()
		for _, item := range items {
			item.Category = noCategory
			err := delivery.itemHandlers.UpdateItem(ctx, item)
			if err != nil {
				delivery.logger.Error(fmt.Sprintf("error on update item: %v", err))
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted success"})
		return
	}
	for _, item := range items {
		item.Category = noCategory
		err := delivery.itemHandlers.UpdateItem(ctx, item)
		if err != nil {
			delivery.logger.Error(fmt.Sprintf("error on update item: %v", err))
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted success"})
}
