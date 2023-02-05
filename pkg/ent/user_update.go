// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/department"
	"github.com/suyuan32/simple-admin-core/pkg/ent/post"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(u uint8) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(u)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(u *uint8) *UserUpdate {
	if u != nil {
		uu.SetStatus(*u)
	}
	return uu
}

// AddStatus adds u to the "status" field.
func (uu *UserUpdate) AddStatus(u int8) *UserUpdate {
	uu.mutation.AddStatus(u)
	return uu
}

// ClearStatus clears the value of the "status" field.
func (uu *UserUpdate) ClearStatus() *UserUpdate {
	uu.mutation.ClearStatus()
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetDescription sets the "description" field.
func (uu *UserUpdate) SetDescription(s string) *UserUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDescription(s *string) *UserUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// ClearDescription clears the value of the "description" field.
func (uu *UserUpdate) ClearDescription() *UserUpdate {
	uu.mutation.ClearDescription()
	return uu
}

// SetHomePath sets the "home_path" field.
func (uu *UserUpdate) SetHomePath(s string) *UserUpdate {
	uu.mutation.SetHomePath(s)
	return uu
}

// SetNillableHomePath sets the "home_path" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHomePath(s *string) *UserUpdate {
	if s != nil {
		uu.SetHomePath(*s)
	}
	return uu
}

