// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// User defines model for User.
type User struct {
	Email    *openapi_types.Email `json:"email,omitempty"`
	Id       *uint                `json:"id,omitempty"`
	Password *string              `json:"password,omitempty"`
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// UpdateUserByIDJSONRequestBody defines body for UpdateUserByID for application/json ContentType.
type UpdateUserByIDJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получение списка всех пользователей
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Создание нового пользователя
	// (POST /users)
	CreateUser(ctx echo.Context) error
	// Удаление пользователя по его ID
	// (DELETE /users/{id})
	DeleteUserByID(ctx echo.Context, id int) error
	// Редактирование полей пользователя по его ID
	// (PATCH /users/{id})
	UpdateUserByID(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// DeleteUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUserByID(ctx, id)
	return err
}

// UpdateUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateUserByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.CreateUser)
	router.DELETE(baseURL+"/users/:id", wrapper.DeleteUserByID)
	router.PATCH(baseURL+"/users/:id", wrapper.UpdateUserByID)

}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsers400Response struct {
}

func (response GetUsers400Response) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetUsers500Response struct {
}

func (response GetUsers500Response) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type CreateUserRequestObject struct {
	Body *CreateUserJSONRequestBody
}

type CreateUserResponseObject interface {
	VisitCreateUserResponse(w http.ResponseWriter) error
}

type CreateUser201JSONResponse User

func (response CreateUser201JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser400Response struct {
}

func (response CreateUser400Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type CreateUser409Response struct {
}

func (response CreateUser409Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type CreateUser500Response struct {
}

func (response CreateUser500Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type DeleteUserByIDRequestObject struct {
	Id int `json:"id"`
}

type DeleteUserByIDResponseObject interface {
	VisitDeleteUserByIDResponse(w http.ResponseWriter) error
}

type DeleteUserByID204Response struct {
}

func (response DeleteUserByID204Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteUserByID400Response struct {
}

func (response DeleteUserByID400Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type DeleteUserByID404Response struct {
}

func (response DeleteUserByID404Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type DeleteUserByID500Response struct {
}

func (response DeleteUserByID500Response) VisitDeleteUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type UpdateUserByIDRequestObject struct {
	Id   int `json:"id"`
	Body *UpdateUserByIDJSONRequestBody
}

type UpdateUserByIDResponseObject interface {
	VisitUpdateUserByIDResponse(w http.ResponseWriter) error
}

type UpdateUserByID200JSONResponse User

func (response UpdateUserByID200JSONResponse) VisitUpdateUserByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUserByID400Response struct {
}

func (response UpdateUserByID400Response) VisitUpdateUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type UpdateUserByID404Response struct {
}

func (response UpdateUserByID404Response) VisitUpdateUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type UpdateUserByID500Response struct {
}

func (response UpdateUserByID500Response) VisitUpdateUserByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Получение списка всех пользователей
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Создание нового пользователя
	// (POST /users)
	CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error)
	// Удаление пользователя по его ID
	// (DELETE /users/{id})
	DeleteUserByID(ctx context.Context, request DeleteUserByIDRequestObject) (DeleteUserByIDResponseObject, error)
	// Редактирование полей пользователя по его ID
	// (PATCH /users/{id})
	UpdateUserByID(ctx context.Context, request UpdateUserByIDRequestObject) (UpdateUserByIDResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateUser operation middleware
func (sh *strictHandler) CreateUser(ctx echo.Context) error {
	var request CreateUserRequestObject

	var body CreateUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUser(ctx.Request().Context(), request.(CreateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreateUserResponseObject); ok {
		return validResponse.VisitCreateUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUserByID operation middleware
func (sh *strictHandler) DeleteUserByID(ctx echo.Context, id int) error {
	var request DeleteUserByIDRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUserByID(ctx.Request().Context(), request.(DeleteUserByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUserByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUserByIDResponseObject); ok {
		return validResponse.VisitDeleteUserByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// UpdateUserByID operation middleware
func (sh *strictHandler) UpdateUserByID(ctx echo.Context, id int) error {
	var request UpdateUserByIDRequestObject

	request.Id = id

	var body UpdateUserByIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateUserByID(ctx.Request().Context(), request.(UpdateUserByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateUserByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateUserByIDResponseObject); ok {
		return validResponse.VisitUpdateUserByIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
