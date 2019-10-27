// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package api

import (
	"net/http"

	"sewadong-service/pkg/models"
	"sewadong-service/pkg/service"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

const (
	getAllUserEndpoint               = ""
	getUserByEmailOrUsernameEndpoint = "/identifier/{identifier}"
	getUserByUserIDEndpoint          = "/{userid}"
	createUserEndpoint               = ""
	updateUserEndpoint               = ""
	loginEndpoint                    = "/login"

	getRoleByRoleIDEndpoint = "/{roleid}"
	getAllRoleEndpoint      = ""

	getAllProductEndpoint          = ""
	getProductByProductIDEndpoint  = "/{productid}"
	getProductByCategoryIDEndpoint = "/category/{categoryid}"
	getProductLikeNameEndpoint     = "/identifier/{identifier}"
	createProductEndpoint          = ""
	updateProductEndpoint          = ""
	deleteProductEndpoint          = ""

	getAllCategoryEndpoint          = ""
	getCategoryByCategoryIDEndpoint = "/{categoryid}"
	createCategoryEndpoint          = ""
	updateCategoryEndpoint          = ""
	deleteCategoryEndpoint          = ""
)

// AddRouteUser is setup restful route
func AddRouteUser(service *service.Service, basePath string) *restful.WebService {
	webService := new(restful.WebService)
	webService.Path(basePath + "/user").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		ApiVersion("0.0.1").
		Doc("User")

	tags := []string{"User"}

	webService.Route(webService.POST(createUserEndpoint).To(service.CreateUser).
		Reads(models.UserRequest{}).
		Notes("Create user \n \n Create user").
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), models.CreateUserResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Create User"))

	webService.Route(webService.GET(getAllUserEndpoint).To(service.GetAllUser).
		Notes("Get all user \n \n Retrieves all user").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetUsersResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get All User"))

	webService.Route(webService.GET(getUserByUserIDEndpoint).To(service.GetUserByID).
		Param(webService.PathParameter("userid", "the id of user")).
		Notes("Get user by user id \n \n Retrieves user").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetUserResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get User by UserID"))

	webService.Route(webService.GET(getUserByEmailOrUsernameEndpoint).To(service.GetUserByEmailOrUsername).
		Param(webService.PathParameter("identifier", "the user email or username")).
		Notes("Get user by email or username \n \n Retrieves user").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetUserResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get User by Email or Username"))

	webService.Route(webService.POST(loginEndpoint).To(service.Login).
		Reads(models.LoginRequest{}).
		Notes("Get user by email or username and password \n \n Retrieves user").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.LoginResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get User by Email or Username and Password"))

	webService.Route(webService.PUT(updateUserEndpoint).To(service.UpdateUser).
		Reads(models.UserRequest{}).
		Notes("Update user \n \n Update user").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.UpdateUserResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Update User"))

	return webService
}

// AddRouteRole is setup restful route
func AddRouteRole(service *service.Service, basePath string) *restful.WebService {
	webService := new(restful.WebService)
	webService.Path(basePath + "/role").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		ApiVersion("0.0.1").
		Doc("Role")

	tags := []string{"Role"}

	webService.Route(webService.GET(getAllRoleEndpoint).To(service.GetAllRole).
		Notes("Get all role \n \n Retrieves all role").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetRolesResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get All Role"))

	webService.Route(webService.GET(getRoleByRoleIDEndpoint).To(service.GetRoleByID).
		Param(webService.PathParameter("roleid", "the id of role")).
		Notes("Get role by role id \n \n Retrieves role").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetRoleResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get role by RoleID"))

	return webService
}

// AddRouteProduct is setup restful route
func AddRouteProduct(service *service.Service, basePath string) *restful.WebService {
	webService := new(restful.WebService)
	webService.Path(basePath + "/product").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		ApiVersion("0.0.1").
		Doc("Product")

	tags := []string{"Product"}

	webService.Route(webService.POST(createProductEndpoint).To(service.CreateProduct).
		Reads(models.ProductRequest{}).
		Notes("Create Product \n \n Create Product").
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), models.CreateProductResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Create Product"))

	webService.Route(webService.GET(getAllProductEndpoint).To(service.GetAllProduct).
		Notes("Get all Product \n \n Retrieves all Product").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetProductsResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get All Product"))

	webService.Route(webService.GET(getProductByProductIDEndpoint).To(service.GetProductByProductID).
		Param(webService.PathParameter("productid", "the id of product")).
		Notes("Get product by product id \n \n Retrieves product").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetProductResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get Product by ProductID"))

	webService.Route(webService.GET(getProductByCategoryIDEndpoint).To(service.GetProductByCategoryID).
		Param(webService.PathParameter("categoryid", "the product category id")).
		Notes("Get Product by category id \n \n Retrieves Product").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetProductResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get Product by Category ID"))

	webService.Route(webService.GET(getProductLikeNameEndpoint).To(service.GetProductLikeName).
		Param(webService.PathParameter("identifier", "the product name identifier")).
		Notes("Get product by identifier \n \n Retrieves product").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.ProductResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get Product by name identifier"))

	webService.Route(webService.PUT(updateProductEndpoint).To(service.UpdateProduct).
		Reads(models.ProductRequest{}).
		Notes("Update product \n \n Update product").
		Returns(http.StatusAccepted, http.StatusText(http.StatusOK), models.UpdateProductResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Update Product"))

	webService.Route(webService.DELETE(deleteProductEndpoint).To(service.DeleteProduct).
		Reads(models.ProductRequest{}).
		Notes("Delete product \n \n Delete product").
		Returns(http.StatusNoContent, http.StatusText(http.StatusOK), models.DeleteProductResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Delete Product"))

	return webService
}

// AddRouteCategory is setup restful route
func AddRouteCategory(service *service.Service, basePath string) *restful.WebService {
	webService := new(restful.WebService)
	webService.Path(basePath + "/category").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		ApiVersion("0.0.1").
		Doc("Category")

	tags := []string{"Category"}

	webService.Route(webService.POST(createCategoryEndpoint).To(service.CreateCategory).
		Reads(models.CategoryRequest{}).
		Notes("Create Category \n \n Create Category").
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), models.CreateCategoryResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Create Category"))

	webService.Route(webService.GET(getAllCategoryEndpoint).To(service.GetAllCategory).
		Notes("Get all Category \n \n Retrieves all Category").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetCategoriesResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get All Category"))

	webService.Route(webService.GET(getCategoryByCategoryIDEndpoint).To(service.GetCategoryByCategoryID).
		Param(webService.PathParameter("categoryid", "the id of category")).
		Notes("Get Category by Category id \n \n Retrieves Category").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), models.GetCategoryResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Get Category by CategoryID"))

	webService.Route(webService.PUT(updateCategoryEndpoint).To(service.UpdateCategory).
		Reads(models.CategoryRequest{}).
		Notes("Update Category \n \n Update Category").
		Returns(http.StatusAccepted, http.StatusText(http.StatusOK), models.UpdateCategoryResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Update Category"))

	webService.Route(webService.DELETE(deleteCategoryEndpoint).To(service.DeleteCategory).
		Reads(models.CategoryRequest{}).
		Notes("Delete Category \n \n Delete Category").
		Returns(http.StatusNoContent, http.StatusText(http.StatusOK), models.DeleteCategoryResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), models.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), models.ErrorResponse{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Delete Category"))

	return webService
}
