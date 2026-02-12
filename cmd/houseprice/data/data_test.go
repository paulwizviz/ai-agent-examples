package data

import (
	"aiagent/internal/sqlops"
	"bytes"
	"fmt"
	"reflect"
	"time"
)

func Example_parseDate() {

	result, _ := ParseDate("1997-09-26")
	fmt.Println(reflect.DeepEqual(result, time.Date(1997, time.Month(9), 26, 0, 0, 0, 0, time.UTC)))

	// Output:
	// true
}

func Example_parsePrice() {
	price, _ := ParsePrice("275000")
	fmt.Println(price)

	// Output:
	// 275000
}

func Example_processCSV() {
	content := []byte(`unique_id,price_paid,deed_date,postcode,property_type,new_build,estate_type,saon,paon,street,locality,town,district,county,transaction_category,linked_data_uri
"C78C5077-AC66-4D54-B903-C0EA2921E449","187500","1999-08-13","AL1 1DH","T","N","F","","78","HOLYWELL HILL","","ST. ALBANS","ST ALBANS","HERTFORDSHIRE","A","http://landregistry.data.gov.uk/data/ppi/transaction/C78C5077-AC66-4D54-B903-C0EA2921E449/current"`)
	reader := bytes.NewReader(content)

	result, _ := ProcessCSV(reader)
	fmt.Println(result)

	// Output:
	// [{C78C5077-AC66-4D54-B903-C0EA2921E449 187500 0001-01-01 00:00:00 +0000 UTC AL1 1DH T N F  78 HOLYWELL HILL  ST. ALBANS ST ALBANS HERTFORDSHIRE A http://landregistry.data.gov.uk/data/ppi/transaction/C78C5077-AC66-4D54-B903-C0EA2921E449/current}]
}

func Example_createTable() {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := CreateTable(db); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Table Created")

	data := []House{
		{
			UniqueID: "1",
			Price:    uint64(1000),
			DeedDate: func() time.Time {
				date, _ := ParseDate("2026-1-10")
				return date
			}(),
			PostCode:     "Post Code",
			PropertyType: "Property Type",
			NewBuild:     "New Build",
			EstateType:   "Estate Type",
			SAON:         "saon",
			PAON:         "paon",
			Street:       "street",
			Locality:     "locality",
			Town:         "town",
			District:     "district",
			County:       "county",
			Category:     "category",
			URI:          "url",
		},
	}

	if err := PersistData(db, data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data persisted")

	results, err := ListAll(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(results)

	// Output:
	// Table Created
	// Data persisted
	// [{1 1000 2026-01-10 00:00:00 +0000 UTC Post Code Property Type New Build Estate Type saon paon street locality town district county category url}]
}
