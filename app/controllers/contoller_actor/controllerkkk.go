package contoller_actor

//
//type ControllerActorInterface interface {
//	CreateActor(ctx context.Context, req RequestActor) (origin.DefaultResponse, int, origin.DefaultResponse)
//	GetActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse)
//	GetAllActor(ctx context.Context, page uint64, username string) (origin.DefaultResponse, int, origin.DefaultResponse)
//	UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (origin.DefaultResponse, int, origin.DefaultResponse)
//	DeleteActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse)
//	ActivateActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse)
//	DeactivateActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse)
//
//	LoginActor(ctx context.Context, req RequestActor, agent string) (origin.DefaultResponse, int, origin.DefaultResponse)
//}
//
//type actorControllerStruct struct {
//	actorRepository RepositoryActorInterface
//}
//
//func (c actorControllerStruct) CreateActor(ctx context.Context, req RequestActor) (origin.DefaultResponse, int, origin.DefaultResponse) {

//}
//
//func (c actorControllerStruct) GetActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse) {
//	var actorRepo model.Actor
//	var response origin.DefaultResponse
//	var errorMessage origin.DefaultResponse
//
//	//get data by id
//	status, err := c.actorRepository.GetActorById(ctx, id, &actorRepo)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//	response = origin.DefaultSuccessResponseWithMessage("Get repository-entity_actor", status, actorRepo)
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) GetAllActor(ctx context.Context, page uint64, username string) (origin.DefaultResponse, int, origin.DefaultResponse) {
//	var actorRepo []model.Actor
//	var actorCountRepo model.Actor
//	var response origin.DefaultResponse
//	var errorMessage origin.DefaultResponse
//
//	var limit uint64 = 30
//	status, err := c.actorRepository.GetCountRowsActor(ctx, &actorCountRepo)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	status, err = c.actorRepository.GetAllActor(ctx, page, limit, username, &actorRepo)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	resMessage := FindAllActor{
//		Page:       page,
//		PerPage:    uint64(len(actorRepo)),
//		TotalPages: uint64(math.Ceil(float64(actorCountRepo.Total) / float64(limit))),
//		Data:       actorRepo,
//	}
//
//	response = origin.DefaultSuccessResponseWithMessage("Get all repository-entity_actor", status, resMessage)
//
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (origin.DefaultResponse, int, origin.DefaultResponse) {
//	var actorRepo model.Actor
//	var response origin.DefaultResponse
//	var errorMessage origin.DefaultResponse
//
//	//check authorization
//	if id == 1 {
//		errorMessage = origin.DefaultErrorResponseWithMessage("not authorization update", http.StatusUnauthorized)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//update data by id
//	status, err := c.actorRepository.UpdateActorById(ctx, id, req)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	//repo
//	status, err = c.actorRepository.GetActorById(ctx, id, &actorRepo)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	response = origin.DefaultSuccessResponseWithMessage("Get repository-entity_actor", status, actorRepo)
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) DeleteActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse) {
//	var response origin.DefaultResponse
//	var errorMessage origin.DefaultResponse
//	//check authorization
//	if id == 1 {
//		errorMessage = origin.DefaultErrorResponseWithMessage("not authorization delete", http.StatusUnauthorized)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//repo
//	status, err := c.actorRepository.DeleteActorById(ctx, id)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//	response = origin.DefaultSuccessResponseWithMessage("delete repository-entity_actor", status, "true")
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) ActivateActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse) {
//	var response origin.DefaultResponse
//	var errorMessage origin.DefaultResponse
//
//	//repo
//	status, err := c.actorRepository.ActivateActorById(ctx, id)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	response = origin.DefaultSuccessResponseWithMessage("delete repository-entity_actor", status, "true")
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) DeactivateActorById(ctx context.Context, id uint64) (origin.DefaultResponse, int, origin.DefaultResponse) {
//	var response origin.DefaultResponse
//	var errorMessage origin.DefaultResponse
//	//check authorization
//	if id == 1 {
//		errorMessage = origin.DefaultErrorResponseWithMessage("not authorization deactivate", http.StatusUnauthorized)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//repo
//	status, err := c.actorRepository.DeactivateActorById(ctx, id)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//	response = origin.DefaultSuccessResponseWithMessage("delete repository-entity_actor", status, "true")
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) LoginActor(ctx context.Context, req RequestActor, agent string) (origin.DefaultResponse, int, origin.DefaultResponse) {
//	var actorRepo model.Actor
//	var response origin.DefaultResponse
//	var errorMessage origin.DefaultResponse
//
//	status, err := c.actorRepository.LoginActor(ctx, req, &actorRepo)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(req.Password))
//	if err != nil {
//		// invalid password
//		errorMessage = origin.DefaultErrorResponseWithMessage("invalid username & password", status)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//check access
//	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
//		errorMessage = origin.DefaultErrorResponseWithMessage("account not activated", status)
//		return response, http.StatusForbidden, errorMessage
//	}
//
//	tokenJWTAccess, _, status, err := c.actorRepository.GenerateJWT(actorRepo, agent)
//	if err != nil {
//		errorMessage = origin.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, http.StatusForbidden, errorMessage
//	}
//
//	response = origin.DefaultSuccessResponseWithMessage("login success", status, tokenJWTAccess)
//	return response, status, errorMessage
//}
