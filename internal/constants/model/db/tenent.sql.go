// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: tenent.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const createTenent = `-- name: CreateTenent :exec
INSERT INTO tenants (
domain_id,
tenant_name,
service_id

) VALUES (
 $1,$2,$3
)
`

type CreateTenentParams struct {
	DomainID   uuid.UUID `json:"domain_id"`
	TenantName string    `json:"tenant_name"`
	ServiceID  uuid.UUID `json:"service_id"`
}

func (q *Queries) CreateTenent(ctx context.Context, arg CreateTenentParams) error {
	_, err := q.db.Exec(ctx, createTenent, arg.DomainID, arg.TenantName, arg.ServiceID)
	return err
}

const getTenentWithNameAndServiceId = `-- name: GetTenentWithNameAndServiceId :one
SELECT id, status, tenant_name, service_id, deleted_at, created_at, updated_at, domain_id, inherit FROM tenants WHERE 
tenant_name = $1 AND service_id = $2 AND deleted_at IS NULL
`

type GetTenentWithNameAndServiceIdParams struct {
	TenantName string    `json:"tenant_name"`
	ServiceID  uuid.UUID `json:"service_id"`
}

func (q *Queries) GetTenentWithNameAndServiceId(ctx context.Context, arg GetTenentWithNameAndServiceIdParams) (Tenant, error) {
	row := q.db.QueryRow(ctx, getTenentWithNameAndServiceId, arg.TenantName, arg.ServiceID)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.TenantName,
		&i.ServiceID,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DomainID,
		&i.Inherit,
	)
	return i, err
}

const isPermissionExistsInTenant = `-- name: IsPermissionExistsInTenant :one
SELECT count_rows() FROM permissions p join tenants t on p.tenant_id=t.id WHERE p.name =$1 and p.service_id=$2 and t.tenant_name=$3 and p.deleted_at IS NULL and p.deleted_at IS NULL
`

type IsPermissionExistsInTenantParams struct {
	Name       string    `json:"name"`
	ServiceID  uuid.UUID `json:"service_id"`
	TenantName string    `json:"tenant_name"`
}

func (q *Queries) IsPermissionExistsInTenant(ctx context.Context, arg IsPermissionExistsInTenantParams) (interface{}, error) {
	row := q.db.QueryRow(ctx, isPermissionExistsInTenant, arg.Name, arg.ServiceID, arg.TenantName)
	var count_rows interface{}
	err := row.Scan(&count_rows)
	return count_rows, err
}

const tenantRegisterPermission = `-- name: TenantRegisterPermission :one
INSERT INTO permissions (name,description,statement,service_id,tenant_id)
SELECT $1,$2,$3,$4,t.id from tenants t where t.tenant_name=$5 and t.deleted_at is null
RETURNING permissions.id,permissions.statement,permissions.description,permissions.name,permissions.created_at,permissions.service_id, $5::string as tenant
`

type TenantRegisterPermissionParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Statement   pgtype.JSON `json:"statement"`
	ServiceID   uuid.UUID   `json:"service_id"`
	TenantName  string      `json:"tenant_name"`
}

type TenantRegisterPermissionRow struct {
	ID          uuid.UUID   `json:"id"`
	Statement   pgtype.JSON `json:"statement"`
	Description string      `json:"description"`
	Name        string      `json:"name"`
	CreatedAt   time.Time   `json:"created_at"`
	ServiceID   uuid.UUID   `json:"service_id"`
	Tenant      string      `json:"tenant"`
}

func (q *Queries) TenantRegisterPermission(ctx context.Context, arg TenantRegisterPermissionParams) (TenantRegisterPermissionRow, error) {
	row := q.db.QueryRow(ctx, tenantRegisterPermission,
		arg.Name,
		arg.Description,
		arg.Statement,
		arg.ServiceID,
		arg.TenantName,
	)
	var i TenantRegisterPermissionRow
	err := row.Scan(
		&i.ID,
		&i.Statement,
		&i.Description,
		&i.Name,
		&i.CreatedAt,
		&i.ServiceID,
		&i.Tenant,
	)
	return i, err
}
