package RetencionController
import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/e-capture/ECMRetencionDocumental/structurs/response"
	"github.com/e-capture/MAJOSystem/db"
	"github.com/e-capture/MAJOSystem/log"
	"net/http"
	"github.com/e-capture/ECMRetencionDocumental/models/Retencion"
)
func Index(c echo.Context) error {

	db, err := db.Open()
	defer db.Close()

	result, err := Retencion.Index(db)
	if err != nil {
		text := fmt.Sprintf("Error DB, error en la consulta a la base de datos., %s: ", err)
		log.Write(text)
		return err
	}

	response := response.ResponseRetencion{}
	response.Error = "false"
	response.Data = result
	response.Message = "La consulta se realizo exitosamente!"

	return c.JSON(http.StatusOK, &response)
}


