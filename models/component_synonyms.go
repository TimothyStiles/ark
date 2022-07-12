// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ComponentSynonym is an object representing the database table.
type ComponentSynonym struct {
	CompsynID        int64       `boil:"compsyn_id" json:"compsyn_id" toml:"compsyn_id" yaml:"compsyn_id"`
	ComponentID      int64       `boil:"component_id" json:"component_id" toml:"component_id" yaml:"component_id"`
	ComponentSynonym null.String `boil:"component_synonym" json:"component_synonym,omitempty" toml:"component_synonym" yaml:"component_synonym,omitempty"`
	SynType          null.String `boil:"syn_type" json:"syn_type,omitempty" toml:"syn_type" yaml:"syn_type,omitempty"`

	R *componentSynonymR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L componentSynonymL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ComponentSynonymColumns = struct {
	CompsynID        string
	ComponentID      string
	ComponentSynonym string
	SynType          string
}{
	CompsynID:        "compsyn_id",
	ComponentID:      "component_id",
	ComponentSynonym: "component_synonym",
	SynType:          "syn_type",
}

var ComponentSynonymTableColumns = struct {
	CompsynID        string
	ComponentID      string
	ComponentSynonym string
	SynType          string
}{
	CompsynID:        "component_synonyms.compsyn_id",
	ComponentID:      "component_synonyms.component_id",
	ComponentSynonym: "component_synonyms.component_synonym",
	SynType:          "component_synonyms.syn_type",
}

// Generated where

var ComponentSynonymWhere = struct {
	CompsynID        whereHelperint64
	ComponentID      whereHelperint64
	ComponentSynonym whereHelpernull_String
	SynType          whereHelpernull_String
}{
	CompsynID:        whereHelperint64{field: "\"component_synonyms\".\"compsyn_id\""},
	ComponentID:      whereHelperint64{field: "\"component_synonyms\".\"component_id\""},
	ComponentSynonym: whereHelpernull_String{field: "\"component_synonyms\".\"component_synonym\""},
	SynType:          whereHelpernull_String{field: "\"component_synonyms\".\"syn_type\""},
}

// ComponentSynonymRels is where relationship names are stored.
var ComponentSynonymRels = struct {
	Component string
}{
	Component: "Component",
}

// componentSynonymR is where relationships are stored.
type componentSynonymR struct {
	Component *ComponentSequence `boil:"Component" json:"Component" toml:"Component" yaml:"Component"`
}

// NewStruct creates a new relationship struct
func (*componentSynonymR) NewStruct() *componentSynonymR {
	return &componentSynonymR{}
}

func (r *componentSynonymR) GetComponent() *ComponentSequence {
	if r == nil {
		return nil
	}
	return r.Component
}

// componentSynonymL is where Load methods for each relationship are stored.
type componentSynonymL struct{}

var (
	componentSynonymAllColumns            = []string{"compsyn_id", "component_id", "component_synonym", "syn_type"}
	componentSynonymColumnsWithoutDefault = []string{"compsyn_id", "component_id"}
	componentSynonymColumnsWithDefault    = []string{"component_synonym", "syn_type"}
	componentSynonymPrimaryKeyColumns     = []string{"compsyn_id"}
	componentSynonymGeneratedColumns      = []string{}
)

type (
	// ComponentSynonymSlice is an alias for a slice of pointers to ComponentSynonym.
	// This should almost always be used instead of []ComponentSynonym.
	ComponentSynonymSlice []*ComponentSynonym
	// ComponentSynonymHook is the signature for custom ComponentSynonym hook methods
	ComponentSynonymHook func(context.Context, boil.ContextExecutor, *ComponentSynonym) error

	componentSynonymQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	componentSynonymType                 = reflect.TypeOf(&ComponentSynonym{})
	componentSynonymMapping              = queries.MakeStructMapping(componentSynonymType)
	componentSynonymPrimaryKeyMapping, _ = queries.BindMapping(componentSynonymType, componentSynonymMapping, componentSynonymPrimaryKeyColumns)
	componentSynonymInsertCacheMut       sync.RWMutex
	componentSynonymInsertCache          = make(map[string]insertCache)
	componentSynonymUpdateCacheMut       sync.RWMutex
	componentSynonymUpdateCache          = make(map[string]updateCache)
	componentSynonymUpsertCacheMut       sync.RWMutex
	componentSynonymUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var componentSynonymAfterSelectHooks []ComponentSynonymHook

var componentSynonymBeforeInsertHooks []ComponentSynonymHook
var componentSynonymAfterInsertHooks []ComponentSynonymHook

var componentSynonymBeforeUpdateHooks []ComponentSynonymHook
var componentSynonymAfterUpdateHooks []ComponentSynonymHook

var componentSynonymBeforeDeleteHooks []ComponentSynonymHook
var componentSynonymAfterDeleteHooks []ComponentSynonymHook

var componentSynonymBeforeUpsertHooks []ComponentSynonymHook
var componentSynonymAfterUpsertHooks []ComponentSynonymHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ComponentSynonym) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ComponentSynonym) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ComponentSynonym) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ComponentSynonym) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ComponentSynonym) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ComponentSynonym) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ComponentSynonym) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ComponentSynonym) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ComponentSynonym) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentSynonymAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddComponentSynonymHook registers your hook function for all future operations.
func AddComponentSynonymHook(hookPoint boil.HookPoint, componentSynonymHook ComponentSynonymHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		componentSynonymAfterSelectHooks = append(componentSynonymAfterSelectHooks, componentSynonymHook)
	case boil.BeforeInsertHook:
		componentSynonymBeforeInsertHooks = append(componentSynonymBeforeInsertHooks, componentSynonymHook)
	case boil.AfterInsertHook:
		componentSynonymAfterInsertHooks = append(componentSynonymAfterInsertHooks, componentSynonymHook)
	case boil.BeforeUpdateHook:
		componentSynonymBeforeUpdateHooks = append(componentSynonymBeforeUpdateHooks, componentSynonymHook)
	case boil.AfterUpdateHook:
		componentSynonymAfterUpdateHooks = append(componentSynonymAfterUpdateHooks, componentSynonymHook)
	case boil.BeforeDeleteHook:
		componentSynonymBeforeDeleteHooks = append(componentSynonymBeforeDeleteHooks, componentSynonymHook)
	case boil.AfterDeleteHook:
		componentSynonymAfterDeleteHooks = append(componentSynonymAfterDeleteHooks, componentSynonymHook)
	case boil.BeforeUpsertHook:
		componentSynonymBeforeUpsertHooks = append(componentSynonymBeforeUpsertHooks, componentSynonymHook)
	case boil.AfterUpsertHook:
		componentSynonymAfterUpsertHooks = append(componentSynonymAfterUpsertHooks, componentSynonymHook)
	}
}

