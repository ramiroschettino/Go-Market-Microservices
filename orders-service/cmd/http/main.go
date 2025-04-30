package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	orderpb "github.com/ramiroschettino/Go-Market-Microservices/orders-service/proto"
	"google.golang.org/grpc"
)

type OrderRequest struct {
	ProductID          int64   `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductPrice       float64 `json:"product_price"`
	Quantity           int32   `json:"quantity"`
}

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ No se pudo conectar al server gRPC de órdenes: %v", err)
	}
	defer conn.Close()

	orderClient := orderpb.NewOrderServiceClient(conn)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		var req OrderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "❌ Error al parsear JSON", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		orderResp, err := orderClient.CreateOrder(ctx, &orderpb.CreateOrderRequest{
			ProductId:          req.ProductID,
			ProductName:        req.ProductName,
			ProductDescription: req.ProductDescription,
			ProductPrice:       req.ProductPrice,
			Quantity:           req.Quantity,
		})
		if err != nil {
			http.Error(w, "❌ Error al crear orden: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orderResp)
	})

	log.Println("🌐 Servidor HTTP escuchando en http://localhost:8080 ...")
	http.ListenAndServe(":8080", nil)
}
