// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/Sujitnapa29/app/ent/adult"
	"github.com/Sujitnapa29/app/ent/books"
	"github.com/Sujitnapa29/app/ent/customer"
	"github.com/Sujitnapa29/app/ent/kid"
	"github.com/Sujitnapa29/app/ent/predicate"
	"github.com/Sujitnapa29/app/ent/room"
	"github.com/Sujitnapa29/app/ent/roomamount"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// BooksQuery is the builder for querying Books entities.
type BooksQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Books
	// eager-loading edges.
	withAdult      *AdultQuery
	withKid        *KidQuery
	withRoomamount *RoomamountQuery
	withRoom       *RoomQuery
	withCustomer   *CustomerQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (bq *BooksQuery) Where(ps ...predicate.Books) *BooksQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BooksQuery) Limit(limit int) *BooksQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BooksQuery) Offset(offset int) *BooksQuery {
	bq.offset = &offset
	return bq
}

// Order adds an order step to the query.
func (bq *BooksQuery) Order(o ...OrderFunc) *BooksQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryAdult chains the current query on the adult edge.
func (bq *BooksQuery) QueryAdult() *AdultQuery {
	query := &AdultQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(books.Table, books.FieldID, bq.sqlQuery()),
			sqlgraph.To(adult.Table, adult.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, books.AdultTable, books.AdultColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryKid chains the current query on the kid edge.
func (bq *BooksQuery) QueryKid() *KidQuery {
	query := &KidQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(books.Table, books.FieldID, bq.sqlQuery()),
			sqlgraph.To(kid.Table, kid.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, books.KidTable, books.KidColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoomamount chains the current query on the roomamount edge.
func (bq *BooksQuery) QueryRoomamount() *RoomamountQuery {
	query := &RoomamountQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(books.Table, books.FieldID, bq.sqlQuery()),
			sqlgraph.To(roomamount.Table, roomamount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, books.RoomamountTable, books.RoomamountColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoom chains the current query on the room edge.
func (bq *BooksQuery) QueryRoom() *RoomQuery {
	query := &RoomQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(books.Table, books.FieldID, bq.sqlQuery()),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, books.RoomTable, books.RoomColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomer chains the current query on the customer edge.
func (bq *BooksQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(books.Table, books.FieldID, bq.sqlQuery()),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, books.CustomerTable, books.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Books entity in the query. Returns *NotFoundError when no books was found.
func (bq *BooksQuery) First(ctx context.Context) (*Books, error) {
	bs, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(bs) == 0 {
		return nil, &NotFoundError{books.Label}
	}
	return bs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BooksQuery) FirstX(ctx context.Context) *Books {
	b, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return b
}

// FirstID returns the first Books id in the query. Returns *NotFoundError when no id was found.
func (bq *BooksQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{books.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (bq *BooksQuery) FirstXID(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Books entity in the query, returns an error if not exactly one entity was returned.
func (bq *BooksQuery) Only(ctx context.Context) (*Books, error) {
	bs, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(bs) {
	case 1:
		return bs[0], nil
	case 0:
		return nil, &NotFoundError{books.Label}
	default:
		return nil, &NotSingularError{books.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BooksQuery) OnlyX(ctx context.Context) *Books {
	b, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// OnlyID returns the only Books id in the query, returns an error if not exactly one id was returned.
func (bq *BooksQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = &NotSingularError{books.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BooksQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BooksSlice.
func (bq *BooksQuery) All(ctx context.Context) ([]*Books, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BooksQuery) AllX(ctx context.Context) []*Books {
	bs, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return bs
}

// IDs executes the query and returns a list of Books ids.
func (bq *BooksQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(books.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BooksQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BooksQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BooksQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BooksQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BooksQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BooksQuery) Clone() *BooksQuery {
	return &BooksQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]OrderFunc{}, bq.order...),
		unique:     append([]string{}, bq.unique...),
		predicates: append([]predicate.Books{}, bq.predicates...),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

//  WithAdult tells the query-builder to eager-loads the nodes that are connected to
// the "adult" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BooksQuery) WithAdult(opts ...func(*AdultQuery)) *BooksQuery {
	query := &AdultQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withAdult = query
	return bq
}

//  WithKid tells the query-builder to eager-loads the nodes that are connected to
// the "kid" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BooksQuery) WithKid(opts ...func(*KidQuery)) *BooksQuery {
	query := &KidQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withKid = query
	return bq
}

//  WithRoomamount tells the query-builder to eager-loads the nodes that are connected to
// the "roomamount" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BooksQuery) WithRoomamount(opts ...func(*RoomamountQuery)) *BooksQuery {
	query := &RoomamountQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withRoomamount = query
	return bq
}

//  WithRoom tells the query-builder to eager-loads the nodes that are connected to
// the "room" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BooksQuery) WithRoom(opts ...func(*RoomQuery)) *BooksQuery {
	query := &RoomQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withRoom = query
	return bq
}

//  WithCustomer tells the query-builder to eager-loads the nodes that are connected to
// the "customer" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BooksQuery) WithCustomer(opts ...func(*CustomerQuery)) *BooksQuery {
	query := &CustomerQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withCustomer = query
	return bq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Checkin time.Time `json:"Checkin,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Books.Query().
//		GroupBy(books.FieldCheckin).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BooksQuery) GroupBy(field string, fields ...string) *BooksGroupBy {
	group := &BooksGroupBy{config: bq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Checkin time.Time `json:"Checkin,omitempty"`
//	}
//
//	client.Books.Query().
//		Select(books.FieldCheckin).
//		Scan(ctx, &v)
//
func (bq *BooksQuery) Select(field string, fields ...string) *BooksSelect {
	selector := &BooksSelect{config: bq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(), nil
	}
	return selector
}

func (bq *BooksQuery) prepareQuery(ctx context.Context) error {
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BooksQuery) sqlAll(ctx context.Context) ([]*Books, error) {
	var (
		nodes       = []*Books{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [5]bool{
			bq.withAdult != nil,
			bq.withKid != nil,
			bq.withRoomamount != nil,
			bq.withRoom != nil,
			bq.withCustomer != nil,
		}
	)
	if bq.withAdult != nil || bq.withKid != nil || bq.withRoomamount != nil || bq.withRoom != nil || bq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, books.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Books{config: bq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bq.withAdult; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Books)
		for i := range nodes {
			if fk := nodes[i].adult_books; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(adult.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "adult_books" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Adult = n
			}
		}
	}

	if query := bq.withKid; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Books)
		for i := range nodes {
			if fk := nodes[i].kid_books; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(kid.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "kid_books" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Kid = n
			}
		}
	}

	if query := bq.withRoomamount; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Books)
		for i := range nodes {
			if fk := nodes[i].roomamount_books; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(roomamount.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "roomamount_books" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Roomamount = n
			}
		}
	}

	if query := bq.withRoom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Books)
		for i := range nodes {
			if fk := nodes[i].room_books; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(room.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_books" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Room = n
			}
		}
	}

	if query := bq.withCustomer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Books)
		for i := range nodes {
			if fk := nodes[i].customer_books; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(customer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_books" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = n
			}
		}
	}

	return nodes, nil
}

func (bq *BooksQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BooksQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bq *BooksQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   books.Table,
			Columns: books.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: books.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
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

func (bq *BooksQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(books.Table)
	selector := builder.Select(t1.Columns(books.Columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(books.Columns...)...)
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BooksGroupBy is the builder for group-by Books entities.
type BooksGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BooksGroupBy) Aggregate(fns ...AggregateFunc) *BooksGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scan the result into the given value.
func (bgb *BooksGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bgb *BooksGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BooksGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bgb *BooksGroupBy) StringsX(ctx context.Context) []string {
	v, err := bgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bgb *BooksGroupBy) StringX(ctx context.Context) string {
	v, err := bgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BooksGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bgb *BooksGroupBy) IntsX(ctx context.Context) []int {
	v, err := bgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bgb *BooksGroupBy) IntX(ctx context.Context) int {
	v, err := bgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BooksGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bgb *BooksGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bgb *BooksGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BooksGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bgb *BooksGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (bgb *BooksGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bgb *BooksGroupBy) BoolX(ctx context.Context) bool {
	v, err := bgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bgb *BooksGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bgb.sqlQuery().Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BooksGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql
	columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
	columns = append(columns, bgb.fields...)
	for _, fn := range bgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(bgb.fields...)
}

// BooksSelect is the builder for select fields of Books entities.
type BooksSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (bs *BooksSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := bs.path(ctx)
	if err != nil {
		return err
	}
	bs.sql = query
	return bs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bs *BooksSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BooksSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bs *BooksSelect) StringsX(ctx context.Context) []string {
	v, err := bs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bs *BooksSelect) StringX(ctx context.Context) string {
	v, err := bs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BooksSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bs *BooksSelect) IntsX(ctx context.Context) []int {
	v, err := bs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bs *BooksSelect) IntX(ctx context.Context) int {
	v, err := bs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BooksSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bs *BooksSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bs *BooksSelect) Float64X(ctx context.Context) float64 {
	v, err := bs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BooksSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bs *BooksSelect) BoolsX(ctx context.Context) []bool {
	v, err := bs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (bs *BooksSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{books.Label}
	default:
		err = fmt.Errorf("ent: BooksSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bs *BooksSelect) BoolX(ctx context.Context) bool {
	v, err := bs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bs *BooksSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bs.sqlQuery().Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bs *BooksSelect) sqlQuery() sql.Querier {
	selector := bs.sql
	selector.Select(selector.Columns(bs.fields...)...)
	return selector
}
