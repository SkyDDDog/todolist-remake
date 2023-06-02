package rpc

import (
	"context"

	"todolist-remake/idl/pb/task"
	"todolist-remake/pkg/errno"
)

func TaskCreate(ctx context.Context, req *task.CreateRequest) (*task.Task, error) {
	resp, err := TaskClient.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}

func TaskUpdate(ctx context.Context, req *task.UpdateRequest) (*task.Task, error) {
	resp, err := TaskClient.Update(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}

func TaskDelete(ctx context.Context, req *task.DeleteRequest) error {
	resp, err := TaskClient.Delete(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return nil
}

func TaskGetList(ctx context.Context, req *task.GetListRequest) ([]*task.Task, error) {
	resp, err := TaskClient.GetList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}

func TaskDone(ctx context.Context, req *task.DoneTaskRequest) (*task.Task, error) {
	resp, err := TaskClient.DoneTask(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}

func TaskUnDone(ctx context.Context, req *task.UnDoneTaskRequest) (*task.Task, error) {
	resp, err := TaskClient.UnDoneTask(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}
