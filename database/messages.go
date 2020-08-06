package database

import (
	"context"
	"log"
	"time"
)

type Thread struct {
	ID     int
	Doctor *Doctor
}

type Message struct {
	Content string
	Sent    time.Time
}

const clientThreadList = `
SELECT d.first_name, d.last_name, d.nickname, mt.id FROM doctors d
JOIN message_thread mt on d.id = mt.doctor_id
JOIN messages m on mt.msg_id = m.id
WHERE mt.client_id = $1
ORDER BY datetime
`

// There is no inter-client messages, only client to doctor/doctor to client.
func (d *Database) ListClientThreads(ctx context.Context, user *User) []*Thread {
	var threadlist []*Thread

	defer delete(d.valid, user.Token)

	// validate token
	t, ok := d.valid[user.Token]
	if !ok || t.Before(time.Now()) {
		return threadlist
	}

	rows, err := d.db.QueryContext(ctx, clientThreadList, user.id)
	if err != nil {
		log.Print(err)
		return nil
	}

	for rows.Next() {
		var id int

		doc := &Doctor{}
		if e := rows.Scan(&doc.First, &doc.Last, &doc.Username, &id); e != nil {
			break
		}

		thread := &Thread{
			Doctor: doc,
			ID:     id,
		}

		threadlist = append(threadlist, thread)
	}

	return threadlist
}

const listMessages = `
SELECT mt.datetime, m.content 
FROM doctors d
JOIN message_thread mt on d.id = mt.doctor_id
JOIN users u on mt.client_id = u.id
JOIN messages m on mt.msg_id = m.id
WHERE d.nickname = $1
AND mt.client_id = $2
ORDER BY datetime
`

func (d *Database) ListThread(ctx context.Context, user *User, doctor *Doctor) []*Message {
	var msglist []*Message

	defer delete(d.valid, user.Token)

	// validate token
	t, ok := d.valid[user.Token]
	if !ok || t.Before(time.Now()) {
		return msglist
	}

	rows, err := d.db.QueryContext(ctx, listMessages, doctor.Username, user.id)
	if err != nil {
		log.Print(err)
		return nil
	}

	for rows.Next() {
		var content string
		var datetime time.Time
		if e := rows.Scan(&datetime, &content); e != nil {
			break
		}

		msg := &Message{
			Content: content,
			Sent:    datetime,
		}

		msglist = append(msglist, msg)
	}

	return msglist
}

const messageSend = `
INSERT into messages(content)
VALUES($1)
`

const messageLink = `INSERT into message_thread (doctor_id, client_id, msg_id, datetime) values ($1, $2, $3, $4)`

func (d *Database) SendMessage(ctx context.Context, user *User, doctor *Doctor, msg *Message) error {
	defer delete(d.valid, user.Token)

	// validate token
	t, ok := d.valid[user.Token]
	if !ok || t.Before(time.Now()) {
		return nil
	}

	res, err := d.db.ExecContext(ctx, messageSend, msg.Content)
	if err != nil {
		return err
	}

	id, err := res.RowsAffected()
	if err != nil {
		return err
	}

	_, err = d.db.ExecContext(ctx, messageLink, doctor.ID, user.id, id, time.Now())
	return err
}
