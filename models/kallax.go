// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package models

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

type modelSaveFunc func(*kallax.Store) error

// NewAsset returns a new instance of Asset.
func NewAsset() (record *Asset) {
	return new(Asset)
}

// GetID returns the primary key of the model.
func (r *Asset) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Asset) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil
	case "symbol":
		return &r.Symbol, nil
	case "name":
		return &r.Name, nil
	case "is_fiat":
		return &r.IsFiat, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Asset: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Asset) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "symbol":
		return r.Symbol, nil
	case "name":
		return r.Name, nil
	case "is_fiat":
		return r.IsFiat, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Asset: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Asset) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Asset has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Asset) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Asset has no relationships")
}

// AssetStore is the entity to access the records of the type Asset
// in the database.
type AssetStore struct {
	*kallax.Store
}

// NewAssetStore creates a new instance of AssetStore
// using a SQL database.
func NewAssetStore(db *sql.DB) *AssetStore {
	return &AssetStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *AssetStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *AssetStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *AssetStore) Debug() *AssetStore {
	return &AssetStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *AssetStore) DebugWith(logger kallax.LoggerFunc) *AssetStore {
	return &AssetStore{s.Store.DebugWith(logger)}
}

// Insert inserts a Asset in the database. A non-persisted object is
// required for this operation.
func (s *AssetStore) Insert(record *Asset) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	return s.Store.Insert(Schema.Asset.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *AssetStore) Update(record *Asset, cols ...kallax.SchemaField) (updated int64, err error) {
	record.SetSaving(true)
	defer record.SetSaving(false)

	return s.Store.Update(Schema.Asset.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *AssetStore) Save(record *Asset) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *AssetStore) Delete(record *Asset) error {
	return s.Store.Delete(Schema.Asset.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *AssetStore) Find(q *AssetQuery) (*AssetResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewAssetResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *AssetStore) MustFind(q *AssetQuery) *AssetResultSet {
	return NewAssetResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *AssetStore) Count(q *AssetQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *AssetStore) MustCount(q *AssetQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *AssetStore) FindOne(q *AssetQuery) (*Asset, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *AssetStore) FindAll(q *AssetQuery) ([]*Asset, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *AssetStore) MustFindOne(q *AssetQuery) *Asset {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Asset with the data in the database and
// makes it writable.
func (s *AssetStore) Reload(record *Asset) error {
	return s.Store.Reload(Schema.Asset.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *AssetStore) Transaction(callback func(*AssetStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&AssetStore{store})
	})
}

// AssetQuery is the object used to create queries for the Asset
// entity.
type AssetQuery struct {
	*kallax.BaseQuery
}

// NewAssetQuery returns a new instance of AssetQuery.
func NewAssetQuery() *AssetQuery {
	return &AssetQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Asset.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *AssetQuery) Select(columns ...kallax.SchemaField) *AssetQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *AssetQuery) SelectNot(columns ...kallax.SchemaField) *AssetQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *AssetQuery) Copy() *AssetQuery {
	return &AssetQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *AssetQuery) Order(cols ...kallax.ColumnOrder) *AssetQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *AssetQuery) BatchSize(size uint64) *AssetQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *AssetQuery) Limit(n uint64) *AssetQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *AssetQuery) Offset(n uint64) *AssetQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *AssetQuery) Where(cond kallax.Condition) *AssetQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *AssetQuery) FindByID(v ...kallax.ULID) *AssetQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Asset.ID, values...))
}

// FindBySymbol adds a new filter to the query that will require that
// the Symbol property is equal to the passed value.
func (q *AssetQuery) FindBySymbol(v string) *AssetQuery {
	return q.Where(kallax.Eq(Schema.Asset.Symbol, v))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *AssetQuery) FindByName(v string) *AssetQuery {
	return q.Where(kallax.Eq(Schema.Asset.Name, v))
}

// FindByIsFiat adds a new filter to the query that will require that
// the IsFiat property is equal to the passed value.
func (q *AssetQuery) FindByIsFiat(v bool) *AssetQuery {
	return q.Where(kallax.Eq(Schema.Asset.IsFiat, v))
}

