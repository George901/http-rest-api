package apiserver

type APIServer struct {

}

func New() *APIServer {
	return &APIServer{}
}

func (server *APIServer) Start() error {
	return nil
}
