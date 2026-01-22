package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`
	KategoriID int    `json:"kategori_id"`
}

type Kategori struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

type ProdukResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	Harga        int    `json:"harga"`
	Stok         int    `json:"stok"`
	KategoriID   int    `json:"kategori_id"`
	NamaKategori string `json:"nama_kategori"`
}

var produk = []Produk{
	{ID: 1, Nama: "Laptop", Harga: 15000000, Stok: 10, KategoriID: 1},
	{ID: 2, Nama: "Smartphone", Harga: 8000000, Stok: 25, KategoriID: 2},
	{ID: 3, Nama: "Tablet", Harga: 5000000, Stok: 15, KategoriID: 1},
}

var kategori = []Kategori{
	{ID: 1, Nama: "Elektronik", Deskripsi: "Kategori Elektronik"},
	{ID: 2, Nama: "Pakaian", Deskripsi: "Kategori Pakaian"},
	{ID: 3, Nama: "Aksesori", Deskripsi: "Kategori Aksesori"},
}

// Helper function untuk mendapatkan nama kategori berdasarkan ID
func getNamaKategori(kategoriID int) string {
	for _, k := range kategori {
		if k.ID == kategoriID {
			return k.Nama
		}
	}
	return ""
}

// PRODUK
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			produkResponse := ProdukResponse{
				ID:           p.ID,
				Nama:         p.Nama,
				Harga:        p.Harga,
				Stok:         p.Stok,
				KategoriID:   p.KategoriID,
				NamaKategori: getNamaKategori(p.KategoriID),
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produkResponse)
			return
		}
	}

	http.Error(w, "Produk not found", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid produk ID", http.StatusBadRequest)
		return
	}

	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk[i])
			return
		}
	}

	http.Error(w, "Produk not found", http.StatusNotFound)
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Produk deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Produk not found", http.StatusNotFound)

}

// KATEGORI

func getKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func postKategori(w http.ResponseWriter, r *http.Request) {
	var kategoriBaru Kategori
	err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	kategoriBaru.ID = len(kategori) + 1
	kategori = append(kategori, kategoriBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategoriBaru)
}

func updateKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid kategori ID", http.StatusBadRequest)
		return
	}

	var updateKategori Kategori
	err = json.NewDecoder(r.Body).Decode(&updateKategori)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	for i := range kategori {
		if kategori[i].ID == id {
			updateKategori.ID = id
			kategori[i] = updateKategori

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(kategori[i])
			return
		}
	}

	http.Error(w, "Kategori not found", http.StatusNotFound)
}

func deleteKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid kategori ID", http.StatusBadRequest)
		return
	}

	for i, k := range kategori {
		if k.ID == id {
			kategori = append(kategori[:i], kategori[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Kategori deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Kategori not found", http.StatusNotFound)
}

func getKategoriByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid kategori ID", http.StatusBadRequest)
		return
	}

	for _, k := range kategori {
		if k.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(k)
			return
		}
	}

	http.Error(w, "Kategori not found", http.StatusNotFound)
}

func main() {

	// PRODUK

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// Membuat array response dengan nama kategori
			var produkResponses []ProdukResponse
			for _, p := range produk {
				produkResponse := ProdukResponse{
					ID:           p.ID,
					Nama:         p.Nama,
					Harga:        p.Harga,
					Stok:         p.Stok,
					KategoriID:   p.KategoriID,
					NamaKategori: getNamaKategori(p.KategoriID),
				}
				produkResponses = append(produkResponses, produkResponse)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produkResponses)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid data", http.StatusBadRequest)
				return
			}

			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		}

	})

	// KATEGORI

	// GET localhost:8080/api/kategori
	// POST localhost:8080/api/kategori
	http.HandleFunc("/api/kategori", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getKategori(w, r)
		} else if r.Method == "POST" {
			postKategori(w, r)
		}
	})
	// PUT localhost:8080/api/kategori/{id}
	// DELETE localhost:8080/api/kategori/{id}
	// GET localhost:8080/api/kategori/{id}
	http.HandleFunc("/api/kategori/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getKategoriByID(w, r)
		} else if r.Method == "PUT" {
			updateKategori(w, r)
		} else if r.Method == "DELETE" {
			deleteKategori(w, r)
		}
	})

	// ============================================
	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal running server")
	}

}
