package database

import (
  "fmt"
  "database/sql"
)

type Activity struct {
  activity_id int
  spots int
  activity_type string
  room string
  speaker string
  topic string
  description string
  time string
  day int
}

func GetAllActivities() (activities []Activity, err error) {
  query := `
  SELECT *
  FROM activities
  `

  rows, err := DB.Query(query)
  if err != nil {
    return nil, fmt.Errorf("could not retrieve activities: %v", err)
  }
  defer rows.Close()

  for rows.Next() {
    var a Activity
    err := rows.Scan(&a.activity_id, &a.spots, &a.activity_type, &a.room, &a.speaker, &a.topic, &a.description, &a.time, &a.day)
    if err != nil {
      return nil, fmt.Errorf("could not scan activity: %v", err)
    }
    activities = append(activities, a)
  }

  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("row iteration error: %v", err)
  }

  return activities, nil
}

func GetActivity(id int) (a Activity, err error) {
  query := `
  SELECT *
  FROM activities
  WHERE activities.id = $1
  `

  err = DB.QueryRow(query, id).Scan(&a.activity_id, &a.spots, &a.activity_type, &a.room, &a.speaker, &a.topic, &a.description, &a.time, &a.day)
  if err != nil {
    if err == sql.ErrNoRows {
      return a, fmt.Errorf("No activity found with id: %v\n", id)
    }
    return a, fmt.Errorf("could not retrieve activity: %v\n", err)
  }

  return a, nil
}

func (a Activity) String() string {
  return fmt.Sprintf("id: %v | spots: %v | day: %v | time: %v\nroom: %v | type: %v\nspeaker: %v | topic %v\ndescription: %v", 
    a.activity_id,
    a.spots,
    a.day,
    a.time,
    a.room,
    a.activity_type,
    a.speaker,
    a.topic,
    a.description,
    )
}
