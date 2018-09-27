package master

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/master/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MasterManager) AddCompany(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	company := model.Company{}
	if err := render.Bind(r, &company); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&company); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	companyID, err := mm.db.AddCompany(ctx, &company)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: companyID})
}

func (mm *MasterManager) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	company := model.Company{}
	if err := render.Bind(r, &company); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&company); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateCompany(ctx, &company)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: true})
}

func (mm *MasterManager) GetCompanyByID(w http.ResponseWriter, r *http.Request) {

	request := model.GetChildrensReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetCompanyByID(request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MasterManager) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := mm.db.GetAllCompanies()
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: companies})
}
