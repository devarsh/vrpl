package misc

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/misc/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MiscManager) AddCity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	city := model.City{}
	if err := render.Bind(r, &city); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&city); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	cityID, err := mm.db.AddCity(ctx, &city)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: cityID})
}

func (mm *MiscManager) UpdateCity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	city := model.City{}
	if err := render.Bind(r, &city); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&city); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateCity(ctx, &city)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Status: true})
}

func (mm *MiscManager) GetAllCitiesByStateID(w http.ResponseWriter, r *http.Request) {
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
	res, err := mm.db.GetCitiesByStateId(ctx, request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MiscManager) GetCityByID(w http.ResponseWriter, r *http.Request) {
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
	res, err := mm.db.GetCityByID(ctx, request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}
