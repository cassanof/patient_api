package controllers

import "github.com/elleven11/patient_api/api/middlewares"

func (srv *Server) initRoutes() {

	srv.Router.HandleFunc("/", middlewares.SetJSON(srv.Home)).Methods("GET")
	srv.Router.HandleFunc("/login", middlewares.SetJSON(srv.Login)).Methods("POST")

	// Users
	srv.Router.HandleFunc("/users", middlewares.SetJSON(srv.CreateUser)).Methods("POST")
	srv.Router.HandleFunc("/users", middlewares.SetJSON(srv.GetUsers)).Methods("GET")
	srv.Router.HandleFunc("/users/{id}", middlewares.SetJSON(srv.GetUser)).Methods("GET")
	srv.Router.HandleFunc("/users/{id}", middlewares.SetJSON(middlewares.SetAuth(srv.UpdateUser))).Methods("PUT")
	srv.Router.HandleFunc("/users/{id}", middlewares.SetAuth(srv.DeleteUser)).Methods("DELETE")

	// Patients
	srv.Router.HandleFunc("/patients", middlewares.SetJSON(middlewares.SetAuth(srv.CreatePatient))).Methods("POST")
	// can only access /patients
	srv.Router.HandleFunc("/patients", middlewares.SetJSON(middlewares.SetAuth(srv.GetPatients))).Methods("GET")
	srv.Router.HandleFunc("/patients/{id}", middlewares.SetJSON(middlewares.SetAuth(srv.GetPatient))).Methods("GET")
	srv.Router.HandleFunc("/patients/{id}", middlewares.SetJSON(middlewares.SetAuth(srv.UpdatePatient))).Methods("PUT")
	srv.Router.HandleFunc("/patients/{id}", middlewares.SetAuth(srv.DeleteUser)).Methods("DELETE")

}
