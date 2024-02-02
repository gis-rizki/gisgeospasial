package gisgeospasial

import (
	"fmt"
	"testing"
)

var dbname = "gismongodb"
var collname = "projectgis"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Point{
		Coordinates: []float64{
			106.84221779677222, -6.484133450448681,
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
				{106.82553814331038, -6.476148372482314},
				{106.82497554919382, -6.47605132314213},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "gis", "gis")
	coordinates := Point{
		Coordinates: []float64{
			106.83767535154254, -6.4844685488837115,
		},
	}
	datagedung := Near(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Point{
		Coordinates: []float64{
			106.8381783906836, -6.485953728953618,
		},
	}
	datagedung := NearSphere(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{106.84072267772467, -6.476259383974224},
			{106.8399156012594, -6.476168843906208},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
