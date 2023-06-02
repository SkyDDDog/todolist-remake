package main

import (
	"context"

	"todolist-remake/cmd/task/pack"
	"todolist-remake/cmd/task/service"
	"todolist-remake/idl/pb/task"
	"todolist-remake/pkg/errno"
)

type TaskServiceImpl struct {
	task.UnimplementedTaskServiceServer
}

func NewTaskServiceImpl() *TaskServiceImpl {
	return &TaskServiceImpl{}
}

func (*TaskServiceImpl) Create(ctx context.Context, req *task.CreateRequest) (resp *task.CreateResponse, err error) {
	resp = new(task.CreateResponse)

	info, err := service.NewTaskService(ctx).Create(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTask(info)
	return resp, nil
}

func (*TaskServiceImpl) Update(ctx context.Context, req *task.UpdateRequest) (resp *task.UpdateResponse, err error) {
	resp = new(task.UpdateResponse)

	info, err := service.NewTaskService(ctx).Update(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTask(info)
	return resp, nil
}

func (*TaskServiceImpl) Delete(ctx context.Context, req *task.DeleteRequest) (resp *task.DeleteResponse, err error) {
	resp = new(task.DeleteResponse)

	err = service.NewTaskService(ctx).Delete(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

func (*TaskServiceImpl) GetList(ctx context.Context, req *task.GetListRequest) (resp *task.GetListResponse, err error) {
	resp = new(task.GetListResponse)

	list, err := service.NewTaskService(ctx).GetList(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTaskList(list)
	return resp, nil
}

func (*TaskServiceImpl) DoneTask(ctx context.Context, req *task.DoneTaskRequest) (resp *task.DoneTaskResponse, err error) {
	resp = new(task.DoneTaskResponse)

	info, err := service.NewTaskService(ctx).DoneTask(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, err
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTask(info)
	return resp, nil
}

func (*TaskServiceImpl) UnDoneTask(ctx context.Context, req *task.UnDoneTaskRequest) (resp *task.UnDoneTaskResponse, err error) {
	resp = new(task.UnDoneTaskResponse)

	info, err := service.NewTaskService(ctx).UnDoneTask(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, err
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTask(info)
	return resp, nil
}
