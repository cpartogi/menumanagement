package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	swaggerDocs "github.com/cpartogi/menumanagement/api/docs"
	account "github.com/cpartogi/menumanagement/internal/menu"
	"github.com/cpartogi/menumanagement/internal/menu/request"
	mdl "github.com/cpartogi/menumanagement/pkg/common/middleware"
	"github.com/cpartogi/menumanagement/pkg/common/util"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

type HTTPHandler struct {
	service account.Service
}

// NewHandler register http handler for Account
func NewHandler(r *httprouter.Router, as account.Service) {
	swaggerDocs.SwaggerInfo.Title = "Menu API"
	swaggerDocs.SwaggerInfo.Description = "API For Menu"
	swaggerDocs.SwaggerInfo.Version = "1.0"
	handler := &HTTPHandler{
		service: as,
	}

	//menus
	r.Handler("GET", "/v1/menus", mdl.Public(handler.GetMenus))
	r.Handler("GET", "/v1/menu/:id", mdl.Public(handler.GetMenuDetail))
	r.Handler("POST", "/v1/menu", mdl.Public(handler.CreateMenu))
	r.Handler("PUT", "/v1/menu/:id", mdl.Public(handler.UpdateMenu))
	r.Handler("DELETE", "/v1/menu/:id", mdl.Public(handler.DeleteMenu))

	r.Handler("GET", "/swagger/*path", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition"
	))
}

// GetMenus godoc
// @Summary Get Menus List
// @Description Get Menus List
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param name query string false "search by menu name"
// @Success 200 {object} response.SwaggerMenuList
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/menus [get]
func (ah *HTTPHandler) GetMenus(w http.ResponseWriter, r *http.Request, nex http.HandlerFunc) {
	ctx := r.Context()
	queryValues := r.URL.Query()
	filter := make(map[string]string)

	if queryValues.Get("name") != "" {
		filter["name"] = queryValues.Get("name")
	}

	result, err := ah.service.GetMenus(ctx, filter)
	if err != nil {
		util.ResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.ResponseJSON(w, http.StatusOK, "", result)
}

// GetMenuDetail godoc
// @Summary Get Menu Detail
// @Description Get Menu Detail
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param menu_id path string true "Menu ID"
// @Success 200 {object} response.SwaggerMenuDetail
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/menu/{menu_id} [get]
func (ah *HTTPHandler) GetMenuDetail(w http.ResponseWriter, r *http.Request, nex http.HandlerFunc) {
	ctx := r.Context()

	params := httprouter.ParamsFromContext(ctx)

	idUri := params.ByName("id")
	id, parseErr := strconv.ParseInt(idUri, 10, 64)

	var err error

	// ! Abort if there's an error occured while parsing URI ID to int
	if parseErr != nil {
		log.Println("Failed to convert id:", parseErr)
		util.ResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	result, err := ah.service.GetMenuDetail(&id)
	if err != nil {
		util.ResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	util.ResponseJSON(w, http.StatusOK, "", result)
}

// CreateMenu godoc
// @Summary Add Menu
// @Description Add Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param request body request.MenuRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Failure 504 {object} response.Base
// @Router /v1/menu [POST]
func (ah *HTTPHandler) CreateMenu(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userId := int64(0)

	var data request.MenuRequest
	var err error

	// * Decode payload
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Println("Failed decode body", err)
		util.ResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// * Fill the rest automatically
	mMenu := data.FormatMenu(&userId)

	// * All done, time to store the data
	err = ah.service.CreateMenu(&mMenu)
	if err != nil {
		log.Println("Failed to store data:", err)
		util.ResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.ResponseJSON(w, http.StatusOK, "", nil)
}

// Update Menu godoc
// @Summary Update Menu
// @Description Update Menu
// @Tags Menu
// @Produce json
// @Param menu_id path string true "menu id"
// @Param request body request.MenuRequest true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Failure 504 {object} response.Base
// @Router /v1/menu/{menu_id} [put]
// Register handles HTTP request to update account
func (ah *HTTPHandler) UpdateMenu(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()

	userId := int64(0)

	var err error
	var data request.MenuRequest

	params := httprouter.ParamsFromContext(ctx)
	idUri := params.ByName("id")
	id, parseErr := strconv.ParseInt(idUri, 10, 64)

	// ! Abort if there's an error occured while parsing URI ID to int
	if parseErr != nil {
		log.Println("Failed to convert id:", parseErr)
		util.ResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// * Decode payload
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Println("Failed to convert id:", err)
		util.ResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// * Fill the rest automatically
	mMenu := data.FormatMenu(&userId)

	//check data exist
	_, err = ah.service.GetMenuDetail(&id)
	if err != nil {
		util.ResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// * All done, time to store the data
	err = ah.service.UpdateMenu(&id, &mMenu)
	if err != nil {
		log.Println("Failed update menu:", err)
		util.ResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.ResponseJSON(w, http.StatusOK, "", nil)
}

// Delete Menu godoc
// @Summary Delete menu
// @Description Delete menu
// @Tags Menu
// @Produce json
// @Param menu_id path string true "Menu Id"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Failure 504 {object} response.Base
// @Router /v1/menu/{menu_id} [delete]
// Register handles HTTP request to delete account
func (ah *HTTPHandler) DeleteMenu(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()

	var err error

	params := httprouter.ParamsFromContext(ctx)
	idUri := params.ByName("id")
	id, parseErr := strconv.ParseInt(idUri, 10, 64)

	// ! Abort if there's an error occured while parsing URI ID to int
	if parseErr != nil {
		log.Println("Failed to convert id:", parseErr)
		util.ResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check data exist
	_, err = ah.service.GetMenuDetail(&id)
	if err != nil {
		util.ResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// delete data
	err = ah.service.DeleteMenu(&id)
	if err != nil {
		log.Println("Failed delete menu:", err)
		util.ResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.ResponseJSON(w, http.StatusOK, "", nil)
}
