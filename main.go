package main
import (
	"flag"
	//"io"
	"io/ioutil"
	"log"

// 	"net/http"
// 	"github.com/gin-gonic/gin"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)
const (
	GEO_FILE = "gis-wards.geojson"
)
func main() {
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	//read cmdline
	var configFile string
	flag.StringVar(&configFile, "configFile", "config.yml", "Provid config file path,  e.g c:/my/dir/eftconf.yml")
	flag.Parse()
	//Read config
// 	var config Config
// 	if err := config.readConfig(configFile); err != nil {
// 		log.Fatalf("Error reading config file :", configFile, err)
// 	} else {
// 		log.Println("Config: ", config)
// 		log.Println("Check log file for details :", config.AppConfig.LumberjackLogConfig.Filename)
// 	}
	// Load in our geojson file into a feature collection
	b, _ := ioutil.ReadFile(GEO_FILE)
	wardFeatures, _ := geojson.UnmarshalFeatureCollection(b)

	p1 := orb.Point{-79.2498779296875,43.89195472686543}
	getWards(wardFeatures, p1)
	p2 := orb.Point{-79.36248779296874,43.847403373019226}
	getWards(wardFeatures, p2)
	//helloHandler := func(w http.ResponseWriter, req *http.Request) {
	//	io.WriteString(w, "Hello, world!\n")
	//}
	//
	//http.HandleFunc("/hello", helloHandler)
	//log.Println("Listing for requests at http://localhost:8000/hello")
	//log.Fatal(http.ListenAndServe(":8000", nil))
}//main

func getWards(fc *geojson.FeatureCollection, point orb.Point) bool {
	for _, feature := range fc.Features {
		// Try on a MultiPolygon to begin
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if isMulti {
			if planar.MultiPolygonContains(multiPoly, point) {

				return true
			}
		} else {
			// Fallback to Polygon
			polygon, isPoly := feature.Geometry.(orb.Polygon)
			if isPoly {
				if planar.PolygonContains(polygon, point) {
					//log.Println("Polygon has points" , polygon)
					log.Println("Point :" , point)
					log.Println("At Ward:" , feature.Properties)
					return true
				}
			}
		}
	}//for
	return false
}
