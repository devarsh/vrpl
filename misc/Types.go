package misc

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/misc/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func CustomAddType(mm *MiscManager, group string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		types := model.Types{}
		if err := render.Bind(r, &types); err != nil {
			render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
			return
		}
		types.GroupNm = group
		if err := customErr.ValidateStruct(&types); err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
			return
		}
		ID, err := mm.db.AddType(ctx, &types)
		if err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
			return
		}
		render.Render(w, r, &resp.Resp{Result: ID})
	}
}

func CustomUpdateType(mm *MiscManager, group string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		types := model.Types{}
		if err := render.Bind(r, &types); err != nil {
			render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
			return
		}
		types.GroupNm = group
		if err := customErr.ValidateStruct(&types); err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
			return
		}
		err := mm.db.UpdateType(ctx, &types)
		if err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
			return
		}
		render.Render(w, r, &resp.Resp{Status: true})
	}
}

func CustomGetAllType(mm *MiscManager, group string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		res, err := mm.db.GetTypesByGroupNm(ctx, group)
		if err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
			return
		}
		render.Render(w, r, &resp.Resp{Result: res})
	}
}

func CustomGetTypeByID(mm *MiscManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		res, err := mm.db.GetTypeById(ctx, request.ParentID)
		if err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
			return
		}
		render.Render(w, r, &resp.Resp{Result: res})
	}
}
