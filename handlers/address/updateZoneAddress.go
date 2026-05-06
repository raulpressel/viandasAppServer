package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

func UpdateZoneAddress(rw http.ResponseWriter, r *http.Request) {

	jsonData := `{
		"type": "FeatureCollection",
		"features": [
		  {
			"type": "Feature",
			"properties": {
				"zona": "1"
			},
			"geometry": {
			  "coordinates": [
				[
				  [
					-60.54131767854109,
					-31.768541062265214
				  ],
				  [
					-60.52411287555728,
					-31.77128213176924
				  ],
				  [
					-60.517405256128626,
					-31.740052583733387
				  ],
				  [
					-60.51748163477768,
					-31.73829441486324
				  ],
				  [
					-60.51821138449719,
					-31.73616213109061
				  ],
				  [
					-60.517815479142726,
					-31.733058372480016
				  ],
				  [
					-60.53604483964561,
					-31.727701822563425
				  ],
				  [
					-60.53935240818477,
					-31.734528381108504
				  ],
				  [
					-60.53847071182254,
					-31.74192631321565
				  ],
				  [
					-60.53879413228731,
					-31.743617017362276
				  ],
				  [
					-60.535921397140584,
					-31.74377071620941
				  ],
				  [
					-60.54131767854109,
					-31.768541062265214
				  ]
				]
			  ],
			  "type": "Polygon"
			}
		  },
		  {
			"type": "Feature",
			"properties": {
				"zona": "2"
			},
			"geometry": {
			  "coordinates": [
				[
				  [
					-60.519478319365405,
					-31.77710162675495
				  ],
				  [
					-60.51173072818588,
					-31.740918758555182
				  ],
				  [
					-60.51698626215031,
					-31.738476127997927
				  ],
				  [
					-60.516244520147225,
					-31.73314529499612
				  ],
				  [
					-60.51863796835501,
					-31.731249054264744
				  ],
				  [
					-60.52984610169034,
					-31.72848902236008
				  ],
				  [
					-60.531932483369886,
					-31.726368028664282
				  ],
				  [
					-60.53658879470039,
					-31.724533030774445
				  ],
				  [
					-60.53888846459732,
					-31.72670114191768
				  ],
				  [
					-60.53853402021308,
					-31.728393743490663
				  ],
				  [
					-60.54101403014374,
					-31.733071292428278
				  ],
				  [
					-60.540423289503536,
					-31.735152762450333
				  ],
				  [
					-60.54170044411555,
					-31.74402585466631
				  ],
				  [
					-60.541087983099715,
					-31.75269035270761
				  ],
				  [
					-60.54476957331488,
					-31.773739578874512
				  ],
				  [
					-60.519478319365405,
					-31.77710162675495
				  ]
				]
			  ],
			  "type": "Polygon"
			}
		  },
		  {
			"type": "Feature",
			"properties": {
				"zona": "3"
			},
			"geometry": {
			  "coordinates": [
				[
				  [
					-60.531546902000755,
					-31.72019982075185
				  ],
				  [
					-60.53487219257771,
					-31.720877091815126
				  ],
				  [
					-60.53582053829096,
					-31.72304303975993
				  ],
				  [
					-60.53916557196061,
					-31.723075514073535
				  ],
				  [
					-60.54229742468294,
					-31.72286642431697
				  ],
				  [
					-60.544344266525286,
					-31.723715569787686
				  ],
				  [
					-60.548600018117725,
					-31.727298283812928
				  ],
				  [
					-60.553035272390204,
					-31.733044435588006
				  ],
				  [
					-60.54470967021457,
					-31.748569857125148
				  ],
				  [
					-60.55109297686924,
					-31.765970625260195
				  ],
				  [
					-60.55067083251052,
					-31.773973943838236
				  ],
				  [
					-60.549466399512255,
					-31.775970579934757
				  ],
				  [
					-60.51844620637226,
					-31.780338581227497
				  ],
				  [
					-60.5176231771559,
					-31.777625878058856
				  ],
				  [
					-60.51443208429497,
					-31.777216322524318
				  ],
				  [
					-60.51118076978382,
					-31.773446026370486
				  ],
				  [
					-60.50679114809377,
					-31.774128645752008
				  ],
				  [
					-60.490959051994395,
					-31.772836292761966
				  ],
				  [
					-60.47645491623686,
					-31.7599107069533
				  ],
				  [
					-60.4833182205676,
					-31.735184381564963
				  ],
				  [
					-60.488156026443335,
					-31.731052719707833
				  ],
				  [
					-60.497181910378046,
					-31.72826665124221
				  ],
				  [
					-60.531546902000755,
					-31.72019982075185
				  ]
				]
			  ],
			  "type": "Polygon"
			}
		  }
		]
	  }`

	// Crear una estructura para deserializar el JSON
	featureCollection := geojson.NewFeatureCollection()

	// Deserializar el JSON en la estructura utilizando la función UnmarshalJSON
	err := json.Unmarshal([]byte(jsonData), &featureCollection)
	if err != nil {
		fmt.Println("Error al deserializar el JSON:", err)
		return
	}

	//punto afuera de todo -31.826276538032122, -60.5191173949662
	//punto zona3 -31.767715423514908, -60.51518525249191

	longitude := -60.51518525249191
	latitude := -31.767715423514908

	// Crear un punto orb.Point con las coordenadas
	point := orb.Point{longitude, latitude}

	point.Lat()

	// Acceder al primer feature en el FeatureCollection
	if len(featureCollection.Features) > 0 {

		for _, feature := range featureCollection.Features {
			//firstFeature := featureCollection.Features[0]

			pp := feature.Properties

			zona := pp.MustString("zona")

			fmt.Println(zona)

			//test := feature.Geometry

			//valid := feature.Geometry.Bound().Contains(point)

			inside := planar.PolygonContains(feature.Geometry.Bound().ToPolygon(), point)

			poli := feature.Geometry.Bound().ToPolygon()

			valid := poli.Bound().Contains(point)

			fmt.Println(valid)

			//polygon.UnmarshalJSON([]byte()))

			//feature.UnmarshalJSON([]byte(polygon))

			/* tt := test.GeoJSONType()

			fmt.Println(tt) */

			fmt.Println(inside)
		}
	}

	/* point := Point{-60.52746073130738, -31.741579596859083}
	polygon := []Point{
		{-60.5366184, -31.7245728}, {-60.5408081, -31.7262413}, {-60.5384746, -31.7285973}, {-60.5410066, -31.7331447},
		{-60.5400087, -31.735195}, {-60.5404218, -31.7371554}, {-60.5409609, -31.7396503}, {-60.5417468, -31.7431761},
		{-60.5410494, -31.7557481}, {-60.5446543, -31.7738457}, {-60.5314364, -31.7751683}, {-60.5251278, -31.7761168},
		{-60.5199261, -31.7768502}, {-60.5175138, -31.7673679}, {-60.5166889, -31.7645694}, {-60.5164446, -31.762339},
		{-60.5161629, -31.7611395}, {-60.5158222, -31.759431}, {-60.5145948, -31.7549545}, {-60.5144203, -31.7533123},
		{-60.5141521, -31.7523042}, {-60.5138625, -31.7509311}, {-60.5133072, -31.7484804}, {-60.5130208, -31.7470215},
		{-60.5128991, -31.7462008}, {-60.5126486, -31.74538}, {-60.5121834, -31.7430826}, {-60.5117248, -31.7408791},
		{-60.5168396, -31.7385092}, {-60.5162967, -31.7331587}, {-60.5186967, -31.7312546}, {-60.5226847, -31.7303541},
		{-60.5263205, -31.7293698}, {-60.5298009, -31.7284764}, {-60.5320163, -31.7262501}, {-60.5346298, -31.7251958},
		{-60.5366184, -31.7245728},
	}

	isInside := pointInPolygon(point, polygon)

	fmt.Println(isInside) */

}

// Acceder a la geometría del feature
/*		geometry := firstFeature.Geometry

	// Verificar el tipo de geometría
	if geometry.Type() == geojson.GeometryTypePolygon {
		polygon, _ := geometry.Polygon()

		// Acceder a las coordenadas del polígono
		coordinates := polygon[0]

		// Imprimir las coordenadas del polígono
		fmt.Println("Coordenadas del polígono:", coordinates)
	}
}
*/

type Point struct {
	X float64
	Y float64
}

func pointInPolygon(point Point, polygon []Point) bool {
	x := point.X
	y := point.Y
	n := len(polygon)
	inside := false

	p1x := polygon[0].X
	p1y := polygon[0].Y
	for i := 0; i <= n; i++ {
		p2x := polygon[i%n].X
		p2y := polygon[i%n].Y

		if y > min(p1y, p2y) {
			if y <= max(p1y, p2y) {
				if x <= max(p1x, p2x) {
					if p1y != p2y {
						xinters := (y-p1y)*(p2x-p1x)/(p2y-p1y) + p1x
						if p1x == p2x || x <= xinters {
							inside = !inside
						}
					}
				}
			}
		}
		p1x, p1y = p2x, p2y
	}

	return inside
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
