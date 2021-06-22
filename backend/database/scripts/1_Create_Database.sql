USE GoSandBox;

CREATE TABLE movies(
    id varchar(36) not null,
    title varchar(100) not null,
    year integer not null,
    PRIMARY KEY (id)
);

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;