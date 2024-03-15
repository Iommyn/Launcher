package handle

// VersionResponse Возвращает статус валидный лаунчер версии или нет
// swagger:response launcherVersionResponse
type launcherVersionResponseWrapper struct {
	Body struct {
		*launcherVersionResponse
	}
}
