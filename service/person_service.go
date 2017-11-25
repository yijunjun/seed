package service

import (
	"seed/pb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// PersonService 实现person服务
type PersonService struct {
}

// Get 根据id获取对象
func (ps *PersonService) Get(context.Context, *pb.IdRequest) (*pb.Person, error){
return nil, nil
}

// Delete 根据id删除对象
func (ps *PersonService) Delete(context.Context, *pb.IdRequest) (*pb.Person, error){
return nil, nil
}

// Add 插入对象
func (ps *PersonService) Add(context.Context, *pb.Person) (*pb.Person, error){
return nil, nil
}

// Update 更新对象
func (ps *PersonService) Update(context.Context, *pb.Person) (*pb.Person, error){
return nil, nil
}

// List 根据id获取对象列表
func (ps *PersonService) List(context.Context, *pb.IdListRequest) (*pb.PersonList, error){
return nil, nil
}
// SearchByName 根据名称搜索对象列表
func (ps *PersonService) SearchByName(context.Context, *pb.NameListRequest) (*pb.PersonList, error){
return nil, nil
}

func init(){
	regService(func(s *grpc.Server){
		pb.RegisterPersonServiceServer(s, &PersonService{})
	})
}