DROP TABLE IF EXISTS students;

CREATE TABLE IF NOT EXISTS students (
    ID			BIGINT			PRIMARY KEY	AUTO_INCREMENT, 
    FirstName	VARCHAR(60) 	NOT NULL, 
    MiddleName	VARCHAR(60), 
    LastName	VARCHAR(60)     NOT NULL, 
    BirthDate	DATE,
    Gender		TINYINT         DEFAULT	'UNK',
    Phone		VARCHAR(15),
    Address		VARCHAR(200),
    Country		CHAR(3)
)