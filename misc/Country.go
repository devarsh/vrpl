package misc

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/misc/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MiscManager) AddCountry(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	country := model.Country{}
	if err := render.Bind(r, &country); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&country); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	countryID, err := mm.db.AddCountry(ctx, &country)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: countryID})
}

func (mm *MiscManager) UpdateCountry(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	country := model.Country{}
	if err := render.Bind(r, &country); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&country); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateCountry(ctx, &country)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Status: true})
}

func (mm *MiscManager) GetAllCountries(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	countries, err := mm.db.GetAllCountries(ctx)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: countries})
}

func (mm *MiscManager) GetCountryByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	request := model.GetChildrensReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetCountryByID(ctx, request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}
