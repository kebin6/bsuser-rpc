// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kebin6/bsuser-rpc/ent/bsuser"
	"github.com/kebin6/bsuser-rpc/ent/predicate"
)

// BsuserQuery is the builder for querying Bsuser entities.
type BsuserQuery struct {
	config
	ctx         *QueryContext
	order       []bsuser.OrderOption
	inters      []Interceptor
	predicates  []predicate.Bsuser
	withInviter *BsuserQuery
	withInvitee *BsuserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BsuserQuery builder.
func (bq *BsuserQuery) Where(ps ...predicate.Bsuser) *BsuserQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BsuserQuery) Limit(limit int) *BsuserQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BsuserQuery) Offset(offset int) *BsuserQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BsuserQuery) Unique(unique bool) *BsuserQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BsuserQuery) Order(o ...bsuser.OrderOption) *BsuserQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryInviter chains the current query on the "inviter" edge.
func (bq *BsuserQuery) QueryInviter() *BsuserQuery {
	query := (&BsuserClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bsuser.Table, bsuser.FieldID, selector),
			sqlgraph.To(bsuser.Table, bsuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bsuser.InviterTable, bsuser.InviterColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInvitee chains the current query on the "invitee" edge.
func (bq *BsuserQuery) QueryInvitee() *BsuserQuery {
	query := (&BsuserClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bsuser.Table, bsuser.FieldID, selector),
			sqlgraph.To(bsuser.Table, bsuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bsuser.InviteeTable, bsuser.InviteeColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bsuser entity from the query.
// Returns a *NotFoundError when no Bsuser was found.
func (bq *BsuserQuery) First(ctx context.Context) (*Bsuser, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bsuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BsuserQuery) FirstX(ctx context.Context) *Bsuser {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bsuser ID from the query.
// Returns a *NotFoundError when no Bsuser ID was found.
func (bq *BsuserQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bsuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BsuserQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bsuser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bsuser entity is found.
// Returns a *NotFoundError when no Bsuser entities are found.
func (bq *BsuserQuery) Only(ctx context.Context) (*Bsuser, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bsuser.Label}
	default:
		return nil, &NotSingularError{bsuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BsuserQuery) OnlyX(ctx context.Context) *Bsuser {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bsuser ID in the query.
// Returns a *NotSingularError when more than one Bsuser ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BsuserQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bsuser.Label}
	default:
		err = &NotSingularError{bsuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BsuserQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bsusers.
func (bq *BsuserQuery) All(ctx context.Context) ([]*Bsuser, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bsuser, *BsuserQuery]()
	return withInterceptors[[]*Bsuser](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BsuserQuery) AllX(ctx context.Context) []*Bsuser {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bsuser IDs.
func (bq *BsuserQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(bsuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BsuserQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BsuserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BsuserQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BsuserQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BsuserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BsuserQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BsuserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BsuserQuery) Clone() *BsuserQuery {
	if bq == nil {
		return nil
	}
	return &BsuserQuery{
		config:      bq.config,
		ctx:         bq.ctx.Clone(),
		order:       append([]bsuser.OrderOption{}, bq.order...),
		inters:      append([]Interceptor{}, bq.inters...),
		predicates:  append([]predicate.Bsuser{}, bq.predicates...),
		withInviter: bq.withInviter.Clone(),
		withInvitee: bq.withInvitee.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithInviter tells the query-builder to eager-load the nodes that are connected to
// the "inviter" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BsuserQuery) WithInviter(opts ...func(*BsuserQuery)) *BsuserQuery {
	query := (&BsuserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withInviter = query
	return bq
}

// WithInvitee tells the query-builder to eager-load the nodes that are connected to
// the "invitee" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BsuserQuery) WithInvitee(opts ...func(*BsuserQuery)) *BsuserQuery {
	query := (&BsuserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withInvitee = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bsuser.Query().
//		GroupBy(bsuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BsuserQuery) GroupBy(field string, fields ...string) *BsuserGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BsuserGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = bsuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Bsuser.Query().
//		Select(bsuser.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BsuserQuery) Select(fields ...string) *BsuserSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BsuserSelect{BsuserQuery: bq}
	sbuild.label = bsuser.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BsuserSelect configured with the given aggregations.
func (bq *BsuserQuery) Aggregate(fns ...AggregateFunc) *BsuserSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BsuserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !bsuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BsuserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bsuser, error) {
	var (
		nodes       = []*Bsuser{}
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withInviter != nil,
			bq.withInvitee != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bsuser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bsuser{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withInviter; query != nil {
		if err := bq.loadInviter(ctx, query, nodes, nil,
			func(n *Bsuser, e *Bsuser) { n.Edges.Inviter = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withInvitee; query != nil {
		if err := bq.loadInvitee(ctx, query, nodes,
			func(n *Bsuser) { n.Edges.Invitee = []*Bsuser{} },
			func(n *Bsuser, e *Bsuser) { n.Edges.Invitee = append(n.Edges.Invitee, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BsuserQuery) loadInviter(ctx context.Context, query *BsuserQuery, nodes []*Bsuser, init func(*Bsuser), assign func(*Bsuser, *Bsuser)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Bsuser)
	for i := range nodes {
		fk := nodes[i].InvitedBy
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(bsuser.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invited_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BsuserQuery) loadInvitee(ctx context.Context, query *BsuserQuery, nodes []*Bsuser, init func(*Bsuser), assign func(*Bsuser, *Bsuser)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Bsuser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(bsuser.FieldInvitedBy)
	}
	query.Where(predicate.Bsuser(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bsuser.InviteeColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.InvitedBy
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "invited_by" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BsuserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BsuserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bsuser.Table, bsuser.Columns, sqlgraph.NewFieldSpec(bsuser.FieldID, field.TypeUint64))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bsuser.FieldID)
		for i := range fields {
			if fields[i] != bsuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if bq.withInviter != nil {
			_spec.Node.AddColumnOnce(bsuser.FieldInvitedBy)
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BsuserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bsuser.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = bsuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BsuserGroupBy is the group-by builder for Bsuser entities.
type BsuserGroupBy struct {
	selector
	build *BsuserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BsuserGroupBy) Aggregate(fns ...AggregateFunc) *BsuserGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BsuserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BsuserQuery, *BsuserGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BsuserGroupBy) sqlScan(ctx context.Context, root *BsuserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BsuserSelect is the builder for selecting fields of Bsuser entities.
type BsuserSelect struct {
	*BsuserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BsuserSelect) Aggregate(fns ...AggregateFunc) *BsuserSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BsuserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BsuserQuery, *BsuserSelect](ctx, bs.BsuserQuery, bs, bs.inters, v)
}

func (bs *BsuserSelect) sqlScan(ctx context.Context, root *BsuserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}