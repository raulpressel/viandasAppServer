package handlers

import (
	"strconv"
	"viandasApp/dtos"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

// Poligonos de las zonas de envio. Es la unica definicion: el front los obtiene de
// /app/setting/getZonePolygons para validar el domicilio en el mapa.
// IDZone tiene que coincidir con el ID de la tabla zones.
var ZonePolygons = []dtos.ZonePolygon{
	{
		IDZone: 1,
		Paths: []dtos.LatLng{
			{Lat: -31.7963851, Lng: -60.5295354},
			{Lat: -31.7166208, Lng: -60.5125489},
			{Lat: -31.7183366, Lng: -60.5202307},
			{Lat: -31.7183366, Lng: -60.5282344},
			{Lat: -31.7167303, Lng: -60.5332984},
			{Lat: -31.7140288, Lng: -60.5354442},
			{Lat: -31.7047555, Lng: -60.5556144},
			{Lat: -31.703295, Lng: -60.5693473},
			{Lat: -31.7114732, Lng: -60.5779304},
			{Lat: -31.7856285, Lng: -60.5961265},
			{Lat: -31.7963851, Lng: -60.5295354},
		},
	},
	{
		IDZone: 2,
		Paths: []dtos.LatLng{
			{Lat: -31.7166208, Lng: -60.5125489},
			{Lat: -31.7963851, Lng: -60.5295354},
			{Lat: -31.8027336, Lng: -60.4655275},
			{Lat: -31.7833647, Lng: -60.450078},
			{Lat: -31.7055031, Lng: -60.4336066},
			{Lat: -31.7025822, Lng: -60.4740328},
			{Lat: -31.7021076, Lng: -60.4826588},
			{Lat: -31.7040792, Lng: -60.4927868},
			{Lat: -31.7070001, Lng: -60.5009837},
			{Lat: -31.7105415, Lng: -60.5061335},
			{Lat: -31.7166208, Lng: -60.5125489},
		},
	},
}

// GetZoneByLatLng devuelve la zona que contiene el punto. Devuelve false si las
// coordenadas no son validas o si el punto esta fuera de todas las zonas.
func GetZoneByLatLng(lat string, lng string) (int, bool) {

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return 0, false
	}

	longitude, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		return 0, false
	}

	point := orb.Point{longitude, latitude}

	for _, zone := range ZonePolygons {

		ring := make(orb.Ring, len(zone.Paths))
		for i, p := range zone.Paths {
			ring[i] = orb.Point{p.Lng, p.Lat}
		}

		if planar.PolygonContains(orb.Polygon{ring}, point) {
			return zone.IDZone, true
		}
	}

	return 0, false
}
