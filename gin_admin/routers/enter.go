package routers

import "project/allfunc/gin_admin/routers/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
