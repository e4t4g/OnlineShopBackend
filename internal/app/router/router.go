/*
 * Backend for Online Shop
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package router

import (
	"OnlineShopBackend/internal/delivery"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/thinkerou/favicon"
	"go.uber.org/zap"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

type Router struct {
	router   *gin.Engine
	delivery *delivery.Delivery
	logger   *zap.Logger
}

// NewRouter returns a new router.
func NewRouter(delivery *delivery.Delivery, logger *zap.Logger) *Router {
	logger.Debug("Enter in NewRouter()")
	router := gin.Default()
	router.Use(favicon.New("./favicon.ico"))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"PUT", "GET", "POST"},
	}))
	routes := Routes{
		{
			"Index",
			http.MethodGet,
			"/",
			delivery.Index,
		},

		{
			"CreateCategory",
			http.MethodPost,
			"/categories",
			delivery.CreateCategory,
		},

		{
			"CreateItem",
			http.MethodPost,
			"/items",
			delivery.CreateItem,
		},

		{
			"GetItem",
			http.MethodGet,
			"/items/:itemID",
			delivery.GetItem,
		},

		{
			"UpdateItem",
			http.MethodPut,
			"/items/:itemID",
			delivery.UpdateItem,
		},

		{
			"UploadFile",
			http.MethodPost,
			"/items/:itemID/upload",
			delivery.UploadFile,
		},

		{
			"GetCart",
			http.MethodGet,
			"/cart/:userID",
			delivery.GetCart,
		},

		{
			"GetCategoryList",
			http.MethodGet,
			"/items/categories",
			delivery.GetCategoryList,
		},

		{
			"ItemsList",
			http.MethodGet,
			"/items",
			delivery.ItemsList,
		},

		{
			"SearchLine",
			http.MethodGet,
			"/search/:searchRequest",
			delivery.SearchLine,
		},

		{
			"CreateUser",
			http.MethodPost,
			"/user/create",
			delivery.CreateUser,
		},

		{
			"LoginUser",
			http.MethodPost,
			"/user/login",
			delivery.LoginUser,
		},

		{
			"LogoutUser",
			http.MethodPost,
			"/user/logout",
			delivery.LogoutUser,
		},
		{
			"LoginUserGoogle",
			http.MethodGet,
			"/user/login/google",
			delivery.LoginUserGoogle,
		},

		{
			"LoginUserYandex",
			http.MethodGet,
			"/user/login/yandex",
			delivery.LoginUserYandex,
		},

		{
			"callbackGoogle",
			http.MethodGet,
			"/user/callbackGoogle",
			delivery.CallbackGoogle,
		},

		{
			"callbackYandex",
			http.MethodPost,
			"/user/callbackYandex",
			delivery.CallbackYandex,
		},
	}

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return &Router{router: router, delivery: delivery, logger: logger}
}

func (router *Router) Run(port string) error {
	router.logger.Debug(fmt.Sprintf("Enter in router Run(), port: %s", port))
	return router.router.Run(port)
}