// One returns a single componentSynonym record from the query.
func (q componentSynonymQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ComponentSynonym, error) {
	o := &ComponentSynonym{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for component_synonyms")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ComponentSynonym records from the query.
func (q componentSynonymQuery) All(ctx context.Context, exec boil.ContextExecutor) (ComponentSynonymSlice, error) {
	var o []*ComponentSynonym

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ComponentSynonym slice")
	}

	if len(componentSynonymAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ComponentSynonym records in the query.
func (q componentSynonymQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count component_synonyms rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q componentSynonymQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if component_synonyms exists")
	}

	return count > 0, nil
}

// Component pointed to by the foreign key.
func (o *ComponentSynonym) Component(mods ...qm.QueryMod) componentSequenceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"component_id\" = ?", o.ComponentID),
	}

	queryMods = append(queryMods, mods...)

	return ComponentSequences(queryMods...)
}

// LoadComponent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (componentSynonymL) LoadComponent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeComponentSynonym interface{}, mods queries.Applicator) error {
	var slice []*ComponentSynonym
	var object *ComponentSynonym

	if singular {
		object = maybeComponentSynonym.(*ComponentSynonym)
	} else {
		slice = *maybeComponentSynonym.(*[]*ComponentSynonym)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &componentSynonymR{}
		}
		args = append(args, object.ComponentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &componentSynonymR{}
			}

			for _, a := range args {
				if a == obj.ComponentID {
					continue Outer
				}
			}

			args = append(args, obj.ComponentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`component_sequences`),
		qm.WhereIn(`component_sequences.component_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ComponentSequence")
	}

	var resultSlice []*ComponentSequence
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ComponentSequence")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for component_sequences")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for component_sequences")
	}

	if len(componentSynonymAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Component = foreign
		if foreign.R == nil {
			foreign.R = &componentSequenceR{}
		}
		foreign.R.ComponentComponentSynonyms = append(foreign.R.ComponentComponentSynonyms, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ComponentID == foreign.ComponentID {
				local.R.Component = foreign
				if foreign.R == nil {
					foreign.R = &componentSequenceR{}
				}
				foreign.R.ComponentComponentSynonyms = append(foreign.R.ComponentComponentSynonyms, local)
				break
			}
		}
	}

	return nil
}

// SetComponent of the componentSynonym to the related item.
// Sets o.R.Component to related.
// Adds o to related.R.ComponentComponentSynonyms.
func (o *ComponentSynonym) SetComponent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ComponentSequence) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"component_synonyms\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"component_id"}),
		strmangle.WhereClause("\"", "\"", 0, componentSynonymPrimaryKeyColumns),
	)
	values := []interface{}{related.ComponentID, o.CompsynID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ComponentID = related.ComponentID
	if o.R == nil {
		o.R = &componentSynonymR{
			Component: related,
		}
	} else {
		o.R.Component = related
	}

	if related.R == nil {
		related.R = &componentSequenceR{
			ComponentComponentSynonyms: ComponentSynonymSlice{o},
		}
	} else {
		related.R.ComponentComponentSynonyms = append(related.R.ComponentComponentSynonyms, o)
	}

	return nil
}

// ComponentSynonyms retrieves all the records using an executor.
func ComponentSynonyms(mods ...qm.QueryMod) componentSynonymQuery {
	mods = append(mods, qm.From("\"component_synonyms\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"component_synonyms\".*"})
	}

	return componentSynonymQuery{q}
}

// FindComponentSynonym retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindComponentSynonym(ctx context.Context, exec boil.ContextExecutor, compsynID int64, selectCols ...string) (*ComponentSynonym, error) {
	componentSynonymObj := &ComponentSynonym{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"component_synonyms\" where \"compsyn_id\"=?", sel,
	)

	q := queries.Raw(query, compsynID)

	err := q.Bind(ctx, exec, componentSynonymObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from component_synonyms")
	}

	if err = componentSynonymObj.doAfterSelectHooks(ctx, exec); err != nil {
		return componentSynonymObj, err
	}

	return componentSynonymObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ComponentSynonym) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no component_synonyms provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(componentSynonymColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	componentSynonymInsertCacheMut.RLock()
	cache, cached := componentSynonymInsertCache[key]
	componentSynonymInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			componentSynonymAllColumns,
			componentSynonymColumnsWithDefault,
			componentSynonymColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(componentSynonymType, componentSynonymMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(componentSynonymType, componentSynonymMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"component_synonyms\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"component_synonyms\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into component_synonyms")
	}

	if !cached {
		componentSynonymInsertCacheMut.Lock()
		componentSynonymInsertCache[key] = cache
		componentSynonymInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ComponentSynonym.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ComponentSynonym) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	componentSynonymUpdateCacheMut.RLock()
	cache, cached := componentSynonymUpdateCache[key]
	componentSynonymUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			componentSynonymAllColumns,
			componentSynonymPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update component_synonyms, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"component_synonyms\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, componentSynonymPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(componentSynonymType, componentSynonymMapping, append(wl, componentSynonymPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update component_synonyms row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for component_synonyms")
	}

	if !cached {
		componentSynonymUpdateCacheMut.Lock()
		componentSynonymUpdateCache[key] = cache
		componentSynonymUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q componentSynonymQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for component_synonyms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for component_synonyms")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ComponentSynonymSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), componentSynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"component_synonyms\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, componentSynonymPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in componentSynonym slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all componentSynonym")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ComponentSynonym) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no component_synonyms provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(componentSynonymColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	componentSynonymUpsertCacheMut.RLock()
	cache, cached := componentSynonymUpsertCache[key]
	componentSynonymUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			componentSynonymAllColumns,
			componentSynonymColumnsWithDefault,
			componentSynonymColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			componentSynonymAllColumns,
			componentSynonymPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert component_synonyms, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(componentSynonymPrimaryKeyColumns))
			copy(conflict, componentSynonymPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"component_synonyms\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(componentSynonymType, componentSynonymMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(componentSynonymType, componentSynonymMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert component_synonyms")
	}

	if !cached {
		componentSynonymUpsertCacheMut.Lock()
		componentSynonymUpsertCache[key] = cache
		componentSynonymUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ComponentSynonym record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ComponentSynonym) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ComponentSynonym provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), componentSynonymPrimaryKeyMapping)
	sql := "DELETE FROM \"component_synonyms\" WHERE \"compsyn_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from component_synonyms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for component_synonyms")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q componentSynonymQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no componentSynonymQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from component_synonyms")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for component_synonyms")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ComponentSynonymSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(componentSynonymBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), componentSynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"component_synonyms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, componentSynonymPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from componentSynonym slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for component_synonyms")
	}

	if len(componentSynonymAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ComponentSynonym) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindComponentSynonym(ctx, exec, o.CompsynID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ComponentSynonymSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ComponentSynonymSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), componentSynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"component_synonyms\".* FROM \"component_synonyms\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, componentSynonymPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ComponentSynonymSlice")
	}

	*o = slice

	return nil
}

// ComponentSynonymExists checks if the ComponentSynonym row exists.
func ComponentSynonymExists(ctx context.Context, exec boil.ContextExecutor, compsynID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"component_synonyms\" where \"compsyn_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, compsynID)
	}
	row := exec.QueryRowContext(ctx, sql, compsynID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if component_synonyms exists")
	}

	return exists, nil
}