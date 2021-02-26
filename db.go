package punksranking

import (
	"database/sql"
)

type Store struct {
	Driver string
	DSN    string
}

var db *sql.DB

func SetupDB(configPath string) error {
	conf, err := LoadConfig(configPath)
	if err != nil {
		return err
	}
	s := &Store{
		Driver: conf.SQLDriver,
		DSN:    conf.SQLDSN,
	}

	db, err = s.connect()
	if err != nil {
		return err
	}

	err = s.prepareSQL()

	return err
}

func (s *Store) connect() (*sql.DB, error) {

	var err error
	if db, err = sql.Open(s.Driver, s.DSN); err != nil {

		return nil, err
	}
	// do we have permission to access the table?
	_, err = db.Query("SELECT `id` FROM  punks LIMIT 1")
	if err != nil {
		return nil, err
	}
	return db, err
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

var statements map[string]*sql.Stmt

func (s *Store) prepareSQL() error {

	if statements == nil {
		statements = make(map[string]*sql.Stmt)
	}

	if stmt, err := db.Prepare(
		"INSERT INTO punks (`id`, `sex`, `type`, `skin`, `type_skin`, `slots`, `att1`, `att2`, `att3`, `att4`, `att5`, `att6`, `att7`, `type_rare`, `att_count`, `att1_score`, `att2_score`, `att3_score`, `att4_score`, `att5_score`, `att6_score`, `att7_score`, `min`, `avg`, `rank`, `category`) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"); err != nil {
		return err
	} else {
		statements["insert"] = stmt
	}

	if stmt, err := db.Prepare("delete from punks;"); err != nil {
		return err
	} else {
		statements["delete"] = stmt
	}

	if stmt, err := db.Prepare(
		"INSERT INTO attributes (`id`, `name`, `total`) VALUES (?, ?, ?) "); err != nil {
		return err
	} else {
		statements["insert_att"] = stmt
	}

	if stmt, err := db.Prepare("delete from punk_attributes;"); err != nil {
		return err
	} else {
		statements["delete_punk_att"] = stmt
	}

	if stmt, err := db.Prepare("delete from attributes;"); err != nil {
		return err
	} else {
		statements["delete_att"] = stmt
	}

	// insert the attributes
	if stmt, err := db.Prepare(
		"INSERT INTO  punk_attributes (punk_id, attribute_id, score)  VAlUES (?, ?, ?)  "); err != nil {
		return err
	} else {
		statements["insert_pa"] = stmt
	}

	if stmt, err := db.Prepare(
		"select * from  punks "); err != nil {
		return err
	} else {
		statements["select_punks"] = stmt
	}

	// calculate the scores
	if stmt, err := db.Prepare(
		`update punks set skin_score = ((select total from attributes where name = skin)/100),	
							type_score = ((select total from attributes where name = type)/100), 
							skin_score = ((select total from attributes where name = skin)/100), 
							att1_score = ((select total from attributes where name = att1)/100),	
							att2_score = ((select total from attributes where name = att2)/100), 
							att3_score = ((select total from attributes where name = att3)/100),	
							att4_score = ((select total from attributes where name = att4)/100), 
							att5_score = ((select total from attributes where name = att5)/100), 
							att6_score = ((select total from attributes where name = att6)/100),	
							att7_score = ((select total from attributes where name = att7)/100),
							att_count_score = ((select total from attributes where name = CONCAT (slots, "Att"))/100)

WHERE 1`,
	); err != nil {
		return err
	} else {
		statements["calculate"] = stmt
	}

	// populates the attributes table with the 0Att, 1Att, 2Att, 3Att ... 7Att attributes
	// the "number of attributes" attribute is a meta-attribute
	if stmt, err := db.Prepare(`insert into attributes (id, name, total)
	select ROW_NUMBER() OVER (
		ORDER BY slots
		) + 98 as id, concat ( ROW_NUMBER() OVER (
			ORDER BY slots
		) -1 , 'Att' ) as name, count(*) as c
		from punks group by slots order by slots;`); err != nil {
		return err
	} else {
		statements["insert_count_att"] = stmt
	}
	// todo: instead of type, use skin?
	// populate the punk_attributes table
	if stmt, err := db.Prepare(
		`INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.skin = a.name AND p.id = punks.id), skin_score FROM punks;`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_skin"] = stmt
	}

	if stmt, err := db.Prepare(
		`INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.type = a.name AND p.id = punks.id), type_score FROM punks;`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_type"] = stmt
	}

	if stmt, err := db.Prepare(
		`INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.att1 = a.name AND p.id = punks.id), att1_score FROM punks;`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_att1"] = stmt
	}

	if stmt, err := db.Prepare(
		`
INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.att2 = a.name AND p.id = punks.id), att2_score FROM punks;

`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_att2"] = stmt
	}

	if stmt, err := db.Prepare(
		`
INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.att3 = a.name AND p.id = punks.id), att3_score FROM punks;

`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_att3"] = stmt
	}

	if stmt, err := db.Prepare(
		`
INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.att4 = a.name AND p.id = punks.id), att4_score FROM punks;

`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_att4"] = stmt
	}

	if stmt, err := db.Prepare(
		`
INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.att5 = a.name AND p.id = punks.id), att5_score FROM punks;
`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_att5"] = stmt
	}

	if stmt, err := db.Prepare(
		`
INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.att6 = a.name AND p.id = punks.id), att6_score FROM punks;
`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_att6"] = stmt
	}

	if stmt, err := db.Prepare(
		`
INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id, (select a.id from attributes a, punks p where p.att7 = a.name AND p.id = punks.id), att7_score FROM punks;`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attributes_att7"] = stmt
	}

	if stmt, err := db.Prepare(
		`INSERT INTO punk_attributes (punk_id, attribute_id, score)
SELECT id as punk_id, (select a.id from attributes a where a.name = concat (p.slots, 'Att')) as attribute_id,  (select a.total/100 from attributes a where a.name = concat (p.slots, 'Att')) as total  FROM punks p where 1;`,
	); err != nil {
		return err
	} else {
		statements["insert_punk_attribute_attributes"] = stmt
	}

	// Calculate the category scores
	/*
		if stmt, err := db.Prepare(
			`CREATE TEMPORARY TABLE new_tbl select count(*) as total , category from punks group by category;`,
		); err != nil {
			return err
		} else {
			statements["insert_punk_cat_temp"] = stmt
		}

		if stmt, err := db.Prepare(
			`UPDATE punks SET category_score = (select total from new_tbl where category =  punks.category) / 100 ;`,
		); err != nil {
			return err
		} else {
			statements["insert_punk_cat_score"] = stmt
		}

		if stmt, err := db.Prepare(
			`DROP TEMPORARY TABLE new_tbl;`,
		); err != nil {
			return err
		} else {
			statements["insert_punk_cat_temp_del"] = stmt
		}
	*/

	return nil
}
