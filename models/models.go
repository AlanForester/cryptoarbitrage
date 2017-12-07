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
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Asset) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
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
	switch field {
	case "BasePairs":
		return new(Pair), nil
	case "QuotePairs":
		return new(Pair), nil
	case "Balances":
		return new(UserBalance), nil

	}
	return nil, fmt.Errorf("kallax: model Asset has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Asset) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "BasePairs":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.BasePairs = make([]*Pair, len(records))
		for i, record := range records {
			rel, ok := record.(*Pair)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.BasePairs[i] = rel
		}
		return nil
	case "QuotePairs":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.QuotePairs = make([]*Pair, len(records))
		for i, record := range records {
			rel, ok := record.(*Pair)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.QuotePairs[i] = rel
		}
		return nil
	case "Balances":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Balances = make([]*UserBalance, len(records))
		for i, record := range records {
			rel, ok := record.(*UserBalance)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Balances[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model Asset has no relationship %s", field)
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

func (s *AssetStore) relationshipRecords(record *Asset) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.BasePairs {
		r := record.BasePairs[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("base_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&PairStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.QuotePairs {
		r := record.QuotePairs[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("quote_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&PairStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Balances {
		r := record.Balances[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("asset_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&UserBalanceStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

// Insert inserts a Asset in the database. A non-persisted object is
// required for this operation.
func (s *AssetStore) Insert(record *Asset) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			if err := s.Insert(Schema.Asset.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

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

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			updated, err = s.Update(Schema.Asset.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

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

// RemoveBasePairs removes the given items of the BasePairs field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `BasePairs` is not empty. This method clears the
// the elements of BasePairs in a model, it does not retrieve them to know
// what relationships the model has.
func (s *AssetStore) RemoveBasePairs(record *Asset, deleted ...*Pair) error {
	var updated []*Pair
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.BasePairs
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Pair.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.BasePairs = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Pair.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Pair.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.BasePairs {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.BasePairs = updated
	return nil
}

// RemoveQuotePairs removes the given items of the QuotePairs field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `QuotePairs` is not empty. This method clears the
// the elements of QuotePairs in a model, it does not retrieve them to know
// what relationships the model has.
func (s *AssetStore) RemoveQuotePairs(record *Asset, deleted ...*Pair) error {
	var updated []*Pair
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.QuotePairs
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Pair.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.QuotePairs = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Pair.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Pair.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.QuotePairs {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.QuotePairs = updated
	return nil
}

// RemoveBalances removes the given items of the Balances field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Balances` is not empty. This method clears the
// the elements of Balances in a model, it does not retrieve them to know
// what relationships the model has.
func (s *AssetStore) RemoveBalances(record *Asset, deleted ...*UserBalance) error {
	var updated []*UserBalance
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Balances
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.UserBalance.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Balances = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.UserBalance.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.UserBalance.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Balances {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Balances = updated
	return nil
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

func (q *AssetQuery) WithBasePairs(cond kallax.Condition) *AssetQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "BasePairs", kallax.OneToMany, cond)
	return q
}

func (q *AssetQuery) WithQuotePairs(cond kallax.Condition) *AssetQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "QuotePairs", kallax.OneToMany, cond)
	return q
}

func (q *AssetQuery) WithBalances(cond kallax.Condition) *AssetQuery {
	q.AddRelation(Schema.UserBalance.BaseSchema, "Balances", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *AssetQuery) FindByID(v ...kallax.NumericID) *AssetQuery {
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

// NewDifference returns a new instance of Difference.
func NewDifference() (record *Difference) {
	return new(Difference)
}

// GetID returns the primary key of the model.
func (r *Difference) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Difference) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "pair_id":
		return types.Nullable(kallax.VirtualColumn("pair_id", r, new(kallax.NumericID))), nil
	case "base_id":
		return types.Nullable(kallax.VirtualColumn("base_id", r, new(kallax.NumericID))), nil
	case "quote_id":
		return types.Nullable(kallax.VirtualColumn("quote_id", r, new(kallax.NumericID))), nil
	case "delta":
		return &r.Delta, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Difference: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Difference) Value(col string) (interface{}, error) {
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
	case "delta":
		return r.Delta, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Difference: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Difference) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Pair":
		return new(Pair), nil
	case "BaseExchange":
		return new(Exchange), nil
	case "QuoteExchange":
		return new(Exchange), nil

	}
	return nil, fmt.Errorf("kallax: model Difference has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Difference) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Pair":
		val, ok := rel.(*Pair)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Pair", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Pair = val
		}

		return nil
	case "BaseExchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship BaseExchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.BaseExchange = val
		}

		return nil
	case "QuoteExchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship QuoteExchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.QuoteExchange = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model Difference has no relationship %s", field)
}

// DifferenceStore is the entity to access the records of the type Difference
// in the database.
type DifferenceStore struct {
	*kallax.Store
}

// NewDifferenceStore creates a new instance of DifferenceStore
// using a SQL database.
func NewDifferenceStore(db *sql.DB) *DifferenceStore {
	return &DifferenceStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *DifferenceStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *DifferenceStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *DifferenceStore) Debug() *DifferenceStore {
	return &DifferenceStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *DifferenceStore) DebugWith(logger kallax.LoggerFunc) *DifferenceStore {
	return &DifferenceStore{s.Store.DebugWith(logger)}
}

func (s *DifferenceStore) inverseRecords(record *Difference) []modelSaveFunc {
	var result []modelSaveFunc

	if record.Pair != nil && !record.Pair.IsSaving() {
		record.AddVirtualColumn("pair_id", record.Pair.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&PairStore{store}).Save(record.Pair)
			return err
		})
	}

	if record.BaseExchange != nil && !record.BaseExchange.IsSaving() {
		record.AddVirtualColumn("base_id", record.BaseExchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.BaseExchange)
			return err
		})
	}

	if record.QuoteExchange != nil && !record.QuoteExchange.IsSaving() {
		record.AddVirtualColumn("quote_id", record.QuoteExchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.QuoteExchange)
			return err
		})
	}

	return result
}

// Insert inserts a Difference in the database. A non-persisted object is
// required for this operation.
func (s *DifferenceStore) Insert(record *Difference) error {
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

			if err := s.Insert(Schema.Difference.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Difference.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *DifferenceStore) Update(record *Difference, cols ...kallax.SchemaField) (updated int64, err error) {
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

			updated, err = s.Update(Schema.Difference.BaseSchema, record, cols...)
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

	return s.Store.Update(Schema.Difference.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *DifferenceStore) Save(record *Difference) (updated bool, err error) {
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
func (s *DifferenceStore) Delete(record *Difference) error {
	return s.Store.Delete(Schema.Difference.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *DifferenceStore) Find(q *DifferenceQuery) (*DifferenceResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewDifferenceResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *DifferenceStore) MustFind(q *DifferenceQuery) *DifferenceResultSet {
	return NewDifferenceResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *DifferenceStore) Count(q *DifferenceQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *DifferenceStore) MustCount(q *DifferenceQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *DifferenceStore) FindOne(q *DifferenceQuery) (*Difference, error) {
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
func (s *DifferenceStore) FindAll(q *DifferenceQuery) ([]*Difference, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *DifferenceStore) MustFindOne(q *DifferenceQuery) *Difference {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Difference with the data in the database and
// makes it writable.
func (s *DifferenceStore) Reload(record *Difference) error {
	return s.Store.Reload(Schema.Difference.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *DifferenceStore) Transaction(callback func(*DifferenceStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&DifferenceStore{store})
	})
}

// DifferenceQuery is the object used to create queries for the Difference
// entity.
type DifferenceQuery struct {
	*kallax.BaseQuery
}

// NewDifferenceQuery returns a new instance of DifferenceQuery.
func NewDifferenceQuery() *DifferenceQuery {
	return &DifferenceQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Difference.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *DifferenceQuery) Select(columns ...kallax.SchemaField) *DifferenceQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *DifferenceQuery) SelectNot(columns ...kallax.SchemaField) *DifferenceQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *DifferenceQuery) Copy() *DifferenceQuery {
	return &DifferenceQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *DifferenceQuery) Order(cols ...kallax.ColumnOrder) *DifferenceQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *DifferenceQuery) BatchSize(size uint64) *DifferenceQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *DifferenceQuery) Limit(n uint64) *DifferenceQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *DifferenceQuery) Offset(n uint64) *DifferenceQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *DifferenceQuery) Where(cond kallax.Condition) *DifferenceQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *DifferenceQuery) WithPair() *DifferenceQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "Pair", kallax.OneToOne, nil)
	return q
}

func (q *DifferenceQuery) WithBaseExchange() *DifferenceQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "BaseExchange", kallax.OneToOne, nil)
	return q
}

