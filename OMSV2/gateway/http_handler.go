package main

import (
	"net/http"
	"errors"

	common "github.com/SebastianFM1/Go-projects/OMSV2/commons"
	pb "github.com/SebastianFM1/Go-projects/OMSV2/commons/api"
)

type handler struct{
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler{
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux){
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)

}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var req pb.CreateOrderRequest
	if err := common.ReadJSON(r, &req); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Sobrescribe el customerID desde la URL
	req.CustomerID = customerID

	if err := validateItems(req.Items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.client.CreateOrder(r.Context(), &req)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusOK, o)
}


func validateItems(items []*pb.ItemsWithQuantity) error{
	if len(items) == 0{
		return common.ErrNoItems
	}

	for _, i := range items{
		if i.ID == "" {
			return errors.New("item id cannot be empty")
		}
		if i.Quantity <= 0 {
			return errors.New("items must have a valid quantity")

	}
	}

	return nil
}
