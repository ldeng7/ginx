package types

import (
	"context"
	"errors"

	"github.com/ldeng7/ginx/gormx"
	"github.com/ldeng7/go-mysql-datatypes/mysqldatatypes/spatial"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var errScanType = errors.New("Failed to scan a geometry field: invalid data type")

type Point struct {
	spatial.Point
}

func (p Point) GormDataType() string {
	return "geometry"
}

func (p *Point) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return p.Decode(bs)
}

func (p Point) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(p.Encode())}
}

type LineString struct {
	spatial.LineString
}

func (l LineString) GormDataType() string {
	return "geometry"
}

func (l *LineString) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return l.Decode(bs)
}

func (l LineString) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(l.Encode())}
}

type Polygon struct {
	spatial.Polygon
}

func (p Polygon) GormDataType() string {
	return "geometry"
}

func (p *Polygon) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return p.Decode(bs)
}

func (p Polygon) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(p.Encode())}
}

type MultiPoint struct {
	spatial.MultiPoint
}

func (mp MultiPoint) GormDataType() string {
	return "geometry"
}

func (mp *MultiPoint) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return mp.Decode(bs)
}

func (mp MultiPoint) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(mp.Encode())}
}

type MultiLineString struct {
	spatial.MultiLineString
}

func (ml MultiLineString) GormDataType() string {
	return "geometry"
}

func (ml *MultiLineString) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return ml.Decode(bs)
}

func (ml MultiLineString) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(ml.Encode())}
}

type MultiPolygon struct {
	spatial.MultiPolygon
}

func (mp MultiPolygon) GormDataType() string {
	return "geometry"
}

func (mp *MultiPolygon) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return mp.Decode(bs)
}

func (mp MultiPolygon) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(mp.Encode())}
}

type GeometryCollection struct {
	spatial.GeometryCollection
}

func (c GeometryCollection) GormDataType() string {
	return "geometry"
}

func (c *GeometryCollection) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return c.Decode(bs)
}

func (c GeometryCollection) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(c.Encode())}
}

type GenericGeometry struct {
	spatial.GenericGeometry
}

func (g GenericGeometry) GormDataType() string {
	return "geometry"
}

func (g *GenericGeometry) Scan(v any) error {
	bs, ok := v.([]byte)
	if !ok {
		return errScanType
	}
	return g.Decode(bs)
}

func (g GenericGeometry) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	return clause.Expr{SQL: gormx.BytesToSql(g.Encode())}
}