// AssetResultSet is the set of results returned by a query to the
// database.
type AssetResultSet struct {
	ResultSet kallax.ResultSet
	last      *Asset
	lastErr   error
}

// NewAssetResultSet creates a new result set for rows of the type
// Asset.
func NewAssetResultSet(rs kallax.ResultSet) *AssetResultSet {
	return &AssetResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *AssetResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Asset.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Asset)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Asset")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *AssetResultSet) Get() (*Asset, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *AssetResultSet) ForEach(fn func(*Asset) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *AssetResultSet) All() ([]*Asset, error) {
	var result []*Asset
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *AssetResultSet) One() (*Asset, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *AssetResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *AssetResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewExchange returns a new instance of Exchange.
func NewExchange() (record *Exchange) {
	return new(Exchange)
}

// GetID returns the primary key of the model.
func (r *Exchange) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Exchange) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil
	case "code":
		return &r.Code, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Exchange: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Exchange) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "code":
		return r.Code, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Exchange: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Exchange) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Exchange has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Exchange) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Exchange has no relationships")
}

// ExchangeStore is the entity to access the records of the type Exchange
// in the database.
type ExchangeStore struct {
	*kallax.Store
}

// NewExchangeStore creates a new instance of ExchangeStore
// using a SQL database.
func NewExchangeStore(db *sql.DB) *ExchangeStore {
	return &ExchangeStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *ExchangeStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *ExchangeStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *ExchangeStore) Debug() *ExchangeStore {
	return &ExchangeStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *ExchangeStore) DebugWith(logger kallax.LoggerFunc) *ExchangeStore {
	return &ExchangeStore{s.Store.DebugWith(logger)}
}

// Insert inserts a Exchange in the database. A non-persisted object is
// required for this operation.
func (s *ExchangeStore) Insert(record *Exchange) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	return s.Store.Insert(Schema.Exchange.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *ExchangeStore) Update(record *Exchange, cols ...kallax.SchemaField) (updated int64, err error) {
	record.SetSaving(true)
	defer record.SetSaving(false)

	return s.Store.Update(Schema.Exchange.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *ExchangeStore) Save(record *Exchange) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *ExchangeStore) Delete(record *Exchange) error {
	return s.Store.Delete(Schema.Exchange.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *ExchangeStore) Find(q *ExchangeQuery) (*ExchangeResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewExchangeResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *ExchangeStore) MustFind(q *ExchangeQuery) *ExchangeResultSet {
	return NewExchangeResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *ExchangeStore) Count(q *ExchangeQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *ExchangeStore) MustCount(q *ExchangeQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *ExchangeStore) FindOne(q *ExchangeQuery) (*Exchange, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *ExchangeStore) FindAll(q *ExchangeQuery) ([]*Exchange, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *ExchangeStore) MustFindOne(q *ExchangeQuery) *Exchange {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Exchange with the data in the database and
// makes it writable.
func (s *ExchangeStore) Reload(record *Exchange) error {
	return s.Store.Reload(Schema.Exchange.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *ExchangeStore) Transaction(callback func(*ExchangeStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&ExchangeStore{store})
	})
}

// ExchangeQuery is the object used to create queries for the Exchange
// entity.
type ExchangeQuery struct {
	*kallax.BaseQuery
}

// NewExchangeQuery returns a new instance of ExchangeQuery.
func NewExchangeQuery() *ExchangeQuery {
	return &ExchangeQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Exchange.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *ExchangeQuery) Select(columns ...kallax.SchemaField) *ExchangeQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *ExchangeQuery) SelectNot(columns ...kallax.SchemaField) *ExchangeQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *ExchangeQuery) Copy() *ExchangeQuery {
	return &ExchangeQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *ExchangeQuery) Order(cols ...kallax.ColumnOrder) *ExchangeQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *ExchangeQuery) BatchSize(size uint64) *ExchangeQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *ExchangeQuery) Limit(n uint64) *ExchangeQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *ExchangeQuery) Offset(n uint64) *ExchangeQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *ExchangeQuery) Where(cond kallax.Condition) *ExchangeQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *ExchangeQuery) FindByID(v ...kallax.ULID) *ExchangeQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Exchange.ID, values...))
}

