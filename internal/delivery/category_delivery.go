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
	"net/http"

	"github.com/gin-gonic/gin"
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

// GetCategoryList - get a list of categories
func (delivery *Delivery) GetCategoryList(c *gin.Context) {
	delivery.logger.Debug("Enter in delivery GetCategoryList()")
	list, err := delivery.categoryHandlers.GetCategoryList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, list)
}
