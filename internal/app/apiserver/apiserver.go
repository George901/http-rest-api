package apiserver

type APIServer struct {
	Config *Config
}

func New(config *Config) *APIServer {
	return &APIServer{
		Config: config,
	}
}

func (server *APIServer) Start() error {
	return nil
}
