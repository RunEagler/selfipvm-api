package controller

import (
	"github.com/goadesign/goa"
	"selfipvm-api/app"
	"selfipvm-api/repository"
)

func CommonController(service *goa.Service,activityRepository repository.ActivityRepository){

	// Mount "day_task" controller
	c := NewActivityController(service,activityRepository)
	app.MountActivityController(service, c)
}



