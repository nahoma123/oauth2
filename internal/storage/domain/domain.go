package domain

import (
	"2f-authorization/internal/constants/dbinstance"
	errors "2f-authorization/internal/constants/error"
	"2f-authorization/internal/constants/error/sqlcerr"
	"2f-authorization/internal/constants/model/db"
	"2f-authorization/internal/constants/model/dto"
	"2f-authorization/internal/storage"
	"2f-authorization/platform/logger"

	"context"

	"go.uber.org/zap"
)

type domain struct {
	db  dbinstance.DBInstance
	log logger.Logger
}

func Init(db dbinstance.DBInstance, log logger.Logger) storage.Domain {
	return &domain{
		db:  db,
		log: log,
	}
}

func (d *domain) CreateDomain(ctx context.Context, param dto.Domain) (*dto.Domain, error) {
	domain, err := d.db.CreateDomain(ctx, db.CreateDomainParams{
		Name:      param.Name,
		ServiceID: param.ServiceID,
	})
	if err != nil {
		err := errors.ErrWriteError.Wrap(err, "could not create domain")
		d.log.Error(ctx, "unable to create domain ", zap.Error(err), zap.Any("domain", param))
		return &dto.Domain{}, err
	}
	return &dto.Domain{
		ID:        domain.ID,
		Name:      domain.Name,
		ServiceID: domain.ID,
		CreatedAt: domain.CreatedAt,
	}, nil

}
func (d *domain) IsDomainExist(ctx context.Context, param dto.Domain) (bool, error) {

	_, err := d.db.IsDomainExist(ctx, db.IsDomainExistParams{ServiceID: param.ServiceID, Name: param.Name})
	if err != nil {
		if sqlcerr.Is(err, sqlcerr.ErrNoRows) {
			return false, nil
		} else {
			err := errors.ErrReadError.Wrap(err, "could not  read domain")
			d.log.Error(ctx, "unable to read the domain", zap.Error(err), zap.Any("domain ", param))
			return false, err

		}

	}
	return true, nil
}
