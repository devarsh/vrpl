package master

import (
	"context"
	"fmt"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/master/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/go-chi/render"
	"net/http"
)

func (mm *MasterManager) AddHoliday(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	holiday := model.Holiday{}
	if err := render.Bind(r, &holiday); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&holiday); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetHolidayByDate(holiday.Date)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	if len(res) > 0 {
		myerr := customErr.ErrInvalidRequestWithValueError(fmt.Errorf("Holidays for the specified date already exists"), res)
		render.Render(w, r, &resp.Resp{Result: myerr})
		return
	}
	holidayID, err := mm.db.AddHoliday(ctx, &holiday)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: holidayID})
}

func (mm *MasterManager) UpdateHoliday(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	holiday := model.Holiday{}
	if err := render.Bind(r, &holiday); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&holiday); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	err := mm.db.UpdateHoliday(ctx, &holiday)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: true})
}

func (mm *MasterManager) GetHolidayByID(w http.ResponseWriter, r *http.Request) {

	request := model.GetChildrensReq{}
	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(&request); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	res, err := mm.db.GetHolidayByID(request.ParentID)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: res})
}

func (mm *MasterManager) GetAllHolidays(w http.ResponseWriter, r *http.Request) {
	holidays, err := mm.db.GetAllHolidays()
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: holidays})
}
