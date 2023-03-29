-- MySQL
CREATE TABLE IF NOT EXISTS `tb_student` (
	`id` varchar(5) NOT NULL,
	`name` varchar(255) NOT NULL,
	`age` int(11) NOT NULL,
	`grade` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- PostgreSQL
CREATE TABLE IF NOT EXISTS tb_student (
	id VARCHAR(5) NOT NULL,
	name VARCHAR(255) NOT NULL,
	age INTEGER NOT NULL,
	grade INTEGER NOT NULL
);

-- MySQL
INSERT INTO `tb_student` (`id`, `name`, `age`, `grade`) VALUES
	('B001', 'Jason Bourne', 20, 1),
	('B002', 'James Bond', 21, 2),
	('E001', 'Jack Ryan', 22, 3),
	('W001', 'John Wick', 23, 4);

-- PostgreSQL
INSERT INTO tb_student (id, name, age, grade) VALUES
	('B001', 'Jason Bourne', 20, 1),
	('B002', 'James Bond', 21, 2),
	('E001', 'Jack Ryan', 22, 3),
	('W001', 'John Wick', 23, 4);

-- MySQL
ALTER TABLE `tb_student` ADD PRIMARY KEY (`id`);

-- PostgreSQL
ALTER TABLE tb_student ADD PRIMARY KEY (id);