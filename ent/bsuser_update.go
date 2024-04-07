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
	"github.com/kebin6/bsuser-rpc/ent/bsuser"
	"github.com/kebin6/bsuser-rpc/ent/predicate"
)

// BsuserUpdate is the builder for updating Bsuser entities.
type BsuserUpdate struct {
	config
	hooks    []Hook
	mutation *BsuserMutation
}

// Where appends a list predicates to the BsuserUpdate builder.
func (bu *BsuserUpdate) Where(ps ...predicate.Bsuser) *BsuserUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BsuserUpdate) SetUpdatedAt(t time.Time) *BsuserUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetStatus sets the "status" field.
func (bu *BsuserUpdate) SetStatus(u uint8) *BsuserUpdate {
	bu.mutation.ResetStatus()
	bu.mutation.SetStatus(u)
	return bu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillableStatus(u *uint8) *BsuserUpdate {
	if u != nil {
		bu.SetStatus(*u)
	}
	return bu
}

// AddStatus adds u to the "status" field.
func (bu *BsuserUpdate) AddStatus(u int8) *BsuserUpdate {
	bu.mutation.AddStatus(u)
	return bu
}

// ClearStatus clears the value of the "status" field.
func (bu *BsuserUpdate) ClearStatus() *BsuserUpdate {
	bu.mutation.ClearStatus()
	return bu
}

// SetName sets the "name" field.
func (bu *BsuserUpdate) SetName(s string) *BsuserUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillableName(s *string) *BsuserUpdate {
	if s != nil {
		bu.SetName(*s)
	}
	return bu
}

// SetMobile sets the "mobile" field.
func (bu *BsuserUpdate) SetMobile(s string) *BsuserUpdate {
	bu.mutation.SetMobile(s)
	return bu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillableMobile(s *string) *BsuserUpdate {
	if s != nil {
		bu.SetMobile(*s)
	}
	return bu
}

// SetPwd sets the "pwd" field.
func (bu *BsuserUpdate) SetPwd(s string) *BsuserUpdate {
	bu.mutation.SetPwd(s)
	return bu
}

// SetNillablePwd sets the "pwd" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillablePwd(s *string) *BsuserUpdate {
	if s != nil {
		bu.SetPwd(*s)
	}
	return bu
}

// SetTotalAmount sets the "total_amount" field.
func (bu *BsuserUpdate) SetTotalAmount(f float64) *BsuserUpdate {
	bu.mutation.ResetTotalAmount()
	bu.mutation.SetTotalAmount(f)
	return bu
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillableTotalAmount(f *float64) *BsuserUpdate {
	if f != nil {
		bu.SetTotalAmount(*f)
	}
	return bu
}

// AddTotalAmount adds f to the "total_amount" field.
func (bu *BsuserUpdate) AddTotalAmount(f float64) *BsuserUpdate {
	bu.mutation.AddTotalAmount(f)
	return bu
}

// SetValidAmount sets the "valid_amount" field.
func (bu *BsuserUpdate) SetValidAmount(f float64) *BsuserUpdate {
	bu.mutation.ResetValidAmount()
	bu.mutation.SetValidAmount(f)
	return bu
}

// SetNillableValidAmount sets the "valid_amount" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillableValidAmount(f *float64) *BsuserUpdate {
	if f != nil {
		bu.SetValidAmount(*f)
	}
	return bu
}

// AddValidAmount adds f to the "valid_amount" field.
func (bu *BsuserUpdate) AddValidAmount(f float64) *BsuserUpdate {
	bu.mutation.AddValidAmount(f)
	return bu
}

// SetInviteCode sets the "invite_code" field.
func (bu *BsuserUpdate) SetInviteCode(s string) *BsuserUpdate {
	bu.mutation.SetInviteCode(s)
	return bu
}

// SetNillableInviteCode sets the "invite_code" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillableInviteCode(s *string) *BsuserUpdate {
	if s != nil {
		bu.SetInviteCode(*s)
	}
	return bu
}

// SetInvitedBy sets the "invited_by" field.
func (bu *BsuserUpdate) SetInvitedBy(u uint64) *BsuserUpdate {
	bu.mutation.SetInvitedBy(u)
	return bu
}

// SetNillableInvitedBy sets the "invited_by" field if the given value is not nil.
func (bu *BsuserUpdate) SetNillableInvitedBy(u *uint64) *BsuserUpdate {
	if u != nil {
		bu.SetInvitedBy(*u)
	}
	return bu
}

