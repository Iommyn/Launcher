package handle

// VersionResponse Возвращает статус валидный лаунчер версии или нет
// swagger:response launcherVersionResponse
type launcherVersionResponseWrapper struct {
	// in:body
	Body struct {
		// Экземпляр launcherVersionResponse
		*launcherVersionResponse
	}
}
