package www

func (s *Service) SetRoutes() {

	// User
	s.HTTP.Router().POST("/www/user/create", s.userCreate)
	s.HTTP.Router().POST("/www/token/create", s.tokenCreate)

}