func (q *DifferenceQuery) WithQuoteExchange() *DifferenceQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "QuoteExchange", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *DifferenceQuery) FindByID(v ...kallax.NumericID) *DifferenceQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Difference.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *DifferenceQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *DifferenceQuery {
	return q.Where(cond(Schema.Difference.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *DifferenceQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *DifferenceQuery {
	return q.Where(cond(Schema.Difference.UpdatedAt, v))
}

// FindByPair adds a new filter to the query that will require that
// the foreign key of Pair is equal to the passed value.
func (q *DifferenceQuery) FindByPair(v kallax.NumericID) *DifferenceQuery {
	return q.Where(kallax.Eq(Schema.Difference.PairFK, v))
}

// FindByBaseExchange adds a new filter to the query that will require that
// the foreign key of BaseExchange is equal to the passed value.
func (q *DifferenceQuery) FindByBaseExchange(v kallax.NumericID) *DifferenceQuery {
	return q.Where(kallax.Eq(Schema.Difference.BaseExchangeFK, v))
}

// FindByQuoteExchange adds a new filter to the query that will require that
// the foreign key of QuoteExchange is equal to the passed value.
func (q *DifferenceQuery) FindByQuoteExchange(v kallax.NumericID) *DifferenceQuery {
	return q.Where(kallax.Eq(Schema.Difference.QuoteExchangeFK, v))
}

// FindByDelta adds a new filter to the query that will require that
// the Delta property is equal to the passed value.
func (q *DifferenceQuery) FindByDelta(cond kallax.ScalarCond, v float32) *DifferenceQuery {
	return q.Where(cond(Schema.Difference.Delta, v))
}

// DifferenceResultSet is the set of results returned by a query to the
// database.
type DifferenceResultSet struct {
	ResultSet kallax.ResultSet
	last      *Difference
	lastErr   error
}

// NewDifferenceResultSet creates a new result set for rows of the type
// Difference.
func NewDifferenceResultSet(rs kallax.ResultSet) *DifferenceResultSet {
	return &DifferenceResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *DifferenceResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Difference.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Difference)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Difference")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *DifferenceResultSet) Get() (*Difference, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *DifferenceResultSet) ForEach(fn func(*Difference) error) error {
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
func (rs *DifferenceResultSet) All() ([]*Difference, error) {
	var result []*Difference
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
func (rs *DifferenceResultSet) One() (*Difference, error) {
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
func (rs *DifferenceResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *DifferenceResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewExchange returns a new instance of Exchange.
func NewExchange() (record *Exchange) {
	return new(Exchange)
}

// GetID returns the primary key of the model.
func (r *Exchange) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Exchange) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "symbol":
		return &r.Symbol, nil
	case "name":
		return &r.Name, nil
	case "is_active":
		return &r.IsActive, nil
	case "is_used_api":
		return &r.IsUsedAPI, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Exchange: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Exchange) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "symbol":
		return r.Symbol, nil
	case "name":
		return r.Name, nil
	case "is_active":
		return r.IsActive, nil
	case "is_used_api":
		return r.IsUsedAPI, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Exchange: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Exchange) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Markets":
		return new(Market), nil
	case "Assets":
		return new(ExchangeAsset), nil
	case "BaseDifferences":
		return new(Difference), nil
	case "QuoteDifferences":
		return new(Difference), nil
	case "Orders":
		return new(Order), nil
	case "Prices":
		return new(Price), nil
	case "Trades":
		return new(Trade), nil
	case "Balances":
		return new(UserBalance), nil

	}
	return nil, fmt.Errorf("kallax: model Exchange has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Exchange) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Markets":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Markets = make([]*Market, len(records))
		for i, record := range records {
			rel, ok := record.(*Market)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Markets[i] = rel
		}
		return nil
	case "Assets":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Assets = make([]*ExchangeAsset, len(records))
		for i, record := range records {
			rel, ok := record.(*ExchangeAsset)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Assets[i] = rel
		}
		return nil
	case "BaseDifferences":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.BaseDifferences = make([]*Difference, len(records))
		for i, record := range records {
			rel, ok := record.(*Difference)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.BaseDifferences[i] = rel
		}
		return nil
	case "QuoteDifferences":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.QuoteDifferences = make([]*Difference, len(records))
		for i, record := range records {
			rel, ok := record.(*Difference)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.QuoteDifferences[i] = rel
		}
		return nil
	case "Orders":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Orders = make([]*Order, len(records))
		for i, record := range records {
			rel, ok := record.(*Order)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Orders[i] = rel
		}
		return nil
	case "Prices":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Prices = make([]*Price, len(records))
		for i, record := range records {
			rel, ok := record.(*Price)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Prices[i] = rel
		}
		return nil
	case "Trades":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Trades = make([]*Trade, len(records))
		for i, record := range records {
			rel, ok := record.(*Trade)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Trades[i] = rel
		}
		return nil
	case "Balances":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Balances = make([]*UserBalance, len(records))
		for i, record := range records {
			rel, ok := record.(*UserBalance)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Balances[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model Exchange has no relationship %s", field)
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

func (s *ExchangeStore) relationshipRecords(record *Exchange) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Markets {
		r := record.Markets[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("exchange_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&MarketStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Assets {
		r := record.Assets[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("exchange_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&ExchangeAssetStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.BaseDifferences {
		r := record.BaseDifferences[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("base_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&DifferenceStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.QuoteDifferences {
		r := record.QuoteDifferences[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("quote_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&DifferenceStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Orders {
		r := record.Orders[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("exchange_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&OrderStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Prices {
		r := record.Prices[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("exchange_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&PriceStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Trades {
		r := record.Trades[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("exchange_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&TradeStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Balances {
		r := record.Balances[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("exchange_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&UserBalanceStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

// Insert inserts a Exchange in the database. A non-persisted object is
// required for this operation.
func (s *ExchangeStore) Insert(record *Exchange) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			if err := s.Insert(Schema.Exchange.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

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

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			updated, err = s.Update(Schema.Exchange.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

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

// RemoveMarkets removes the given items of the Markets field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Markets` is not empty. This method clears the
// the elements of Markets in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemoveMarkets(record *Exchange, deleted ...*Market) error {
	var updated []*Market
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Markets
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Market.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Markets = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Market.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Market.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Markets {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Markets = updated
	return nil
}

// RemoveAssets removes the given items of the Assets field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Assets` is not empty. This method clears the
// the elements of Assets in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemoveAssets(record *Exchange, deleted ...*ExchangeAsset) error {
	var updated []*ExchangeAsset
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Assets
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.ExchangeAsset.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Assets = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.ExchangeAsset.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.ExchangeAsset.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Assets {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Assets = updated
	return nil
}

// RemoveBaseDifferences removes the given items of the BaseDifferences field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `BaseDifferences` is not empty. This method clears the
// the elements of BaseDifferences in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemoveBaseDifferences(record *Exchange, deleted ...*Difference) error {
	var updated []*Difference
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.BaseDifferences
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Difference.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.BaseDifferences = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Difference.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Difference.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.BaseDifferences {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.BaseDifferences = updated
	return nil
}

// RemoveQuoteDifferences removes the given items of the QuoteDifferences field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `QuoteDifferences` is not empty. This method clears the
// the elements of QuoteDifferences in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemoveQuoteDifferences(record *Exchange, deleted ...*Difference) error {
	var updated []*Difference
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.QuoteDifferences
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Difference.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.QuoteDifferences = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Difference.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Difference.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.QuoteDifferences {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.QuoteDifferences = updated
	return nil
}

// RemoveOrders removes the given items of the Orders field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Orders` is not empty. This method clears the
// the elements of Orders in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemoveOrders(record *Exchange, deleted ...*Order) error {
	var updated []*Order
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Orders
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Order.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Orders = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Order.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Order.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Orders {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Orders = updated
	return nil
}

// RemovePrices removes the given items of the Prices field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Prices` is not empty. This method clears the
// the elements of Prices in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemovePrices(record *Exchange, deleted ...*Price) error {
	var updated []*Price
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Prices
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Price.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Prices = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Price.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Price.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Prices {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Prices = updated
	return nil
}

// RemoveTrades removes the given items of the Trades field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Trades` is not empty. This method clears the
// the elements of Trades in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemoveTrades(record *Exchange, deleted ...*Trade) error {
	var updated []*Trade
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Trades
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Trade.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Trades = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Trade.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Trade.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Trades {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Trades = updated
	return nil
}

// RemoveBalances removes the given items of the Balances field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Balances` is not empty. This method clears the
// the elements of Balances in a model, it does not retrieve them to know
// what relationships the model has.
func (s *ExchangeStore) RemoveBalances(record *Exchange, deleted ...*UserBalance) error {
	var updated []*UserBalance
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Balances
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.UserBalance.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Balances = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.UserBalance.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.UserBalance.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Balances {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Balances = updated
	return nil
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

func (q *ExchangeQuery) WithMarkets(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.Market.BaseSchema, "Markets", kallax.OneToMany, cond)
	return q
}

func (q *ExchangeQuery) WithAssets(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.ExchangeAsset.BaseSchema, "Assets", kallax.OneToMany, cond)
	return q
}

func (q *ExchangeQuery) WithBaseDifferences(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.Difference.BaseSchema, "BaseDifferences", kallax.OneToMany, cond)
	return q
}

func (q *ExchangeQuery) WithQuoteDifferences(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.Difference.BaseSchema, "QuoteDifferences", kallax.OneToMany, cond)
	return q
}

func (q *ExchangeQuery) WithOrders(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.Order.BaseSchema, "Orders", kallax.OneToMany, cond)
	return q
}

func (q *ExchangeQuery) WithPrices(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.Price.BaseSchema, "Prices", kallax.OneToMany, cond)
	return q
}

func (q *ExchangeQuery) WithTrades(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.Trade.BaseSchema, "Trades", kallax.OneToMany, cond)
	return q
}

func (q *ExchangeQuery) WithBalances(cond kallax.Condition) *ExchangeQuery {
	q.AddRelation(Schema.UserBalance.BaseSchema, "Balances", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *ExchangeQuery) FindByID(v ...kallax.NumericID) *ExchangeQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Exchange.ID, values...))
}

// FindBySymbol adds a new filter to the query that will require that
// the Symbol property is equal to the passed value.
func (q *ExchangeQuery) FindBySymbol(v string) *ExchangeQuery {
	return q.Where(kallax.Eq(Schema.Exchange.Symbol, v))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *ExchangeQuery) FindByName(v string) *ExchangeQuery {
	return q.Where(kallax.Eq(Schema.Exchange.Name, v))
}

// FindByIsActive adds a new filter to the query that will require that
// the IsActive property is equal to the passed value.
func (q *ExchangeQuery) FindByIsActive(v bool) *ExchangeQuery {
	return q.Where(kallax.Eq(Schema.Exchange.IsActive, v))
}

// FindByIsUsedAPI adds a new filter to the query that will require that
// the IsUsedAPI property is equal to the passed value.
func (q *ExchangeQuery) FindByIsUsedAPI(v bool) *ExchangeQuery {
	return q.Where(kallax.Eq(Schema.Exchange.IsUsedAPI, v))
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

// NewExchangeAsset returns a new instance of ExchangeAsset.
func NewExchangeAsset() (record *ExchangeAsset) {
	return new(ExchangeAsset)
}

// GetID returns the primary key of the model.
func (r *ExchangeAsset) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *ExchangeAsset) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "asset_id":
		return types.Nullable(kallax.VirtualColumn("asset_id", r, new(kallax.NumericID))), nil
	case "exchange_id":
		return types.Nullable(kallax.VirtualColumn("exchange_id", r, new(kallax.NumericID))), nil
	case "transaction_fee":
		return &r.TransactionFee, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in ExchangeAsset: %s", col)
	}
}

// Value returns the value of the given column.
func (r *ExchangeAsset) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "asset_id":
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
	case "transaction_fee":
		return r.TransactionFee, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in ExchangeAsset: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *ExchangeAsset) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Asset":
		return new(Asset), nil
	case "Exchange":
		return new(Exchange), nil

	}
	return nil, fmt.Errorf("kallax: model ExchangeAsset has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *ExchangeAsset) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Asset":
		val, ok := rel.(*Asset)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Asset", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Asset = val
		}

		return nil
	case "Exchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Exchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Exchange = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model ExchangeAsset has no relationship %s", field)
}

// ExchangeAssetStore is the entity to access the records of the type ExchangeAsset
// in the database.
type ExchangeAssetStore struct {
	*kallax.Store
}

// NewExchangeAssetStore creates a new instance of ExchangeAssetStore
// using a SQL database.
func NewExchangeAssetStore(db *sql.DB) *ExchangeAssetStore {
	return &ExchangeAssetStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *ExchangeAssetStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *ExchangeAssetStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *ExchangeAssetStore) Debug() *ExchangeAssetStore {
	return &ExchangeAssetStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *ExchangeAssetStore) DebugWith(logger kallax.LoggerFunc) *ExchangeAssetStore {
	return &ExchangeAssetStore{s.Store.DebugWith(logger)}
}

func (s *ExchangeAssetStore) inverseRecords(record *ExchangeAsset) []modelSaveFunc {
	var result []modelSaveFunc

	if record.Asset != nil && !record.Asset.IsSaving() {
		record.AddVirtualColumn("asset_id", record.Asset.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&AssetStore{store}).Save(record.Asset)
			return err
		})
	}

	if record.Exchange != nil && !record.Exchange.IsSaving() {
		record.AddVirtualColumn("exchange_id", record.Exchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.Exchange)
			return err
		})
	}

	return result
}

// Insert inserts a ExchangeAsset in the database. A non-persisted object is
// required for this operation.
func (s *ExchangeAssetStore) Insert(record *ExchangeAsset) error {
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

			if err := s.Insert(Schema.ExchangeAsset.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.ExchangeAsset.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *ExchangeAssetStore) Update(record *ExchangeAsset, cols ...kallax.SchemaField) (updated int64, err error) {
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

			updated, err = s.Update(Schema.ExchangeAsset.BaseSchema, record, cols...)
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

	return s.Store.Update(Schema.ExchangeAsset.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *ExchangeAssetStore) Save(record *ExchangeAsset) (updated bool, err error) {
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
func (s *ExchangeAssetStore) Delete(record *ExchangeAsset) error {
	return s.Store.Delete(Schema.ExchangeAsset.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *ExchangeAssetStore) Find(q *ExchangeAssetQuery) (*ExchangeAssetResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewExchangeAssetResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *ExchangeAssetStore) MustFind(q *ExchangeAssetQuery) *ExchangeAssetResultSet {
	return NewExchangeAssetResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *ExchangeAssetStore) Count(q *ExchangeAssetQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *ExchangeAssetStore) MustCount(q *ExchangeAssetQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *ExchangeAssetStore) FindOne(q *ExchangeAssetQuery) (*ExchangeAsset, error) {
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
func (s *ExchangeAssetStore) FindAll(q *ExchangeAssetQuery) ([]*ExchangeAsset, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *ExchangeAssetStore) MustFindOne(q *ExchangeAssetQuery) *ExchangeAsset {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the ExchangeAsset with the data in the database and
// makes it writable.
func (s *ExchangeAssetStore) Reload(record *ExchangeAsset) error {
	return s.Store.Reload(Schema.ExchangeAsset.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *ExchangeAssetStore) Transaction(callback func(*ExchangeAssetStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&ExchangeAssetStore{store})
	})
}

// ExchangeAssetQuery is the object used to create queries for the ExchangeAsset
// entity.
type ExchangeAssetQuery struct {
	*kallax.BaseQuery
}

// NewExchangeAssetQuery returns a new instance of ExchangeAssetQuery.
func NewExchangeAssetQuery() *ExchangeAssetQuery {
	return &ExchangeAssetQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.ExchangeAsset.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *ExchangeAssetQuery) Select(columns ...kallax.SchemaField) *ExchangeAssetQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *ExchangeAssetQuery) SelectNot(columns ...kallax.SchemaField) *ExchangeAssetQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *ExchangeAssetQuery) Copy() *ExchangeAssetQuery {
	return &ExchangeAssetQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *ExchangeAssetQuery) Order(cols ...kallax.ColumnOrder) *ExchangeAssetQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *ExchangeAssetQuery) BatchSize(size uint64) *ExchangeAssetQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *ExchangeAssetQuery) Limit(n uint64) *ExchangeAssetQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *ExchangeAssetQuery) Offset(n uint64) *ExchangeAssetQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *ExchangeAssetQuery) Where(cond kallax.Condition) *ExchangeAssetQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *ExchangeAssetQuery) WithAsset() *ExchangeAssetQuery {
	q.AddRelation(Schema.Asset.BaseSchema, "Asset", kallax.OneToOne, nil)
	return q
}

func (q *ExchangeAssetQuery) WithExchange() *ExchangeAssetQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "Exchange", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *ExchangeAssetQuery) FindByID(v ...kallax.NumericID) *ExchangeAssetQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.ExchangeAsset.ID, values...))
}

// FindByAsset adds a new filter to the query that will require that
// the foreign key of Asset is equal to the passed value.
func (q *ExchangeAssetQuery) FindByAsset(v kallax.NumericID) *ExchangeAssetQuery {
	return q.Where(kallax.Eq(Schema.ExchangeAsset.AssetFK, v))
}

// FindByExchange adds a new filter to the query that will require that
// the foreign key of Exchange is equal to the passed value.
func (q *ExchangeAssetQuery) FindByExchange(v kallax.NumericID) *ExchangeAssetQuery {
	return q.Where(kallax.Eq(Schema.ExchangeAsset.ExchangeFK, v))
}

// FindByTransactionFee adds a new filter to the query that will require that
// the TransactionFee property is equal to the passed value.
func (q *ExchangeAssetQuery) FindByTransactionFee(cond kallax.ScalarCond, v float32) *ExchangeAssetQuery {
	return q.Where(cond(Schema.ExchangeAsset.TransactionFee, v))
}

// ExchangeAssetResultSet is the set of results returned by a query to the
// database.
type ExchangeAssetResultSet struct {
	ResultSet kallax.ResultSet
	last      *ExchangeAsset
	lastErr   error
}

// NewExchangeAssetResultSet creates a new result set for rows of the type
// ExchangeAsset.
func NewExchangeAssetResultSet(rs kallax.ResultSet) *ExchangeAssetResultSet {
	return &ExchangeAssetResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *ExchangeAssetResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.ExchangeAsset.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*ExchangeAsset)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *ExchangeAsset")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *ExchangeAssetResultSet) Get() (*ExchangeAsset, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *ExchangeAssetResultSet) ForEach(fn func(*ExchangeAsset) error) error {
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
func (rs *ExchangeAssetResultSet) All() ([]*ExchangeAsset, error) {
	var result []*ExchangeAsset
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
func (rs *ExchangeAssetResultSet) One() (*ExchangeAsset, error) {
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
func (rs *ExchangeAssetResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *ExchangeAssetResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewMarket returns a new instance of Market.
func NewMarket() (record *Market) {
	return new(Market)
}

// GetID returns the primary key of the model.
func (r *Market) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Market) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "pair_id":
		return types.Nullable(kallax.VirtualColumn("pair_id", r, new(kallax.NumericID))), nil
	case "exchange_id":
		return types.Nullable(kallax.VirtualColumn("exchange_id", r, new(kallax.NumericID))), nil
	case "is_active":
		return &r.IsActive, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Market: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Market) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
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
	case "is_active":
		return r.IsActive, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Market: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Market) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Pair":
		return new(Pair), nil
	case "Exchange":
		return new(Exchange), nil
	case "Orders":
		return new(Order), nil
	case "Prices":
		return new(Price), nil
	case "Trades":
		return new(Trade), nil

	}
	return nil, fmt.Errorf("kallax: model Market has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Market) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Pair":
		val, ok := rel.(*Pair)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Pair", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Pair = val
		}

		return nil
	case "Exchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Exchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Exchange = val
		}

		return nil
	case "Orders":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Orders = make([]*Order, len(records))
		for i, record := range records {
			rel, ok := record.(*Order)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Orders[i] = rel
		}
		return nil
	case "Prices":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Prices = make([]*Price, len(records))
		for i, record := range records {
			rel, ok := record.(*Price)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Prices[i] = rel
		}
		return nil
	case "Trades":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Trades = make([]*Trade, len(records))
		for i, record := range records {
			rel, ok := record.(*Trade)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Trades[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model Market has no relationship %s", field)
}

// MarketStore is the entity to access the records of the type Market
// in the database.
type MarketStore struct {
	*kallax.Store
}

// NewMarketStore creates a new instance of MarketStore
// using a SQL database.
func NewMarketStore(db *sql.DB) *MarketStore {
	return &MarketStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *MarketStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *MarketStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *MarketStore) Debug() *MarketStore {
	return &MarketStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *MarketStore) DebugWith(logger kallax.LoggerFunc) *MarketStore {
	return &MarketStore{s.Store.DebugWith(logger)}
}

func (s *MarketStore) relationshipRecords(record *Market) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Orders {
		r := record.Orders[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("market_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&OrderStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Prices {
		r := record.Prices[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("market_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&PriceStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Trades {
		r := record.Trades[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("market_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&TradeStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

func (s *MarketStore) inverseRecords(record *Market) []modelSaveFunc {
	var result []modelSaveFunc

	if record.Pair != nil && !record.Pair.IsSaving() {
		record.AddVirtualColumn("pair_id", record.Pair.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&PairStore{store}).Save(record.Pair)
			return err
		})
	}

	if record.Exchange != nil && !record.Exchange.IsSaving() {
		record.AddVirtualColumn("exchange_id", record.Exchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.Exchange)
			return err
		})
	}

	return result
}

// Insert inserts a Market in the database. A non-persisted object is
// required for this operation.
func (s *MarketStore) Insert(record *Market) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	records := s.relationshipRecords(record)

	inverseRecords := s.inverseRecords(record)

	if len(records) > 0 || len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Market.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Market.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *MarketStore) Update(record *Market, cols ...kallax.SchemaField) (updated int64, err error) {
	record.SetSaving(true)
	defer record.SetSaving(false)

	records := s.relationshipRecords(record)

	inverseRecords := s.inverseRecords(record)

	if len(records) > 0 || len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Market.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Market.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *MarketStore) Save(record *Market) (updated bool, err error) {
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
func (s *MarketStore) Delete(record *Market) error {
	return s.Store.Delete(Schema.Market.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *MarketStore) Find(q *MarketQuery) (*MarketResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewMarketResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *MarketStore) MustFind(q *MarketQuery) *MarketResultSet {
	return NewMarketResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *MarketStore) Count(q *MarketQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *MarketStore) MustCount(q *MarketQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *MarketStore) FindOne(q *MarketQuery) (*Market, error) {
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
func (s *MarketStore) FindAll(q *MarketQuery) ([]*Market, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *MarketStore) MustFindOne(q *MarketQuery) *Market {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Market with the data in the database and
// makes it writable.
func (s *MarketStore) Reload(record *Market) error {
	return s.Store.Reload(Schema.Market.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *MarketStore) Transaction(callback func(*MarketStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&MarketStore{store})
	})
}

// RemoveOrders removes the given items of the Orders field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Orders` is not empty. This method clears the
// the elements of Orders in a model, it does not retrieve them to know
// what relationships the model has.
func (s *MarketStore) RemoveOrders(record *Market, deleted ...*Order) error {
	var updated []*Order
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Orders
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Order.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Orders = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Order.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Order.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Orders {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Orders = updated
	return nil
}

// RemovePrices removes the given items of the Prices field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Prices` is not empty. This method clears the
// the elements of Prices in a model, it does not retrieve them to know
// what relationships the model has.
func (s *MarketStore) RemovePrices(record *Market, deleted ...*Price) error {
	var updated []*Price
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Prices
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Price.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Prices = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Price.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Price.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Prices {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Prices = updated
	return nil
}

// RemoveTrades removes the given items of the Trades field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Trades` is not empty. This method clears the
// the elements of Trades in a model, it does not retrieve them to know
// what relationships the model has.
func (s *MarketStore) RemoveTrades(record *Market, deleted ...*Trade) error {
	var updated []*Trade
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Trades
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Trade.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Trades = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Trade.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Trade.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Trades {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Trades = updated
	return nil
}

// MarketQuery is the object used to create queries for the Market
// entity.
type MarketQuery struct {
	*kallax.BaseQuery
}

// NewMarketQuery returns a new instance of MarketQuery.
func NewMarketQuery() *MarketQuery {
	return &MarketQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Market.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *MarketQuery) Select(columns ...kallax.SchemaField) *MarketQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *MarketQuery) SelectNot(columns ...kallax.SchemaField) *MarketQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *MarketQuery) Copy() *MarketQuery {
	return &MarketQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *MarketQuery) Order(cols ...kallax.ColumnOrder) *MarketQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *MarketQuery) BatchSize(size uint64) *MarketQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *MarketQuery) Limit(n uint64) *MarketQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *MarketQuery) Offset(n uint64) *MarketQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *MarketQuery) Where(cond kallax.Condition) *MarketQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *MarketQuery) WithPair() *MarketQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "Pair", kallax.OneToOne, nil)
	return q
}

