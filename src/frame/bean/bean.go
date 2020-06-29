package bean

var controllerMap = make(map[string]Controller)

func Register(uri string, bean Controller) {
	controller := controllerMap[uri]
	if controller != nil {
		panic("uri [" + uri + "]不可以重复")
	}
	controllerMap[uri] = bean
}

func GetControllerByUri(uri string) Controller {
	controller := controllerMap[uri]
	return controller
}