// SetRoleID sets the "role_id" field.
func (uu *UserUpdate) SetRoleID(u uint64) *UserUpdate {
	uu.mutation.ResetRoleID()
	uu.mutation.SetRoleID(u)
	return uu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRoleID(u *uint64) *UserUpdate {
	if u != nil {
		uu.SetRoleID(*u)
	}
	return uu
}

// AddRoleID adds u to the "role_id" field.
func (uu *UserUpdate) AddRoleID(u int64) *UserUpdate {
	uu.mutation.AddRoleID(u)
	return uu
}

// ClearRoleID clears the value of the "role_id" field.
func (uu *UserUpdate) ClearRoleID() *UserUpdate {
	uu.mutation.ClearRoleID()
	return uu
}

// SetMobile sets the "mobile" field.
func (uu *UserUpdate) SetMobile(s string) *UserUpdate {
	uu.mutation.SetMobile(s)
	return uu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMobile(s *string) *UserUpdate {
	if s != nil {
		uu.SetMobile(*s)
	}
	return uu
}

// ClearMobile clears the value of the "mobile" field.
func (uu *UserUpdate) ClearMobile() *UserUpdate {
	uu.mutation.ClearMobile()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetDepartmentID sets the "department_id" field.
func (uu *UserUpdate) SetDepartmentID(u uint64) *UserUpdate {
	uu.mutation.SetDepartmentID(u)
	return uu
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDepartmentID(u *uint64) *UserUpdate {
	if u != nil {
		uu.SetDepartmentID(*u)
	}
	return uu
}

// ClearDepartmentID clears the value of the "department_id" field.
func (uu *UserUpdate) ClearDepartmentID() *UserUpdate {
	uu.mutation.ClearDepartmentID()
	return uu
}

// SetPostID sets the "post_id" field.
func (uu *UserUpdate) SetPostID(u uint64) *UserUpdate {
	uu.mutation.SetPostID(u)
	return uu
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePostID(u *uint64) *UserUpdate {
	if u != nil {
		uu.SetPostID(*u)
	}
	return uu
}

// ClearPostID clears the value of the "post_id" field.
func (uu *UserUpdate) ClearPostID() *UserUpdate {
	uu.mutation.ClearPostID()
	return uu
}

// SetDepartment sets the "department" edge to the Department entity.
func (uu *UserUpdate) SetDepartment(d *Department) *UserUpdate {
	return uu.SetDepartmentID(d.ID)
}

// SetPost sets the "post" edge to the Post entity.
func (uu *UserUpdate) SetPost(p *Post) *UserUpdate {
	return uu.SetPostID(p.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearDepartment clears the "department" edge to the Department entity.
func (uu *UserUpdate) ClearDepartment() *UserUpdate {
	uu.mutation.ClearDepartment()
	return uu
}

// ClearPost clears the "post" edge to the Post entity.
func (uu *UserUpdate) ClearPost() *UserUpdate {
	uu.mutation.ClearPost()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeUint8, value)
	}
	if uu.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeUint8)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uu.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uu.mutation.HomePath(); ok {
		_spec.SetField(user.FieldHomePath, field.TypeString, value)
	}
	if value, ok := uu.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint64, value)
	}
	if uu.mutation.RoleIDCleared() {
		_spec.ClearField(user.FieldRoleID, field.TypeUint64)
	}
	if value, ok := uu.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uu.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if uu.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PostTable,
			Columns: []string{user.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PostTable,
			Columns: []string{user.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(u uint8) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(u *uint8) *UserUpdateOne {
	if u != nil {
		uuo.SetStatus(*u)
	}
	return uuo
}

// AddStatus adds u to the "status" field.
func (uuo *UserUpdateOne) AddStatus(u int8) *UserUpdateOne {
	uuo.mutation.AddStatus(u)
	return uuo
}

// ClearStatus clears the value of the "status" field.
func (uuo *UserUpdateOne) ClearStatus() *UserUpdateOne {
	uuo.mutation.ClearStatus()
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UserUpdateOne) SetDescription(s string) *UserUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDescription(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// ClearDescription clears the value of the "description" field.
func (uuo *UserUpdateOne) ClearDescription() *UserUpdateOne {
	uuo.mutation.ClearDescription()
	return uuo
}

// SetHomePath sets the "home_path" field.
func (uuo *UserUpdateOne) SetHomePath(s string) *UserUpdateOne {
	uuo.mutation.SetHomePath(s)
	return uuo
}

// SetNillableHomePath sets the "home_path" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHomePath(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetHomePath(*s)
	}
	return uuo
}

// SetRoleID sets the "role_id" field.
func (uuo *UserUpdateOne) SetRoleID(u uint64) *UserUpdateOne {
	uuo.mutation.ResetRoleID()
	uuo.mutation.SetRoleID(u)
	return uuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRoleID(u *uint64) *UserUpdateOne {
	if u != nil {
		uuo.SetRoleID(*u)
	}
	return uuo
}

// AddRoleID adds u to the "role_id" field.
func (uuo *UserUpdateOne) AddRoleID(u int64) *UserUpdateOne {
	uuo.mutation.AddRoleID(u)
	return uuo
}

// ClearRoleID clears the value of the "role_id" field.
func (uuo *UserUpdateOne) ClearRoleID() *UserUpdateOne {
	uuo.mutation.ClearRoleID()
	return uuo
}

// SetMobile sets the "mobile" field.
func (uuo *UserUpdateOne) SetMobile(s string) *UserUpdateOne {
	uuo.mutation.SetMobile(s)
	return uuo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMobile(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMobile(*s)
	}
	return uuo
}

// ClearMobile clears the value of the "mobile" field.
func (uuo *UserUpdateOne) ClearMobile() *UserUpdateOne {
	uuo.mutation.ClearMobile()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetDepartmentID sets the "department_id" field.
func (uuo *UserUpdateOne) SetDepartmentID(u uint64) *UserUpdateOne {
	uuo.mutation.SetDepartmentID(u)
	return uuo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDepartmentID(u *uint64) *UserUpdateOne {
	if u != nil {
		uuo.SetDepartmentID(*u)
	}
	return uuo
}

// ClearDepartmentID clears the value of the "department_id" field.
func (uuo *UserUpdateOne) ClearDepartmentID() *UserUpdateOne {
	uuo.mutation.ClearDepartmentID()
	return uuo
}

// SetPostID sets the "post_id" field.
func (uuo *UserUpdateOne) SetPostID(u uint64) *UserUpdateOne {
	uuo.mutation.SetPostID(u)
	return uuo
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePostID(u *uint64) *UserUpdateOne {
	if u != nil {
		uuo.SetPostID(*u)
	}
	return uuo
}

// ClearPostID clears the value of the "post_id" field.
func (uuo *UserUpdateOne) ClearPostID() *UserUpdateOne {
	uuo.mutation.ClearPostID()
	return uuo
}

// SetDepartment sets the "department" edge to the Department entity.
func (uuo *UserUpdateOne) SetDepartment(d *Department) *UserUpdateOne {
	return uuo.SetDepartmentID(d.ID)
}

// SetPost sets the "post" edge to the Post entity.
func (uuo *UserUpdateOne) SetPost(p *Post) *UserUpdateOne {
	return uuo.SetPostID(p.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearDepartment clears the "department" edge to the Department entity.
func (uuo *UserUpdateOne) ClearDepartment() *UserUpdateOne {
	uuo.mutation.ClearDepartment()
	return uuo
}

// ClearPost clears the "post" edge to the Post entity.
func (uuo *UserUpdateOne) ClearPost() *UserUpdateOne {
	uuo.mutation.ClearPost()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeUint8, value)
	}
	if uuo.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeUint8)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uuo.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uuo.mutation.HomePath(); ok {
		_spec.SetField(user.FieldHomePath, field.TypeString, value)
	}
	if value, ok := uuo.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint64, value)
	}
	if uuo.mutation.RoleIDCleared() {
		_spec.ClearField(user.FieldRoleID, field.TypeUint64)
	}
	if value, ok := uuo.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uuo.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if uuo.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.DepartmentTable,
			Columns: []string{user.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PostTable,
			Columns: []string{user.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.PostTable,
			Columns: []string{user.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
