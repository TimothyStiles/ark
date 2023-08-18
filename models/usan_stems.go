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

// UsanStem is an object representing the database table.
type UsanStem struct {
	UsanStemID int64       `boil:"usan_stem_id" json:"usan_stem_id" toml:"usan_stem_id" yaml:"usan_stem_id"`
	Stem       string      `boil:"stem" json:"stem" toml:"stem" yaml:"stem"`
	Subgroup   string      `boil:"subgroup" json:"subgroup" toml:"subgroup" yaml:"subgroup"`
	Annotation null.String `boil:"annotation" json:"annotation,omitempty" toml:"annotation" yaml:"annotation,omitempty"`
	StemClass  null.String `boil:"stem_class" json:"stem_class,omitempty" toml:"stem_class" yaml:"stem_class,omitempty"`
	MajorClass null.String `boil:"major_class" json:"major_class,omitempty" toml:"major_class" yaml:"major_class,omitempty"`
	WhoExtra   null.Int16  `boil:"who_extra" json:"who_extra,omitempty" toml:"who_extra" yaml:"who_extra,omitempty"`

	R *usanStemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L usanStemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UsanStemColumns = struct {
	UsanStemID string
	Stem       string
	Subgroup   string
	Annotation string
	StemClass  string
	MajorClass string
	WhoExtra   string
}{
	UsanStemID: "usan_stem_id",
	Stem:       "stem",
	Subgroup:   "subgroup",
	Annotation: "annotation",
	StemClass:  "stem_class",
	MajorClass: "major_class",
	WhoExtra:   "who_extra",
}

var UsanStemTableColumns = struct {
	UsanStemID string
	Stem       string
	Subgroup   string
	Annotation string
	StemClass  string
	MajorClass string
	WhoExtra   string
}{
	UsanStemID: "usan_stems.usan_stem_id",
	Stem:       "usan_stems.stem",
	Subgroup:   "usan_stems.subgroup",
	Annotation: "usan_stems.annotation",
	StemClass:  "usan_stems.stem_class",
	MajorClass: "usan_stems.major_class",
	WhoExtra:   "usan_stems.who_extra",
}

// Generated where

var UsanStemWhere = struct {
	UsanStemID whereHelperint64
	Stem       whereHelperstring
	Subgroup   whereHelperstring
	Annotation whereHelpernull_String
	StemClass  whereHelpernull_String
	MajorClass whereHelpernull_String
	WhoExtra   whereHelpernull_Int16
}{
	UsanStemID: whereHelperint64{field: "\"usan_stems\".\"usan_stem_id\""},
	Stem:       whereHelperstring{field: "\"usan_stems\".\"stem\""},
	Subgroup:   whereHelperstring{field: "\"usan_stems\".\"subgroup\""},
	Annotation: whereHelpernull_String{field: "\"usan_stems\".\"annotation\""},
	StemClass:  whereHelpernull_String{field: "\"usan_stems\".\"stem_class\""},
	MajorClass: whereHelpernull_String{field: "\"usan_stems\".\"major_class\""},
	WhoExtra:   whereHelpernull_Int16{field: "\"usan_stems\".\"who_extra\""},
}

// UsanStemRels is where relationship names are stored.
var UsanStemRels = struct {
}{}

// usanStemR is where relationships are stored.
type usanStemR struct {
}

// NewStruct creates a new relationship struct
func (*usanStemR) NewStruct() *usanStemR {
	return &usanStemR{}
}

// usanStemL is where Load methods for each relationship are stored.
type usanStemL struct{}

var (
	usanStemAllColumns            = []string{"usan_stem_id", "stem", "subgroup", "annotation", "stem_class", "major_class", "who_extra"}
	usanStemColumnsWithoutDefault = []string{"usan_stem_id", "stem", "subgroup"}
	usanStemColumnsWithDefault    = []string{"annotation", "stem_class", "major_class", "who_extra"}
	usanStemPrimaryKeyColumns     = []string{"usan_stem_id"}
	usanStemGeneratedColumns      = []string{}
)

type (
	// UsanStemSlice is an alias for a slice of pointers to UsanStem.
	// This should almost always be used instead of []UsanStem.
	UsanStemSlice []*UsanStem
	// UsanStemHook is the signature for custom UsanStem hook methods
	UsanStemHook func(context.Context, boil.ContextExecutor, *UsanStem) error

	usanStemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	usanStemType                 = reflect.TypeOf(&UsanStem{})
	usanStemMapping              = queries.MakeStructMapping(usanStemType)
	usanStemPrimaryKeyMapping, _ = queries.BindMapping(usanStemType, usanStemMapping, usanStemPrimaryKeyColumns)
	usanStemInsertCacheMut       sync.RWMutex
	usanStemInsertCache          = make(map[string]insertCache)
	usanStemUpdateCacheMut       sync.RWMutex
	usanStemUpdateCache          = make(map[string]updateCache)
	usanStemUpsertCacheMut       sync.RWMutex
	usanStemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var usanStemAfterSelectHooks []UsanStemHook

var usanStemBeforeInsertHooks []UsanStemHook
var usanStemAfterInsertHooks []UsanStemHook

var usanStemBeforeUpdateHooks []UsanStemHook
var usanStemAfterUpdateHooks []UsanStemHook

var usanStemBeforeDeleteHooks []UsanStemHook
var usanStemAfterDeleteHooks []UsanStemHook

var usanStemBeforeUpsertHooks []UsanStemHook
var usanStemAfterUpsertHooks []UsanStemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UsanStem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UsanStem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UsanStem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UsanStem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UsanStem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UsanStem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UsanStem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UsanStem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UsanStem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usanStemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUsanStemHook registers your hook function for all future operations.
func AddUsanStemHook(hookPoint boil.HookPoint, usanStemHook UsanStemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		usanStemAfterSelectHooks = append(usanStemAfterSelectHooks, usanStemHook)
	case boil.BeforeInsertHook:
		usanStemBeforeInsertHooks = append(usanStemBeforeInsertHooks, usanStemHook)
	case boil.AfterInsertHook:
		usanStemAfterInsertHooks = append(usanStemAfterInsertHooks, usanStemHook)
	case boil.BeforeUpdateHook:
		usanStemBeforeUpdateHooks = append(usanStemBeforeUpdateHooks, usanStemHook)
	case boil.AfterUpdateHook:
		usanStemAfterUpdateHooks = append(usanStemAfterUpdateHooks, usanStemHook)
	case boil.BeforeDeleteHook:
		usanStemBeforeDeleteHooks = append(usanStemBeforeDeleteHooks, usanStemHook)
	case boil.AfterDeleteHook:
		usanStemAfterDeleteHooks = append(usanStemAfterDeleteHooks, usanStemHook)
	case boil.BeforeUpsertHook:
		usanStemBeforeUpsertHooks = append(usanStemBeforeUpsertHooks, usanStemHook)
	case boil.AfterUpsertHook:
		usanStemAfterUpsertHooks = append(usanStemAfterUpsertHooks, usanStemHook)
	}
}

// One returns a single usanStem record from the query.
func (q usanStemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UsanStem, error) {
	o := &UsanStem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for usan_stems")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UsanStem records from the query.
func (q usanStemQuery) All(ctx context.Context, exec boil.ContextExecutor) (UsanStemSlice, error) {
	var o []*UsanStem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UsanStem slice")
	}

	if len(usanStemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UsanStem records in the query.
func (q usanStemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count usan_stems rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q usanStemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if usan_stems exists")
	}

	return count > 0, nil
}

// UsanStems retrieves all the records using an executor.
func UsanStems(mods ...qm.QueryMod) usanStemQuery {
	mods = append(mods, qm.From("\"usan_stems\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"usan_stems\".*"})
	}

	return usanStemQuery{q}
}

// FindUsanStem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUsanStem(ctx context.Context, exec boil.ContextExecutor, usanStemID int64, selectCols ...string) (*UsanStem, error) {
	usanStemObj := &UsanStem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"usan_stems\" where \"usan_stem_id\"=?", sel,
	)

	q := queries.Raw(query, usanStemID)

	err := q.Bind(ctx, exec, usanStemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from usan_stems")
	}

	if err = usanStemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return usanStemObj, err
	}

	return usanStemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UsanStem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no usan_stems provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(usanStemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	usanStemInsertCacheMut.RLock()
	cache, cached := usanStemInsertCache[key]
	usanStemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			usanStemAllColumns,
			usanStemColumnsWithDefault,
			usanStemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(usanStemType, usanStemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(usanStemType, usanStemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"usan_stems\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"usan_stems\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into usan_stems")
	}

	if !cached {
		usanStemInsertCacheMut.Lock()
		usanStemInsertCache[key] = cache
		usanStemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UsanStem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UsanStem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	usanStemUpdateCacheMut.RLock()
	cache, cached := usanStemUpdateCache[key]
	usanStemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			usanStemAllColumns,
			usanStemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update usan_stems, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"usan_stems\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, usanStemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(usanStemType, usanStemMapping, append(wl, usanStemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update usan_stems row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for usan_stems")
	}

	if !cached {
		usanStemUpdateCacheMut.Lock()
		usanStemUpdateCache[key] = cache
		usanStemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q usanStemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for usan_stems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for usan_stems")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UsanStemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usanStemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"usan_stems\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, usanStemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in usanStem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all usanStem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UsanStem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no usan_stems provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(usanStemColumnsWithDefault, o)

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

	usanStemUpsertCacheMut.RLock()
	cache, cached := usanStemUpsertCache[key]
	usanStemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			usanStemAllColumns,
			usanStemColumnsWithDefault,
			usanStemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			usanStemAllColumns,
			usanStemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert usan_stems, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(usanStemPrimaryKeyColumns))
			copy(conflict, usanStemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"usan_stems\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(usanStemType, usanStemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(usanStemType, usanStemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert usan_stems")
	}

	if !cached {
		usanStemUpsertCacheMut.Lock()
		usanStemUpsertCache[key] = cache
		usanStemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UsanStem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UsanStem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UsanStem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), usanStemPrimaryKeyMapping)
	sql := "DELETE FROM \"usan_stems\" WHERE \"usan_stem_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from usan_stems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for usan_stems")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q usanStemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no usanStemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from usan_stems")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for usan_stems")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UsanStemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(usanStemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usanStemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"usan_stems\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, usanStemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from usanStem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for usan_stems")
	}

	if len(usanStemAfterDeleteHooks) != 0 {
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
func (o *UsanStem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUsanStem(ctx, exec, o.UsanStemID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UsanStemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UsanStemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usanStemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"usan_stems\".* FROM \"usan_stems\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, usanStemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UsanStemSlice")
	}

	*o = slice

	return nil
}

// UsanStemExists checks if the UsanStem row exists.
func UsanStemExists(ctx context.Context, exec boil.ContextExecutor, usanStemID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"usan_stems\" where \"usan_stem_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, usanStemID)
	}
	row := exec.QueryRowContext(ctx, sql, usanStemID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if usan_stems exists")
	}

	return exists, nil
}