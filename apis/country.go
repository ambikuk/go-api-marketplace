package apis

import (
	"strconv"

	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// countryService specifies the interface for the country service needed by countryResource.
	countryService interface {
		Get(rs app.RequestScope, id int) (*models.Country, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Country, error)
		Count(rs app.RequestScope) (int, error)
	}

	// countryResource defines the handlers for the CRUD APIs.
	countryResource struct {
		service countryService
	}
)

// ServeCountry sets up the routing of country endpoints and the corresponding handlers.
func ServeCountryResource(rg *routing.RouteGroup, service countryService) {
	r := &countryResource{service}
	rg.Get("/countries/<id>", r.get)
	rg.Get("/countries", r.query)
}

func (r *countryResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *countryResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}