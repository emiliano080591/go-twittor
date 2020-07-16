package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/emiliano080591/go-twittor/bd"
	"github.com/emiliano080591/go-twittor/models"
)

/*SubirBanner sube una imagen al servidor*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoResgitro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner en la BD! "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaGeneral{
		Status:  true,
		Message: "Se subio el banner satisfactoriamente",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
