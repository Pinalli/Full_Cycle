package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)
type Route struct{
	ID string 'json:"RouteId"'
	ClinteID string 'json:"ClientId"'
	Positions []Positions 'json:"Position"'
}

type Positions struct{
	Lat float64 'json:"Lat"'
	long float64 'json:"Long"'
}

func (r *Route) LoadPositions() error {
	if(r.ID == ""){
		return errors.New("ID is required")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if(err != nil){
		return err
	}

	defer f.close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if(err != nil){
			return nil
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if(err != nil){
			return nil		
	}
	r.Positions = append(r.Positions, Positions{
	 Lat: lat,
	 Long: long
	})	
}
	return nil
}

type PartRoutePosition struct  {
	ID string 'json:"RouteId"'
	ClintID string  'json:"ClientId"'
	Position []float64 'json:"Position"'
	finished bool 'json:"Finished"'
}

func (r *Route) ExportJsonPositions()(string, err){
	var route PartialRoutePosition
	var result []string
	total:= len(r.Positions)

	for k, v := range r.Positions{
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		json, err := json.Marshal(route)
		if(total-1 == k){
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if(err != nil){ 
		     return nil, err
	   }
	   result = append(result, string(jsonRoute))
	}
	return result, nil
}