// ClearInvitedBy clears the value of the "invited_by" field.
func (bu *BsuserUpdate) ClearInvitedBy() *BsuserUpdate {
	bu.mutation.ClearInvitedBy()
	return bu
}

// SetInviterID sets the "inviter" edge to the Bsuser entity by ID.
func (bu *BsuserUpdate) SetInviterID(id uint64) *BsuserUpdate {
	bu.mutation.SetInviterID(id)
	return bu
}

// SetNillableInviterID sets the "inviter" edge to the Bsuser entity by ID if the given value is not nil.
func (bu *BsuserUpdate) SetNillableInviterID(id *uint64) *BsuserUpdate {
	if id != nil {
		bu = bu.SetInviterID(*id)
	}
	return bu
}

// SetInviter sets the "inviter" edge to the Bsuser entity.
func (bu *BsuserUpdate) SetInviter(b *Bsuser) *BsuserUpdate {
	return bu.SetInviterID(b.ID)
}

// AddInviteeIDs adds the "invitee" edge to the Bsuser entity by IDs.
func (bu *BsuserUpdate) AddInviteeIDs(ids ...uint64) *BsuserUpdate {
	bu.mutation.AddInviteeIDs(ids...)
	return bu
}

// AddInvitee adds the "invitee" edges to the Bsuser entity.
func (bu *BsuserUpdate) AddInvitee(b ...*Bsuser) *BsuserUpdate {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.AddInviteeIDs(ids...)
}

// Mutation returns the BsuserMutation object of the builder.
func (bu *BsuserUpdate) Mutation() *BsuserMutation {
	return bu.mutation
}

// ClearInviter clears the "inviter" edge to the Bsuser entity.
func (bu *BsuserUpdate) ClearInviter() *BsuserUpdate {
	bu.mutation.ClearInviter()
	return bu
}

// ClearInvitee clears all "invitee" edges to the Bsuser entity.
func (bu *BsuserUpdate) ClearInvitee() *BsuserUpdate {
	bu.mutation.ClearInvitee()
	return bu
}

// RemoveInviteeIDs removes the "invitee" edge to Bsuser entities by IDs.
func (bu *BsuserUpdate) RemoveInviteeIDs(ids ...uint64) *BsuserUpdate {
	bu.mutation.RemoveInviteeIDs(ids...)
	return bu
}

