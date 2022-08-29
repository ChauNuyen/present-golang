package commentbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

type DeleteCommentStore interface {
	FindDataByCondition(
		ctx context.Context,
		data *commentmodel.CommentDelete,
		moreKeys ...string,
	) (*commentmodel.Comment, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteCommentBiz struct {
	store DeleteCommentStore
}

func NewDeleteCommentBiz(store DeleteCommentStore) *deleteCommentBiz {
	return &deleteCommentBiz{store: store}
}

func (biz *deleteCommentBiz) DeleteComment(ctx context.Context, data *commentmodel.CommentDelete) error {
	oldData, err := biz.store.FindDataByCondition(ctx, data)

	if err != nil {
		return common.ErrCannotGetEntity(commentmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(commentmodel.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, data.CommentId); err != nil {
		return common.ErrCannotDeleteEntity(commentmodel.EntityName, err)
	}

	return nil
}
