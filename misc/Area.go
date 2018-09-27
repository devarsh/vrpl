package misc

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/misc/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MiscManager) AddArea(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	area := model.Area{}
	if err := render.Bind(r, &area); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&area); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	areaID, err := mm.db.AddArea(ctx, &area)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: areaID})
}

func (mm *MiscManager) UpdateArea(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	area := model.Area{}
	if err := render.Bind(r, &area); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&area); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateArea(ctx, &area)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Status: true})
}

func (mm *MiscManager) GetAllAreasByCityID(w http.ResponseWriter, r *http.Request) {
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
	res, err := mm.db.GetAreasByCityId(ctx, request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MiscManager) GetAreaByID(w http.ResponseWriter, r *http.Request) {
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
	res, err := mm.db.GetAreaById(ctx, request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}
