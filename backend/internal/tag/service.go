package tag

import (
	"context"

	"github.com/repeale/fp-go"

	"github.com/imartinezalberte/go-todo-list/backend/internal/db_utils"
	"github.com/imartinezalberte/go-todo-list/backend/server"
)

type (
	Service interface {
		GetTags(context.Context, TagsQuery) (TagsResponseDTO, error)
		GetTag(context.Context, TagQuery) (TagResponseDTO, error)
		AddTag(context.Context, AddTagCommand) (TagResponseDTO, error)
		UpdateTag(context.Context, UpdateTagCommand) (TagResponseDTO, error)
		DelTag(context.Context, DelTagCommand) error
	}

	service struct {
		r Repo
	}
)

func NewService(r Repo) Service {
	return &service{r}
}

func Execute[T server.Commander](
	svc Service,
	ctx context.Context,
	commander T,
) (result any, err error) {
	cmd, err := commander.ToCmd()
	if err != nil {
		return result, err
	}

	switch cmd := any(cmd).(type) {
	case TagQuery:
		return svc.GetTag(ctx, cmd)
	case TagsQuery:
		return svc.GetTags(ctx, cmd)
	case AddTagCommand:
		return svc.AddTag(ctx, cmd)
	case UpdateTagCommand:
		return svc.UpdateTag(ctx, cmd)
	case DelTagCommand:
		return result, svc.DelTag(ctx, cmd)
	}

	return result, nil
}

func (s *service) GetTags(ctx context.Context, query TagsQuery) (TagsResponseDTO, error) {
	r, err := s.r.GetTags(ctx, query)
	if err != nil {
		return TagsResponseDTO{}, err
	}

	return fp.Map(
		func(t Tag) TagResponseDTO { return TagResponseDTO{db_utils.Model{ID: t.ID}} },
	)(
		r,
	), nil
}

func (s *service) GetTag(ctx context.Context, query TagQuery) (TagResponseDTO, error) {
	r, err := s.r.GetTag(ctx, query)
	if err != nil {
		return TagResponseDTO{}, err
	}

	return TagResponseDTO(r), nil
}

func (s *service) AddTag(ctx context.Context, cmd AddTagCommand) (TagResponseDTO, error) {
	r, err := s.r.AddTag(ctx, cmd.Tag)
	if err != nil {
		return TagResponseDTO{}, err
	}

	return TagResponseDTO(r), nil
}

func (s *service) UpdateTag(ctx context.Context, cmd UpdateTagCommand) (TagResponseDTO, error) {
	r, err := s.r.UpdateTag(ctx, cmd.Tag)
	if err != nil {
		return TagResponseDTO{}, err
	}

	return s.GetTag(ctx, TagQuery{r.ID})
}

func (s *service) DelTag(
	ctx context.Context,
	cmd DelTagCommand,
) error {
	return s.r.DelTag(ctx, Tag{ID: cmd.ID})
}
