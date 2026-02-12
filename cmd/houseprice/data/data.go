package data

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

type House struct {
	UniqueID     string
	Price        uint64
	DeedDate     time.Time
	PostCode     string
	PropertyType string
	NewBuild     string
	EstateType   string
	SAON         string
	PAON         string
	Street       string
	Locality     string
	Town         string
	District     string
	County       string
	Category     string
	URI          string
}

func ParseDate(date string) (time.Time, error) {

	yymmdd := strings.Split(date, "-")
	yy, err := strconv.Atoi(yymmdd[0])
	if err != nil {
		return time.Time{}, err
	}

	mm, err := strconv.Atoi(yymmdd[1])
	if err != nil {
		return time.Time{}, err
	}

	dd, err := strconv.Atoi(yymmdd[2])
	if err != nil {
		return time.Time{}, err
	}

	t := time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)

	return t, nil
}

func ParsePrice(price string) (uint64, error) {
	p, err := strconv.ParseUint(price, 10, 0)
	if err != nil {
		return 0, err
	}
	return p, nil
}

func ProcessCSV(reader *bytes.Reader) ([]House, error) {
	csvr := csv.NewReader(reader)
	csvr.Read() // Remove holder
	txns := []House{}
loop:
	for {
		rec, err := csvr.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Println(err)
			continue loop
		}
		house := House{}
		house.UniqueID = rec[0]
		p, err := ParsePrice(rec[1])
		if err != nil {
			continue loop
		}
		house.Price = p
		d, err := ParseDate(rec[2])
		if err != nil {
			house.DeedDate = d
		}
		house.PostCode = rec[3]
		house.PropertyType = rec[4]
		house.NewBuild = rec[5]
		house.EstateType = rec[6]
		house.SAON = rec[7]
		house.PAON = rec[8]
		house.Street = rec[9]
		house.Locality = rec[10]
		house.Town = rec[11]
		house.District = rec[12]
		house.County = rec[13]
		house.Category = rec[14]
		house.URI = rec[15]
		txns = append(txns, house)
	}
	return txns, nil
}

const (
	tblName      = "property_txn"
	uniqueID     = "id"
	price        = "price"
	deedDate     = "deed_date"
	postCode     = "postcode"
	propertyType = "property_type"
	newBuild     = "new_build"
	estateType   = "estate_type"
	saon         = "saon"
	paqn         = "paqn"
	street       = "street"
	locality     = "locality"
	town         = "town"
	district     = "district"
	county       = "county"
	category     = "category"
	uri          = "uri"
)

func CreateTable(db *sql.DB) error {
	rawSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %[1]s(
%[2]s TEXT,%[3]s TEXT,%[4]s TEXT,%[5]s TEXT,%[6]s TEXT,%[7]s TEXT,%[8]s TEXT,%[9]s TEXT,%[10]s TEXT,%[11]s TEXT,%[12]s TEXT,%[13]s TEXT,%[14]s TEXT,%[15]s TEXT,%[16]s TEXT,%[17]s TEXT)`, tblName, uniqueID, price, deedDate, postCode, propertyType, newBuild, estateType, saon, paqn, street, locality, town, district, county, category, uri)
	_, err := db.Exec(rawSQL)
	if err != nil {
		return err
	}
	return nil
}

func PersistData(db *sql.DB, data []House) error {
	rawSQL := fmt.Sprintf(`INSERT INTO %[1]s(
%[2]s,%[3]s,%[4]s,%[5]s,%[6]s,%[7]s,%[8]s,%[9]s,%[10]s,%[11]s,%[12]s,%[13]s,%[14]s,%[15]s,%[16]s,%[17]s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`, tblName, uniqueID, price, deedDate, postCode, propertyType, newBuild, estateType, saon, paqn, street, locality, town, district, county, category, uri)

	stmt, err := db.Prepare(rawSQL)
	if err != nil {
		return err
	}

	for _, d := range data {
		stmt.Exec(d.UniqueID,
			d.Price,
			d.DeedDate,
			d.PostCode,
			d.PropertyType,
			d.NewBuild,
			d.EstateType,
			d.SAON,
			d.PAON,
			d.Street,
			d.Locality,
			d.Town,
			d.District,
			d.County,
			d.Category,
			d.URI)
	}

	return nil
}

func ListAll(db *sql.DB) ([]House, error) {

	rawSQL := fmt.Sprintf("SELECT * FROM %s", tblName)
	stmt, err := db.Prepare(rawSQL)
	if err != nil {
		return nil, nil
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txns []House
	for rows.Next() {
		txn := House{}
		var deedDate string
		rows.Scan(&txn.UniqueID,
			&txn.Price,
			&deedDate,
			&txn.PostCode,
			&txn.PropertyType,
			&txn.NewBuild,
			&txn.EstateType,
			&txn.SAON,
			&txn.PAON,
			&txn.Street,
			&txn.Locality,
			&txn.Town,
			&txn.District,
			&txn.County,
			&txn.Category,
			&txn.URI)
		txn.DeedDate, err = time.Parse("2006-01-02 15:04:05 -0700 MST", deedDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		txns = append(txns, txn)
	}

	return txns, nil
}
