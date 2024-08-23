package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "96HW/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
    items := make([]models.Item, 0, len(models.Items))
    for _, item := range models.Items {
        items = append(items, item)
    }
    json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
    var item models.Item
    json.NewDecoder(r.Body).Decode(&item)
    item.ID = len(models.Items) + 1
    models.Items[item.ID] = item
    json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }

    var item models.Item
    json.NewDecoder(r.Body).Decode(&item)
    item.ID = id
    models.Items[id] = item
    json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }

    delete(models.Items, id)
    w.WriteHeader(http.StatusNoContent)
}
