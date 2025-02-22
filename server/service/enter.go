package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Booking"
	"github.com/flipped-aurora/gin-vue-admin/server/service/EmailMarketing"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	EmailmarketingServiceGroup EmailMarketing.ServiceGroup
	BookingServiceGroup        Booking.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
