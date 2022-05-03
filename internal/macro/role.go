//go:generate go run github.com/dmarkham/enumer -type Role -trimprefix Role -transform upper -sql -json -output role_string.go
package macro

type Role int

const (
	_ Role = iota
	RoleD1
	RoleD2
	RoleD3
	RoleD4
	RoleMT
	RoleST
	RoleH1
	RoleH2
)

func (i Role) Values() []string {
	return RoleStrings()
}