func (q *MarketQuery) WithExchange() *MarketQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "Exchange", kallax.OneToOne, nil)
	return q
}

func (q *MarketQuery) WithOrders(cond kallax.Condition) *MarketQuery {
	q.AddRelation(Schema.Order.BaseSchema, "Orders", kallax.OneToMany, cond)
	return q
}

func (q *MarketQuery) WithPrices(cond kallax.Condition) *MarketQuery {
	q.AddRelation(Schema.Price.BaseSchema, "Prices", kallax.OneToMany, cond)
	return q
}

func (q *MarketQuery) WithTrades(cond kallax.Condition) *MarketQuery {
	q.AddRelation(Schema.Trade.BaseSchema, "Trades", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *MarketQuery) FindByID(v ...kallax.NumericID) *MarketQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Market.ID, values...))
}

// FindByPair adds a new filter to the query that will require that
// the foreign key of Pair is equal to the passed value.
func (q *MarketQuery) FindByPair(v kallax.NumericID) *MarketQuery {
	return q.Where(kallax.Eq(Schema.Market.PairFK, v))
}

// FindByExchange adds a new filter to the query that will require that
// the foreign key of Exchange is equal to the passed value.
func (q *MarketQuery) FindByExchange(v kallax.NumericID) *MarketQuery {
	return q.Where(kallax.Eq(Schema.Market.ExchangeFK, v))
}

// FindByIsActive adds a new filter to the query that will require that
// the IsActive property is equal to the passed value.
func (q *MarketQuery) FindByIsActive(v bool) *MarketQuery {
	return q.Where(kallax.Eq(Schema.Market.IsActive, v))
}

// MarketResultSet is the set of results returned by a query to the
// database.
type MarketResultSet struct {
	ResultSet kallax.ResultSet
	last      *Market
	lastErr   error
}

// NewMarketResultSet creates a new result set for rows of the type
// Market.
func NewMarketResultSet(rs kallax.ResultSet) *MarketResultSet {
	return &MarketResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *MarketResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Market.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Market)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Market")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *MarketResultSet) Get() (*Market, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *MarketResultSet) ForEach(fn func(*Market) error) error {
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
func (rs *MarketResultSet) All() ([]*Market, error) {
	var result []*Market
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
func (rs *MarketResultSet) One() (*Market, error) {
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
func (rs *MarketResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *MarketResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewOrder returns a new instance of Order.
func NewOrder() (record *Order) {
	return new(Order)
}

// GetID returns the primary key of the model.
func (r *Order) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Order) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "user_id":
		return types.Nullable(kallax.VirtualColumn("user_id", r, new(kallax.NumericID))), nil
	case "exchange_id":
		return types.Nullable(kallax.VirtualColumn("exchange_id", r, new(kallax.NumericID))), nil
	case "pair_id":
		return types.Nullable(kallax.VirtualColumn("pair_id", r, new(kallax.NumericID))), nil
	case "market_id":
		return types.Nullable(kallax.VirtualColumn("market_id", r, new(kallax.NumericID))), nil
	case "order_type":
		return &r.OrderType, nil
	case "open_price":
		return &r.OpenPrice, nil
	case "close_price":
		return &r.ClosePrice, nil
	case "ordered_volume":
		return &r.OrderedVolume, nil
	case "swapped_volume":
		return &r.SwappedVolume, nil
	case "is_closed":
		return &r.IsClosed, nil
	case "stop_loss":
		return &r.StopLoss, nil
	case "take_profit":
		return &r.TakeProfit, nil
	case "buy_fee":
		return &r.BuyFee, nil
	case "sell_fee":
		return &r.SellFee, nil
	case "delta":
		return &r.Delta, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Order: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Order) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "user_id":
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
	case "pair_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "market_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "order_type":
		return r.OrderType, nil
	case "open_price":
		return r.OpenPrice, nil
	case "close_price":
		return r.ClosePrice, nil
	case "ordered_volume":
		return r.OrderedVolume, nil
	case "swapped_volume":
		return r.SwappedVolume, nil
	case "is_closed":
		return r.IsClosed, nil
	case "stop_loss":
		return r.StopLoss, nil
	case "take_profit":
		return r.TakeProfit, nil
	case "buy_fee":
		return r.BuyFee, nil
	case "sell_fee":
		return r.SellFee, nil
	case "delta":
		return r.Delta, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Order: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Order) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "User":
		return new(User), nil
	case "Exchange":
		return new(Exchange), nil
	case "Pair":
		return new(Pair), nil
	case "Market":
		return new(Market), nil
	case "Trades":
		return new(Trade), nil

	}
	return nil, fmt.Errorf("kallax: model Order has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Order) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "User":
		val, ok := rel.(*User)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship User", rel)
		}
		if !val.GetID().IsEmpty() {
			r.User = val
		}

		return nil
	case "Exchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Exchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Exchange = val
		}

		return nil
	case "Pair":
		val, ok := rel.(*Pair)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Pair", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Pair = val
		}

		return nil
	case "Market":
		val, ok := rel.(*Market)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Market", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Market = val
		}

		return nil
	case "Trades":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Trades = make([]*Trade, len(records))
		for i, record := range records {
			rel, ok := record.(*Trade)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Trades[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model Order has no relationship %s", field)
}

// OrderStore is the entity to access the records of the type Order
// in the database.
type OrderStore struct {
	*kallax.Store
}

// NewOrderStore creates a new instance of OrderStore
// using a SQL database.
func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *OrderStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *OrderStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *OrderStore) Debug() *OrderStore {
	return &OrderStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *OrderStore) DebugWith(logger kallax.LoggerFunc) *OrderStore {
	return &OrderStore{s.Store.DebugWith(logger)}
}

