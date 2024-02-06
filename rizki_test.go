package gisgeospasial

import (
	"fmt"
	"testing"
)

var dbname = "gismongodb"
var collname = "projectgis"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{106.82570304536097, -6.477059201501504},
				{106.82606286126702, -6.477081901134298},
				{106.82614853172089, -6.476724381786013},
				{106.82578300445192, -6.476667632660593},
				{106.82570304536097, -6.477059201501504},
			},
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{106.82570304536097, -6.477059201501504},
				{106.82606286126702, -6.477081901134298},
				{106.82614853172089, -6.476724381786013},
				{106.82578300445192, -6.476667632660593},
				{106.82570304536097, -6.477059201501504},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "gismongodb", "projectgis")
	coordinates := Point{
		Coordinates: []float64{
			106.83165780963839, -6.4737440161616036,
		},
	}
	datagedung := Near(mconn, "projectgis", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gismongodb")
	coordinates := Point{
		Coordinates: []float64{
			106.82523088801918, -6.476722204156317,
		},
	}
	datagedung := NearSphere(mconn, "projectgis", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gismongodb")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{106.83882411639951, -6.477228063790221},
			{106.83879808167569, -6.477803638454645},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
