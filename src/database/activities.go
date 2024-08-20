package database

import (
  "fmt"
  "database/sql"

  _ "github.com/lib/pq"
)

type Activity struct {
  Activity_id int
  Spots int
  Activity_type string
  Room string
  Speaker string
  Topic string
  Description string
  Time string
  Day int
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
    err := rows.Scan(
      &a.Activity_id,
      &a.Spots,
      &a.Activity_type,
      &a.Room,
      &a.Speaker,
      &a.Topic,
      &a.Description,
      &a.Time,
      &a.Day,
    )

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

  err = DB.QueryRow(query, id).Scan(
    &a.Activity_id,
    &a.Spots,
    &a.Activity_type,
    &a.Room,
    &a.Speaker,
    &a.Topic,
    &a.Description,
    &a.Time,
    &a.Day,
  )

  if err != nil {
    if err == sql.ErrNoRows {
      return a, fmt.Errorf("No activity found with id: %v\n", id)
    }
    return a, fmt.Errorf("could not retrieve activity: %v\n", err)
  }

  return a, nil
}

func CreateActivity(a Activity) (int, error) {
  query := `
  INSERT INTO activities
  (spots, activity_type, room, speaker, topic, description, time, day)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING id
  `
  var id int
  err := DB.QueryRow(query, a.Spots, a.Activity_type, a.Room, a.Speaker, a.Topic, a.Description, a.Time, a.Day).Scan(&id)
  if err != nil {
    return 0, fmt.Errorf("could not create activity: %v", err)
  }
  return id, nil
}

func (a Activity) String() string {
  return fmt.Sprintf("id: %v | spots: %v | day: %v | time: %v\nroom: %v | type: %v\nspeaker: %v | topic %v\ndescription: %v",
    a.Activity_id,
    a.Spots,
    a.Day,
    a.Time,
    a.Room,
    a.Activity_type,
    a.Speaker,
    a.Topic,
    a.Description,
  )
}