func (s *OrderStore) relationshipRecords(record *Order) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Trades {
		r := record.Trades[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("order_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&TradeStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

func (s *OrderStore) inverseRecords(record *Order) []modelSaveFunc {
	var result []modelSaveFunc

	if record.User != nil && !record.User.IsSaving() {
		record.AddVirtualColumn("user_id", record.User.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&UserStore{store}).Save(record.User)
			return err
		})
	}

	if record.Exchange != nil && !record.Exchange.IsSaving() {
		record.AddVirtualColumn("exchange_id", record.Exchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.Exchange)
			return err
		})
	}

	if record.Pair != nil && !record.Pair.IsSaving() {
		record.AddVirtualColumn("pair_id", record.Pair.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&PairStore{store}).Save(record.Pair)
			return err
		})
	}

	if record.Market != nil && !record.Market.IsSaving() {
		record.AddVirtualColumn("market_id", record.Market.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&MarketStore{store}).Save(record.Market)
			return err
		})
	}

	return result
}

// Insert inserts a Order in the database. A non-persisted object is
// required for this operation.
func (s *OrderStore) Insert(record *Order) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	records := s.relationshipRecords(record)

	inverseRecords := s.inverseRecords(record)

	if len(records) > 0 || len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Order.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Order.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *OrderStore) Update(record *Order, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	records := s.relationshipRecords(record)

	inverseRecords := s.inverseRecords(record)

	if len(records) > 0 || len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Order.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Order.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *OrderStore) Save(record *Order) (updated bool, err error) {
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
func (s *OrderStore) Delete(record *Order) error {
	return s.Store.Delete(Schema.Order.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *OrderStore) Find(q *OrderQuery) (*OrderResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewOrderResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *OrderStore) MustFind(q *OrderQuery) *OrderResultSet {
	return NewOrderResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *OrderStore) Count(q *OrderQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *OrderStore) MustCount(q *OrderQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *OrderStore) FindOne(q *OrderQuery) (*Order, error) {
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
func (s *OrderStore) FindAll(q *OrderQuery) ([]*Order, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *OrderStore) MustFindOne(q *OrderQuery) *Order {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Order with the data in the database and
// makes it writable.
func (s *OrderStore) Reload(record *Order) error {
	return s.Store.Reload(Schema.Order.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *OrderStore) Transaction(callback func(*OrderStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&OrderStore{store})
	})
}

// RemoveTrades removes the given items of the Trades field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Trades` is not empty. This method clears the
// the elements of Trades in a model, it does not retrieve them to know
// what relationships the model has.
func (s *OrderStore) RemoveTrades(record *Order, deleted ...*Trade) error {
	var updated []*Trade
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Trades
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Trade.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Trades = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Trade.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Trade.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Trades {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Trades = updated
	return nil
}

// OrderQuery is the object used to create queries for the Order
// entity.
type OrderQuery struct {
	*kallax.BaseQuery
}

// NewOrderQuery returns a new instance of OrderQuery.
func NewOrderQuery() *OrderQuery {
	return &OrderQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Order.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *OrderQuery) Select(columns ...kallax.SchemaField) *OrderQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *OrderQuery) SelectNot(columns ...kallax.SchemaField) *OrderQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *OrderQuery) Copy() *OrderQuery {
	return &OrderQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *OrderQuery) Order(cols ...kallax.ColumnOrder) *OrderQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *OrderQuery) BatchSize(size uint64) *OrderQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *OrderQuery) Limit(n uint64) *OrderQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *OrderQuery) Offset(n uint64) *OrderQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *OrderQuery) Where(cond kallax.Condition) *OrderQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *OrderQuery) WithUser() *OrderQuery {
	q.AddRelation(Schema.User.BaseSchema, "User", kallax.OneToOne, nil)
	return q
}

func (q *OrderQuery) WithExchange() *OrderQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "Exchange", kallax.OneToOne, nil)
	return q
}

func (q *OrderQuery) WithPair() *OrderQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "Pair", kallax.OneToOne, nil)
	return q
}

func (q *OrderQuery) WithMarket() *OrderQuery {
	q.AddRelation(Schema.Market.BaseSchema, "Market", kallax.OneToOne, nil)
	return q
}

func (q *OrderQuery) WithTrades(cond kallax.Condition) *OrderQuery {
	q.AddRelation(Schema.Trade.BaseSchema, "Trades", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *OrderQuery) FindByID(v ...kallax.NumericID) *OrderQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Order.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *OrderQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *OrderQuery {
	return q.Where(cond(Schema.Order.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *OrderQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *OrderQuery {
	return q.Where(cond(Schema.Order.UpdatedAt, v))
}

// FindByUser adds a new filter to the query that will require that
// the foreign key of User is equal to the passed value.
func (q *OrderQuery) FindByUser(v kallax.NumericID) *OrderQuery {
	return q.Where(kallax.Eq(Schema.Order.UserFK, v))
}

// FindByExchange adds a new filter to the query that will require that
// the foreign key of Exchange is equal to the passed value.
func (q *OrderQuery) FindByExchange(v kallax.NumericID) *OrderQuery {
	return q.Where(kallax.Eq(Schema.Order.ExchangeFK, v))
}

// FindByPair adds a new filter to the query that will require that
// the foreign key of Pair is equal to the passed value.
func (q *OrderQuery) FindByPair(v kallax.NumericID) *OrderQuery {
	return q.Where(kallax.Eq(Schema.Order.PairFK, v))
}

// FindByMarket adds a new filter to the query that will require that
// the foreign key of Market is equal to the passed value.
func (q *OrderQuery) FindByMarket(v kallax.NumericID) *OrderQuery {
	return q.Where(kallax.Eq(Schema.Order.MarketFK, v))
}

// FindByOrderType adds a new filter to the query that will require that
// the OrderType property is equal to the passed value.
func (q *OrderQuery) FindByOrderType(v string) *OrderQuery {
	return q.Where(kallax.Eq(Schema.Order.OrderType, v))
}

// FindByOpenPrice adds a new filter to the query that will require that
// the OpenPrice property is equal to the passed value.
func (q *OrderQuery) FindByOpenPrice(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.OpenPrice, v))
}

// FindByClosePrice adds a new filter to the query that will require that
// the ClosePrice property is equal to the passed value.
func (q *OrderQuery) FindByClosePrice(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.ClosePrice, v))
}

// FindByOrderedVolume adds a new filter to the query that will require that
// the OrderedVolume property is equal to the passed value.
func (q *OrderQuery) FindByOrderedVolume(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.OrderedVolume, v))
}

// FindBySwappedVolume adds a new filter to the query that will require that
// the SwappedVolume property is equal to the passed value.
func (q *OrderQuery) FindBySwappedVolume(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.SwappedVolume, v))
}

// FindByIsClosed adds a new filter to the query that will require that
// the IsClosed property is equal to the passed value.
func (q *OrderQuery) FindByIsClosed(v bool) *OrderQuery {
	return q.Where(kallax.Eq(Schema.Order.IsClosed, v))
}

// FindByStopLoss adds a new filter to the query that will require that
// the StopLoss property is equal to the passed value.
func (q *OrderQuery) FindByStopLoss(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.StopLoss, v))
}

// FindByTakeProfit adds a new filter to the query that will require that
// the TakeProfit property is equal to the passed value.
func (q *OrderQuery) FindByTakeProfit(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.TakeProfit, v))
}

// FindByBuyFee adds a new filter to the query that will require that
// the BuyFee property is equal to the passed value.
func (q *OrderQuery) FindByBuyFee(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.BuyFee, v))
}

// FindBySellFee adds a new filter to the query that will require that
// the SellFee property is equal to the passed value.
func (q *OrderQuery) FindBySellFee(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.SellFee, v))
}

// FindByDelta adds a new filter to the query that will require that
// the Delta property is equal to the passed value.
func (q *OrderQuery) FindByDelta(cond kallax.ScalarCond, v float32) *OrderQuery {
	return q.Where(cond(Schema.Order.Delta, v))
}

// OrderResultSet is the set of results returned by a query to the
// database.
type OrderResultSet struct {
	ResultSet kallax.ResultSet
	last      *Order
	lastErr   error
}

