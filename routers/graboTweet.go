package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/emiliano080591/go-twittor/bd"
	"github.com/emiliano080591/go-twittor/models"
)

/*GraboTweet inserta un tweet*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el tweet "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet ", 400)
		return
	}

	resp := models.RespuestaGeneral{
		Status:  true,
		Message: "Se inserto el tweet satisfactoriamente",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