// FindByCode adds a new filter to the query that will require that
// the Code property is equal to the passed value.
func (q *ExchangeQuery) FindByCode(v string) *ExchangeQuery {
	return q.Where(kallax.Eq(Schema.Exchange.Code, v))
}

// ExchangeResultSet is the set of results returned by a query to the
// database.
type ExchangeResultSet struct {
	ResultSet kallax.ResultSet
	last      *Exchange
	lastErr   error
}

// NewExchangeResultSet creates a new result set for rows of the type
// Exchange.
func NewExchangeResultSet(rs kallax.ResultSet) *ExchangeResultSet {
	return &ExchangeResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *ExchangeResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Exchange.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Exchange)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Exchange")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *ExchangeResultSet) Get() (*Exchange, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *ExchangeResultSet) ForEach(fn func(*Exchange) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *ExchangeResultSet) All() ([]*Exchange, error) {
	var result []*Exchange
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *ExchangeResultSet) One() (*Exchange, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *ExchangeResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *ExchangeResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewPair returns a new instance of Pair.
func NewPair() (record *Pair) {
	return new(Pair)
}

// GetID returns the primary key of the model.
func (r *Pair) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Pair) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil
	case "symbol":
		return &r.Symbol, nil
	case "base_id":
		return types.Nullable(kallax.VirtualColumn("base_id", r, new(kallax.ULID))), nil
	case "quote_id":
		return types.Nullable(kallax.VirtualColumn("quote_id", r, new(kallax.ULID))), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Pair: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Pair) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "symbol":
		return r.Symbol, nil
	case "base_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "quote_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Pair: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Pair) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "BaseId":
		return new(Asset), nil
	case "QuoteId":
		return new(Asset), nil

	}
	return nil, fmt.Errorf("kallax: model Pair has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Pair) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "BaseId":
		val, ok := rel.(*Asset)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship BaseId", rel)
		}
		if !val.GetID().IsEmpty() {
			r.BaseId = val
		}

		return nil
	case "QuoteId":
		val, ok := rel.(*Asset)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship QuoteId", rel)
		}
		if !val.GetID().IsEmpty() {
			r.QuoteId = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model Pair has no relationship %s", field)
}

// PairStore is the entity to access the records of the type Pair
// in the database.
type PairStore struct {
	*kallax.Store
}

// NewPairStore creates a new instance of PairStore
// using a SQL database.
func NewPairStore(db *sql.DB) *PairStore {
	return &PairStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *PairStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *PairStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *PairStore) Debug() *PairStore {
	return &PairStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *PairStore) DebugWith(logger kallax.LoggerFunc) *PairStore {
	return &PairStore{s.Store.DebugWith(logger)}
}

func (s *PairStore) inverseRecords(record *Pair) []modelSaveFunc {
	var result []modelSaveFunc

	if record.BaseId != nil && !record.BaseId.IsSaving() {
		record.AddVirtualColumn("base_id", record.BaseId.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&AssetStore{store}).Save(record.BaseId)
			return err
		})
	}

	if record.QuoteId != nil && !record.QuoteId.IsSaving() {
		record.AddVirtualColumn("quote_id", record.QuoteId.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&AssetStore{store}).Save(record.QuoteId)
			return err
		})
	}

	return result
}