// NewOrderResultSet creates a new result set for rows of the type
// Order.
func NewOrderResultSet(rs kallax.ResultSet) *OrderResultSet {
	return &OrderResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *OrderResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Order.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Order)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Order")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *OrderResultSet) Get() (*Order, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *OrderResultSet) ForEach(fn func(*Order) error) error {
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
func (rs *OrderResultSet) All() ([]*Order, error) {
	var result []*Order
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
func (rs *OrderResultSet) One() (*Order, error) {
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
func (rs *OrderResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *OrderResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewPair returns a new instance of Pair.
func NewPair() (record *Pair) {
	return new(Pair)
}

// GetID returns the primary key of the model.
func (r *Pair) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Pair) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "symbol":
		return &r.Symbol, nil
	case "base_id":
		return types.Nullable(kallax.VirtualColumn("base_id", r, new(kallax.NumericID))), nil
	case "quote_id":
		return types.Nullable(kallax.VirtualColumn("quote_id", r, new(kallax.NumericID))), nil

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
	case "Base":
		return new(Asset), nil
	case "Quote":
		return new(Asset), nil
	case "Markets":
		return new(Market), nil
	case "Differences":
		return new(Difference), nil
	case "Orders":
		return new(Order), nil
	case "Prices":
		return new(Price), nil
	case "Trades":
		return new(Trade), nil

	}
	return nil, fmt.Errorf("kallax: model Pair has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Pair) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Base":
		val, ok := rel.(*Asset)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Base", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Base = val
		}

		return nil
	case "Quote":
		val, ok := rel.(*Asset)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Quote", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Quote = val
		}

		return nil
	case "Markets":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Markets = make([]*Market, len(records))
		for i, record := range records {
			rel, ok := record.(*Market)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Markets[i] = rel
		}
		return nil
	case "Differences":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Differences = make([]*Difference, len(records))
		for i, record := range records {
			rel, ok := record.(*Difference)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Differences[i] = rel
		}
		return nil
	case "Orders":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Orders = make([]*Order, len(records))
		for i, record := range records {
			rel, ok := record.(*Order)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Orders[i] = rel
		}
		return nil
	case "Prices":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Prices = make([]*Price, len(records))
		for i, record := range records {
			rel, ok := record.(*Price)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Prices[i] = rel
		}
		return nil
	case "Trades":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Trades = make([]*Trade, len(records))
		for i, record := range records {
			rel, ok := record.(*Trade)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Trades[i] = rel
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

func (s *PairStore) relationshipRecords(record *Pair) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Markets {
		r := record.Markets[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("pair_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&MarketStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Differences {
		r := record.Differences[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("pair_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&DifferenceStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Orders {
		r := record.Orders[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("pair_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&OrderStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Prices {
		r := record.Prices[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("pair_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&PriceStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Trades {
		r := record.Trades[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("pair_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&TradeStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

func (s *PairStore) inverseRecords(record *Pair) []modelSaveFunc {
	var result []modelSaveFunc

	if record.Base != nil && !record.Base.IsSaving() {
		record.AddVirtualColumn("base_id", record.Base.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&AssetStore{store}).Save(record.Base)
			return err
		})
	}

	if record.Quote != nil && !record.Quote.IsSaving() {
		record.AddVirtualColumn("quote_id", record.Quote.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&AssetStore{store}).Save(record.Quote)
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

	records := s.relationshipRecords(record)

	inverseRecords := s.inverseRecords(record)

	if len(records) > 0 || len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Pair.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
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

	records := s.relationshipRecords(record)

	inverseRecords := s.inverseRecords(record)

	if len(records) > 0 || len(inverseRecords) > 0 {
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

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
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

// RemoveMarkets removes the given items of the Markets field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Markets` is not empty. This method clears the
// the elements of Markets in a model, it does not retrieve them to know
// what relationships the model has.
func (s *PairStore) RemoveMarkets(record *Pair, deleted ...*Market) error {
	var updated []*Market
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Markets
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Market.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Markets = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Market.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Market.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Markets {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Markets = updated
	return nil
}

// RemoveDifferences removes the given items of the Differences field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Differences` is not empty. This method clears the
// the elements of Differences in a model, it does not retrieve them to know
// what relationships the model has.
func (s *PairStore) RemoveDifferences(record *Pair, deleted ...*Difference) error {
	var updated []*Difference
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Differences
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Difference.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Differences = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Difference.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Difference.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Differences {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Differences = updated
	return nil
}

// RemoveOrders removes the given items of the Orders field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Orders` is not empty. This method clears the
// the elements of Orders in a model, it does not retrieve them to know
// what relationships the model has.
func (s *PairStore) RemoveOrders(record *Pair, deleted ...*Order) error {
	var updated []*Order
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Orders
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Order.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Orders = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Order.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Order.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Orders {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Orders = updated
	return nil
}

// RemovePrices removes the given items of the Prices field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Prices` is not empty. This method clears the
// the elements of Prices in a model, it does not retrieve them to know
// what relationships the model has.
func (s *PairStore) RemovePrices(record *Pair, deleted ...*Price) error {
	var updated []*Price
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Prices
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Price.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Prices = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Price.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Price.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Prices {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Prices = updated
	return nil
}

// RemoveTrades removes the given items of the Trades field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Trades` is not empty. This method clears the
// the elements of Trades in a model, it does not retrieve them to know
// what relationships the model has.
func (s *PairStore) RemoveTrades(record *Pair, deleted ...*Trade) error {
	var updated []*Trade
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Trades
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Trade.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Trades = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Trade.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Trade.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Trades {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Trades = updated
	return nil
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

func (q *PairQuery) WithBase() *PairQuery {
	q.AddRelation(Schema.Asset.BaseSchema, "Base", kallax.OneToOne, nil)
	return q
}

func (q *PairQuery) WithQuote() *PairQuery {
	q.AddRelation(Schema.Asset.BaseSchema, "Quote", kallax.OneToOne, nil)
	return q
}

func (q *PairQuery) WithMarkets(cond kallax.Condition) *PairQuery {
	q.AddRelation(Schema.Market.BaseSchema, "Markets", kallax.OneToMany, cond)
	return q
}

func (q *PairQuery) WithDifferences(cond kallax.Condition) *PairQuery {
	q.AddRelation(Schema.Difference.BaseSchema, "Differences", kallax.OneToMany, cond)
	return q
}

func (q *PairQuery) WithOrders(cond kallax.Condition) *PairQuery {
	q.AddRelation(Schema.Order.BaseSchema, "Orders", kallax.OneToMany, cond)
	return q
}

func (q *PairQuery) WithPrices(cond kallax.Condition) *PairQuery {
	q.AddRelation(Schema.Price.BaseSchema, "Prices", kallax.OneToMany, cond)
	return q
}

func (q *PairQuery) WithTrades(cond kallax.Condition) *PairQuery {
	q.AddRelation(Schema.Trade.BaseSchema, "Trades", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PairQuery) FindByID(v ...kallax.NumericID) *PairQuery {
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

// FindByBase adds a new filter to the query that will require that
// the foreign key of Base is equal to the passed value.
func (q *PairQuery) FindByBase(v kallax.NumericID) *PairQuery {
	return q.Where(kallax.Eq(Schema.Pair.BaseFK, v))
}

// FindByQuote adds a new filter to the query that will require that
// the foreign key of Quote is equal to the passed value.
func (q *PairQuery) FindByQuote(v kallax.NumericID) *PairQuery {
	return q.Where(kallax.Eq(Schema.Pair.QuoteFK, v))
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
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Price) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "pair_id":
		return types.Nullable(kallax.VirtualColumn("pair_id", r, new(kallax.NumericID))), nil
	case "exchange_id":
		return types.Nullable(kallax.VirtualColumn("exchange_id", r, new(kallax.NumericID))), nil
	case "market_id":
		return types.Nullable(kallax.VirtualColumn("market_id", r, new(kallax.NumericID))), nil
	case "price":
		return &r.Price, nil
	case "pair_symbols":
		return types.Slice(&r.PairSymbols), nil
	case "exchange_symbols":
		return types.Slice(&r.ExchangeSymbols), nil

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
	case "market_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "price":
		return r.Price, nil
	case "pair_symbols":
		return types.Slice(r.PairSymbols), nil
	case "exchange_symbols":
		return types.Slice(r.ExchangeSymbols), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Price: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Price) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Pair":
		return new(Pair), nil
	case "Exchange":
		return new(Exchange), nil
	case "Market":
		return new(Market), nil

	}
	return nil, fmt.Errorf("kallax: model Price has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Price) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Pair":
		val, ok := rel.(*Pair)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Pair", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Pair = val
		}

		return nil
	case "Exchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Exchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Exchange = val
		}

		return nil
	case "Market":
		val, ok := rel.(*Market)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Market", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Market = val
		}

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

	if record.Pair != nil && !record.Pair.IsSaving() {
		record.AddVirtualColumn("pair_id", record.Pair.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&PairStore{store}).Save(record.Pair)
			return err
		})
	}

	if record.Exchange != nil && !record.Exchange.IsSaving() {
		record.AddVirtualColumn("exchange_id", record.Exchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.Exchange)
			return err
		})
	}

	if record.Market != nil && !record.Market.IsSaving() {
		record.AddVirtualColumn("market_id", record.Market.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&MarketStore{store}).Save(record.Market)
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

func (q *PriceQuery) WithPair() *PriceQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "Pair", kallax.OneToOne, nil)
	return q
}

func (q *PriceQuery) WithExchange() *PriceQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "Exchange", kallax.OneToOne, nil)
	return q
}

func (q *PriceQuery) WithMarket() *PriceQuery {
	q.AddRelation(Schema.Market.BaseSchema, "Market", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PriceQuery) FindByID(v ...kallax.NumericID) *PriceQuery {
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

// FindByPair adds a new filter to the query that will require that
// the foreign key of Pair is equal to the passed value.
func (q *PriceQuery) FindByPair(v kallax.NumericID) *PriceQuery {
	return q.Where(kallax.Eq(Schema.Price.PairFK, v))
}

// FindByExchange adds a new filter to the query that will require that
// the foreign key of Exchange is equal to the passed value.
func (q *PriceQuery) FindByExchange(v kallax.NumericID) *PriceQuery {
	return q.Where(kallax.Eq(Schema.Price.ExchangeFK, v))
}

// FindByMarket adds a new filter to the query that will require that
// the foreign key of Market is equal to the passed value.
func (q *PriceQuery) FindByMarket(v kallax.NumericID) *PriceQuery {
	return q.Where(kallax.Eq(Schema.Price.MarketFK, v))
}

// FindByPrice adds a new filter to the query that will require that
// the Price property is equal to the passed value.
func (q *PriceQuery) FindByPrice(cond kallax.ScalarCond, v float32) *PriceQuery {
	return q.Where(cond(Schema.Price.Price, v))
}

// FindByPairSymbols adds a new filter to the query that will require that
// the PairSymbols property contains all the passed values; if no passed values,
// it will do nothing.
func (q *PriceQuery) FindByPairSymbols(v ...string) *PriceQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.Price.PairSymbols, values...))
}

// FindByExchangeSymbols adds a new filter to the query that will require that
// the ExchangeSymbols property contains all the passed values; if no passed values,
// it will do nothing.
func (q *PriceQuery) FindByExchangeSymbols(v ...string) *PriceQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.Price.ExchangeSymbols, values...))
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

// NewTrade returns a new instance of Trade.
func NewTrade() (record *Trade) {
	return new(Trade)
}

// GetID returns the primary key of the model.
func (r *Trade) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Trade) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "user_id":
		return types.Nullable(kallax.VirtualColumn("user_id", r, new(kallax.NumericID))), nil
	case "exchange_id":
		return types.Nullable(kallax.VirtualColumn("exchange_id", r, new(kallax.NumericID))), nil
	case "pair_id":
		return types.Nullable(kallax.VirtualColumn("pair_id", r, new(kallax.NumericID))), nil
	case "market_id":
		return types.Nullable(kallax.VirtualColumn("market_id", r, new(kallax.NumericID))), nil
	case "order_id":
		return types.Nullable(kallax.VirtualColumn("order_id", r, new(kallax.NumericID))), nil
	case "type":
		return &r.Type, nil
	case "volume":
		return &r.Volume, nil
	case "price":
		return &r.Price, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Trade: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Trade) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "user_id":
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
	case "pair_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "market_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "order_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "type":
		return r.Type, nil
	case "volume":
		return r.Volume, nil
	case "price":
		return r.Price, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Trade: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Trade) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "User":
		return new(User), nil
	case "Exchange":
		return new(Exchange), nil
	case "Pair":
		return new(Pair), nil
	case "Market":
		return new(Market), nil
	case "Order":
		return new(Order), nil

	}
	return nil, fmt.Errorf("kallax: model Trade has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Trade) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "User":
		val, ok := rel.(*User)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship User", rel)
		}
		if !val.GetID().IsEmpty() {
			r.User = val
		}

		return nil
	case "Exchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Exchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Exchange = val
		}

		return nil
	case "Pair":
		val, ok := rel.(*Pair)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Pair", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Pair = val
		}

		return nil
	case "Market":
		val, ok := rel.(*Market)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Market", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Market = val
		}

		return nil
	case "Order":
		val, ok := rel.(*Order)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Order", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Order = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model Trade has no relationship %s", field)
}

// TradeStore is the entity to access the records of the type Trade
// in the database.
type TradeStore struct {
	*kallax.Store
}

// NewTradeStore creates a new instance of TradeStore
// using a SQL database.
func NewTradeStore(db *sql.DB) *TradeStore {
	return &TradeStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *TradeStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *TradeStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *TradeStore) Debug() *TradeStore {
	return &TradeStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *TradeStore) DebugWith(logger kallax.LoggerFunc) *TradeStore {
	return &TradeStore{s.Store.DebugWith(logger)}
}

func (s *TradeStore) inverseRecords(record *Trade) []modelSaveFunc {
	var result []modelSaveFunc

	if record.User != nil && !record.User.IsSaving() {
		record.AddVirtualColumn("user_id", record.User.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&UserStore{store}).Save(record.User)
			return err
		})
	}

	if record.Exchange != nil && !record.Exchange.IsSaving() {
		record.AddVirtualColumn("exchange_id", record.Exchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.Exchange)
			return err
		})
	}

	if record.Pair != nil && !record.Pair.IsSaving() {
		record.AddVirtualColumn("pair_id", record.Pair.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&PairStore{store}).Save(record.Pair)
			return err
		})
	}

	if record.Market != nil && !record.Market.IsSaving() {
		record.AddVirtualColumn("market_id", record.Market.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&MarketStore{store}).Save(record.Market)
			return err
		})
	}

	if record.Order != nil && !record.Order.IsSaving() {
		record.AddVirtualColumn("order_id", record.Order.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&OrderStore{store}).Save(record.Order)
			return err
		})
	}

	return result
}

