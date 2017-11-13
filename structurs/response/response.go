package response

import (
	"github.com/e-capture/ECMRetencionDocumental/models/Retencion"
)

type Response struct {
	Error		string	`json:"error"`
	Data		string	`json:"data"`
	Message		string	`json:"message"`
}

type ResponseRetencion struct {
	Error		string						`json:"error"`
	Data		*Retencion.T_Doctype		`json:"data"`
	Message		string						`json:"message"`
}

type ModelRetencion struct {
	Error		string						`json:"error"`
	Data		*Retencion.Doctypes			`json:"data"`
	Message		string						`json:"message"`
}
