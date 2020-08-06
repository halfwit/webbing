package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type ApptRequest struct {
	Block  time.Time
	Doctor *Doctor
	User   *User
	Booked bool
	Token  string
	ID     uint
}

type ApptResponse struct {
	ID     int
	Block  time.Time
	Doctor *Doctor
}

// We use the nickname to guard against id errors
const bookAppt = `UPDATE appointments
SET client_id = $1
WHERE id = $2 AND username = $3
`

func (d *Database) BookAppointment(ctx context.Context, apt *ApptRequest) error {
	defer delete(d.valid, apt.Token)

	t, ok := d.valid[apt.Token]
	if !ok || t.Before(time.Now()) {
		return errors.New("request token invalid")
	}

	_, err := d.db.ExecContext(ctx, bookAppt, apt.User.id, apt.ID, apt.Doctor.Username)
	return err
}

const apptAdd = `INSERT into appointments
(username, date, time) values ($1, $2, $3)
`

// Add an appointment a client can purchase
func (d *Database) AddAppointment(ctx context.Context, apt *ApptRequest) error {
	defer delete(d.valid, apt.Token)

	t, ok := d.valid[apt.Token]
	if !ok || t.Before(time.Now()) {
		return errors.New("request token invalid")
	}

	year, month, day := apt.Block.Date()
	hour, min, _ := apt.Block.Clock()

	if min%15 > 0 {
		return errors.New("blocks must be 15 minute sessions, starting at either 0, 15, 30, or 45 past the hour")
	}

	date := fmt.Sprintf("%d-%d-%d", year, month, day)
	when := fmt.Sprintf("%d:%d:00", hour, min)

	_, err := d.db.ExecContext(ctx, apptAdd, apt.Doctor.Username, date, when)
	return err
}

const findAvail = `SELECT date, time, id FROM appointments
WHERE username = $1
`

// From a search for doctors, this will be shown on the purchase a block page
func (d *Database) ListAvailable(ctx context.Context, apt *ApptRequest) []*ApptResponse {
	var appts []*ApptResponse

	defer delete(d.valid, apt.Token)

	t, ok := d.valid[apt.Token]
	if !ok || t.Before(time.Now()) {
		return nil
	}

	rows, err := d.db.QueryContext(ctx, findAvail, apt.Doctor.Username)
	if err != nil {
		return nil
	}

	for rows.Next() {
		var tie, date time.Time

		appt := ApptResponse{}

		if e := rows.Scan(&date, &tie, &appt.ID); e != nil {
			break
		}

		date.Add(time.Duration(tie.Nanosecond()))
		appt.Block = date

		appts = append(appts, &appt)
	}

	return appts
}

const findClientBookings = `SELECT a.date, a.time, a.id, a.username FROM appointments a
JOIN users on a.client_id = users.id 
WHERE users.name = $1
`

// Show any upcoming bookings on a user profile
func (d *Database) ListClientBookings(ctx context.Context, user *User) []*ApptResponse {
	var appts []*ApptResponse

	defer delete(d.valid, user.Token)

	t, ok := d.valid[user.Token]
	if !ok || t.Before(time.Now()) {
		return nil
	}

	rows, err := d.db.QueryContext(ctx, findClientBookings, user.Username)
	if err != nil {
		log.Print(err)
		return nil
	}

	for rows.Next() {
		var date, tie time.Time

		appt := ApptResponse{}
		if e := rows.Scan(&date, &tie, &appt.ID); e != nil {
			log.Print(e)
			break
		}

		date.Add(time.Duration(tie.Nanosecond()))
		appt.Block = date

		appts = append(appts, &appt)
	}

	return appts
}

const findDoctorBookings = `SELECT a.date, a.time, a.id, a.username FROM appointments a
JOIN doctors on a.doctor_id = doctors.id 
WHERE doctors.nickname = $1
`

func (d *Database) ListDoctorBookings(ctx context.Context, doctor *Doctor) []*ApptResponse {
	var appts []*ApptResponse

	defer delete(d.valid, doctor.Token)

	t, ok := d.valid[doctor.Token]
	if !ok || t.Before(time.Now()) {
		return nil
	}

	rows, err := d.db.QueryContext(ctx, findDoctorBookings, doctor.Username)
	if err != nil {
		return nil
	}

	for rows.Next() {
		var date, tie time.Time

		appt := ApptResponse{}
		if e := rows.Scan(&date, &tie, &appt.ID); e != nil {
			break
		}

		date.Add(time.Duration(tie.Nanosecond()))
		appt.Block = date

		appts = append(appts, &appt)
	}

	return appts
}