// Insert inserts a Trade in the database. A non-persisted object is
// required for this operation.
func (s *TradeStore) Insert(record *Trade) error {
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

			if err := s.Insert(Schema.Trade.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Trade.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *TradeStore) Update(record *Trade, cols ...kallax.SchemaField) (updated int64, err error) {
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

			updated, err = s.Update(Schema.Trade.BaseSchema, record, cols...)
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

	return s.Store.Update(Schema.Trade.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *TradeStore) Save(record *Trade) (updated bool, err error) {
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
func (s *TradeStore) Delete(record *Trade) error {
	return s.Store.Delete(Schema.Trade.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *TradeStore) Find(q *TradeQuery) (*TradeResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewTradeResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *TradeStore) MustFind(q *TradeQuery) *TradeResultSet {
	return NewTradeResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *TradeStore) Count(q *TradeQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *TradeStore) MustCount(q *TradeQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *TradeStore) FindOne(q *TradeQuery) (*Trade, error) {
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
func (s *TradeStore) FindAll(q *TradeQuery) ([]*Trade, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *TradeStore) MustFindOne(q *TradeQuery) *Trade {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Trade with the data in the database and
// makes it writable.
func (s *TradeStore) Reload(record *Trade) error {
	return s.Store.Reload(Schema.Trade.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *TradeStore) Transaction(callback func(*TradeStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&TradeStore{store})
	})
}

// TradeQuery is the object used to create queries for the Trade
// entity.
type TradeQuery struct {
	*kallax.BaseQuery
}

// NewTradeQuery returns a new instance of TradeQuery.
func NewTradeQuery() *TradeQuery {
	return &TradeQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Trade.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *TradeQuery) Select(columns ...kallax.SchemaField) *TradeQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *TradeQuery) SelectNot(columns ...kallax.SchemaField) *TradeQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *TradeQuery) Copy() *TradeQuery {
	return &TradeQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *TradeQuery) Order(cols ...kallax.ColumnOrder) *TradeQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *TradeQuery) BatchSize(size uint64) *TradeQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *TradeQuery) Limit(n uint64) *TradeQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *TradeQuery) Offset(n uint64) *TradeQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *TradeQuery) Where(cond kallax.Condition) *TradeQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *TradeQuery) WithUser() *TradeQuery {
	q.AddRelation(Schema.User.BaseSchema, "User", kallax.OneToOne, nil)
	return q
}

func (q *TradeQuery) WithExchange() *TradeQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "Exchange", kallax.OneToOne, nil)
	return q
}

func (q *TradeQuery) WithPair() *TradeQuery {
	q.AddRelation(Schema.Pair.BaseSchema, "Pair", kallax.OneToOne, nil)
	return q
}

func (q *TradeQuery) WithMarket() *TradeQuery {
	q.AddRelation(Schema.Market.BaseSchema, "Market", kallax.OneToOne, nil)
	return q
}

func (q *TradeQuery) WithOrder() *TradeQuery {
	q.AddRelation(Schema.Order.BaseSchema, "Order", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *TradeQuery) FindByID(v ...kallax.NumericID) *TradeQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Trade.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *TradeQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *TradeQuery {
	return q.Where(cond(Schema.Trade.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *TradeQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *TradeQuery {
	return q.Where(cond(Schema.Trade.UpdatedAt, v))
}

// FindByUser adds a new filter to the query that will require that
// the foreign key of User is equal to the passed value.
func (q *TradeQuery) FindByUser(v kallax.NumericID) *TradeQuery {
	return q.Where(kallax.Eq(Schema.Trade.UserFK, v))
}

// FindByExchange adds a new filter to the query that will require that
// the foreign key of Exchange is equal to the passed value.
func (q *TradeQuery) FindByExchange(v kallax.NumericID) *TradeQuery {
	return q.Where(kallax.Eq(Schema.Trade.ExchangeFK, v))
}

// FindByPair adds a new filter to the query that will require that
// the foreign key of Pair is equal to the passed value.
func (q *TradeQuery) FindByPair(v kallax.NumericID) *TradeQuery {
	return q.Where(kallax.Eq(Schema.Trade.PairFK, v))
}

// FindByMarket adds a new filter to the query that will require that
// the foreign key of Market is equal to the passed value.
func (q *TradeQuery) FindByMarket(v kallax.NumericID) *TradeQuery {
	return q.Where(kallax.Eq(Schema.Trade.MarketFK, v))
}

// FindByOrder adds a new filter to the query that will require that
// the foreign key of Order is equal to the passed value.
func (q *TradeQuery) FindByOrder(v kallax.NumericID) *TradeQuery {
	return q.Where(kallax.Eq(Schema.Trade.OrderFK, v))
}

// FindByType adds a new filter to the query that will require that
// the Type property is equal to the passed value.
func (q *TradeQuery) FindByType(v string) *TradeQuery {
	return q.Where(kallax.Eq(Schema.Trade.Type, v))
}

// FindByVolume adds a new filter to the query that will require that
// the Volume property is equal to the passed value.
func (q *TradeQuery) FindByVolume(cond kallax.ScalarCond, v float32) *TradeQuery {
	return q.Where(cond(Schema.Trade.Volume, v))
}

// FindByPrice adds a new filter to the query that will require that
// the Price property is equal to the passed value.
func (q *TradeQuery) FindByPrice(cond kallax.ScalarCond, v float32) *TradeQuery {
	return q.Where(cond(Schema.Trade.Price, v))
}

// TradeResultSet is the set of results returned by a query to the
// database.
type TradeResultSet struct {
	ResultSet kallax.ResultSet
	last      *Trade
	lastErr   error
}

// NewTradeResultSet creates a new result set for rows of the type
// Trade.
func NewTradeResultSet(rs kallax.ResultSet) *TradeResultSet {
	return &TradeResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *TradeResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Trade.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Trade)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Trade")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *TradeResultSet) Get() (*Trade, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *TradeResultSet) ForEach(fn func(*Trade) error) error {
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
func (rs *TradeResultSet) All() ([]*Trade, error) {
	var result []*Trade
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
func (rs *TradeResultSet) One() (*Trade, error) {
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
func (rs *TradeResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *TradeResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewUser returns a new instance of User.
func NewUser() (record *User) {
	return new(User)
}

// GetID returns the primary key of the model.
func (r *User) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *User) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "email":
		return &r.Email, nil
	case "password":
		return &r.Password, nil
	case "last_login":
		return &r.LastLogin, nil
	case "subscribe_to":
		return &r.SubscribeTo, nil
	case "role":
		return &r.Role, nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// Value returns the value of the given column.
func (r *User) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "email":
		return r.Email, nil
	case "password":
		return r.Password, nil
	case "last_login":
		return r.LastLogin, nil
	case "subscribe_to":
		return r.SubscribeTo, nil
	case "role":
		return r.Role, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *User) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Balances":
		return new(UserBalance), nil
	case "Orders":
		return new(Order), nil
	case "Trades":
		return new(Trade), nil

	}
	return nil, fmt.Errorf("kallax: model User has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *User) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Balances":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Balances = make([]*UserBalance, len(records))
		for i, record := range records {
			rel, ok := record.(*UserBalance)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Balances[i] = rel
		}
		return nil
	case "Orders":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Orders = make([]*Order, len(records))
		for i, record := range records {
			rel, ok := record.(*Order)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Orders[i] = rel
		}
		return nil
	case "Trades":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Trades = make([]*Trade, len(records))
		for i, record := range records {
			rel, ok := record.(*Trade)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Trades[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model User has no relationship %s", field)
}

// UserStore is the entity to access the records of the type User
// in the database.
type UserStore struct {
	*kallax.Store
}

// NewUserStore creates a new instance of UserStore
// using a SQL database.
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *UserStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *UserStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *UserStore) Debug() *UserStore {
	return &UserStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *UserStore) DebugWith(logger kallax.LoggerFunc) *UserStore {
	return &UserStore{s.Store.DebugWith(logger)}
}

func (s *UserStore) relationshipRecords(record *User) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Balances {
		r := record.Balances[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("user_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&UserBalanceStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Orders {
		r := record.Orders[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("user_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&OrderStore{store}).Save(r)
				return err
			})
		}
	}

	for i := range record.Trades {
		r := record.Trades[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("user_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&TradeStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

// Insert inserts a User in the database. A non-persisted object is
// required for this operation.
func (s *UserStore) Insert(record *User) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.LastLogin = record.LastLogin.Truncate(time.Microsecond)
	record.SubscribeTo = record.SubscribeTo.Truncate(time.Microsecond)
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			if err := s.Insert(Schema.User.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.User.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *UserStore) Update(record *User, cols ...kallax.SchemaField) (updated int64, err error) {
	record.LastLogin = record.LastLogin.Truncate(time.Microsecond)
	record.SubscribeTo = record.SubscribeTo.Truncate(time.Microsecond)
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			updated, err = s.Update(Schema.User.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.User.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *UserStore) Save(record *User) (updated bool, err error) {
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
func (s *UserStore) Delete(record *User) error {
	return s.Store.Delete(Schema.User.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *UserStore) Find(q *UserQuery) (*UserResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewUserResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *UserStore) MustFind(q *UserQuery) *UserResultSet {
	return NewUserResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *UserStore) Count(q *UserQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *UserStore) MustCount(q *UserQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *UserStore) FindOne(q *UserQuery) (*User, error) {
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
func (s *UserStore) FindAll(q *UserQuery) ([]*User, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *UserStore) MustFindOne(q *UserQuery) *User {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the User with the data in the database and
// makes it writable.
func (s *UserStore) Reload(record *User) error {
	return s.Store.Reload(Schema.User.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *UserStore) Transaction(callback func(*UserStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&UserStore{store})
	})
}

// RemoveBalances removes the given items of the Balances field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Balances` is not empty. This method clears the
// the elements of Balances in a model, it does not retrieve them to know
// what relationships the model has.
func (s *UserStore) RemoveBalances(record *User, deleted ...*UserBalance) error {
	var updated []*UserBalance
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Balances
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.UserBalance.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Balances = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.UserBalance.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.UserBalance.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Balances {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Balances = updated
	return nil
}

// RemoveOrders removes the given items of the Orders field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Orders` is not empty. This method clears the
// the elements of Orders in a model, it does not retrieve them to know
// what relationships the model has.
func (s *UserStore) RemoveOrders(record *User, deleted ...*Order) error {
	var updated []*Order
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Orders
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Order.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Orders = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Order.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Order.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Orders {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Orders = updated
	return nil
}

// RemoveTrades removes the given items of the Trades field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Trades` is not empty. This method clears the
// the elements of Trades in a model, it does not retrieve them to know
// what relationships the model has.
func (s *UserStore) RemoveTrades(record *User, deleted ...*Trade) error {
	var updated []*Trade
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Trades
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Trade.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Trades = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Trade.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Trade.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Trades {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Trades = updated
	return nil
}

// UserQuery is the object used to create queries for the User
// entity.
type UserQuery struct {
	*kallax.BaseQuery
}

// NewUserQuery returns a new instance of UserQuery.
func NewUserQuery() *UserQuery {
	return &UserQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.User.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *UserQuery) Select(columns ...kallax.SchemaField) *UserQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *UserQuery) SelectNot(columns ...kallax.SchemaField) *UserQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *UserQuery) Copy() *UserQuery {
	return &UserQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *UserQuery) Order(cols ...kallax.ColumnOrder) *UserQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *UserQuery) BatchSize(size uint64) *UserQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *UserQuery) Limit(n uint64) *UserQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *UserQuery) Offset(n uint64) *UserQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *UserQuery) Where(cond kallax.Condition) *UserQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *UserQuery) WithBalances(cond kallax.Condition) *UserQuery {
	q.AddRelation(Schema.UserBalance.BaseSchema, "Balances", kallax.OneToMany, cond)
	return q
}

func (q *UserQuery) WithOrders(cond kallax.Condition) *UserQuery {
	q.AddRelation(Schema.Order.BaseSchema, "Orders", kallax.OneToMany, cond)
	return q
}

func (q *UserQuery) WithTrades(cond kallax.Condition) *UserQuery {
	q.AddRelation(Schema.Trade.BaseSchema, "Trades", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *UserQuery) FindByID(v ...kallax.NumericID) *UserQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.User.ID, values...))
}

// FindByEmail adds a new filter to the query that will require that
// the Email property is equal to the passed value.
func (q *UserQuery) FindByEmail(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Email, v))
}

// FindByPassword adds a new filter to the query that will require that
// the Password property is equal to the passed value.
func (q *UserQuery) FindByPassword(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Password, v))
}

// FindByLastLogin adds a new filter to the query that will require that
// the LastLogin property is equal to the passed value.
func (q *UserQuery) FindByLastLogin(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.LastLogin, v))
}

// FindBySubscribeTo adds a new filter to the query that will require that
// the SubscribeTo property is equal to the passed value.
func (q *UserQuery) FindBySubscribeTo(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.SubscribeTo, v))
}

