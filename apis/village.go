package apis

import (
	"github.com/ambikuk/go-api-marketplace/app"
	"github.com/ambikuk/go-api-marketplace/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// villageService specifies the interface for the village service needed by villageResource.
	villageService interface {
		Query(rs app.RequestScope, offset, limit int, districtID int) ([]models.Village, error)
		Count(rs app.RequestScope, districtID int) (int, error)
	}

	// villageResource defines the handlers for the CRUD APIs.
	villageResource struct {
		service villageService
	}
)

// ServeVillageResource sets up the routing of village endpoints and the corresponding handlers.
func ServeVillageResource(rg *routing.RouteGroup, service villageService) {
	r := &villageResource{service}
	rg.Get("/villages", r.query)
}

func (r *villageResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	districtID := parseInt(c.Query("district_id"), 0)
	count, err := r.service.Count(rs, districtID)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit(), districtID)
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}