// Insert inserts a Pair in the database. A non-persisted object is
// required for this operation.
func (s *PairStore) Insert(record *Pair) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Pair.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Pair.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *PairStore) Update(record *Pair, cols ...kallax.SchemaField) (updated int64, err error) {
	record.SetSaving(true)
	defer record.SetSaving(false)

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Pair.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Pair.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *PairStore) Save(record *Pair) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *PairStore) Delete(record *Pair) error {
	return s.Store.Delete(Schema.Pair.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *PairStore) Find(q *PairQuery) (*PairResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewPairResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *PairStore) MustFind(q *PairQuery) *PairResultSet {
	return NewPairResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *PairStore) Count(q *PairQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *PairStore) MustCount(q *PairQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *PairStore) FindOne(q *PairQuery) (*Pair, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *PairStore) FindAll(q *PairQuery) ([]*Pair, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *PairStore) MustFindOne(q *PairQuery) *Pair {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Pair with the data in the database and
// makes it writable.
func (s *PairStore) Reload(record *Pair) error {
	return s.Store.Reload(Schema.Pair.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *PairStore) Transaction(callback func(*PairStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&PairStore{store})
	})
}

// PairQuery is the object used to create queries for the Pair
// entity.
type PairQuery struct {
	*kallax.BaseQuery
}

// NewPairQuery returns a new instance of PairQuery.
func NewPairQuery() *PairQuery {
	return &PairQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Pair.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *PairQuery) Select(columns ...kallax.SchemaField) *PairQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *PairQuery) SelectNot(columns ...kallax.SchemaField) *PairQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *PairQuery) Copy() *PairQuery {
	return &PairQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *PairQuery) Order(cols ...kallax.ColumnOrder) *PairQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *PairQuery) BatchSize(size uint64) *PairQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *PairQuery) Limit(n uint64) *PairQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *PairQuery) Offset(n uint64) *PairQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *PairQuery) Where(cond kallax.Condition) *PairQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *PairQuery) WithBaseId() *PairQuery {
	q.AddRelation(Schema.Asset.BaseSchema, "BaseId", kallax.OneToOne, nil)
	return q
}

func (q *PairQuery) WithQuoteId() *PairQuery {
	q.AddRelation(Schema.Asset.BaseSchema, "QuoteId", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PairQuery) FindByID(v ...kallax.ULID) *PairQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Pair.ID, values...))
}

// FindBySymbol adds a new filter to the query that will require that
// the Symbol property is equal to the passed value.
func (q *PairQuery) FindBySymbol(v string) *PairQuery {
	return q.Where(kallax.Eq(Schema.Pair.Symbol, v))
}

// FindByBaseId adds a new filter to the query that will require that
// the foreign key of BaseId is equal to the passed value.
func (q *PairQuery) FindByBaseId(v kallax.ULID) *PairQuery {
	return q.Where(kallax.Eq(Schema.Pair.BaseIdFK, v))
}

// FindByQuoteId adds a new filter to the query that will require that
// the foreign key of QuoteId is equal to the passed value.
func (q *PairQuery) FindByQuoteId(v kallax.ULID) *PairQuery {
	return q.Where(kallax.Eq(Schema.Pair.QuoteIdFK, v))
}

// PairResultSet is the set of results returned by a query to the
// database.
type PairResultSet struct {
	ResultSet kallax.ResultSet
	last      *Pair
	lastErr   error
}

// NewPairResultSet creates a new result set for rows of the type
// Pair.
func NewPairResultSet(rs kallax.ResultSet) *PairResultSet {
	return &PairResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *PairResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Pair.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Pair)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Pair")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *PairResultSet) Get() (*Pair, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *PairResultSet) ForEach(fn func(*Pair) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *PairResultSet) All() ([]*Pair, error) {
	var result []*Pair
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *PairResultSet) One() (*Pair, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *PairResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *PairResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewPrice returns a new instance of Price.
func NewPrice() (record *Price) {
	return new(Price)
}

// GetID returns the primary key of the model.
func (r *Price) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Price) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "pair_id":
		return types.Nullable(kallax.VirtualColumn("pair_id", r, new(kallax.ULID))), nil
	case "exchange_id":
		return types.Nullable(kallax.VirtualColumn("exchange_id", r, new(kallax.ULID))), nil
	case "price":
		return &r.Price, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Price: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Price) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "pair_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "exchange_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "price":
		return r.Price, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Price: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Price) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "PairId":
		return new(Pair), nil
	case "ExchangeId":
		return new(Exchange), nil

	}
	return nil, fmt.Errorf("kallax: model Price has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Price) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "PairId":
		val, ok := rel.(*Pair)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship PairId", rel)
		}
		r.PairId = *val
		return nil
	case "ExchangeId":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship ExchangeId", rel)
		}
		r.ExchangeId = *val
		return nil

	}
	return fmt.Errorf("kallax: model Price has no relationship %s", field)
}