// RemoveInvitee removes "invitee" edges to Bsuser entities.
func (bu *BsuserUpdate) RemoveInvitee(b ...*Bsuser) *BsuserUpdate {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.RemoveInviteeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BsuserUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BsuserUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BsuserUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BsuserUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BsuserUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := bsuser.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

func (bu *BsuserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bsuser.Table, bsuser.Columns, sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(bsuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(bsuser.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := bu.mutation.AddedStatus(); ok {
		_spec.AddField(bsuser.FieldStatus, field.TypeUint8, value)
	}
	if bu.mutation.StatusCleared() {
		_spec.ClearField(bsuser.FieldStatus, field.TypeUint8)
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(bsuser.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.Mobile(); ok {
		_spec.SetField(bsuser.FieldMobile, field.TypeString, value)
	}
	if value, ok := bu.mutation.Pwd(); ok {
		_spec.SetField(bsuser.FieldPwd, field.TypeString, value)
	}
	if value, ok := bu.mutation.TotalAmount(); ok {
		_spec.SetField(bsuser.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedTotalAmount(); ok {
		_spec.AddField(bsuser.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.ValidAmount(); ok {
		_spec.SetField(bsuser.FieldValidAmount, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedValidAmount(); ok {
		_spec.AddField(bsuser.FieldValidAmount, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.InviteCode(); ok {
		_spec.SetField(bsuser.FieldInviteCode, field.TypeString, value)
	}
	if bu.mutation.InviterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bsuser.InviterTable,
			Columns: []string{bsuser.InviterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.InviterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bsuser.InviterTable,
			Columns: []string{bsuser.InviterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.InviteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bsuser.InviteeTable,
			Columns: []string{bsuser.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedInviteeIDs(); len(nodes) > 0 && !bu.mutation.InviteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bsuser.InviteeTable,
			Columns: []string{bsuser.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.InviteeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bsuser.InviteeTable,
			Columns: []string{bsuser.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bsuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BsuserUpdateOne is the builder for updating a single Bsuser entity.
type BsuserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BsuserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BsuserUpdateOne) SetUpdatedAt(t time.Time) *BsuserUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetStatus sets the "status" field.
func (buo *BsuserUpdateOne) SetStatus(u uint8) *BsuserUpdateOne {
	buo.mutation.ResetStatus()
	buo.mutation.SetStatus(u)
	return buo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableStatus(u *uint8) *BsuserUpdateOne {
	if u != nil {
		buo.SetStatus(*u)
	}
	return buo
}

// AddStatus adds u to the "status" field.
func (buo *BsuserUpdateOne) AddStatus(u int8) *BsuserUpdateOne {
	buo.mutation.AddStatus(u)
	return buo
}

// ClearStatus clears the value of the "status" field.
func (buo *BsuserUpdateOne) ClearStatus() *BsuserUpdateOne {
	buo.mutation.ClearStatus()
	return buo
}

// SetName sets the "name" field.
func (buo *BsuserUpdateOne) SetName(s string) *BsuserUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableName(s *string) *BsuserUpdateOne {
	if s != nil {
		buo.SetName(*s)
	}
	return buo
}

// SetMobile sets the "mobile" field.
func (buo *BsuserUpdateOne) SetMobile(s string) *BsuserUpdateOne {
	buo.mutation.SetMobile(s)
	return buo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableMobile(s *string) *BsuserUpdateOne {
	if s != nil {
		buo.SetMobile(*s)
	}
	return buo
}

// SetPwd sets the "pwd" field.
func (buo *BsuserUpdateOne) SetPwd(s string) *BsuserUpdateOne {
	buo.mutation.SetPwd(s)
	return buo
}

// SetNillablePwd sets the "pwd" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillablePwd(s *string) *BsuserUpdateOne {
	if s != nil {
		buo.SetPwd(*s)
	}
	return buo
}

// SetTotalAmount sets the "total_amount" field.
func (buo *BsuserUpdateOne) SetTotalAmount(f float64) *BsuserUpdateOne {
	buo.mutation.ResetTotalAmount()
	buo.mutation.SetTotalAmount(f)
	return buo
}

// SetNillableTotalAmount sets the "total_amount" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableTotalAmount(f *float64) *BsuserUpdateOne {
	if f != nil {
		buo.SetTotalAmount(*f)
	}
	return buo
}

// AddTotalAmount adds f to the "total_amount" field.
func (buo *BsuserUpdateOne) AddTotalAmount(f float64) *BsuserUpdateOne {
	buo.mutation.AddTotalAmount(f)
	return buo
}

// SetValidAmount sets the "valid_amount" field.
func (buo *BsuserUpdateOne) SetValidAmount(f float64) *BsuserUpdateOne {
	buo.mutation.ResetValidAmount()
	buo.mutation.SetValidAmount(f)
	return buo
}

// SetNillableValidAmount sets the "valid_amount" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableValidAmount(f *float64) *BsuserUpdateOne {
	if f != nil {
		buo.SetValidAmount(*f)
	}
	return buo
}

// AddValidAmount adds f to the "valid_amount" field.
func (buo *BsuserUpdateOne) AddValidAmount(f float64) *BsuserUpdateOne {
	buo.mutation.AddValidAmount(f)
	return buo
}

// SetInviteCode sets the "invite_code" field.
func (buo *BsuserUpdateOne) SetInviteCode(s string) *BsuserUpdateOne {
	buo.mutation.SetInviteCode(s)
	return buo
}

// SetNillableInviteCode sets the "invite_code" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableInviteCode(s *string) *BsuserUpdateOne {
	if s != nil {
		buo.SetInviteCode(*s)
	}
	return buo
}

// SetInvitedBy sets the "invited_by" field.
func (buo *BsuserUpdateOne) SetInvitedBy(u uint64) *BsuserUpdateOne {
	buo.mutation.SetInvitedBy(u)
	return buo
}

// SetNillableInvitedBy sets the "invited_by" field if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableInvitedBy(u *uint64) *BsuserUpdateOne {
	if u != nil {
		buo.SetInvitedBy(*u)
	}
	return buo
}

// ClearInvitedBy clears the value of the "invited_by" field.
func (buo *BsuserUpdateOne) ClearInvitedBy() *BsuserUpdateOne {
	buo.mutation.ClearInvitedBy()
	return buo
}

// SetInviterID sets the "inviter" edge to the Bsuser entity by ID.
func (buo *BsuserUpdateOne) SetInviterID(id uint64) *BsuserUpdateOne {
	buo.mutation.SetInviterID(id)
	return buo
}

// SetNillableInviterID sets the "inviter" edge to the Bsuser entity by ID if the given value is not nil.
func (buo *BsuserUpdateOne) SetNillableInviterID(id *uint64) *BsuserUpdateOne {
	if id != nil {
		buo = buo.SetInviterID(*id)
	}
	return buo
}

// SetInviter sets the "inviter" edge to the Bsuser entity.
func (buo *BsuserUpdateOne) SetInviter(b *Bsuser) *BsuserUpdateOne {
	return buo.SetInviterID(b.ID)
}

// AddInviteeIDs adds the "invitee" edge to the Bsuser entity by IDs.
func (buo *BsuserUpdateOne) AddInviteeIDs(ids ...uint64) *BsuserUpdateOne {
	buo.mutation.AddInviteeIDs(ids...)
	return buo
}

// AddInvitee adds the "invitee" edges to the Bsuser entity.
func (buo *BsuserUpdateOne) AddInvitee(b ...*Bsuser) *BsuserUpdateOne {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.AddInviteeIDs(ids...)
}

// Mutation returns the BsuserMutation object of the builder.
func (buo *BsuserUpdateOne) Mutation() *BsuserMutation {
	return buo.mutation
}

// ClearInviter clears the "inviter" edge to the Bsuser entity.
func (buo *BsuserUpdateOne) ClearInviter() *BsuserUpdateOne {
	buo.mutation.ClearInviter()
	return buo
}

// ClearInvitee clears all "invitee" edges to the Bsuser entity.
func (buo *BsuserUpdateOne) ClearInvitee() *BsuserUpdateOne {
	buo.mutation.ClearInvitee()
	return buo
}

// RemoveInviteeIDs removes the "invitee" edge to Bsuser entities by IDs.
func (buo *BsuserUpdateOne) RemoveInviteeIDs(ids ...uint64) *BsuserUpdateOne {
	buo.mutation.RemoveInviteeIDs(ids...)
	return buo
}

// RemoveInvitee removes "invitee" edges to Bsuser entities.
func (buo *BsuserUpdateOne) RemoveInvitee(b ...*Bsuser) *BsuserUpdateOne {
	ids := make([]uint64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.RemoveInviteeIDs(ids...)
}

// Where appends a list predicates to the BsuserUpdate builder.
func (buo *BsuserUpdateOne) Where(ps ...predicate.Bsuser) *BsuserUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BsuserUpdateOne) Select(field string, fields ...string) *BsuserUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bsuser entity.
func (buo *BsuserUpdateOne) Save(ctx context.Context) (*Bsuser, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BsuserUpdateOne) SaveX(ctx context.Context) *Bsuser {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BsuserUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BsuserUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BsuserUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := bsuser.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

func (buo *BsuserUpdateOne) sqlSave(ctx context.Context) (_node *Bsuser, err error) {
	_spec := sqlgraph.NewUpdateSpec(bsuser.Table, bsuser.Columns, sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bsuser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bsuser.FieldID)
		for _, f := range fields {
			if !bsuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bsuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(bsuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(bsuser.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := buo.mutation.AddedStatus(); ok {
		_spec.AddField(bsuser.FieldStatus, field.TypeUint8, value)
	}
	if buo.mutation.StatusCleared() {
		_spec.ClearField(bsuser.FieldStatus, field.TypeUint8)
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(bsuser.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.Mobile(); ok {
		_spec.SetField(bsuser.FieldMobile, field.TypeString, value)
	}
	if value, ok := buo.mutation.Pwd(); ok {
		_spec.SetField(bsuser.FieldPwd, field.TypeString, value)
	}
	if value, ok := buo.mutation.TotalAmount(); ok {
		_spec.SetField(bsuser.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(bsuser.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.ValidAmount(); ok {
		_spec.SetField(bsuser.FieldValidAmount, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedValidAmount(); ok {
		_spec.AddField(bsuser.FieldValidAmount, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.InviteCode(); ok {
		_spec.SetField(bsuser.FieldInviteCode, field.TypeString, value)
	}
	if buo.mutation.InviterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bsuser.InviterTable,
			Columns: []string{bsuser.InviterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.InviterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bsuser.InviterTable,
			Columns: []string{bsuser.InviterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.InviteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bsuser.InviteeTable,
			Columns: []string{bsuser.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedInviteeIDs(); len(nodes) > 0 && !buo.mutation.InviteeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bsuser.InviteeTable,
			Columns: []string{bsuser.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.InviteeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bsuser.InviteeTable,
			Columns: []string{bsuser.InviteeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bsuser{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bsuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