// FindByRole adds a new filter to the query that will require that
// the Role property is equal to the passed value.
func (q *UserQuery) FindByRole(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Role, v))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *UserQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *UserQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.UpdatedAt, v))
}

// UserResultSet is the set of results returned by a query to the
// database.
type UserResultSet struct {
	ResultSet kallax.ResultSet
	last      *User
	lastErr   error
}

// NewUserResultSet creates a new result set for rows of the type
// User.
func NewUserResultSet(rs kallax.ResultSet) *UserResultSet {
	return &UserResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *UserResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.User.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*User)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *User")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *UserResultSet) Get() (*User, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *UserResultSet) ForEach(fn func(*User) error) error {
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
func (rs *UserResultSet) All() ([]*User, error) {
	var result []*User
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
func (rs *UserResultSet) One() (*User, error) {
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
func (rs *UserResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *UserResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewUserBalance returns a new instance of UserBalance.
func NewUserBalance() (record *UserBalance) {
	return new(UserBalance)
}

// GetID returns the primary key of the model.
func (r *UserBalance) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *UserBalance) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "user_id":
		return types.Nullable(kallax.VirtualColumn("user_id", r, new(kallax.NumericID))), nil
	case "exchange_id":
		return types.Nullable(kallax.VirtualColumn("exchange_id", r, new(kallax.NumericID))), nil
	case "asset_id":
		return types.Nullable(kallax.VirtualColumn("asset_id", r, new(kallax.NumericID))), nil
	case "volume":
		return &r.Volume, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in UserBalance: %s", col)
	}
}

// Value returns the value of the given column.
func (r *UserBalance) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "user_id":
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
	case "asset_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "volume":
		return r.Volume, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in UserBalance: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *UserBalance) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "User":
		return new(User), nil
	case "Exchange":
		return new(Exchange), nil
	case "Asset":
		return new(Asset), nil

	}
	return nil, fmt.Errorf("kallax: model UserBalance has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *UserBalance) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "User":
		val, ok := rel.(*User)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship User", rel)
		}
		if !val.GetID().IsEmpty() {
			r.User = val
		}

		return nil
	case "Exchange":
		val, ok := rel.(*Exchange)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Exchange", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Exchange = val
		}

		return nil
	case "Asset":
		val, ok := rel.(*Asset)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Asset", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Asset = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model UserBalance has no relationship %s", field)
}

// UserBalanceStore is the entity to access the records of the type UserBalance
// in the database.
type UserBalanceStore struct {
	*kallax.Store
}

// NewUserBalanceStore creates a new instance of UserBalanceStore
// using a SQL database.
func NewUserBalanceStore(db *sql.DB) *UserBalanceStore {
	return &UserBalanceStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *UserBalanceStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *UserBalanceStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *UserBalanceStore) Debug() *UserBalanceStore {
	return &UserBalanceStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *UserBalanceStore) DebugWith(logger kallax.LoggerFunc) *UserBalanceStore {
	return &UserBalanceStore{s.Store.DebugWith(logger)}
}

func (s *UserBalanceStore) inverseRecords(record *UserBalance) []modelSaveFunc {
	var result []modelSaveFunc

	if record.User != nil && !record.User.IsSaving() {
		record.AddVirtualColumn("user_id", record.User.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&UserStore{store}).Save(record.User)
			return err
		})
	}

	if record.Exchange != nil && !record.Exchange.IsSaving() {
		record.AddVirtualColumn("exchange_id", record.Exchange.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&ExchangeStore{store}).Save(record.Exchange)
			return err
		})
	}

	if record.Asset != nil && !record.Asset.IsSaving() {
		record.AddVirtualColumn("asset_id", record.Asset.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&AssetStore{store}).Save(record.Asset)
			return err
		})
	}

	return result
}

// Insert inserts a UserBalance in the database. A non-persisted object is
// required for this operation.
func (s *UserBalanceStore) Insert(record *UserBalance) error {
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

			if err := s.Insert(Schema.UserBalance.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.UserBalance.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *UserBalanceStore) Update(record *UserBalance, cols ...kallax.SchemaField) (updated int64, err error) {
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

			updated, err = s.Update(Schema.UserBalance.BaseSchema, record, cols...)
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

	return s.Store.Update(Schema.UserBalance.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *UserBalanceStore) Save(record *UserBalance) (updated bool, err error) {
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
func (s *UserBalanceStore) Delete(record *UserBalance) error {
	return s.Store.Delete(Schema.UserBalance.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *UserBalanceStore) Find(q *UserBalanceQuery) (*UserBalanceResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewUserBalanceResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *UserBalanceStore) MustFind(q *UserBalanceQuery) *UserBalanceResultSet {
	return NewUserBalanceResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *UserBalanceStore) Count(q *UserBalanceQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *UserBalanceStore) MustCount(q *UserBalanceQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *UserBalanceStore) FindOne(q *UserBalanceQuery) (*UserBalance, error) {
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
func (s *UserBalanceStore) FindAll(q *UserBalanceQuery) ([]*UserBalance, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *UserBalanceStore) MustFindOne(q *UserBalanceQuery) *UserBalance {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the UserBalance with the data in the database and
// makes it writable.
func (s *UserBalanceStore) Reload(record *UserBalance) error {
	return s.Store.Reload(Schema.UserBalance.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *UserBalanceStore) Transaction(callback func(*UserBalanceStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&UserBalanceStore{store})
	})
}

// UserBalanceQuery is the object used to create queries for the UserBalance
// entity.
type UserBalanceQuery struct {
	*kallax.BaseQuery
}

// NewUserBalanceQuery returns a new instance of UserBalanceQuery.
func NewUserBalanceQuery() *UserBalanceQuery {
	return &UserBalanceQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.UserBalance.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *UserBalanceQuery) Select(columns ...kallax.SchemaField) *UserBalanceQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *UserBalanceQuery) SelectNot(columns ...kallax.SchemaField) *UserBalanceQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *UserBalanceQuery) Copy() *UserBalanceQuery {
	return &UserBalanceQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *UserBalanceQuery) Order(cols ...kallax.ColumnOrder) *UserBalanceQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *UserBalanceQuery) BatchSize(size uint64) *UserBalanceQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *UserBalanceQuery) Limit(n uint64) *UserBalanceQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *UserBalanceQuery) Offset(n uint64) *UserBalanceQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *UserBalanceQuery) Where(cond kallax.Condition) *UserBalanceQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *UserBalanceQuery) WithUser() *UserBalanceQuery {
	q.AddRelation(Schema.User.BaseSchema, "User", kallax.OneToOne, nil)
	return q
}

func (q *UserBalanceQuery) WithExchange() *UserBalanceQuery {
	q.AddRelation(Schema.Exchange.BaseSchema, "Exchange", kallax.OneToOne, nil)
	return q
}

func (q *UserBalanceQuery) WithAsset() *UserBalanceQuery {
	q.AddRelation(Schema.Asset.BaseSchema, "Asset", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *UserBalanceQuery) FindByID(v ...kallax.NumericID) *UserBalanceQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.UserBalance.ID, values...))
}

// FindByUser adds a new filter to the query that will require that
// the foreign key of User is equal to the passed value.
func (q *UserBalanceQuery) FindByUser(v kallax.NumericID) *UserBalanceQuery {
	return q.Where(kallax.Eq(Schema.UserBalance.UserFK, v))
}

// FindByExchange adds a new filter to the query that will require that
// the foreign key of Exchange is equal to the passed value.
func (q *UserBalanceQuery) FindByExchange(v kallax.NumericID) *UserBalanceQuery {
	return q.Where(kallax.Eq(Schema.UserBalance.ExchangeFK, v))
}

// FindByAsset adds a new filter to the query that will require that
// the foreign key of Asset is equal to the passed value.
func (q *UserBalanceQuery) FindByAsset(v kallax.NumericID) *UserBalanceQuery {
	return q.Where(kallax.Eq(Schema.UserBalance.AssetFK, v))
}

// FindByVolume adds a new filter to the query that will require that
// the Volume property is equal to the passed value.
func (q *UserBalanceQuery) FindByVolume(cond kallax.ScalarCond, v float32) *UserBalanceQuery {
	return q.Where(cond(Schema.UserBalance.Volume, v))
}

// UserBalanceResultSet is the set of results returned by a query to the
// database.
type UserBalanceResultSet struct {
	ResultSet kallax.ResultSet
	last      *UserBalance
	lastErr   error
}

