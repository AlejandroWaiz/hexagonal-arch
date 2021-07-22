package commonhttp

import (
	"net/http"

	web "github.com/AlejandroWaiz/hexagonal-arch/infrastructure/webResponse"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	web.Succes("Hello there, im working well, go try the rest of the page!.", 200).Send(w)

}
