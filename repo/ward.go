package repo

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
	"io/ioutil"
)

type Ward struct {
	features *geojson.FeatureCollection
}

func (this *Ward) New(geoJsonPath string) error {
	byteJson, err1 := ioutil.ReadFile(geoJsonPath); if err1 != nil {
		return err1
	}
	wardFeatures, err := geojson.UnmarshalFeatureCollection(byteJson); if err != nil {
		//log.Fatalf("Error reading GIS file :", geojson, err)
		return err
	}
	this.features = wardFeatures
	return nil
}

func (this *Ward) GetWards(point orb.Point) (outJson interface{}) {
	type empty struct{}
	outJson = new(empty)
	for _, feature := range this.features.Features {
		// Try on a MultiPolygon to begin
		if multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon); isMulti {
			if planar.MultiPolygonContains(multiPoly, point) {
				return
			}
		}
		// Fallback to Polygon
		if polygon, isPoly := feature.Geometry.(orb.Polygon); isPoly {
			if planar.PolygonContains(polygon, point) {
				//log.Println("Point :", point)
				//log.Println("At Ward:", feature.Properties)
				outJson = feature.Properties
			}
		}
	} //for
	return
}
