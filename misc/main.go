package misc

import (
	miscDB "github.com/devarsh/vrpl/misc/db"
	"github.com/devarsh/vrpl/util"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"net/http"
)

type MiscManager struct {
	db *miscDB.MiscDb
}

var (
	ADD_TYPE        = "ADD_TYPE"        //AddressType :- "Residental, Office"
	CONTACT_TYPE    = "CONTACT_TYPE"    //Contact Types :- Mobile, phone, fax
	GROUP_TYPE      = "GROUP_TYPE"      //"Banks Name"
	MACHINE_TYPE    = "MACHINE_TYPE"    //"NCM/ CCM"
	AMC_TYPE        = "AMC_TYPE"        //"Comphrensive/Non Comphrensive"
	PAY_TERMS       = "PAY_TERMS"       //"Before/After"
	PAY_PERIOD_TYPE = "PAY_PERIOD_TYPE" //"Quartely/Half-yearly/Yearly"
)

func NewMiscManager(mydb *gorm.DB) *MiscManager {
	if mydb == nil {
		panic("Nil Db instance passed")
	}

	mymiscDb := miscDB.NewMiscDb(mydb)
	return &MiscManager{db: mymiscDb}
}

func notFound(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("my not found"))
}

func (mm *MiscManager) Routers() chi.Router {
	r := chi.NewRouter()
	r.Route("/country", func(r chi.Router) {
		r.Get("/", util.GetReqUrlParamsToJson(mm.GetAllCountries))
		r.Post("/", mm.AddCountry)
		r.Put("/", mm.UpdateCountry)
	})
	r.Route("/state", func(r chi.Router) {
		r.Get("/byCountry/{id}", util.GetReqUrlParamsToJson(mm.GetAllStatesByCountryID))
		r.Post("/", mm.AddState)
		r.Put("/", mm.UpdateState)
	})
	r.Route("/city", func(r chi.Router) {
		r.Get("/byState/{id}", util.GetReqUrlParamsToJson(mm.GetAllCitiesByStateID))
		r.Post("/", mm.AddCity)
		r.Put("/", mm.UpdateCity)
		r.NotFound(notFound)
	})
	r.Route("/area", func(r chi.Router) {
		r.Get("/byCity/{id}", util.GetReqUrlParamsToJson(mm.GetAllAreasByCityID))
		r.Post("/", mm.AddArea)
		r.Put("/", mm.UpdateArea)
	})
	r.Route("/address", func(r chi.Router) {
		r.Get("/byClient/{id}", util.GetReqUrlParamsToJson(mm.GetAddressByClientID))
		r.Get("/byCompany/{id}", util.GetReqUrlParamsToJson(mm.GetAddressByCompanyID))
		r.Get("/byEmployee/{id}", util.GetReqUrlParamsToJson(mm.GetAddressByEmployeeID))
		r.Post("/", mm.AddAddress)
		r.Put("/", mm.UpdateAddress)
	})
	r.Route("/contact", func(r chi.Router) {
		r.Get("/byClient/{id}", util.GetReqUrlParamsToJson(mm.GetContactByClientID))
		r.Get("/byCompany/{id}", util.GetReqUrlParamsToJson(mm.GetContactByCompanyID))
		r.Get("/byEmployee/{id}", util.GetReqUrlParamsToJson(mm.GetContactByEmployeeID))
		r.Post("/", mm.AddContact)
		r.Put("/", mm.UpdateContact)
	})
	r.Route("/item", func(r chi.Router) {
		r.Get("/{id}", util.GetReqUrlParamsToJson(mm.GetItemByID))
		r.Get("/", util.GetReqUrlParamsToJson(mm.GetAllItems))
		r.Post("/", mm.AddItem)
		r.Put("/", mm.UpdateItem)
	})
	r.Route("/types", func(r1 chi.Router) {
		r1.Route("/address", func(r2 chi.Router) {
			r2.Get("/", util.GetReqUrlParamsToJson(CustomGetAllType(mm, ADD_TYPE)))
			r2.Post("/", CustomAddType(mm, ADD_TYPE))
			r2.Put("/", CustomUpdateType(mm, ADD_TYPE))
		})
		r1.Route("/contact", func(r2 chi.Router) {
			r2.Get("/", util.GetReqUrlParamsToJson(CustomGetAllType(mm, CONTACT_TYPE)))
			r2.Post("/", CustomAddType(mm, CONTACT_TYPE))
			r2.Put("/", CustomUpdateType(mm, CONTACT_TYPE))
		})
		r1.Route("/group", func(r2 chi.Router) {
			r2.Get("/", util.GetReqUrlParamsToJson(CustomGetAllType(mm, GROUP_TYPE)))
			r2.Post("/", CustomAddType(mm, GROUP_TYPE))
			r2.Put("/", CustomUpdateType(mm, GROUP_TYPE))
		})
		r1.Route("/machine", func(r2 chi.Router) {
			r2.Get("/", util.GetReqUrlParamsToJson(CustomGetAllType(mm, MACHINE_TYPE)))
			r2.Post("/", CustomAddType(mm, MACHINE_TYPE))
			r2.Put("/", CustomUpdateType(mm, MACHINE_TYPE))
		})
		r1.Route("/amc", func(r2 chi.Router) {
			r2.Get("/", util.GetReqUrlParamsToJson(CustomGetAllType(mm, AMC_TYPE)))
			r2.Post("/", CustomAddType(mm, AMC_TYPE))
			r2.Put("/", CustomUpdateType(mm, AMC_TYPE))
		})
		r1.Route("/payTerms", func(r2 chi.Router) {
			r2.Get("/", util.GetReqUrlParamsToJson(CustomGetAllType(mm, PAY_TERMS)))
			r2.Post("/", CustomAddType(mm, PAY_TERMS))
			r2.Put("/", CustomUpdateType(mm, PAY_TERMS))
		})
		r1.Route("/payPeriod", func(r2 chi.Router) {
			r2.Get("/", util.GetReqUrlParamsToJson(CustomGetAllType(mm, PAY_PERIOD_TYPE)))
			r2.Post("/", CustomAddType(mm, PAY_PERIOD_TYPE))
			r2.Put("/", CustomUpdateType(mm, PAY_PERIOD_TYPE))
		})
	})
	return r
}
