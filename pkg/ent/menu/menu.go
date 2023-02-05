// Code generated by ent, DO NOT EDIT.

package menu

import (
	"time"
)

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldMenuLevel holds the string denoting the menu_level field in the database.
	FieldMenuLevel = "menu_level"
	// FieldMenuType holds the string denoting the menu_type field in the database.
	FieldMenuType = "menu_type"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRedirect holds the string denoting the redirect field in the database.
	FieldRedirect = "redirect"
	// FieldComponent holds the string denoting the component field in the database.
	FieldComponent = "component"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldHideMenu holds the string denoting the hide_menu field in the database.
	FieldHideMenu = "hide_menu"
	// FieldHideBreadcrumb holds the string denoting the hide_breadcrumb field in the database.
	FieldHideBreadcrumb = "hide_breadcrumb"
	// FieldCurrentActiveMenu holds the string denoting the current_active_menu field in the database.
	FieldCurrentActiveMenu = "current_active_menu"
	// FieldIgnoreKeepAlive holds the string denoting the ignore_keep_alive field in the database.
	FieldIgnoreKeepAlive = "ignore_keep_alive"
	// FieldHideTab holds the string denoting the hide_tab field in the database.
	FieldHideTab = "hide_tab"
	// FieldFrameSrc holds the string denoting the frame_src field in the database.
	FieldFrameSrc = "frame_src"
	// FieldCarryParam holds the string denoting the carry_param field in the database.
	FieldCarryParam = "carry_param"
	// FieldHideChildrenInMenu holds the string denoting the hide_children_in_menu field in the database.
	FieldHideChildrenInMenu = "hide_children_in_menu"
	// FieldAffix holds the string denoting the affix field in the database.
	FieldAffix = "affix"
	// FieldDynamicLevel holds the string denoting the dynamic_level field in the database.
	FieldDynamicLevel = "dynamic_level"
	// FieldRealPath holds the string denoting the real_path field in the database.
	FieldRealPath = "real_path"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParams holds the string denoting the params edge name in mutations.
	EdgeParams = "params"
	// Table holds the table name of the menu in the database.
	Table = "sys_menus"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "role_menus"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "sys_roles"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_menus"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "sys_menus"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// ParamsTable is the table that holds the params relation/edge.
	ParamsTable = "sys_menu_params"
	// ParamsInverseTable is the table name for the MenuParam entity.
	// It exists in this package in order to avoid circular dependency with the "menuparam" package.
	ParamsInverseTable = "sys_menu_params"
	// ParamsColumn is the table column denoting the params relation/edge.
	ParamsColumn = "menu_params"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSort,
	FieldParentID,
	FieldMenuLevel,
	FieldMenuType,
	FieldPath,
	FieldName,
	FieldRedirect,
	FieldComponent,
	FieldDisabled,
	FieldTitle,
	FieldIcon,
	FieldHideMenu,
	FieldHideBreadcrumb,
	FieldCurrentActiveMenu,
	FieldIgnoreKeepAlive,
	FieldHideTab,
	FieldFrameSrc,
	FieldCarryParam,
	FieldHideChildrenInMenu,
	FieldAffix,
	FieldDynamicLevel,
	FieldRealPath,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"role_id", "menu_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort uint32
	// DefaultPath holds the default value on creation for the "path" field.
	DefaultPath string
	// DefaultRedirect holds the default value on creation for the "redirect" field.
	DefaultRedirect string
	// DefaultComponent holds the default value on creation for the "component" field.
	DefaultComponent string
	// DefaultDisabled holds the default value on creation for the "disabled" field.
	DefaultDisabled bool
	// DefaultHideMenu holds the default value on creation for the "hide_menu" field.
	DefaultHideMenu bool
	// DefaultHideBreadcrumb holds the default value on creation for the "hide_breadcrumb" field.
	DefaultHideBreadcrumb bool
	// DefaultCurrentActiveMenu holds the default value on creation for the "current_active_menu" field.
	DefaultCurrentActiveMenu string
	// DefaultIgnoreKeepAlive holds the default value on creation for the "ignore_keep_alive" field.
	DefaultIgnoreKeepAlive bool
	// DefaultHideTab holds the default value on creation for the "hide_tab" field.
	DefaultHideTab bool
	// DefaultFrameSrc holds the default value on creation for the "frame_src" field.
	DefaultFrameSrc string
	// DefaultCarryParam holds the default value on creation for the "carry_param" field.
	DefaultCarryParam bool
	// DefaultHideChildrenInMenu holds the default value on creation for the "hide_children_in_menu" field.
	DefaultHideChildrenInMenu bool
	// DefaultAffix holds the default value on creation for the "affix" field.
	DefaultAffix bool
	// DefaultDynamicLevel holds the default value on creation for the "dynamic_level" field.
	DefaultDynamicLevel uint32
	// DefaultRealPath holds the default value on creation for the "real_path" field.
	DefaultRealPath string
)
