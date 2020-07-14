package models

/*Tweet captura del body*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
