package routers

import (
	"encoding/json"
	"net/http"

	"github.com/emiliano080591/go-twittor/bd"
	"github.com/emiliano080591/go-twittor/models"
)

/*EliminarTweet permite borrar un tweet determinado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaGeneral{
		Status:  true,
		Message: "Se elimino el tweet satisfactoriamente",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
