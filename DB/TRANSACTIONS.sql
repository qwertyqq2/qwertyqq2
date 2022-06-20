CREATE TABLE TRANSACTIONS (
    Id int NOT NULL AUTO_INCREMENT,
    IdUser int NOT NULL,
    Email varchar(255),
    Currency varchar(255),
    TimeOfCreation varchar(255),
    TimeOfLastChange varchar(255),
    Status varchar(255),
    PRIMARY KEY (Id)
); 
INSERT INTO TRANSACTIONS VALUES (1, 2, "SemahcokOlga@mail.ru", "RUB", "2009-11-10 15:00:00 -0800 PST", "2009-11-10 15:00:00 -0800 PST", "True");
