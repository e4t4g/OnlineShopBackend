![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/GBteammates/OnlineShopBackend/test.yml)
![GitHub language count](https://img.shields.io/github/languages/count/GBteammates/OnlineShopBackend)
![GitHub top language](https://img.shields.io/github/languages/top/GBteammates/OnlineShopBackend)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/GBteammates/OnlineShopBackend)
![GitHub repo file count](https://img.shields.io/github/directory-file-count/GBteammates/OnlineShopBackend)
![GitHub repo size](https://img.shields.io/github/repo-size/GBteammates/OnlineShopBackend)
![GitHub](https://img.shields.io/github/license/GBteammates/OnlineShopBackend)
![GitHub contributors](https://img.shields.io/github/contributors/GBteammates/OnlineShopBackend)
![GitHub last commit](https://img.shields.io/github/last-commit/GBteammates/OnlineShopBackend)

<img src="./static/pictures/header.jpg">
<img align="right" width="50%" src="./static/pictures/favicon.ico">

## Description

Backend service for a small online store, written in the Golang programming language.
The project is an MVP in the form of a monolith. The functionality is implemented as a REST API. Due to its simplicity, wide functionality, and the fact that most contributors have worked with it before, the gin-gonic/gin framework was chosen as the router. In the future, as the load increases, a transition to more performant routers, such as fasthttp, is possible. For the same reasons, the PostgreSQL relational database management system (RDBMS) was chosen as the main database for storing service data. The current implementation provides sufficient capabilities of this RDBMS, and most contributors are familiar with this database, which has accelerated the development process. To improve the responsiveness of the service, a cache based on the Redis key-value database is implemented. The cache is used in the most frequent and heavy operations, such as retrieving a list of categories and products. For logging, the zap logger was chosen for its extensive functionality, high performance, and clear documentation. A combination of Prometheus for collecting metrics and Grafana for visualization is used for service monitoring. Whenever possible, the service development followed the principles of clean architecture. Unit tests have been implemented for most systems, as well as integration tests for the database.

## Scheme:
<img src="./static/pictures/2.jpg">

## Functionality

### For non-administrator users, including those not logged into the system, the following functionality is available:

- Creating/registering a new user (endpoint `/user/create`, method POST)
- Logging in as an existing user (endpoint `/user/login`, method POST)
- Logging in with a Google account (endpoint `/user/login/google`, method GET)
- Logging out (endpoint `/user/logout`, method GET)
- Viewing a list of all products (endpoint `/items/list`, method GET), with options to set offset, limit, and sorting parameters (sorting can be done by name or price, in ascending or descending order; the endpoint is supplemented with parameters like /items/list/?offset=0&limit=10&sortType=name&sortOrder=asc)
- Viewing a list of all product categories (endpoint `categories/list`, method GET)
- Viewing information about a specific item (endpoint `items/{itemID}`, method GET)
- Viewing information about a specific category (endpoint `categories/{categoryID}`, method GET)
- Viewing a list of products in a specific category (endpoint `/items/?param=categoryName&offset=20&limit=10&sort_type=name&sort_order=asc`), with options for sorting and limiting the number of results (sort_type can be name or price, sort_order can be asc or desc, method GET)
- Searching for a specific item (endpoint `/items/?param=searchRequest&offset=20&limit=10&sort_type=name&sort_order=asc`), with options for sorting and limiting the number of results (sort_type can be name or price, sort_order can be asc or desc, method GET)
- Viewing information about the total quantity of products (endpoint `/items/quantity`, method GET)
- Viewing information about the quantity of products in a specific category (endpoint `/items/quantityCat/{categoryName}`, method GET)
- Viewing information about the quantity of products in search results (endpoint `/items/quantitySearch/{searchRequest}`, method GET)

### For logged-in users without administrator rights:

- Viewing user profile (endpoint `/user/profile`, method GET)
- Changing user profile information (endpoint `/user/profile/edit`, method PUT)
- Adding an item to the favorites list (endpoint `/items/addFavItem/`, method POST)
- Viewing items from the favorites list (endpoint `/items/favList?param=userIDt&offset=20&limit=10&sort_type=name&sort_order=asc` (sort_type can be name or price, sort_order can be asc or desc), method GET)
- Removing an item from the favorites list (endpoint `/items/deleteFav/{userID}/{itemID}`, method DELETE)
- A cart is created when a user logs in, but there is also the option to manually create a cart (endpoint `/cart/create/{userID}`, method POST)
- Adding an item to the cart (endpoint `/cart/addItem`, method PUT)
- Removing an item from the cart (endpoint `/cart/delete/{cartID}/{itemID}`, method DELETE)
- Viewing a cart by cart ID (endpoint `/cart/{cartID}`, method GET)
- Viewing a cart by user ID (endpoint `/cart/byUser/{userID}`, method GET)
- Deleting a cart (endpoint `/cart/delete/{cartID}`, method DELETE)
- Creating an order (endpoint /order/create, method POST)
- Viewing order information (endpoint `/order/{orderID}`, method GET)
- Viewing user's order information (endpoint `/order/list/{userID}`, method GET)
- Changing the delivery address in an order (endpoint `/order/changeaddress`, method PATCH)

### For users logged in with administrator rights:

- Changing user role/permissions (endpoint `/user/role/update`, method PUT)
- Creating a new role/permissions (endpoint `/user/createRights`, method POST)
- Viewing a list of roles/permissions (endpoint `/user/rights/list`, method GET)
- Creating a new product category (endpoint `/categories/create`, method POST)
- Modifying an existing product category (endpoint `/categories/{categoryID}`, method PUT)
- Adding an image to an existing category (endpoint `/categories/image/upload/{categoryID}`, method POST)
- Deleting an image from a category (endpoint `/categories/image/delete`, method DELETE)
- Deleting a category (endpoint `/categories/delete/{categoryID}`, method DELETE)
- Creating a new product (endpoint `/items/create`, method POST)
- Modifying an existing product (endpoint `/items/update`, method PUT)
- Adding an image to an existing product (endpoint `/items/image/upload/:itemID`, method POST)
- Deleting an image of a product (endpoint `/items/image/delete?id=25f32441-587a-452d-af8c-b3876ae29d45&name=20221209194557.jpeg`, method DELETE)
- Deleting a product (endpoint `/items/delete/{itemID}`, method DELETE)
- Deleting an order (endpoint `/order/delete/{orderID}`, method DELETE)
- Changing the status of an order (endpoint `/order/changestatus`, method PATCH)
- Getting a list of images for categories and products (endpoint `/images/list`, method GET)

Authentication on the service is done using JWT tokens. The cache is created upon service startup, and user and administrator rights are also created during startup. A user with administrator rights is created as well. The data for creating the administrator is specified through environment variables. By default, these are admin@mail.ru and 12345678. The service is designed to gracefully shut down when necessary.

Thanks to the contributions of [ZavNatalia](https://github.com/ZavNatalia) almost all the functionality of this service can be conveniently tested using a [graphical interface](https://github.com/ZavNatalia/gb-store/tree/feature/new-api)

The service is documented using the library [swaggo](https://github.com/swaggo/swag).

### Minimum requirements for running the service:
- Installed [docker](https://docs.docker.com/engine/install/)
- Installed [docker-compose](https://docs.docker.com/compose/install/)
- Installed [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git?source=post_page---------------------------)
- Installed [make](https://habr.com/ru/post/211751/)
- In order to run the tests, it is necessary to install the programming language. [Golang](https://go.dev/doc/install)

### Deployment instructions:
1. Navigate to the directory where you plan to place the source files.
2. Clone the repository by running the following command in the command line: 
`git clone https://github.com/GBteammates/OnlineShopBackend.git`
3. To run the tests, execute the following command in the command line: `make test`
4. To launch the service with a graphical interface on Windows, run the following command in the command line: 
`make up-win`. After the installation and build process is completed, open any preferred browser and enter http://localhost:3000 in the address bar. To view the monitoring data, enter http://localhost:3001 in the browser's address bar, which will take you to the Grafana login page. The default login is `admin` and the password is `admin`. You will be prompted to change the password upon first login, which you can either change or skip.
5. To launch the service with a graphical interface on Ubuntu, run the following command in the command line: `make up-lin`. After the installation and build process is completed, open any preferred browser and enter http://localhost:3000 in the address bar. To view the monitoring data, enter http://localhost:3001 in the browser's address bar, which will take you to the Grafana login page. The default login is `admin` and the password is `admin`. You will be prompted to change the password upon first login, which you can either change or skip.
6. To launch the service without a graphical interface, run the following command in the command line: `make up`. If the make utility is not installed, use the following command instead: `docker-compose up -d`
7. Once the download, creation, and launch of the necessary containers are completed, you can open the browser and enter http://localhost:8000/docs/swagger/index.html in the address bar. This will open the Swagger documentation page with a graphical interface where you can test the core functions of the service. You can also use the [postman](https://www.postman.com/downloads/) or the [curl](https://curl.se/) commands for testing.

<img src="./static/pictures/swagger.jpg">

To stop the service, type `make down` or `docker-compose down`.

The service is running at: http://cozydragon.online/

The service is distributed under the [MIT](https://mit-license.org/).

# Video:

<img src="./static/pictures/video.gif">


















