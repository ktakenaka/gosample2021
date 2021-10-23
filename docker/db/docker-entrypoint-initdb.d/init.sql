CREATE DATABASE IF NOT EXISTS `gosample_development` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE USER 'gosample'@'%' IDENTIFIED BY 'password';  
GRANT ALL PRIVILEGES ON *.* TO 'gosample'@'%' WITH GRANT OPTION;  