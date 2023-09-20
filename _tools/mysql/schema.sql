CREATE TABLE Users (
    id  INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL ,
    high_score INT DEFAULT 0,
    coin       INT DEFAULT 100
)