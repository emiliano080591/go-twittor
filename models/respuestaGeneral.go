package models

/*RespuestaGeneral da una respuesta de true*/
type RespuestaGeneral struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
