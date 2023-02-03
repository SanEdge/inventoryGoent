// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/renaldyhidayatt/inventorygoent/ent/productmasuk"
)

// ProductMasuk is the model entity for the ProductMasuk schema.
type ProductMasuk struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Qty holds the value of the "qty" field.
	Qty string `json:"qty,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductMasukQuery when eager-loading is set.
	Edges ProductMasukEdges `json:"edges"`
}

// ProductMasukEdges holds the relations/edges for other nodes in the graph.
type ProductMasukEdges struct {
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// Supplier holds the value of the supplier edge.
	Supplier []*Supplier `json:"supplier,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e ProductMasukEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// SupplierOrErr returns the Supplier value or an error if the edge
// was not loaded in eager-loading.
func (e ProductMasukEdges) SupplierOrErr() ([]*Supplier, error) {
	if e.loadedTypes[1] {
		return e.Supplier, nil
	}
	return nil, &NotLoadedError{edge: "supplier"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductMasuk) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productmasuk.FieldID:
			values[i] = new(sql.NullInt64)
		case productmasuk.FieldName, productmasuk.FieldQty:
			values[i] = new(sql.NullString)
		case productmasuk.FieldCreatedAt, productmasuk.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductMasuk", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductMasuk fields.
func (pm *ProductMasuk) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productmasuk.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pm.ID = int(value.Int64)
		case productmasuk.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pm.Name = value.String
			}
		case productmasuk.FieldQty:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field qty", values[i])
			} else if value.Valid {
				pm.Qty = value.String
			}
		case productmasuk.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pm.CreatedAt = value.Time
			}
		case productmasuk.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pm.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryProduct queries the "product" edge of the ProductMasuk entity.
func (pm *ProductMasuk) QueryProduct() *ProductQuery {
	return NewProductMasukClient(pm.config).QueryProduct(pm)
}

// QuerySupplier queries the "supplier" edge of the ProductMasuk entity.
func (pm *ProductMasuk) QuerySupplier() *SupplierQuery {
	return NewProductMasukClient(pm.config).QuerySupplier(pm)
}

// Update returns a builder for updating this ProductMasuk.
// Note that you need to call ProductMasuk.Unwrap() before calling this method if this ProductMasuk
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *ProductMasuk) Update() *ProductMasukUpdateOne {
	return NewProductMasukClient(pm.config).UpdateOne(pm)
}

// Unwrap unwraps the ProductMasuk entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pm *ProductMasuk) Unwrap() *ProductMasuk {
	_tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductMasuk is not a transactional entity")
	}
	pm.config.driver = _tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *ProductMasuk) String() string {
	var builder strings.Builder
	builder.WriteString("ProductMasuk(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pm.ID))
	builder.WriteString("name=")
	builder.WriteString(pm.Name)
	builder.WriteString(", ")
	builder.WriteString("qty=")
	builder.WriteString(pm.Qty)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pm.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProductMasuks is a parsable slice of ProductMasuk.
type ProductMasuks []*ProductMasuk

func (pm ProductMasuks) config(cfg config) {
	for _i := range pm {
		pm[_i].config = cfg
	}
}
