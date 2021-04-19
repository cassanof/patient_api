package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/elleven11/patient_api/api/auth"
	"github.com/elleven11/patient_api/api/models"
	"github.com/elleven11/patient_api/api/responses"
	"github.com/elleven11/patient_api/api/utils"
	"github.com/gorilla/mux"
)

func (srv *Server) CreateUser(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.CreateUser(srv.DB)
	if err != nil {
		fmtError := utils.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, fmtError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", req.Host, req.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (srv *Server) GetUsers(w http.ResponseWriter, req *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(srv.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (srv *Server) GetUser(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(srv.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (srv *Server) UpdateUser(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.TokenExtractTokenID(req)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(srv.DB, uint32(uid))
	if err != nil {
		fmtError := utils.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, fmtError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

func (srv *Server) DeleteUser(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.TokenExtractTokenID(req)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = user.DeleteAUser(srv.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
