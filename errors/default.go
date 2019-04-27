package errors

import "errors"

var (
	ErrCacheKey = errors.New("缓存的key不存在")
	ErrCacheValue = errors.New("缓存的value不存在")
	ErrUpdateForbid	= errors.New("更新条件不允许")

	ErrAllowController = errors.New("不允许访问的controller")
	ErrInterfaceAssert = errors.New("接口类型断言失败")

	ErrAccessUrl = errors.New("不允许访问的url")
)