// PriceStore is the entity to access the records of the type Price
// in the database.
type PriceStore struct {
	*kallax.Store
}

// NewPriceStore creates a new instance of PriceStore
// using a SQL database.
func NewPriceStore(db *sql.DB) *PriceStore {
	return &PriceStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *PriceStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *PriceStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *PriceStore) Debug() *PriceStore {
	return &PriceStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *PriceStore) DebugWith(logger kallax.LoggerFunc) *PriceStore {
	return &PriceStore{s.Store.DebugWith(logger)}
}

func (s *PriceStore) inverseRecords(record *Price) []modelSaveFunc {
	var result []modelSaveFunc

	if !record.PairId.GetID().IsEmpty() && !record.PairId.IsSaving() {
		record.AddVirtualColumn("pair_id", record.PairId.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&PairStore{store}).Save(record.PairId)
			return err
		})
	}

	if !record.ExchangeId.GetID().IsEmpty() && !record.ExchangeId.IsSaving() {
		record.AddVirtualColumn("exchange_id", record.ExchangeId.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.ExchangeId)
			return err
		})
	}

	return result
}

// Insert inserts a Price in the database. A non-persisted object is
// required for this operation.
func (s *PriceStore) Insert(record *Price) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Price.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Price.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *PriceStore) Update(record *Price, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Price.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Price.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *PriceStore) Save(record *Price) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *PriceStore) Delete(record *Price) error {
	return s.Store.Delete(Schema.Price.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *PriceStore) Find(q *PriceQuery) (*PriceResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewPriceResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *PriceStore) MustFind(q *PriceQuery) *PriceResultSet {
	return NewPriceResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *PriceStore) Count(q *PriceQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *PriceStore) MustCount(q *PriceQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *PriceStore) FindOne(q *PriceQuery) (*Price, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *PriceStore) FindAll(q *PriceQuery) ([]*Price, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *PriceStore) MustFindOne(q *PriceQuery) *Price {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Price with the data in the database and
// makes it writable.
func (s *PriceStore) Reload(record *Price) error {
	return s.Store.Reload(Schema.Price.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *PriceStore) Transaction(callback func(*PriceStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&PriceStore{store})
	})
}

// PriceQuery is the object used to create queries for the Price
// entity.
type PriceQuery struct {
	*kallax.BaseQuery
}

// NewPriceQuery returns a new instance of PriceQuery.
func NewPriceQuery() *PriceQuery {
	return &PriceQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Price.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *PriceQuery) Select(columns ...kallax.SchemaField) *PriceQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *PriceQuery) SelectNot(columns ...kallax.SchemaField) *PriceQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *PriceQuery) Copy() *PriceQuery {
	return &PriceQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *PriceQuery) Order(cols ...kallax.ColumnOrder) *PriceQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *PriceQuery) BatchSize(size uint64) *PriceQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *PriceQuery) Limit(n uint64) *PriceQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *PriceQuery) Offset(n uint64) *PriceQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *PriceQuery) Where(cond kallax.Condition) *PriceQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *PriceQuery) WithPairId() *PriceQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "PairId", kallax.OneToOne, nil)
	return q
}

