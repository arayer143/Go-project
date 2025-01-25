package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-project/db"
	"go-project/models"
)

// GetProducts handles retrieving all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var products []models.Product
	cursor, err := db.DB.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		logrus.Errorf("Error fetching products: %v", err)
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &products); err != nil {
		logrus.Errorf("Error decoding products: %v", err)
		http.Error(w, "Error decoding products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProduct handles retrieving a single product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product
	err = db.DB.Collection("products").FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			logrus.Errorf("Error fetching product: %v", err)
			http.Error(w, "Error fetching product", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// CreateProduct handles the creation of a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		logrus.Errorf("Error decoding product JSON: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.DB.Collection("products").InsertOne(ctx, product)
	if err != nil {
		logrus.Errorf("Error inserting product into database: %v", err)
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct handles updating an existing product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		logrus.Errorf("Error decoding product JSON: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	product.ID = id
	product.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.DB.Collection("products").ReplaceOne(ctx, bson.M{"_id": id}, product)
	if err != nil {
		logrus.Errorf("Error updating product: %v", err)
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct handles deleting a product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.DB.Collection("products").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		logrus.Errorf("Error deleting product: %v", err)
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

