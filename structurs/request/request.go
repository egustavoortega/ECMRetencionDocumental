package request



type RequestCreateRoles struct {

	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Privilege_Id	string		`json:"privilege_id"`

}

type RequestShowRoles  struct {

	Id  string   `json:"id"`

}

type RequestUpdateRoles  struct {

	Id  			string   	`json:"id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Privilege_Id	string		`json:"privilege_id"`

}
type RequestdeleteRoles  struct {

	Id  string   `json:"id"`

}