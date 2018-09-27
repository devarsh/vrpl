package util

import (
	"encoding/json"
	jwtModel "github.com/devarsh/vrpl/jwt/model"
	masterModel "github.com/devarsh/vrpl/master/model"
	miscModel "github.com/devarsh/vrpl/misc/model"
	userModel "github.com/devarsh/vrpl/user/model"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetReqUrlParamsToJson(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]string)
		vals := r.URL.Query()
		for key, val := range vals {
			data[key] = val[0]
		}
		rctx := chi.RouteContext(r.Context())
		for index, key := range rctx.URLParams.Keys {
			data[key] = rctx.URLParams.Values[index]
		}
		jsonStr, _ := json.Marshal(data)
		reader := strings.NewReader(string(jsonStr))
		x := ioutil.NopCloser(reader)
		r.Body = x
		handle.ServeHTTP(w, r)
	}
}

func CreateTables(mydb *gorm.DB) {
	if mydb == nil {
		panic("DB not initialized")
	}
	usrTbl := userModel.User{}
	if !mydb.HasTable(usrTbl) {
		mydb.CreateTable(usrTbl)
	}
	jwtTbl := jwtModel.JwtPersist{}
	if !mydb.HasTable(jwtTbl) {
		mydb.CreateTable(jwtTbl).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	}
	typesTbl := miscModel.Types{}
	if !mydb.HasTable(typesTbl) {
		mydb.CreateTable(typesTbl)
	}
	countryTbl := miscModel.Country{}
	if !mydb.HasTable(countryTbl) {
		mydb.CreateTable(countryTbl)
	}
	stateTbl := miscModel.State{}
	if !mydb.HasTable(stateTbl) {
		mydb.CreateTable(stateTbl).
			AddForeignKey("country_id", "country(id)", "CASCADE", "CASCADE")
	}
	cityTbl := miscModel.City{}
	if !mydb.HasTable(cityTbl) {
		mydb.CreateTable(cityTbl).
			AddForeignKey("state_id", "state(id)", "CASCADE", "CASCADE")
	}
	areaTbl := miscModel.Area{}
	if !mydb.HasTable(areaTbl) {
		mydb.CreateTable(areaTbl).
			AddForeignKey("city_id", "city(id)", "CASCADE", "CASCADE")
	}
	itemTbl := miscModel.Item{}
	if !mydb.HasTable(itemTbl) {
		mydb.CreateTable(itemTbl).
			AddForeignKey("machine_type_id", "types(id)", "CASCADE", "CASCADE")
	}
	clientTbl := masterModel.Client{}
	if !mydb.HasTable(clientTbl) {
		mydb.CreateTable(clientTbl).
			AddForeignKey("group_id", "types(id)", "CASCADE", "CASCADE")
	}
	companyTbl := masterModel.Company{}
	if !mydb.HasTable(companyTbl) {
		mydb.CreateTable(companyTbl)
	}
	employeeTbl := masterModel.Employee{}
	if !mydb.HasTable(employeeTbl) {
		mydb.CreateTable(employeeTbl)
	}
	holidayTbl := masterModel.Holiday{}
	if !mydb.HasTable(holidayTbl) {
		mydb.CreateTable(holidayTbl)
	}
	machineTbl := masterModel.Machine{}
	if !mydb.HasTable(machineTbl) {
		mydb.CreateTable(machineTbl).
			AddForeignKey("client_id", "client(id)", "CASCADE", "CASCADE").
			AddForeignKey("item_id", "item(id)", "CASCADE", "CASCADE").
			AddForeignKey("employee_id", "employee(id)", "CASCADE", "CASCADE")
	}
	contactTbl := miscModel.Contact{}
	if !mydb.HasTable(contactTbl) {
		mydb.CreateTable(contactTbl).
			AddForeignKey("contact_type_id", "types(id)", "CASCADE", "CASCADE").
			AddForeignKey("email_type_id", "types(id)", "CASCADE", "CASCADE").
			AddForeignKey("client_id", "client(id)", "CASCADE", "CASCADE").
			AddForeignKey("employee_id", "employee(id)", "CASCADE", "CASCADE").
			AddForeignKey("company_id", "company(id)", "CASCADE", "CASCADE")

	}
	addressTbl := miscModel.Address{}
	if !mydb.HasTable(addressTbl) {
		mydb.CreateTable(addressTbl).
			AddForeignKey("area_id", "area(id)", "CASCADE", "CASCADE").
			AddForeignKey("city_id", "city(id)", "CASCADE", "CASCADE").
			AddForeignKey("state_id", "state(id)", "CASCADE", "CASCADE").
			AddForeignKey("country_id", "country(id)", "CASCADE", "CASCADE").
			AddForeignKey("client_id", "client(id)", "CASCADE", "CASCADE").
			AddForeignKey("address_type_id", "types(id)", "CASCADE", "CASCADE").
			AddForeignKey("company_id", "company(id)", "CASCADE", "CASCADE").
			AddForeignKey("employee_id", "employee(id)", "CASCADE", "CASCADE")
	}

}
