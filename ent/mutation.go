// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/bsuser-rpc/ent/bsuser"
	"github.com/kebin6/bsuser-rpc/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBsuser = "Bsuser"
)

// BsuserMutation represents an operation that mutates the Bsuser nodes in the graph.
type BsuserMutation struct {
	config
	op              Op
	typ             string
	id              *uint64
	created_at      *time.Time
	updated_at      *time.Time
	status          *uint8
	addstatus       *int8
	name            *string
	mobile          *string
	pwd             *string
	total_amount    *float64
	addtotal_amount *float64
	valid_amount    *float64
	addvalid_amount *float64
	invite_code     *string
	clearedFields   map[string]struct{}
	inviter         *uint64
	clearedinviter  bool
	invitee         map[uint64]struct{}
	removedinvitee  map[uint64]struct{}
	clearedinvitee  bool
	done            bool
	oldValue        func(context.Context) (*Bsuser, error)
	predicates      []predicate.Bsuser
}

var _ ent.Mutation = (*BsuserMutation)(nil)

// bsuserOption allows management of the mutation configuration using functional options.
type bsuserOption func(*BsuserMutation)

// newBsuserMutation creates new mutation for the Bsuser entity.
func newBsuserMutation(c config, op Op, opts ...bsuserOption) *BsuserMutation {
	m := &BsuserMutation{
		config:        c,
		op:            op,
		typ:           TypeBsuser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBsuserID sets the ID field of the mutation.
func withBsuserID(id uint64) bsuserOption {
	return func(m *BsuserMutation) {
		var (
			err   error
			once  sync.Once
			value *Bsuser
		)
		m.oldValue = func(ctx context.Context) (*Bsuser, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Bsuser.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBsuser sets the old Bsuser of the mutation.
func withBsuser(node *Bsuser) bsuserOption {
	return func(m *BsuserMutation) {
		m.oldValue = func(context.Context) (*Bsuser, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BsuserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BsuserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Bsuser entities.
func (m *BsuserMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BsuserMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BsuserMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Bsuser.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *BsuserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *BsuserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *BsuserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *BsuserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *BsuserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *BsuserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetStatus sets the "status" field.
func (m *BsuserMutation) SetStatus(u uint8) {
	m.status = &u
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *BsuserMutation) Status() (r uint8, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldStatus(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds u to the "status" field.
func (m *BsuserMutation) AddStatus(u int8) {
	if m.addstatus != nil {
		*m.addstatus += u
	} else {
		m.addstatus = &u
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *BsuserMutation) AddedStatus() (r int8, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ClearStatus clears the value of the "status" field.
func (m *BsuserMutation) ClearStatus() {
	m.status = nil
	m.addstatus = nil
	m.clearedFields[bsuser.FieldStatus] = struct{}{}
}

// StatusCleared returns if the "status" field was cleared in this mutation.
func (m *BsuserMutation) StatusCleared() bool {
	_, ok := m.clearedFields[bsuser.FieldStatus]
	return ok
}

// ResetStatus resets all changes to the "status" field.
func (m *BsuserMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
	delete(m.clearedFields, bsuser.FieldStatus)
}

// SetName sets the "name" field.
func (m *BsuserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *BsuserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *BsuserMutation) ResetName() {
	m.name = nil
}

// SetMobile sets the "mobile" field.
func (m *BsuserMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *BsuserMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ResetMobile resets all changes to the "mobile" field.
func (m *BsuserMutation) ResetMobile() {
	m.mobile = nil
}

// SetPwd sets the "pwd" field.
func (m *BsuserMutation) SetPwd(s string) {
	m.pwd = &s
}

// Pwd returns the value of the "pwd" field in the mutation.
func (m *BsuserMutation) Pwd() (r string, exists bool) {
	v := m.pwd
	if v == nil {
		return
	}
	return *v, true
}

// OldPwd returns the old "pwd" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldPwd(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPwd is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPwd requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPwd: %w", err)
	}
	return oldValue.Pwd, nil
}

// ResetPwd resets all changes to the "pwd" field.
func (m *BsuserMutation) ResetPwd() {
	m.pwd = nil
}

// SetTotalAmount sets the "total_amount" field.
func (m *BsuserMutation) SetTotalAmount(f float64) {
	m.total_amount = &f
	m.addtotal_amount = nil
}

// TotalAmount returns the value of the "total_amount" field in the mutation.
func (m *BsuserMutation) TotalAmount() (r float64, exists bool) {
	v := m.total_amount
	if v == nil {
		return
	}
	return *v, true
}

// OldTotalAmount returns the old "total_amount" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldTotalAmount(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTotalAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTotalAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTotalAmount: %w", err)
	}
	return oldValue.TotalAmount, nil
}

// AddTotalAmount adds f to the "total_amount" field.
func (m *BsuserMutation) AddTotalAmount(f float64) {
	if m.addtotal_amount != nil {
		*m.addtotal_amount += f
	} else {
		m.addtotal_amount = &f
	}
}

// AddedTotalAmount returns the value that was added to the "total_amount" field in this mutation.
func (m *BsuserMutation) AddedTotalAmount() (r float64, exists bool) {
	v := m.addtotal_amount
	if v == nil {
		return
	}
	return *v, true
}

// ResetTotalAmount resets all changes to the "total_amount" field.
func (m *BsuserMutation) ResetTotalAmount() {
	m.total_amount = nil
	m.addtotal_amount = nil
}

// SetValidAmount sets the "valid_amount" field.
func (m *BsuserMutation) SetValidAmount(f float64) {
	m.valid_amount = &f
	m.addvalid_amount = nil
}

// ValidAmount returns the value of the "valid_amount" field in the mutation.
func (m *BsuserMutation) ValidAmount() (r float64, exists bool) {
	v := m.valid_amount
	if v == nil {
		return
	}
	return *v, true
}

// OldValidAmount returns the old "valid_amount" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldValidAmount(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldValidAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldValidAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldValidAmount: %w", err)
	}
	return oldValue.ValidAmount, nil
}

// AddValidAmount adds f to the "valid_amount" field.
func (m *BsuserMutation) AddValidAmount(f float64) {
	if m.addvalid_amount != nil {
		*m.addvalid_amount += f
	} else {
		m.addvalid_amount = &f
	}
}

// AddedValidAmount returns the value that was added to the "valid_amount" field in this mutation.
func (m *BsuserMutation) AddedValidAmount() (r float64, exists bool) {
	v := m.addvalid_amount
	if v == nil {
		return
	}
	return *v, true
}

// ResetValidAmount resets all changes to the "valid_amount" field.
func (m *BsuserMutation) ResetValidAmount() {
	m.valid_amount = nil
	m.addvalid_amount = nil
}

// SetInviteCode sets the "invite_code" field.
func (m *BsuserMutation) SetInviteCode(s string) {
	m.invite_code = &s
}

// InviteCode returns the value of the "invite_code" field in the mutation.
func (m *BsuserMutation) InviteCode() (r string, exists bool) {
	v := m.invite_code
	if v == nil {
		return
	}
	return *v, true
}

// OldInviteCode returns the old "invite_code" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldInviteCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInviteCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInviteCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInviteCode: %w", err)
	}
	return oldValue.InviteCode, nil
}

// ResetInviteCode resets all changes to the "invite_code" field.
func (m *BsuserMutation) ResetInviteCode() {
	m.invite_code = nil
}

// SetInvitedBy sets the "invited_by" field.
func (m *BsuserMutation) SetInvitedBy(u uint64) {
	m.inviter = &u
}

// InvitedBy returns the value of the "invited_by" field in the mutation.
func (m *BsuserMutation) InvitedBy() (r uint64, exists bool) {
	v := m.inviter
	if v == nil {
		return
	}
	return *v, true
}

// OldInvitedBy returns the old "invited_by" field's value of the Bsuser entity.
// If the Bsuser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BsuserMutation) OldInvitedBy(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInvitedBy is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInvitedBy requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInvitedBy: %w", err)
	}
	return oldValue.InvitedBy, nil
}

// ClearInvitedBy clears the value of the "invited_by" field.
func (m *BsuserMutation) ClearInvitedBy() {
	m.inviter = nil
	m.clearedFields[bsuser.FieldInvitedBy] = struct{}{}
}

// InvitedByCleared returns if the "invited_by" field was cleared in this mutation.
func (m *BsuserMutation) InvitedByCleared() bool {
	_, ok := m.clearedFields[bsuser.FieldInvitedBy]
	return ok
}

// ResetInvitedBy resets all changes to the "invited_by" field.
func (m *BsuserMutation) ResetInvitedBy() {
	m.inviter = nil
	delete(m.clearedFields, bsuser.FieldInvitedBy)
}

// SetInviterID sets the "inviter" edge to the Bsuser entity by id.
func (m *BsuserMutation) SetInviterID(id uint64) {
	m.inviter = &id
}

// ClearInviter clears the "inviter" edge to the Bsuser entity.
func (m *BsuserMutation) ClearInviter() {
	m.clearedinviter = true
	m.clearedFields[bsuser.FieldInvitedBy] = struct{}{}
}

// InviterCleared reports if the "inviter" edge to the Bsuser entity was cleared.
func (m *BsuserMutation) InviterCleared() bool {
	return m.InvitedByCleared() || m.clearedinviter
}

// InviterID returns the "inviter" edge ID in the mutation.
func (m *BsuserMutation) InviterID() (id uint64, exists bool) {
	if m.inviter != nil {
		return *m.inviter, true
	}
	return
}

// InviterIDs returns the "inviter" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// InviterID instead. It exists only for internal usage by the builders.
func (m *BsuserMutation) InviterIDs() (ids []uint64) {
	if id := m.inviter; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetInviter resets all changes to the "inviter" edge.
func (m *BsuserMutation) ResetInviter() {
	m.inviter = nil
	m.clearedinviter = false
}

// AddInviteeIDs adds the "invitee" edge to the Bsuser entity by ids.
func (m *BsuserMutation) AddInviteeIDs(ids ...uint64) {
	if m.invitee == nil {
		m.invitee = make(map[uint64]struct{})
	}
	for i := range ids {
		m.invitee[ids[i]] = struct{}{}
	}
}

// ClearInvitee clears the "invitee" edge to the Bsuser entity.
func (m *BsuserMutation) ClearInvitee() {
	m.clearedinvitee = true
}

// InviteeCleared reports if the "invitee" edge to the Bsuser entity was cleared.
func (m *BsuserMutation) InviteeCleared() bool {
	return m.clearedinvitee
}

// RemoveInviteeIDs removes the "invitee" edge to the Bsuser entity by IDs.
func (m *BsuserMutation) RemoveInviteeIDs(ids ...uint64) {
	if m.removedinvitee == nil {
		m.removedinvitee = make(map[uint64]struct{})
	}
	for i := range ids {
		delete(m.invitee, ids[i])
		m.removedinvitee[ids[i]] = struct{}{}
	}
}

// RemovedInvitee returns the removed IDs of the "invitee" edge to the Bsuser entity.
func (m *BsuserMutation) RemovedInviteeIDs() (ids []uint64) {
	for id := range m.removedinvitee {
		ids = append(ids, id)
	}
	return
}

// InviteeIDs returns the "invitee" edge IDs in the mutation.
func (m *BsuserMutation) InviteeIDs() (ids []uint64) {
	for id := range m.invitee {
		ids = append(ids, id)
	}
	return
}

// ResetInvitee resets all changes to the "invitee" edge.
func (m *BsuserMutation) ResetInvitee() {
	m.invitee = nil
	m.clearedinvitee = false
	m.removedinvitee = nil
}

// Where appends a list predicates to the BsuserMutation builder.
func (m *BsuserMutation) Where(ps ...predicate.Bsuser) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the BsuserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *BsuserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Bsuser, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *BsuserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *BsuserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Bsuser).
func (m *BsuserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BsuserMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.created_at != nil {
		fields = append(fields, bsuser.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, bsuser.FieldUpdatedAt)
	}
	if m.status != nil {
		fields = append(fields, bsuser.FieldStatus)
	}
	if m.name != nil {
		fields = append(fields, bsuser.FieldName)
	}
	if m.mobile != nil {
		fields = append(fields, bsuser.FieldMobile)
	}
	if m.pwd != nil {
		fields = append(fields, bsuser.FieldPwd)
	}
	if m.total_amount != nil {
		fields = append(fields, bsuser.FieldTotalAmount)
	}
	if m.valid_amount != nil {
		fields = append(fields, bsuser.FieldValidAmount)
	}
	if m.invite_code != nil {
		fields = append(fields, bsuser.FieldInviteCode)
	}
	if m.inviter != nil {
		fields = append(fields, bsuser.FieldInvitedBy)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BsuserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case bsuser.FieldCreatedAt:
		return m.CreatedAt()
	case bsuser.FieldUpdatedAt:
		return m.UpdatedAt()
	case bsuser.FieldStatus:
		return m.Status()
	case bsuser.FieldName:
		return m.Name()
	case bsuser.FieldMobile:
		return m.Mobile()
	case bsuser.FieldPwd:
		return m.Pwd()
	case bsuser.FieldTotalAmount:
		return m.TotalAmount()
	case bsuser.FieldValidAmount:
		return m.ValidAmount()
	case bsuser.FieldInviteCode:
		return m.InviteCode()
	case bsuser.FieldInvitedBy:
		return m.InvitedBy()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BsuserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case bsuser.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case bsuser.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case bsuser.FieldStatus:
		return m.OldStatus(ctx)
	case bsuser.FieldName:
		return m.OldName(ctx)
	case bsuser.FieldMobile:
		return m.OldMobile(ctx)
	case bsuser.FieldPwd:
		return m.OldPwd(ctx)
	case bsuser.FieldTotalAmount:
		return m.OldTotalAmount(ctx)
	case bsuser.FieldValidAmount:
		return m.OldValidAmount(ctx)
	case bsuser.FieldInviteCode:
		return m.OldInviteCode(ctx)
	case bsuser.FieldInvitedBy:
		return m.OldInvitedBy(ctx)
	}
	return nil, fmt.Errorf("unknown Bsuser field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BsuserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case bsuser.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case bsuser.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case bsuser.FieldStatus:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case bsuser.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case bsuser.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	case bsuser.FieldPwd:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPwd(v)
		return nil
	case bsuser.FieldTotalAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTotalAmount(v)
		return nil
	case bsuser.FieldValidAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValidAmount(v)
		return nil
	case bsuser.FieldInviteCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInviteCode(v)
		return nil
	case bsuser.FieldInvitedBy:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInvitedBy(v)
		return nil
	}
	return fmt.Errorf("unknown Bsuser field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BsuserMutation) AddedFields() []string {
	var fields []string
	if m.addstatus != nil {
		fields = append(fields, bsuser.FieldStatus)
	}
	if m.addtotal_amount != nil {
		fields = append(fields, bsuser.FieldTotalAmount)
	}
	if m.addvalid_amount != nil {
		fields = append(fields, bsuser.FieldValidAmount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BsuserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case bsuser.FieldStatus:
		return m.AddedStatus()
	case bsuser.FieldTotalAmount:
		return m.AddedTotalAmount()
	case bsuser.FieldValidAmount:
		return m.AddedValidAmount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BsuserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case bsuser.FieldStatus:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	case bsuser.FieldTotalAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTotalAmount(v)
		return nil
	case bsuser.FieldValidAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddValidAmount(v)
		return nil
	}
	return fmt.Errorf("unknown Bsuser numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BsuserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(bsuser.FieldStatus) {
		fields = append(fields, bsuser.FieldStatus)
	}
	if m.FieldCleared(bsuser.FieldInvitedBy) {
		fields = append(fields, bsuser.FieldInvitedBy)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BsuserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BsuserMutation) ClearField(name string) error {
	switch name {
	case bsuser.FieldStatus:
		m.ClearStatus()
		return nil
	case bsuser.FieldInvitedBy:
		m.ClearInvitedBy()
		return nil
	}
	return fmt.Errorf("unknown Bsuser nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BsuserMutation) ResetField(name string) error {
	switch name {
	case bsuser.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case bsuser.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case bsuser.FieldStatus:
		m.ResetStatus()
		return nil
	case bsuser.FieldName:
		m.ResetName()
		return nil
	case bsuser.FieldMobile:
		m.ResetMobile()
		return nil
	case bsuser.FieldPwd:
		m.ResetPwd()
		return nil
	case bsuser.FieldTotalAmount:
		m.ResetTotalAmount()
		return nil
	case bsuser.FieldValidAmount:
		m.ResetValidAmount()
		return nil
	case bsuser.FieldInviteCode:
		m.ResetInviteCode()
		return nil
	case bsuser.FieldInvitedBy:
		m.ResetInvitedBy()
		return nil
	}
	return fmt.Errorf("unknown Bsuser field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BsuserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.inviter != nil {
		edges = append(edges, bsuser.EdgeInviter)
	}
	if m.invitee != nil {
		edges = append(edges, bsuser.EdgeInvitee)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BsuserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case bsuser.EdgeInviter:
		if id := m.inviter; id != nil {
			return []ent.Value{*id}
		}
	case bsuser.EdgeInvitee:
		ids := make([]ent.Value, 0, len(m.invitee))
		for id := range m.invitee {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BsuserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedinvitee != nil {
		edges = append(edges, bsuser.EdgeInvitee)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BsuserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case bsuser.EdgeInvitee:
		ids := make([]ent.Value, 0, len(m.removedinvitee))
		for id := range m.removedinvitee {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BsuserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedinviter {
		edges = append(edges, bsuser.EdgeInviter)
	}
	if m.clearedinvitee {
		edges = append(edges, bsuser.EdgeInvitee)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BsuserMutation) EdgeCleared(name string) bool {
	switch name {
	case bsuser.EdgeInviter:
		return m.clearedinviter
	case bsuser.EdgeInvitee:
		return m.clearedinvitee
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BsuserMutation) ClearEdge(name string) error {
	switch name {
	case bsuser.EdgeInviter:
		m.ClearInviter()
		return nil
	}
	return fmt.Errorf("unknown Bsuser unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BsuserMutation) ResetEdge(name string) error {
	switch name {
	case bsuser.EdgeInviter:
		m.ResetInviter()
		return nil
	case bsuser.EdgeInvitee:
		m.ResetInvitee()
		return nil
	}
	return fmt.Errorf("unknown Bsuser edge %s", name)
}
