package master

import (
	masterDB "github.com/devarsh/vrpl/master/db"
	"github.com/devarsh/vrpl/util"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type MasterManager struct {
	db *masterDB.MasterDb
}

func NewMasterManager(mydb *gorm.DB) *MasterManager {
	if mydb == nil {
		panic("Nil Db instance passed")
	}
	mymasterDb := masterDB.NewMasterDb(mydb)
	return &MasterManager{db: mymasterDb}
}

func (mm *MasterManager) Routers() chi.Router {
	r := chi.NewRouter()
	r.Route("/client", func(r chi.Router) {
		r.Get("/{id}", util.GetReqUrlParamsToJson(mm.GetClientByID))
		r.Get("/byGroup/{id}", util.GetReqUrlParamsToJson(mm.GetClientByGroupID))
		r.Get("/byName/{name}", util.GetReqUrlParamsToJson(mm.GetClientByName))
		r.Get("/byTallyName/{name}", util.GetReqUrlParamsToJson(mm.GetClientByTallyName))
		r.Post("/", mm.AddClient)
		r.Put("/", mm.UpdateClient)
	})
	r.Route("/company", func(r chi.Router) {
		r.Get("/", mm.GetAllCompanies)
		r.Get("/{id}", util.GetReqUrlParamsToJson(mm.GetCompanyByID))
		r.Post("/", mm.AddCompany)
		r.Put("/", mm.UpdateCompany)
	})
	r.Route("/employee", func(r chi.Router) {
		r.Get("/", mm.GetAllEmployees)
		r.Get("/{id}", util.GetReqUrlParamsToJson(mm.GetEmployeeByID))
		r.Post("/", mm.AddEmployee)
		r.Put("/", mm.UpdateEmployee)
	})
	r.Route("/holiday", func(r chi.Router) {
		r.Get("/", mm.GetAllHolidays)
		r.Get("/{id}", util.GetReqUrlParamsToJson(mm.GetHolidayByID))
		r.Post("/", mm.AddHoliday)
		r.Put("/", mm.UpdateHoliday)
	})
	r.Route("/machine", func(r chi.Router) {
		r.Get("/byEmployee/{id}", util.GetReqUrlParamsToJson(mm.GetMachinesByEmployeeID))
		r.Get("/byClient/{id}", util.GetReqUrlParamsToJson(mm.GetMachinesByClientID))
		r.Get("/{id}", util.GetReqUrlParamsToJson(mm.GetMachineByID))
		r.Post("/", mm.AddMachine)
		r.Put("/", mm.UpdateMachine)
	})
	return r
}
