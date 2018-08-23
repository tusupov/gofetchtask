package handle

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tusupov/gofetchtask/task"
)

func Task(w http.ResponseWriter, r *http.Request) {

	// Load data from body
	var requestTask task.RequestTask
	err := json.NewDecoder(r.Body).Decode(&requestTask)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create new task (do request)
	t, err := task.NewTask(requestTask)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write task
	writeSuccess(w, t)
}

func List(w http.ResponseWriter, r *http.Request) {

	// Task list
	l := task.StoreList()

	// write list
	writeSuccess(w, l)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	// Get query params
	query := r.URL.Query()
	idStr := query.Get("id")

	if len(idStr) == 0 {
		writeError(w, http.StatusBadRequest, "`id` is param required")
		return
	}

	// Convert id to int
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Delete from store by `id`
	ok := task.StoreDelete(id)
	if !ok {
		writeError(w, http.StatusNotFound, "Not found")
		return
	}

	writeSuccess(w, map[string]string{
		"success": "ok",
	})

}

func writeSuccess(w http.ResponseWriter, result interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)

}

func writeError(w http.ResponseWriter, code int, message string) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
		"code":    code,
	})

}
