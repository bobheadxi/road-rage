package tomtom

type tomTomResp struct {
	FlowSegmentData FlowSegmentData `json:"flowSegmentData"`
}

type FlowSegmentData struct {
	CurrentSpeed       float64     `json:"currentSpeed"`
	FreeFlowSpeed      float64     `json:"freeFlowSpeed"`
	CurrentTravelTime  float64     `json:"currentTravelTime"`
	FreeFlowTravelTime float64     `json:"freeFlowTravelTime"`
	Confidence         float64     `json:"confidence"`
	Coordinates        Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Points []Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

/*
<flowSegmentData xmlns="http://lbs.tomtom.com/services" version="1.0.21-mascoma">
    <frc>FRC2</frc>
    <currentSpeed>41</currentSpeed>
    <freeFlowSpeed>70</freeFlowSpeed>
    <currentTravelTime>153</currentTravelTime>
    <freeFlowTravelTime>90</freeFlowTravelTime>
    <confidence>0.59</confidence>
    <coordinates>
        <coordinate>
            <latitude>52.40476</latitude>
            <longitude>4.844318</longitude>
        </coordinate>
        <coordinate>
            <latitude>52.411312</latitude>
            <longitude>4.8299975</longitude>
        </coordinate>
        <coordinate>
            <latitude>52.415073</latitude>
            <longitude>4.827327</longitude>
        </coordinate>
    </coordinates>
</flowSegmentData>
*/
