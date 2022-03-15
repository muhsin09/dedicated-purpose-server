package handlers

import "net/http"

type IServerHandler interface {
	AddMessage(w http.ResponseWriter, r *http.Request)
	PollMessage(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func (h handler) AddMessage(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h handler) PollMessage(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewServerHandler() IServerHandler {

	return &handler{}
}
