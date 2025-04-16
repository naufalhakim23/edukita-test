package service

type (
	ILearningManagementService interface{}
	LearningManagementService  struct {
		ServiceOption
	}
)

func InitiateLearningManagementService(opt ServiceOption) ILearningManagementService {
	return &LearningManagementService{
		ServiceOption: opt,
	}
}
