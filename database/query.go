package database

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Aux int

const (
	AuxNone Aux = iota
	AuxEmail
	AuxWage
)

var aux = map[Aux]string{
	AuxNone:  "",
	AuxEmail: "d.email",
	AuxWage:  "d.wage",
}

// Query token must be a valid query token
// These are issued by the plugin manager
type Query struct {
	Service string
	Country string
	Query   Aux
	Sort    string
	Token   string
}

// Doctor is a single result from a query
// This likely will be populated with more fields soon
type Doctor struct {
	ID       int
	First    string
	Last     string
	Username string
	// TODO(halfwit) Add Rate field
	//Rate     float32
	Aux interface{}
}

// This can likely be cleaned up now
const queryAux = `
SELECT d.first_name, d.last_name, d.id, `

const queryDoctorStart = `
SELECT d.first_name, d.last_name, d.id`

const queryDoctorAll = `
FROM doctors d`

const queryDoctorNiche = `
FROM doctors d
JOIN doctor_specialties ds ON ds.doctor_id = d.id
JOIN doctor_countries dc ON dc.doctor_id = d.id
JOIN specialties s on s.id = ds.specialty_id
JOIN countries c on c.id = dc.country_id
WHERE s.specialty_name = ?
AND c.official_name = ?`

const queryDoctorCountry = `
FROM doctors d
JOIN doctor_countries dc ON dc.doctor_id = d.id
JOIN countries c ON dc.country_id = c.id
WHERE c.official_name = ?`

const queryDoctorSpecialty = `
FROM doctors d
JOIN doctor_specialties ds ON ds.doctor_id = d.id
JOIN specialties s on s.id = ds.specialty_id
WHERE s.specialty_name = ?`

const queryDoctorEnd = `
ORDER BY d.last_name;`

func (d *Database) RunQuery(ctx context.Context, query *Query) chan *Doctor {
	// Unabashedly blow up this token regardless, if it exists
	defer delete(d.valid, query.Token)

	doctor := make(chan *Doctor)

	// Ensure we have a good token
	t, ok := d.valid[query.Token]
	if !ok || t.Before(time.Now()) {
		defer close(doctor)
		return doctor
	}

	// Make sure we sanitize inputs
	if _, ok := aux[query.Query]; !ok {
		query.Query = AuxNone
	}

	if query.Sort == "" {
		query.Sort = "doctors.last_name"
	}

	go func(ctx context.Context, query *Query, doctor chan *Doctor) {
		hasaux, format, args := buildArgs(ctx, query)

		for doc := range d.listDoctors(ctx, hasaux, format, args...) {
			doctor <- doc
		}

		close(doctor)
	}(ctx, query, doctor)

	return doctor
}

func buildArgs(ctx context.Context, query *Query) (bool, string, []string) {
	var hasaux bool
	var format string
	var args []string

	switch {
	// Specific query
	case query.Query != AuxNone:
		hasaux = true

		switch {
		case query.Service != "" && query.Country != "":
			format = queryAux + aux[query.Query] + queryDoctorNiche + queryDoctorEnd
			args = []string{query.Service, query.Country}
		case query.Service != "":
			format = queryAux + aux[query.Query] + queryDoctorSpecialty + queryDoctorEnd
			args = []string{query.Service}
		case query.Country != "":
			format = queryAux + aux[query.Query] + queryDoctorCountry + queryDoctorEnd
			args = []string{query.Country}
		default:
			format = queryAux + aux[query.Query] + queryDoctorAll + queryDoctorEnd
		}
	case query.Service != "" && query.Country != "":
		format = queryDoctorStart + queryDoctorNiche + queryDoctorEnd
		args = []string{query.Service, query.Country}
	case query.Service != "":
		format = queryDoctorStart + queryDoctorSpecialty + queryDoctorEnd
		args = []string{query.Service}
	case query.Country != "":
		format = queryDoctorStart + queryDoctorCountry + queryDoctorEnd
		args = []string{query.Country}
	default:
		format = queryDoctorStart + queryDoctorAll + queryDoctorEnd
	}

	return hasaux, format, args
}

func (d *Database) listDoctors(ctx context.Context, hasaux bool, query string, args ...string) chan *Doctor {
	docchan := make(chan *Doctor)
	go func(docchan chan *Doctor) {
		query = sanitizeQuery(query, len(args))

		rows, err := d.db.QueryContext(ctx, query, strings.Join(args, " "))
		if err != nil {
			close(docchan)
			return
		}

		for rows.Next() {
			d := Doctor{}
			var err error

			if hasaux {
				err = rows.Scan(&d.First, &d.Last, &d.ID, &d.Aux)
			} else {
				err = rows.Scan(&d.First, &d.Last, &d.ID)
			}

			if err != nil {
				close(docchan)
				return
			}

			docchan <- &d
		}

		close(docchan)
	}(docchan)

	return docchan
}

func sanitizeQuery(query string, n int) string {
	for i := 1; i <= n; i++ {
		n := strings.IndexByte(query, '?')
		if n < 0 {
			break
		}
		query = fmt.Sprintf("%s$%d%s", query[:n], i, query[n+1:])
	}

	return query
}
