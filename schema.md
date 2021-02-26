```mysql
create table attributes
(
	id int unsigned not null,
	name varchar(128) null,
	total int null,
	constraint attributes_id_uindex
		unique (id)
);

alter table attributes
	add primary key (id);

create table punks
(
	id int unsigned not null,
	sex varchar(32) not null,
	type varchar(32) not null,
	skin varchar(32) not null,
	type_skin varchar(32) null,
	slots int not null,
	att1 varchar(32) not null,
	att2 varchar(32) not null,
	att3 varchar(32) not null,
	att4 varchar(32) not null,
	att5 varchar(32) not null,
	att6 varchar(32) not null,
	att7 varchar(32) not null,
	type_rare decimal(19,10) not null,
	att_count int not null,
	att_count_score decimal(19,10) null,
	sex_score decimal(19,10) null,
	type_score decimal(19,10) null,
	skin_score decimal(19,10) null,
	att1_score decimal(19,10) null,
	att2_score decimal(19,10) null,
	att3_score decimal(19,10) null,
	att4_score decimal(19,10) null,
	att5_score decimal(19,10) null,
	att6_score decimal(19,10) null,
	att7_score decimal(19,10) null,
	min decimal(19,10) not null,
	avg decimal(19,10) not null,
	`rank` decimal(19,10) null,
	category varchar(128) null,
	category_score decimal(19,10) null,
	constraint id
		unique (id)
)
collate=utf8mb4_bin;

alter table punks
	add primary key (id);

create table punk_attributes
(
	punk_id int unsigned null,
	attribute_id int unsigned null,
	score decimal(19,10) null,
	constraint punk_attributes_attributes_id_fk
		foreign key (attribute_id) references attributes (id),
	constraint punk_attributes_flat_id_fk
		foreign key (punk_id) references punks (id)
);

create index punk_attributes_punk_id_attribute_id_index
	on punk_attributes (punk_id, attribute_id);


```