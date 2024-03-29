// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/bsuser-rpc/ent/bsuser"
)

// Bsuser is the model entity for the Bsuser schema.
type Bsuser struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 手机
	Mobile string `json:"mobile,omitempty"`
	// 密码
	Pwd string `json:"pwd,omitempty"`
	// 总收益
	TotalAmount string `json:"total_amount,omitempty"`
	// 可提现金额
	ValidAmount string `json:"valid_amount,omitempty"`
	// 分享码
	InviteCode string `json:"invite_code,omitempty"`
	// Inviter ID | 邀请人ID
	InvitedBy uint64 `json:"invited_by,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BsuserQuery when eager-loading is set.
	Edges        BsuserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BsuserEdges holds the relations/edges for other nodes in the graph.
type BsuserEdges struct {
	// Inviter holds the value of the inviter edge.
	Inviter *Bsuser `json:"inviter,omitempty"`
	// Invitee holds the value of the invitee edge.
	Invitee []*Bsuser `json:"invitee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InviterOrErr returns the Inviter value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BsuserEdges) InviterOrErr() (*Bsuser, error) {
	if e.Inviter != nil {
		return e.Inviter, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: bsuser.Label}
	}
	return nil, &NotLoadedError{edge: "inviter"}
}

// InviteeOrErr returns the Invitee value or an error if the edge
// was not loaded in eager-loading.
func (e BsuserEdges) InviteeOrErr() ([]*Bsuser, error) {
	if e.loadedTypes[1] {
		return e.Invitee, nil
	}
	return nil, &NotLoadedError{edge: "invitee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bsuser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bsuser.FieldID, bsuser.FieldStatus, bsuser.FieldInvitedBy:
			values[i] = new(sql.NullInt64)
		case bsuser.FieldName, bsuser.FieldMobile, bsuser.FieldPwd, bsuser.FieldTotalAmount, bsuser.FieldValidAmount, bsuser.FieldInviteCode:
			values[i] = new(sql.NullString)
		case bsuser.FieldCreatedAt, bsuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bsuser fields.
func (b *Bsuser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bsuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = uint64(value.Int64)
		case bsuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case bsuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case bsuser.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = uint8(value.Int64)
			}
		case bsuser.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case bsuser.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				b.Mobile = value.String
			}
		case bsuser.FieldPwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pwd", values[i])
			} else if value.Valid {
				b.Pwd = value.String
			}
		case bsuser.FieldTotalAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value.Valid {
				b.TotalAmount = value.String
			}
		case bsuser.FieldValidAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field valid_amount", values[i])
			} else if value.Valid {
				b.ValidAmount = value.String
			}
		case bsuser.FieldInviteCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invite_code", values[i])
			} else if value.Valid {
				b.InviteCode = value.String
			}
		case bsuser.FieldInvitedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field invited_by", values[i])
			} else if value.Valid {
				b.InvitedBy = uint64(value.Int64)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bsuser.
// This includes values selected through modifiers, order, etc.
func (b *Bsuser) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryInviter queries the "inviter" edge of the Bsuser entity.
func (b *Bsuser) QueryInviter() *BsuserQuery {
	return NewBsuserClient(b.config).QueryInviter(b)
}

// QueryInvitee queries the "invitee" edge of the Bsuser entity.
func (b *Bsuser) QueryInvitee() *BsuserQuery {
	return NewBsuserClient(b.config).QueryInvitee(b)
}

// Update returns a builder for updating this Bsuser.
// Note that you need to call Bsuser.Unwrap() before calling this method if this Bsuser
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bsuser) Update() *BsuserUpdateOne {
	return NewBsuserClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bsuser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bsuser) Unwrap() *Bsuser {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bsuser is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bsuser) String() string {
	var builder strings.Builder
	builder.WriteString("Bsuser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(b.Mobile)
	builder.WriteString(", ")
	builder.WriteString("pwd=")
	builder.WriteString(b.Pwd)
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(b.TotalAmount)
	builder.WriteString(", ")
	builder.WriteString("valid_amount=")
	builder.WriteString(b.ValidAmount)
	builder.WriteString(", ")
	builder.WriteString("invite_code=")
	builder.WriteString(b.InviteCode)
	builder.WriteString(", ")
	builder.WriteString("invited_by=")
	builder.WriteString(fmt.Sprintf("%v", b.InvitedBy))
	builder.WriteByte(')')
	return builder.String()
}

// Bsusers is a parsable slice of Bsuser.
type Bsusers []*Bsuser