// NewUserBalanceResultSet creates a new result set for rows of the type
// UserBalance.
func NewUserBalanceResultSet(rs kallax.ResultSet) *UserBalanceResultSet {
	return &UserBalanceResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *UserBalanceResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.UserBalance.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*UserBalance)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *UserBalance")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *UserBalanceResultSet) Get() (*UserBalance, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *UserBalanceResultSet) ForEach(fn func(*UserBalance) error) error {
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
func (rs *UserBalanceResultSet) All() ([]*UserBalance, error) {
	var result []*UserBalance
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
func (rs *UserBalanceResultSet) One() (*UserBalance, error) {
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
func (rs *UserBalanceResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *UserBalanceResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Asset         *schemaAsset
	Difference    *schemaDifference
	Exchange      *schemaExchange
	ExchangeAsset *schemaExchangeAsset
	Market        *schemaMarket
	Order         *schemaOrder
	Pair          *schemaPair
	Price         *schemaPrice
	Trade         *schemaTrade
	User          *schemaUser
	UserBalance   *schemaUserBalance
}

type schemaAsset struct {
	*kallax.BaseSchema
	ID     kallax.SchemaField
	Symbol kallax.SchemaField
	Name   kallax.SchemaField
	IsFiat kallax.SchemaField
}

type schemaDifference struct {
	*kallax.BaseSchema
	ID              kallax.SchemaField
	CreatedAt       kallax.SchemaField
	UpdatedAt       kallax.SchemaField
	PairFK          kallax.SchemaField
	BaseExchangeFK  kallax.SchemaField
	QuoteExchangeFK kallax.SchemaField
	Delta           kallax.SchemaField
}

type schemaExchange struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	Symbol    kallax.SchemaField
	Name      kallax.SchemaField
	IsActive  kallax.SchemaField
	IsUsedAPI kallax.SchemaField
}

type schemaExchangeAsset struct {
	*kallax.BaseSchema
	ID             kallax.SchemaField
	AssetFK        kallax.SchemaField
	ExchangeFK     kallax.SchemaField
	TransactionFee kallax.SchemaField
}

type schemaMarket struct {
	*kallax.BaseSchema
	ID         kallax.SchemaField
	PairFK     kallax.SchemaField
	ExchangeFK kallax.SchemaField
	IsActive   kallax.SchemaField
}

type schemaOrder struct {
	*kallax.BaseSchema
	ID            kallax.SchemaField
	CreatedAt     kallax.SchemaField
	UpdatedAt     kallax.SchemaField
	UserFK        kallax.SchemaField
	ExchangeFK    kallax.SchemaField
	PairFK        kallax.SchemaField
	MarketFK      kallax.SchemaField
	OrderType     kallax.SchemaField
	OpenPrice     kallax.SchemaField
	ClosePrice    kallax.SchemaField
	OrderedVolume kallax.SchemaField
	SwappedVolume kallax.SchemaField
	IsClosed      kallax.SchemaField
	StopLoss      kallax.SchemaField
	TakeProfit    kallax.SchemaField
	BuyFee        kallax.SchemaField
	SellFee       kallax.SchemaField
	Delta         kallax.SchemaField
}

type schemaPair struct {
	*kallax.BaseSchema
	ID      kallax.SchemaField
	Symbol  kallax.SchemaField
	BaseFK  kallax.SchemaField
	QuoteFK kallax.SchemaField
}

type schemaPrice struct {
	*kallax.BaseSchema
	ID              kallax.SchemaField
	CreatedAt       kallax.SchemaField
	UpdatedAt       kallax.SchemaField
	PairFK          kallax.SchemaField
	ExchangeFK      kallax.SchemaField
	MarketFK        kallax.SchemaField
	Price           kallax.SchemaField
	PairSymbols     kallax.SchemaField
	ExchangeSymbols kallax.SchemaField
}

type schemaTrade struct {
	*kallax.BaseSchema
	ID         kallax.SchemaField
	CreatedAt  kallax.SchemaField
	UpdatedAt  kallax.SchemaField
	UserFK     kallax.SchemaField
	ExchangeFK kallax.SchemaField
	PairFK     kallax.SchemaField
	MarketFK   kallax.SchemaField
	OrderFK    kallax.SchemaField
	Type       kallax.SchemaField
	Volume     kallax.SchemaField
	Price      kallax.SchemaField
}

type schemaUser struct {
	*kallax.BaseSchema
	ID          kallax.SchemaField
	Email       kallax.SchemaField
	Password    kallax.SchemaField
	LastLogin   kallax.SchemaField
	SubscribeTo kallax.SchemaField
	Role        kallax.SchemaField
	CreatedAt   kallax.SchemaField
	UpdatedAt   kallax.SchemaField
}

type schemaUserBalance struct {
	*kallax.BaseSchema
	ID         kallax.SchemaField
	UserFK     kallax.SchemaField
	ExchangeFK kallax.SchemaField
	AssetFK    kallax.SchemaField
	Volume     kallax.SchemaField
}

var Schema = &schema{
	Asset: &schemaAsset{
		BaseSchema: kallax.NewBaseSchema(
			"assets",
			"__asset",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"BasePairs":  kallax.NewForeignKey("base_id", false),
				"QuotePairs": kallax.NewForeignKey("quote_id", false),
				"Balances":   kallax.NewForeignKey("asset_id", false),
			},
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
	Difference: &schemaDifference{
		BaseSchema: kallax.NewBaseSchema(
			"differences",
			"__difference",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Pair":          kallax.NewForeignKey("pair_id", true),
				"BaseExchange":  kallax.NewForeignKey("base_id", true),
				"QuoteExchange": kallax.NewForeignKey("quote_id", true),
			},
			func() kallax.Record {
				return new(Difference)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("pair_id"),
			kallax.NewSchemaField("base_id"),
			kallax.NewSchemaField("quote_id"),
			kallax.NewSchemaField("delta"),
		),
		ID:              kallax.NewSchemaField("id"),
		CreatedAt:       kallax.NewSchemaField("created_at"),
		UpdatedAt:       kallax.NewSchemaField("updated_at"),
		PairFK:          kallax.NewSchemaField("pair_id"),
		BaseExchangeFK:  kallax.NewSchemaField("base_id"),
		QuoteExchangeFK: kallax.NewSchemaField("quote_id"),
		Delta:           kallax.NewSchemaField("delta"),
	},
	Exchange: &schemaExchange{
		BaseSchema: kallax.NewBaseSchema(
			"exchanges",
			"__exchange",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Markets":          kallax.NewForeignKey("exchange_id", false),
				"Assets":           kallax.NewForeignKey("exchange_id", false),
				"BaseDifferences":  kallax.NewForeignKey("base_id", false),
				"QuoteDifferences": kallax.NewForeignKey("quote_id", false),
				"Orders":           kallax.NewForeignKey("exchange_id", false),
				"Prices":           kallax.NewForeignKey("exchange_id", false),
				"Trades":           kallax.NewForeignKey("exchange_id", false),
				"Balances":         kallax.NewForeignKey("exchange_id", false),
			},
			func() kallax.Record {
				return new(Exchange)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("symbol"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("is_active"),
			kallax.NewSchemaField("is_used_api"),
		),
		ID:        kallax.NewSchemaField("id"),
		Symbol:    kallax.NewSchemaField("symbol"),
		Name:      kallax.NewSchemaField("name"),
		IsActive:  kallax.NewSchemaField("is_active"),
		IsUsedAPI: kallax.NewSchemaField("is_used_api"),
	},
	ExchangeAsset: &schemaExchangeAsset{
		BaseSchema: kallax.NewBaseSchema(
			"exchange_assets",
			"__exchangeasset",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Asset":    kallax.NewForeignKey("asset_id", true),
				"Exchange": kallax.NewForeignKey("exchange_id", true),
			},
			func() kallax.Record {
				return new(ExchangeAsset)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("asset_id"),
			kallax.NewSchemaField("exchange_id"),
			kallax.NewSchemaField("transaction_fee"),
		),
		ID:             kallax.NewSchemaField("id"),
		AssetFK:        kallax.NewSchemaField("asset_id"),
		ExchangeFK:     kallax.NewSchemaField("exchange_id"),
		TransactionFee: kallax.NewSchemaField("transaction_fee"),
	},
	Market: &schemaMarket{
		BaseSchema: kallax.NewBaseSchema(
			"markets",
			"__market",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Pair":     kallax.NewForeignKey("pair_id", true),
				"Exchange": kallax.NewForeignKey("exchange_id", true),
				"Orders":   kallax.NewForeignKey("market_id", false),
				"Prices":   kallax.NewForeignKey("market_id", false),
				"Trades":   kallax.NewForeignKey("market_id", false),
			},
			func() kallax.Record {
				return new(Market)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("pair_id"),
			kallax.NewSchemaField("exchange_id"),
			kallax.NewSchemaField("is_active"),
		),
		ID:         kallax.NewSchemaField("id"),
		PairFK:     kallax.NewSchemaField("pair_id"),
		ExchangeFK: kallax.NewSchemaField("exchange_id"),
		IsActive:   kallax.NewSchemaField("is_active"),
	},
	Order: &schemaOrder{
		BaseSchema: kallax.NewBaseSchema(
			"orders",
			"__order",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"User":     kallax.NewForeignKey("user_id", true),
				"Exchange": kallax.NewForeignKey("exchange_id", true),
				"Pair":     kallax.NewForeignKey("pair_id", true),
				"Market":   kallax.NewForeignKey("market_id", true),
				"Trades":   kallax.NewForeignKey("order_id", false),
			},
			func() kallax.Record {
				return new(Order)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("user_id"),
			kallax.NewSchemaField("exchange_id"),
			kallax.NewSchemaField("pair_id"),
			kallax.NewSchemaField("market_id"),
			kallax.NewSchemaField("order_type"),
			kallax.NewSchemaField("open_price"),
			kallax.NewSchemaField("close_price"),
			kallax.NewSchemaField("ordered_volume"),
			kallax.NewSchemaField("swapped_volume"),
			kallax.NewSchemaField("is_closed"),
			kallax.NewSchemaField("stop_loss"),
			kallax.NewSchemaField("take_profit"),
			kallax.NewSchemaField("buy_fee"),
			kallax.NewSchemaField("sell_fee"),
			kallax.NewSchemaField("delta"),
		),
		ID:            kallax.NewSchemaField("id"),
		CreatedAt:     kallax.NewSchemaField("created_at"),
		UpdatedAt:     kallax.NewSchemaField("updated_at"),
		UserFK:        kallax.NewSchemaField("user_id"),
		ExchangeFK:    kallax.NewSchemaField("exchange_id"),
		PairFK:        kallax.NewSchemaField("pair_id"),
		MarketFK:      kallax.NewSchemaField("market_id"),
		OrderType:     kallax.NewSchemaField("order_type"),
		OpenPrice:     kallax.NewSchemaField("open_price"),
		ClosePrice:    kallax.NewSchemaField("close_price"),
		OrderedVolume: kallax.NewSchemaField("ordered_volume"),
		SwappedVolume: kallax.NewSchemaField("swapped_volume"),
		IsClosed:      kallax.NewSchemaField("is_closed"),
		StopLoss:      kallax.NewSchemaField("stop_loss"),
		TakeProfit:    kallax.NewSchemaField("take_profit"),
		BuyFee:        kallax.NewSchemaField("buy_fee"),
		SellFee:       kallax.NewSchemaField("sell_fee"),
		Delta:         kallax.NewSchemaField("delta"),
	},
	Pair: &schemaPair{
		BaseSchema: kallax.NewBaseSchema(
			"pairs",
			"__pair",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Base":        kallax.NewForeignKey("base_id", true),
				"Quote":       kallax.NewForeignKey("quote_id", true),
				"Markets":     kallax.NewForeignKey("pair_id", false),
				"Differences": kallax.NewForeignKey("pair_id", false),
				"Orders":      kallax.NewForeignKey("pair_id", false),
				"Prices":      kallax.NewForeignKey("pair_id", false),
				"Trades":      kallax.NewForeignKey("pair_id", false),
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
		ID:      kallax.NewSchemaField("id"),
		Symbol:  kallax.NewSchemaField("symbol"),
		BaseFK:  kallax.NewSchemaField("base_id"),
		QuoteFK: kallax.NewSchemaField("quote_id"),
	},
	Price: &schemaPrice{
		BaseSchema: kallax.NewBaseSchema(
			"prices",
			"__price",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Pair":     kallax.NewForeignKey("pair_id", true),
				"Exchange": kallax.NewForeignKey("exchange_id", true),
				"Market":   kallax.NewForeignKey("market_id", true),
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
			kallax.NewSchemaField("market_id"),
			kallax.NewSchemaField("price"),
			kallax.NewSchemaField("pair_symbols"),
			kallax.NewSchemaField("exchange_symbols"),
		),
		ID:              kallax.NewSchemaField("id"),
		CreatedAt:       kallax.NewSchemaField("created_at"),
		UpdatedAt:       kallax.NewSchemaField("updated_at"),
		PairFK:          kallax.NewSchemaField("pair_id"),
		ExchangeFK:      kallax.NewSchemaField("exchange_id"),
		MarketFK:        kallax.NewSchemaField("market_id"),
		Price:           kallax.NewSchemaField("price"),
		PairSymbols:     kallax.NewSchemaField("pair_symbols"),
		ExchangeSymbols: kallax.NewSchemaField("exchange_symbols"),
	},
	Trade: &schemaTrade{
		BaseSchema: kallax.NewBaseSchema(
			"trades",
			"__trade",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"User":     kallax.NewForeignKey("user_id", true),
				"Exchange": kallax.NewForeignKey("exchange_id", true),
				"Pair":     kallax.NewForeignKey("pair_id", true),
				"Market":   kallax.NewForeignKey("market_id", true),
				"Order":    kallax.NewForeignKey("order_id", true),
			},
			func() kallax.Record {
				return new(Trade)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("user_id"),
			kallax.NewSchemaField("exchange_id"),
			kallax.NewSchemaField("pair_id"),
			kallax.NewSchemaField("market_id"),
			kallax.NewSchemaField("order_id"),
			kallax.NewSchemaField("type"),
			kallax.NewSchemaField("volume"),
			kallax.NewSchemaField("price"),
		),
		ID:         kallax.NewSchemaField("id"),
		CreatedAt:  kallax.NewSchemaField("created_at"),
		UpdatedAt:  kallax.NewSchemaField("updated_at"),
		UserFK:     kallax.NewSchemaField("user_id"),
		ExchangeFK: kallax.NewSchemaField("exchange_id"),
		PairFK:     kallax.NewSchemaField("pair_id"),
		MarketFK:   kallax.NewSchemaField("market_id"),
		OrderFK:    kallax.NewSchemaField("order_id"),
		Type:       kallax.NewSchemaField("type"),
		Volume:     kallax.NewSchemaField("volume"),
		Price:      kallax.NewSchemaField("price"),
	},
	User: &schemaUser{
		BaseSchema: kallax.NewBaseSchema(
			"users",
			"__user",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Balances": kallax.NewForeignKey("user_id", false),
				"Orders":   kallax.NewForeignKey("user_id", false),
				"Trades":   kallax.NewForeignKey("user_id", false),
			},
			func() kallax.Record {
				return new(User)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("email"),
			kallax.NewSchemaField("password"),
			kallax.NewSchemaField("last_login"),
			kallax.NewSchemaField("subscribe_to"),
			kallax.NewSchemaField("role"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
		),
		ID:          kallax.NewSchemaField("id"),
		Email:       kallax.NewSchemaField("email"),
		Password:    kallax.NewSchemaField("password"),
		LastLogin:   kallax.NewSchemaField("last_login"),
		SubscribeTo: kallax.NewSchemaField("subscribe_to"),
		Role:        kallax.NewSchemaField("role"),
		CreatedAt:   kallax.NewSchemaField("created_at"),
		UpdatedAt:   kallax.NewSchemaField("updated_at"),
	},
	UserBalance: &schemaUserBalance{
		BaseSchema: kallax.NewBaseSchema(
			"user_balances",
			"__userbalance",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"User":     kallax.NewForeignKey("user_id", true),
				"Exchange": kallax.NewForeignKey("exchange_id", true),
				"Asset":    kallax.NewForeignKey("asset_id", true),
			},
			func() kallax.Record {
				return new(UserBalance)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("user_id"),
			kallax.NewSchemaField("exchange_id"),
			kallax.NewSchemaField("asset_id"),
			kallax.NewSchemaField("volume"),
		),
		ID:         kallax.NewSchemaField("id"),
		UserFK:     kallax.NewSchemaField("user_id"),
		ExchangeFK: kallax.NewSchemaField("exchange_id"),
		AssetFK:    kallax.NewSchemaField("asset_id"),
		Volume:     kallax.NewSchemaField("volume"),
	},
}
