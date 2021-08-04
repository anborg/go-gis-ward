package repo

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
	"io/ioutil"
	"log"
)

type Ward struct {
	features *geojson.FeatureCollection
}

func (this *Ward) New(geoJsonPath string) error {

	// Load in our geojson file into a features collection
	b, _ := ioutil.ReadFile(geoJsonPath)
	wardFeatures, _ := geojson.UnmarshalFeatureCollection(b)
	this.features = wardFeatures
	//this = &Ward{
	//	features: wardFeatures,
	//}
	p1 := orb.Point{-79.2498779296875, 43.89195472686543}
	this.GetWards(p1)
	p2 := orb.Point{-79.36248779296874, 43.847403373019226}
	this.GetWards(p2)
	return nil
}

func (this *Ward) GetWards(point orb.Point) (outJson interface{}) {
	type empty struct{}
	outJson = new(empty)
	for _, feature := range this.features.Features {
		// Try on a MultiPolygon to begin
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if isMulti {
			if planar.MultiPolygonContains(multiPoly, point) {
				return
			}
		} else {
			// Fallback to Polygon
			polygon, isPoly := feature.Geometry.(orb.Polygon)
			if isPoly {
				if planar.PolygonContains(polygon, point) {
					//log.Println("Polygon has points" , polygon)
					log.Println("Point :", point)
					log.Println("At Ward:", feature.Properties)
					//outJson, _ = json.Marshal(feature.Properties)
					outJson = feature.Properties
				}
			}
		}
	} //for
	return
}
