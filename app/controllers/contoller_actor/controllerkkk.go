package contoller_actor

//
//type ControllerActorInterface interface {
//	CreateActor(ctx context.Context, req RequestActor) (original.DefaultResponse, int, original.DefaultResponse)
//	GetActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse)
//	GetAllActor(ctx context.Context, page uint64, username string) (original.DefaultResponse, int, original.DefaultResponse)
//	UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (original.DefaultResponse, int, original.DefaultResponse)
//	DeleteActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse)
//	ActivateActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse)
//	DeactivateActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse)
//
//	LoginActor(ctx context.Context, req RequestActor, agent string) (original.DefaultResponse, int, original.DefaultResponse)
//}
//
//type actorControllerStruct struct {
//	actorRepository RepositoryActorInterface
//}
//
//func (c actorControllerStruct) CreateActor(ctx context.Context, req RequestActor) (original.DefaultResponse, int, original.DefaultResponse) {
//	var actorRepo model.Actor
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//
//	//hashing password
//	hashingPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
//	reqActor := RequestActor{
//		Username: req.Username,
//		Password: string(hashingPassword),
//	}
//
//	// create repository-entity_actor
//	status, err := c.actorRepository.CreateActor(ctx, &reqActor)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	//get data
//	status, err = c.actorRepository.GetActorByUsername(ctx, reqActor, &actorRepo)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	//req approval
//	reqApproval := RequestApproval{
//		ID: actorRepo.ID,
//	}
//
//	//create approval
//	status, err = c.actorRepository.CreateApproval(ctx, &reqApproval)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	response = original.DefaultSuccessResponseWithMessage("repository-entity_actor created", status, actorRepo)
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) GetActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse) {
//	var actorRepo model.Actor
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//
//	//get data by id
//	status, err := c.actorRepository.GetActorById(ctx, id, &actorRepo)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//	response = original.DefaultSuccessResponseWithMessage("Get repository-entity_actor", status, actorRepo)
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) GetAllActor(ctx context.Context, page uint64, username string) (original.DefaultResponse, int, original.DefaultResponse) {
//	var actorRepo []model.Actor
//	var actorCountRepo model.Actor
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//
//	var limit uint64 = 30
//	status, err := c.actorRepository.GetCountRowsActor(ctx, &actorCountRepo)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	status, err = c.actorRepository.GetAllActor(ctx, page, limit, username, &actorRepo)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
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
//	response = original.DefaultSuccessResponseWithMessage("Get all repository-entity_actor", status, resMessage)
//
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) UpdateActorById(ctx context.Context, id uint64, req RequestUpdateActor) (original.DefaultResponse, int, original.DefaultResponse) {
//	var actorRepo model.Actor
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//
//	//check authorization
//	if id == 1 {
//		errorMessage = original.DefaultErrorResponseWithMessage("not authorization update", http.StatusUnauthorized)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//update data by id
//	status, err := c.actorRepository.UpdateActorById(ctx, id, req)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	//repo
//	status, err = c.actorRepository.GetActorById(ctx, id, &actorRepo)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	response = original.DefaultSuccessResponseWithMessage("Get repository-entity_actor", status, actorRepo)
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) DeleteActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse) {
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//	//check authorization
//	if id == 1 {
//		errorMessage = original.DefaultErrorResponseWithMessage("not authorization delete", http.StatusUnauthorized)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//repo
//	status, err := c.actorRepository.DeleteActorById(ctx, id)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//	response = original.DefaultSuccessResponseWithMessage("delete repository-entity_actor", status, "true")
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) ActivateActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse) {
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//
//	//repo
//	status, err := c.actorRepository.ActivateActorById(ctx, id)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	response = original.DefaultSuccessResponseWithMessage("delete repository-entity_actor", status, "true")
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) DeactivateActorById(ctx context.Context, id uint64) (original.DefaultResponse, int, original.DefaultResponse) {
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//	//check authorization
//	if id == 1 {
//		errorMessage = original.DefaultErrorResponseWithMessage("not authorization deactivate", http.StatusUnauthorized)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//repo
//	status, err := c.actorRepository.DeactivateActorById(ctx, id)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//	response = original.DefaultSuccessResponseWithMessage("delete repository-entity_actor", status, "true")
//	return response, status, errorMessage
//}
//
//func (c actorControllerStruct) LoginActor(ctx context.Context, req RequestActor, agent string) (original.DefaultResponse, int, original.DefaultResponse) {
//	var actorRepo model.Actor
//	var response original.DefaultResponse
//	var errorMessage original.DefaultResponse
//
//	status, err := c.actorRepository.LoginActor(ctx, req, &actorRepo)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, status, errorMessage
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(actorRepo.Password), []byte(req.Password))
//	if err != nil {
//		// invalid password
//		errorMessage = original.DefaultErrorResponseWithMessage("invalid username & password", status)
//		return response, http.StatusUnauthorized, errorMessage
//	}
//
//	//check access
//	if actorRepo.Verified != "true" && actorRepo.Active != "true" {
//		errorMessage = original.DefaultErrorResponseWithMessage("account not activated", status)
//		return response, http.StatusForbidden, errorMessage
//	}
//
//	tokenJWTAccess, _, status, err := c.actorRepository.GenerateJWT(actorRepo, agent)
//	if err != nil {
//		errorMessage = original.DefaultErrorResponseWithMessage(err.Error(), status)
//		return response, http.StatusForbidden, errorMessage
//	}
//
//	response = original.DefaultSuccessResponseWithMessage("login success", status, tokenJWTAccess)
//	return response, status, errorMessage
//}
