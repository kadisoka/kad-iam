package iam

type Service interface {
	UserService() UserService
	ApplicationService() ApplicationService
}
