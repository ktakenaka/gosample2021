CREATE DATABASE IF NOT EXISTS `gosample2021_development` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS `gosample2021_test` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE USER 'gosample2021'@'%' IDENTIFIED BY 'password';  
GRANT ALL PRIVILEGES ON *.* TO 'gosample2021'@'%' WITH GRANT OPTION;  
