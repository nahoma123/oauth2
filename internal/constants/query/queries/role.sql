-- name: CreateRole :one 
with _tenant as(
    select tenants.id as tenant_id from tenants where tenants.tenant_name=$1 AND tenants.service_id=$2
), _role as(
    insert into roles(name, tenant_id) select $3,tenant_id from _tenant
    returning  id as role_id,name,created_at,status
),_rp as(
insert into role_permissions (role_id,permission_id)
    select role_id,permissions.id from _role, permissions where permissions.id =ANY($4::uuid[])ON CONFLICT DO NOTHING returning id
)select _role.* from _role;

-- name: GetRoleByNameAndTenantName :one 
SELECT roles.id FROM roles join tenants on roles.tenant_id =tenants.id WHERE 
roles.name = $1 AND tenants.tenant_name = $2;


-- name: AssignRole :exec
insert into tenant_users_roles(tenant_id, user_id, role_id)
select tenants.id,users.id,$1
from tenants,users  where tenants.tenant_name=$2
and users.user_id=$3;

-- name: IsRoleAssigned :one 
SELECT count_rows() FROM tenant_users_roles 
WHERE tenant_users_roles.tenant_id in (
    SELECT tenants.id FROM 
    tenants where tenants.tenant_name = $1
)
and tenant_users_roles.user_id in (
    SELECT users.id from users 
    where users.user_id = $2
) and tenant_users_roles.role_id = $3;