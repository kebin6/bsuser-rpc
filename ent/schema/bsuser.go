package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Bsuser holds the schema definition for the Bsuser entity.
type Bsuser struct {
	ent.Schema
}

// Fields of the Bsuser.
func (Bsuser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("mobile").Comment("手机"),
		field.String("pwd").Comment("密码"),
		field.String("total_amount").Comment("总收益"),
		field.String("valid_amount").Comment("可提现金额"),
		field.String("invite_code").Comment("分享码"),
		field.Uint64("invited_by").Optional().Comment("Inviter ID | 邀请人ID"),
	}
}

// Edges of the Bsuser.
func (Bsuser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invitee", Bsuser.Type).From("inviter").Unique().Field("invited_by"),
	}
}

// Indexes of the Bsuser.
func (Bsuser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mobile").Unique().StorageKey("uniq_mobile"),
		index.Fields("invite_code").Unique().StorageKey("uniq_invite_code"),
	}
}

// Annotations of the Bsuser.
func (Bsuser) Annotations() []schema.Annotation {
	withComment := true
	return []schema.Annotation{
		entsql.Annotation{Table: "bs_user", WithComments: &withComment},
	}
}

// Mixin of the Bsuser.
func (Bsuser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}