package controller

import (
	"github.com/goadesign/goa"
	"selfipvm-api/app"
	"selfipvm-api/repository"
	"selfipvm-api/models"
	"net/http"
	"selfipvm-api/common"
	"selfipvm-api/error_message"
	"time"
)

// ActivityController implements the activity resource.
type ActivityController struct {
	*goa.Controller
	activeRepository repository.ActivityRepository
	service          *goa.Service
	request          *http.Request
}

// NewActivityController creates a activity controller.
func NewActivityController(service *goa.Service, activeRepository repository.ActivityRepository) *ActivityController {
	return &ActivityController{
		Controller:       service.NewController("ActivityController"),
		activeRepository: activeRepository,
	}
}

// Entry runs the entry action.
func (c *ActivityController) Entry(ctx *app.EntryActivityContext) error {
	// ActivityController_Entry: start_implement
	// Put your logic here
	c.service = ctx.Service
	c.request = ctx.Request

	statusCode := c.entry(ctx)
	switch statusCode {
	case http.StatusOK:
		return ctx.OK(nil)
	case http.StatusBadRequest:
		return ctx.BadRequest()
	case http.StatusInternalServerError:
		return ctx.InternalServerError()
	default:
		return ctx.InternalServerError()
	}
	// ActivityController_Entry: end_implement
}

func (c *ActivityController) entry(ctx *app.EntryActivityContext) int {

	var err error
	var date *time.Time

	date,err = common.ConvertStringToDate(ctx.Payload.Date)
	if err != nil{
		return http.StatusBadRequest
	}
	err = c.activeRepository.InsertActivity(&models.Activity{
		Minutes:  ctx.Payload.Minutes,
		Content:  common.ToNullString(ctx.Payload.Content),
		Date:     *date,
		Type: ctx.Payload.Type,
	})
	if err != nil {
		common.LogError(c.service,c.request,error_message.ActivityInsert.String(),"err",err)
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

