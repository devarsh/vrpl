package master

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/master/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MasterManager) AddClient(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := model.Client{}
	if err := render.Bind(r, &client); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&client); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	clientID, err := mm.db.AddClient(ctx, &client)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: clientID})
}

func (mm *MasterManager) UpdateClient(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := model.Client{}
	if err := render.Bind(r, &client); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&client); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateClient(ctx, &client)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: true})
}

func (mm *MasterManager) GetClientByID(w http.ResponseWriter, r *http.Request) {

	request := model.GetChildrensReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetClientByID(request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MasterManager) GetClientByGroupID(w http.ResponseWriter, r *http.Request) {

	request := model.GetChildrensReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetClientByGroupID(request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MasterManager) GetClientByName(w http.ResponseWriter, r *http.Request) {
	request := model.QueryReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetClientsByName(request.Name)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MasterManager) GetClientByTallyName(w http.ResponseWriter, r *http.Request) {
	request := model.QueryReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetClientsByTallyName(request.Name)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}
