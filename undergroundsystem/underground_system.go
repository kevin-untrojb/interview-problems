package undergroundsystem
// from https://leetcode.com/problems/design-underground-system

type UndergroundSystem struct {
	currentCheckInMap map[int]CheckInStationTime
	AverageTimeMap    map[string]map[string]AverageUnit
}
type CheckInStationTime struct{
	Station string
	Time int
}
type AverageUnit struct{
	Average float64
	Amount int
}


func Constructor() UndergroundSystem {
	return UndergroundSystem{
		make (map[int]CheckInStationTime),
		make(map[string]map[string]AverageUnit),
	}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	if _,ok := this.currentCheckInMap[id]; ok{
		return
	}
	this.currentCheckInMap[id] = CheckInStationTime{stationName,t}

}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	CheckInStationTime,ok := this.currentCheckInMap[id]
	if !ok{
		return
	}
	delete(this.currentCheckInMap, id)

	totalTime := t-CheckInStationTime.Time
	destinationMap,ok := this.AverageTimeMap[CheckInStationTime.Station]
	if !ok{
		this.AverageTimeMap[CheckInStationTime.Station] = make(map[string]AverageUnit)
		this.AverageTimeMap[CheckInStationTime.Station][stationName]=AverageUnit{float64(totalTime),1}
		return
	}
	currentAvarage,ok :=destinationMap[stationName]
	if !ok{
		this.AverageTimeMap[CheckInStationTime.Station][stationName]=AverageUnit{float64(totalTime),1}
		return
	}
	newAverage:= currentAvarage.Average*float64(currentAvarage.Amount) + float64(totalTime)
	newAverage = newAverage/float64(currentAvarage.Amount+1)
	this.AverageTimeMap[CheckInStationTime.Station][stationName]=AverageUnit{newAverage,currentAvarage.Amount+1}
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	destinationMap,ok := this.AverageTimeMap[startStation]
	if !ok{
		return 0
	}
	currentAvarage,ok :=destinationMap[endStation]
	if !ok{
		return 0
	}
	return currentAvarage.Average
}
