package api

type ProxyApi interface {
	Callback(ctx interface{}) error

	// ProxyForwarder  godoc
	// @Summary Show account
	// @Description get string by ID
	// @ID proxy
	// @Accept  json
	// @Produce  json
	// @Param id path int true "Account ID"
	// @Success 200 {string} Token "1111"
	// @Header 200 {string} Token "qwerty"
	// @Router /proxy/{path} [get]
	ProxyForwarder(ctx interface{}) error
}
