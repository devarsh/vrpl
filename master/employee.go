package master

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/master/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MasterManager) AddEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	employee := model.Employee{}
	if err := render.Bind(r, &employee); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&employee); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	employeeID, err := mm.db.AddEmployee(ctx, &employee)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: employeeID})
}

func (mm *MasterManager) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	employee := model.Employee{}
	if err := render.Bind(r, &employee); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&employee); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateEmployee(ctx, &employee)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: true})
}

func (mm *MasterManager) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	request := model.GetChildrensReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetEmployeeByID(request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MasterManager) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := mm.db.GetAllEmployees()
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: employees})
}
