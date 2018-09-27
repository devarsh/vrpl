package misc

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/misc/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MiscManager) AddItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	item := model.Item{}
	if err := render.Bind(r, &item); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&item); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	itemID, err := mm.db.AddItem(ctx, &item)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: itemID})
}

func (mm *MiscManager) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	item := model.Item{}
	if err := render.Bind(r, &item); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&item); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateItem(ctx, &item)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Status: true})
}

func (mm *MiscManager) GetAllItems(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	items, err := mm.db.GetAllItems(ctx)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: items})
}

func (mm *MiscManager) GetItemByID(w http.ResponseWriter, r *http.Request) {
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
	res, err := mm.db.GetItemByID(ctx, request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}
