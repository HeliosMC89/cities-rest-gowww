package handlers

import (
	"log"
	"net/http"

	"github.com/gowww/router"
	"github.com/heliosmc89/api-rest-gowww/api/models"
	"github.com/heliosmc89/api-rest-gowww/api/repository"
	"github.com/heliosmc89/api-rest-gowww/helpers"
	"github.com/jmoiron/sqlx"
)

type CityHandler struct {
	Log  *log.Logger
	Repo repository.CityInterface
}

func NewCityHandler(logging *log.Logger, db *sqlx.DB) *CityHandler {
	return &CityHandler{
		Log:  logging,
		Repo: repository.NewRepoCity(logging, db),
	}
}

func (c *CityHandler) GetAllCities(w http.ResponseWriter, r *http.Request) {
	cites, err := c.Repo.FindAll()
	if err != nil {
		c.Log.Fatal(err)
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.JSONResponse{
			Error: true,
			Msg:   err.Error(),
			Code:  http.StatusInternalServerError,
			Data:  cites,
		})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.JSONResponse{
		Error: false,
		Msg:   "List of Cities",
		Code:  http.StatusOK,
		Data:  cites,
	})
}

func (c *CityHandler) GetCityByID(w http.ResponseWriter, r *http.Request) {
	id := router.Parameter(r, "id")
	city, err := c.Repo.FindOne(id)
	if err != nil {
		c.Log.Println(err)
		helpers.WriteJSON(w, http.StatusNotFound, helpers.JSONResponse{
			Error: true,
			Msg:   err.Error(),
			Code:  http.StatusNotFound,
			Data:  city,
		})
		return
	}
	helpers.WriteJSON(w, http.StatusNotFound, helpers.JSONResponse{
		Error: false,
		Msg:   "city",
		Code:  http.StatusOK,
		Data:  city,
	})
}

func (c *CityHandler) UpdateCity(w http.ResponseWriter, r *http.Request) {
	id := router.Parameter(r, "id")

	var city models.City
	// Read the body of the request.
	err := helpers.ReadJSON(r, &city)
	if err != nil {
		helpers.WriteJSON(w, http.StatusUnprocessableEntity, helpers.JSONResponse{
			Error: true,
			Msg:   "Body malformed",
			Code:  http.StatusUnprocessableEntity,
			Data:  nil,
		})
		return
	}
	defer r.Body.Close()

	// Query the database.
	updateCity, err := c.Repo.Update(id, &city)
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.JSONResponse{
			Error: true,
			Msg:   "internal error",
			Code:  http.StatusInternalServerError,
			Data:  nil,
		})
		return
	}

	// Send the response
	helpers.WriteJSON(w, http.StatusOK, helpers.JSONResponse{
		Error: false,
		Msg:   "Ok city update",
		Code:  http.StatusOK,
		Data:  updateCity,
	})

}

func (c *CityHandler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	// Get URL parameters with the city ID to delete.
	id := router.Parameter(r, "id")
	// Query the database.
	err := c.Repo.Delete(id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, helpers.JSONResponse{
			Error: true,
			Msg:   "resource not found",
			Code:  http.StatusNotFound,
			Data:  nil,
		})
		return
	}

	// Send the response
	helpers.WriteJSON(w, http.StatusNoContent, helpers.JSONResponse{
		Error: false,
		Msg:   "Ok city deleted",
		Code:  http.StatusNoContent,
		Data:  nil,
	})
}

func (c *CityHandler) CreateCity(w http.ResponseWriter, r *http.Request) {
	var city models.City

	// Read the body of the request.
	err := helpers.ReadJSON(r, &city)
	if err != nil {
		helpers.WriteJSON(w, http.StatusUnprocessableEntity, helpers.JSONResponse{
			Error: true,
			Msg:   "Body malformed",
			Code:  http.StatusUnprocessableEntity,
			Data:  nil,
		})
		return
	}
	defer r.Body.Close()

	// Write to the database
	addResult, err := c.Repo.Create(&city)
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, helpers.JSONResponse{
			Error: true,
			Msg:   "internal error",
			Code:  http.StatusInternalServerError,
			Data:  nil,
		})
		return
	}

	// Format response
	helpers.WriteJSON(w, http.StatusCreated, helpers.JSONResponse{
		Error: false,
		Msg:   "City added",
		Code:  http.StatusCreated,
		Data:  addResult,
	})
}
