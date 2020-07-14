package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/emiliano080591/go-twittor/bd"
	"github.com/emiliano080591/go-twittor/models"
)

//Email sera el email de los usuarios para todos los endpoints
var Email string
//IDUsuario trae el id de los usuarios
var IDUsuario string

/*ProcesoToken extrae los valores del token*/
func ProcesoToken(tk string) (*models.Claim,bool,string,error) {
	miClave := []byte("G@$Tw1tt0r")
	claims:= &models.Claim{}

	splitToken:=strings.Split(tk,"Bearer")
	if len(splitToken) != 2{
		return claims,false,string(""),errors.New("formato de token invalido")
	}

	tk=strings.TrimSpace(splitToken[1])

	tkn,err:=jwt.ParseWithClaims(tk,claims,func(token *jwt.Token)(interface{},error){
		return miClave,nil
	})

	if err == nil {
		_,encontrado,_:=bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado==true {
			Email=claims.Email
			IDUsuario=claims.ID.Hex()
		} 
		return claims,encontrado,IDUsuario,nil
	}

	if !tkn.Valid {
		return claims,false,string(""),errors.New("token invalido")
	}

	return claims,false,string(""),err
}