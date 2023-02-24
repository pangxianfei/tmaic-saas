package cache

import (
	"tmaic/vendors/framework/cache/driver"
)

type Cacher interface {
	driver.ProtoCache
	driver.BasicCache
}