func (q *PriceQuery) WithExchangeId() *PriceQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "ExchangeId", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PriceQuery) FindByID(v ...kallax.ULID) *PriceQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Price.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *PriceQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *PriceQuery {
	return q.Where(cond(Schema.Price.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *PriceQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *PriceQuery {
	return q.Where(cond(Schema.Price.UpdatedAt, v))
}

// FindByPairId adds a new filter to the query that will require that
// the foreign key of PairId is equal to the passed value.
func (q *PriceQuery) FindByPairId(v kallax.ULID) *PriceQuery {
	return q.Where(kallax.Eq(Schema.Price.PairIdFK, v))
}

// FindByExchangeId adds a new filter to the query that will require that
// the foreign key of ExchangeId is equal to the passed value.
func (q *PriceQuery) FindByExchangeId(v kallax.ULID) *PriceQuery {
	return q.Where(kallax.Eq(Schema.Price.ExchangeIdFK, v))
}

// FindByPrice adds a new filter to the query that will require that
// the Price property is equal to the passed value.
func (q *PriceQuery) FindByPrice(cond kallax.ScalarCond, v float32) *PriceQuery {
	return q.Where(cond(Schema.Price.Price, v))
}

// PriceResultSet is the set of results returned by a query to the
// database.
type PriceResultSet struct {
	ResultSet kallax.ResultSet
	last      *Price
	lastErr   error
}

// NewPriceResultSet creates a new result set for rows of the type
// Price.
func NewPriceResultSet(rs kallax.ResultSet) *PriceResultSet {
	return &PriceResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *PriceResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Price.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Price)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Price")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *PriceResultSet) Get() (*Price, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *PriceResultSet) ForEach(fn func(*Price) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *PriceResultSet) All() ([]*Price, error) {
	var result []*Price
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *PriceResultSet) One() (*Price, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *PriceResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *PriceResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Asset    *schemaAsset
	Exchange *schemaExchange
	Pair     *schemaPair
	Price    *schemaPrice
}

type schemaAsset struct {
	*kallax.BaseSchema
	ID     kallax.SchemaField
	Symbol kallax.SchemaField
	Name   kallax.SchemaField
	IsFiat kallax.SchemaField
}

type schemaExchange struct {
	*kallax.BaseSchema
	ID   kallax.SchemaField
	Code kallax.SchemaField
}

type schemaPair struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	Symbol    kallax.SchemaField
	BaseIdFK  kallax.SchemaField
	QuoteIdFK kallax.SchemaField
}

type schemaPrice struct {
	*kallax.BaseSchema
	ID           kallax.SchemaField
	CreatedAt    kallax.SchemaField
	UpdatedAt    kallax.SchemaField
	PairIdFK     kallax.SchemaField
	ExchangeIdFK kallax.SchemaField
	Price        kallax.SchemaField
}

var Schema = &schema{
	Asset: &schemaAsset{
		BaseSchema: kallax.NewBaseSchema(
			"assets",
			"__asset",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Asset)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("symbol"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("is_fiat"),
		),
		ID:     kallax.NewSchemaField("id"),
		Symbol: kallax.NewSchemaField("symbol"),
		Name:   kallax.NewSchemaField("name"),
		IsFiat: kallax.NewSchemaField("is_fiat"),
	},
	Exchange: &schemaExchange{
		BaseSchema: kallax.NewBaseSchema(
			"exchanges",
			"__exchange",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Exchange)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("code"),
		),
		ID:   kallax.NewSchemaField("id"),
		Code: kallax.NewSchemaField("code"),
	},
	Pair: &schemaPair{
		BaseSchema: kallax.NewBaseSchema(
			"pairs",
			"__pair",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"BaseId":  kallax.NewForeignKey("base_id", true),
				"QuoteId": kallax.NewForeignKey("quote_id", true),
			},
			func() kallax.Record {
				return new(Pair)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("symbol"),
			kallax.NewSchemaField("base_id"),
			kallax.NewSchemaField("quote_id"),
		),
		ID:        kallax.NewSchemaField("id"),
		Symbol:    kallax.NewSchemaField("symbol"),
		BaseIdFK:  kallax.NewSchemaField("base_id"),
		QuoteIdFK: kallax.NewSchemaField("quote_id"),
	},
	Price: &schemaPrice{
		BaseSchema: kallax.NewBaseSchema(
			"prices",
			"__price",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"PairId":     kallax.NewForeignKey("pair_id", true),
				"ExchangeId": kallax.NewForeignKey("exchange_id", true),
			},
			func() kallax.Record {
				return new(Price)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("pair_id"),
			kallax.NewSchemaField("exchange_id"),
			kallax.NewSchemaField("price"),
		),
		ID:           kallax.NewSchemaField("id"),
		CreatedAt:    kallax.NewSchemaField("created_at"),
		UpdatedAt:    kallax.NewSchemaField("updated_at"),
		PairIdFK:     kallax.NewSchemaField("pair_id"),
		ExchangeIdFK: kallax.NewSchemaField("exchange_id"),
		Price:        kallax.NewSchemaField("price"),
	},
}
