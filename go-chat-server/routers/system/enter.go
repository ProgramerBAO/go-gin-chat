package system

// 这里对外暴露

type RouterGroup struct {
	UserRouter
}

var UseRouterGroup = new(RouterGroup)
