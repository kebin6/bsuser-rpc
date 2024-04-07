// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BsUserColumns holds the columns for the "bs_user" table.
	BsUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "名称", Default: ""},
		{Name: "mobile", Type: field.TypeString, Comment: "手机", Default: ""},
		{Name: "pwd", Type: field.TypeString, Comment: "密码", Default: ""},
		{Name: "total_amount", Type: field.TypeFloat64, Comment: "总收益", Default: 0},
		{Name: "valid_amount", Type: field.TypeFloat64, Comment: "可提现金额", Default: 0},
		{Name: "invite_code", Type: field.TypeString, Comment: "分享码", Default: ""},
		{Name: "invited_by", Type: field.TypeUint64, Nullable: true, Comment: "Inviter ID | 邀请人ID"},
	}
	// BsUserTable holds the schema information for the "bs_user" table.
	BsUserTable = &schema.Table{
		Name:       "bs_user",
		Columns:    BsUserColumns,
		PrimaryKey: []*schema.Column{BsUserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bs_user_bs_user_invitee",
				Columns:    []*schema.Column{BsUserColumns[10]},
				RefColumns: []*schema.Column{BsUserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "uniq_mobile",
				Unique:  true,
				Columns: []*schema.Column{BsUserColumns[5]},
			},
			{
				Name:    "uniq_invite_code",
				Unique:  true,
				Columns: []*schema.Column{BsUserColumns[9]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BsUserTable,
	}
)

func init() {
	BsUserTable.ForeignKeys[0].RefTable = BsUserTable
	BsUserTable.Annotation = &entsql.Annotation{
		Table: "bs_user",
	}
}
