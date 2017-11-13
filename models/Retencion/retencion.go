package Retencion
import (
	"database/sql"
	"fmt"

)
type Doctypes struct {
    Codigo				int64		`json:"codigo"`
	Name_Serie			string		`json:"name_serie"`
	Name_Subserie		string		`json:"name_subserie"`
	Tipo_Soporte		string		`json:"tipo_soporte"`
	Retencion_Ag		string		`json:"retencion_ag"`
	Retencion_Ac		string		`json:"retencion_ac"`
	Retencion_Ah		string		`json:"retencion_ah"`
	Disposicion_Final	string		`json:"disposicion_final"`
	Digitalizacion		string		`json:"digitalizacion"`
	procedimiento		string		`json:"procedimiento"`

}

type T_Doctype []Doctypes

func Index(db *sql.DB) (*T_Doctype, error){

	query := fmt.Sprintf(`select right('000' + CAST( (CONCAT (b.id,a.id)) as varchar),3)  Codigo2,b.name Serie,a.name Subserie , a.tipo_soporte Soporte,a.retencion_ag,a.retencion_ac,a.retencion_ah,a.disposicion_final,a.digitalizacion,a.Procedimiento
								 from doc_types as a, doc_type_groups as b, doc_storages as c
								 where c.id = b.storage_id
								 and a.typegroup_id=b.id`)

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	data := T_Doctype{}
	for rows.Next(){
		regist := Doctypes{}
		err := rows.Scan(
			&regist.Codigo,
			&regist.Name_Serie,
			&regist.Name_Subserie,
			&regist.Tipo_Soporte,
			&regist.Retencion_Ac,
			&regist.Retencion_Ag,
			&regist.Retencion_Ah,
			&regist.Digitalizacion,
			&regist.Disposicion_Final,
			&regist.procedimiento,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, regist)
	}

	return &data, nil
}


func GetAll(db *sql.DB) (*T_Doctype, error){

	query := fmt.Sprintf(`select name,tipo_soporte,retencion_ag,retencion_ac,retencion_ah,disposicion_final,digitalizacion,Procedimiento
								FROM doc_types`)

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	data := T_Doctype{}
	for rows.Next(){
		regist := Doctypes{}
		err := rows.Scan(
			&regist.Codigo,
			&regist.Name_Serie,
			&regist.Name_Subserie,
			&regist.Tipo_Soporte,
			&regist.Retencion_Ag,
			&regist.Retencion_Ac,
			&regist.Retencion_Ah,
			&regist.Disposicion_Final,
			&regist.Digitalizacion,
			&regist.procedimiento,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, regist)
	}

	return &data, nil
}