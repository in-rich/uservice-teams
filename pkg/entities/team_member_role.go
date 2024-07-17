package entities

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type MemberRole string

const (
	MemberRoleMember MemberRole = "member"
	MemberRoleAdmin  MemberRole = "admin"
)

var _ sql.Scanner = (*MemberRole)(nil)
var _ driver.Valuer = (*MemberRole)(nil)

func (role MemberRole) Valid() bool {
	switch role {
	case MemberRoleMember, MemberRoleAdmin:
		return true
	default:
		return false
	}
}

func (role *MemberRole) Scan(src interface{}) error {
	switch tsrc := src.(type) {
	case string:
		*role = MemberRole(tsrc)
		if !role.Valid() {
			return fmt.Errorf("invalid role: %q", tsrc)
		}
		return nil
	case []byte:
		*role = MemberRole(tsrc)
		if !role.Valid() {
			return fmt.Errorf("invalid role: %q", tsrc)
		}
		return nil
	case nil:
		return fmt.Errorf("scanning nil into MemberRole")
	default:
		return fmt.Errorf("unsupported data type for MemberRole: %T", src)
	}
}

func (role MemberRole) Value() (driver.Value, error) {
	if !role.Valid() {
		return nil, fmt.Errorf("invalid role: %q", role)
	}
	return string(role), nil
}
