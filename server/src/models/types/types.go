package types

import (
  "gopkg.in/mgo.v2/bson"
)

const (
  MALE = iota
  FEMALE
)

type Person struct {
  Name string                   `bson:"name" json:"name"`
  Gender int                    `bson:"gender" json:"gender"`
  Age int                       `bson:"age" json:"age"`
  Height float32                `bson:"height" json:"height"`
  Weight float32                `bson:"weight" json:"weight"`
  HeartRate int                 `bson:"heartrate" json:"heartrate"`
  MaxHeartRate int              `bson:"maxheartrate" json:"maxheartrate"`
}

func (p *Person) Valid() bool {
  return p.Name != "" && (p.Gender == MALE || p.Gender == FEMALE) &&
    p.Age > 0 && p.Height > 0 && p.Weight > 0 && p.HeartRate > 0 &&
    p.MaxHeartRate > 0
}

type Remark struct {
  Speed float64                 `bson:"speed" json:"speed"`
  HeartRate int                 `bson:"heartrate" json:"heartrate"`
  /**
  more remark criterion
  */
}

type Player struct {
  ObjId bson.ObjectId            `bson:"_id,omitempty" json:"_id"`
  Name string                    `bson:"name" json:"name"`
  DetailInfo Person              `bson:"detailinfo" json:"detailinfo"`
  OverallRemark Remark           `bson:"overallremark" json:"overallremark"`
  History string                 `bson:"history" json:"history"`
}

func (p *Player) Valid() bool {
  if p.ObjId.Hex() == "" {
    return p.Name != "" && p.DetailInfo.Valid()
  } else {
    return p.Name != "" && p.DetailInfo.Valid() && p.History != ""
  }
}

func (p1 *Player) Equals(p2 *Player) bool {
  return p1.Name == p2.Name &&
    p1.DetailInfo == p2.DetailInfo && p1.OverallRemark == p2.OverallRemark &&
    p1.History == p2.History
}

type TrainDesc struct {
  TimeStamp int64                 `bson:"timestamp" json:"timestamp"`
  Title string                    `bson:"title" json:"title"`
  Time string                     `bson:"time" json:"time"`
  Place string                    `bson:"place" json:"place"`
  Desc string                     `bson:"desc" json:"desc"`
}

type TrainRecord struct {
  ObjId bson.ObjectId             `bson:"_id,omitempty" json:"_id"`
  Desc TrainDesc                  `bson:"desc" json:"desc"`
  Speed []float64                 `bson:"speed" json:"speed"`
  Distance []float64              `bson:"distance" json:"distance"`
  HeartRate []int                 `bson:"heartrate" json:"heartrate"`
  /**
  more display criterion
  */
}

func (p1 *TrainRecord) Equals(p2 *TrainRecord) bool {
  if len(p1.Speed) != len(p2.Speed) ||
      len(p1.Distance) != len(p2.Distance) {
    return false
  }

  for i, _ := range p1.Speed {
    if p1.Speed[i] != p2.Speed[i] {
      return false
    }
  }

  for i, _ := range p1.Distance {
    if p1.Distance[i] != p2.Distance[i] {
      return false
    }
  }

  return true
}

/**
Raw train record
*/
type GPSData struct {
  Latitude float64                `bson:"latitude" json:"latitude"`
  Longitude float64               `bson:"longitude" json:"longitude"`
  Altitude float64                `bson:"altitude" json:"altitude"`
  Bearing float64                 `bson:"bearing" json:"bearing"`
  Speed float64                   `bson:"speed" json:"speed"`
  Accuracy float64                `bson:"accuracy" json:"accuracy"`
  Time int64                      `bson:"time" json:"time"`
}

type ACCData struct {
  XAcc float64                    `bson:"xacc" json:"xacc"`
  YAcc float64                    `bson:"yacc" json:"yacc"`
  ZAcc float64                    `bson:"zacc" json:"zacc"`
}

type GYROData struct {
  XGyro float64                   `bson:"xgyro" json:"xgyro"`
  YGyro float64                   `bson:"ygyro" json:"ygyro"`
  ZGyro float64                   `bson:"zgyro" json:"zgyro"`
}

type RawTrainRecord struct {
  ObjId bson.ObjectId             `bson:"_id,omitempty" json:"_id"`
  GpsData []GPSData               `bson:"gpsdata" json:"gpsdata"`
  AccData []ACCData               `bson:"accdata" json:"accdata"`
  GyroData []GYROData             `bson:"gyrodata" json:"gyrodata"`
  HeartRateData []int             `bson:"heartratedata" json:"heartratedata"`
}

func (p1 *RawTrainRecord) Equals(p2 *RawTrainRecord) bool {
  if len(p1.GpsData) != len(p2.GpsData) ||
      len(p1.AccData) != len(p2.AccData) || len(p1.GyroData) != len(p2.GyroData) ||
      len(p1.HeartRateData) != len(p2.HeartRateData) {
    return false
  }

  for i, _ := range p1.GpsData {
    if p1.GpsData[i] != p2.GpsData[i] {
      return false
    }
  }

  for i, _ := range p1.AccData {
    if p1.AccData[i] != p2.AccData[i] {
      return false
    }
  }

  for i, _ := range p1.GyroData {
    if p1.GyroData[i] != p2.GyroData[i] {
      return false
    }
  }

  for i, _ := range p1.HeartRateData {
    if p1.HeartRateData[i] != p2.HeartRateData[i] {
      return false
    }
  }

  return true
}

/**
query result
*/
type QueryResult struct {
  Result interface{}  `json:"result"`
  TotalNum int        `json:"total_number"`
  Before int          `json:"before"`        //-1 means this is the first page
  Current int         `json:"current"`
  Next int            `json:"next"`          //-1 means this is the last page
}

type AuthInfo struct {
  Username string     `json:"username"`
  Password string     `json:"password"`
}

func (a *AuthInfo) Valid() bool {
  return a.Username != "" && a.Password != ""
}
