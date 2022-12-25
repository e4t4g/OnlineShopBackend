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
	"go.uber.org/zap"
)

// CreateCategory - create a new category
func (delivery *Delivery) CreateCategory(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery CreateCategory()")
	ctx := c.Request.Context()
	var deliveryCategory handlers.Category
	if err := c.ShouldBindJSON(&deliveryCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if deliveryCategory.Name == "" && deliveryCategory.Description == "" {
		c.JSON(http.StatusBadRequest, "empty category is not correct")
	}

	id, err := delivery.categoryHandlers.CreateCategory(ctx, deliveryCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": id.String()})
}

// UpdateCategory updating category
func (delivery *Delivery) UpdateCategory(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery UpdateCategory()")
	ctx := c.Request.Context()
	id := c.Param("categoryID")
	delivery.logger.Debug(fmt.Sprintf("Category id from request is %v", id))
	var deliveryCategory handlers.Category
	if err := c.ShouldBindJSON(&deliveryCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deliveryCategory.Id = id
	err := delivery.categoryHandlers.UpdateCategory(ctx, deliveryCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.Image = path

	err = delivery.categoryHandlers.UpdateCategory(ctx, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.Contains(category.Image, imageOptions.Name) {
		category.Image = ""
	}
	err = delivery.categoryHandlers.UpdateCategory(ctx, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "delete image success"})
}

// GetCategory - get a specific category by id
func (delivery *Delivery) GetCategory(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery GetCategory()")
	id := c.Param("categoryID")
	delivery.logger.Debug(fmt.Sprintf("Category id from request is %v", id))
	ctx := c.Request.Context()
	category, err := delivery.categoryHandlers.GetCategory(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

//Path Information

// GetCategoryList godoc
// @Summary Get Category List
// @Description Get Category List
// @Accept  json
// @Produce  json
// @ Failure 400 {object}
// @Success 200 {object} list "ok"
// @Router /categories/list [get]
func (delivery *Delivery) GetCategoryList(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery GetCategoryList()")
	list, err := delivery.categoryHandlers.GetCategoryList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, list)
}
