package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}
}

func TestCreateOrder(t *testing.T) {
	store := NewStore()
	handler := createOrderHandler(store)

	body := `{"item": "pizza", "quantity": 2}`
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201 got %d", w.Code)
	}

	var order Order
	json.NewDecoder(w.Body).Decode(&order)

	if order.Item != "pizza" {
		t.Errorf("expected pizza got %s", order.Item)
	}
	if order.Quantity != 2 {
		t.Errorf("expected 2 got %d", order.Quantity)
	}
}

func TestCreateOrder_InvalidBody(t *testing.T) {
	store := NewStore()
	handler := createOrderHandler(store)

	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 got %d", w.Code)
	}
}

func TestGetOrder(t *testing.T) {
	store := NewStore()
	order := store.Add("burger", 1)

	handler := getOrderHandler(store)
	req := httptest.NewRequest(http.MethodGet, "/orders/"+order.ID, nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}
}

func TestGetOrder_NotFound(t *testing.T) {
	store := NewStore()
	handler := getOrderHandler(store)

	req := httptest.NewRequest(http.MethodGet, "/orders/999", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404 got %d", w.Code)
	}
}

func TestListOrders(t *testing.T) {
	store := NewStore()
	store.Add("pizza", 1)
	store.Add("burger", 2)

	handler := listOrdersHandler(store)
	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", w.Code)
	}

	var orders []Order
	json.NewDecoder(w.Body).Decode(&orders)

	if len(orders) != 2 {
		t.Errorf("expected 2 orders got %d", len(orders))
	}
}
