package punksranking

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var attributes = make(map[string]int)

var attIdToPunkId = make(map[string]int)

// 38%
func parsePercent(s string) (result float64) {
	if len(s) == 0 {
		return
	}
	if i := strings.Index(s, "%"); i != -1 {
		s = s[:i]
	}
	result, _ = strconv.ParseFloat(s, 64)
	return
}

func fixString(s string) string {
	if s == "(blank)" {
		return ""
	}
	return s
}

func fixInt(s string) int {
	if result, err := strconv.Atoi(s); err != nil {
		return 0
	} else {
		return result
	}
}

func Zap() {

	statements["delete_punk_att"].Exec()
	statements["delete"].Exec()
	statements["delete_att"].Exec()
}

func addAttributes(record ...string) {
	for i := range record {
		key := fixString(record[i])
		if _, exists := attributes[key]; exists {
			attributes[key]++
		} else if key != "" {
			attributes[key] = 1
		}
	}
}

// Calculate calculates all the scores
func Calculate() error {
	_, err := statements["calculate"].Exec()
	return err
}

// Link links the punks to attributes by linking
func Link() error {

	/*
		_, err := statements["insert_punk_attributes_type"].Exec()
		if err != nil {
			return err
		}
		_, err = statements["insert_punk_attributes_skin"].Exec()
		if err != nil {
			return err
		}

	*/
	_, err := statements["insert_punk_attributes_att1"].Exec()
	if err != nil {
		return err
	}
	_, err = statements["insert_punk_attributes_att2"].Exec()
	if err != nil {
		return err
	}
	_, err = statements["insert_punk_attributes_att3"].Exec()
	if err != nil {
		return err
	}
	_, err = statements["insert_punk_attributes_att4"].Exec()
	if err != nil {
		return err
	}
	_, err = statements["insert_punk_attributes_att5"].Exec()
	if err != nil {
		return err
	}
	_, err = statements["insert_punk_attributes_att6"].Exec()
	if err != nil {
		return err
	}
	_, err = statements["insert_punk_attributes_att7"].Exec()
	if err != nil {
		return err
	}
	/*
		_, err = statements["insert_punk_attribute_attributes"].Exec()
		if err != nil {
			return err
		}

	*/

	if _, err = db.Exec(
		`CREATE TEMPORARY TABLE new_tbl select count(*) as total , category from punks group by category;`,
	); err != nil {
		return err
	}

	if _, err = db.Exec(
		`UPDATE punks SET category_score = (select total from new_tbl where category =  punks.category) / 100 ;`,
	); err != nil {
		return err
	}

	if _, err = db.Exec(
		`DROP TEMPORARY TABLE new_tbl;`,
	); err != nil {
		return err
	}

	return err
}

// ImportAttr inserts the attributes
func ImportAttr() error {
	id := 1
	if len(attributes) > 0 {
		for key, val := range attributes {
			_, err := statements["insert_att"].Exec(
				id,
				key,
				val,
			)
			if err != nil {
				log.Print("insert_att err", id, key, val)
				return err
			}
			attIdToPunkId[key] = id
			id++

		}
	}

	if _, err := statements["insert_count_att"].Exec(); err != nil {
		log.Print("insert_count_att insert err")
		return err
	}

	return nil
}

func Import(file string) error {
	fin, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fin.Close()
	r := csv.NewReader(fin)
	count := 0
	for ; ; count++ {
		record, err := r.Read()
		if count == 0 {
			// skip headers
			continue
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		category := buildCategory(record)

		record[1] = ""

		addAttributes(record[6:12]...) // att1 to att7
		//addAttributes(record[3]) // skin
		//addAttributes(record[3:4]...) //  TSkin

		addAttributes(category) // synthetic attribute

		_, err = statements["insert"].Exec(
			record[0],                // id
			fixString(record[1]),     // sex (not used, already in the type field)
			record[2],                // type
			record[3],                // skin
			record[4],                // type_skin
			record[5],                // slots
			fixString(record[6]),     // arr1
			fixString(record[7]),     // arr2
			fixString(record[8]),     // att3
			fixString(record[9]),     // att4
			fixString(record[10]),    // arr5
			fixString(record[11]),    // att6
			fixString(record[12]),    // att7
			parsePercent(record[13]), // type_rare
			fixInt(record[14]),       // att_count
			parsePercent(record[15]), // att1_score
			parsePercent(record[16]), // att2_score
			parsePercent(record[17]), // att3_score
			parsePercent(record[18]), // att4_score
			parsePercent(record[19]), // att5_score
			parsePercent(record[10]), // att6_score
			parsePercent(record[20]), // att7_score
			parsePercent(record[21]), // min
			parsePercent(record[22]), // avg
			parsePercent(record[23]), // rank
			buildCategory(record),
		)
		if err != nil {
			return err
		}

		fmt.Println(record)
	}
	return nil

}

func buildCategory(record []string) string {
	human := "x"
	nonhumanType := "xx"
	skin := "xxx"
	sex := "x"
	// fix up skin
	if record[2] == "Alien" {
		//	record[3] = ""
		nonhumanType = "AL"
	} else if record[2] == "Ape" {
		//record[3] = ""
		nonhumanType = "AP"
	} else if record[2] == "Zombie" {
		//	record[3] = ""
		nonhumanType = "ZO"
	} else {
		human = "H"
		skin = record[3]
		if record[1] == "Girl" {
			sex = "F"
		} else {
			sex = "M"
		}
	}

	// CATHMxx

	switch record[3] {
	case "Albino":
		skin = "Alb"
	case "Light":
		skin = "lit"
	case "Mid":
		skin = "Mid"
	case "Dark":
		skin = "Drk"

	}

	if nonhumanType == "" {
		log.Print("nuill")
	}

	return "CAT" + human + sex + skin + nonhumanType

}
