package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Booking"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/EmailMarketing"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	EmailmarketingApiGroup EmailMarketing.ApiGroup
	BookingApiGroup        Booking.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
