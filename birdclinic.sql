-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 21, 2019 at 03:58 PM
-- Server version: 10.1.40-MariaDB
-- PHP Version: 7.3.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `birdclinic`
--

-- --------------------------------------------------------

--
-- Table structure for table `disease`
--

CREATE TABLE `disease` (
  `name` varchar(65) NOT NULL,
  `disease_id` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `disease`
--

INSERT INTO `disease` (`name`, `disease_id`) VALUES
('Fowl Cholera', 1000),
('Infectious Coryza', 1001),
('Infectious Bursal disease', 1003),
('Avian infected bronchitis', 1004),
('Oral canker', 1005),
('Tick fever', 1006),
('Fowl typhoid', 1007),
('Heximetiasis', 1008),
('Internal parasites', 1009),
('Coccidiosis', 1010),
('Thrush', 1011),
('Histomomiasis', 1012),
('Marek’s disease', 1013),
('Newcastle disease', 1014);

-- --------------------------------------------------------

--
-- Table structure for table `farmer_account`
--

CREATE TABLE `farmer_account` (
  `id` varchar(128) NOT NULL,
  `username` varchar(20) NOT NULL,
  `email` varchar(65) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `farmer_account`
--

INSERT INTO `farmer_account` (`id`, `username`, `email`, `password`) VALUES
('2c254829-5ee9-4d71-99b2-fe34728c92a1', 'silent wolf', 'cetokola2015@gmail.com', '$2a$10$qJ1mJp9vRmfF8IsD7XPhP.Dova2zg7ifQgsn3yLsiWHLWvYJVoyWe'),
('6154dfd8-d5fd-46e9-b9d6-02cd7baa6a0a', 'cetric', 'cetokola@gmail.com', '$2a$10$mvLYfqo4lt7ak8uZ1Joc.OacSc0BEvG30iCxIMhU7D2riFimUCW6K'),
('ea5969b1-1936-49b8-bb09-c616802a8332', 'alex', 'ce1tokola@gmail.com', '$2a$10$s5k9N34/jisaMLxqs2R7xeCNwcR8k2KaWjx6BKNEB5mamHuIzPbU6');

-- --------------------------------------------------------

--
-- Table structure for table `message`
--

CREATE TABLE `message` (
  `id` varchar(128) NOT NULL,
  `username` varchar(20) NOT NULL,
  `email` varchar(65) NOT NULL,
  `message` text NOT NULL,
  `subject` text NOT NULL,
  `send_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `pro`
--

CREATE TABLE `pro` (
  `id` varchar(128) NOT NULL,
  `username` varchar(20) NOT NULL,
  `first_name` varchar(20) NOT NULL,
  `last_name` varchar(20) NOT NULL,
  `email` varchar(60) NOT NULL,
  `phone` varchar(10) NOT NULL,
  `district` varchar(20) NOT NULL,
  `county` varchar(20) NOT NULL,
  `country` varchar(20) NOT NULL,
  `password` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pro`
--

INSERT INTO `pro` (`id`, `username`, `first_name`, `last_name`, `email`, `phone`, `district`, `county`, `country`, `password`) VALUES
('8a0b15e6-d451-4715-8e31-c22663adb0a9', 'brayo', 'Brian', 'Mukune', 'cetokola2015@gmail.com', '0704145833', 'kisumu', 'kakamega', 'kenya', '$2a$10$1k.MQ1A2rr08hxZUsspxDujZRK/WPQ70fGqzyQyNpliHBxf76ojXG'),
('d459d1cc-def4-4cff-ad30-0a6dc2a17194', 'cetric1', 'cetric', 'okola', 'cetokola@gmail.com', '0704145832', 'kisumu', 'kakamega', 'kenya', '$2a$10$ophOboefdlaralAPqNniLujfpUZ7RVXk5NZvh3jpaTcJq.IJdHJBq'),
('e2f3e99e-8bf5-4660-9f07-ed96426876b2', 'dero', 'Derrick', 'Mugune', 'cetric@gmail.com', '0704145896', 'Mumias', 'Kakamega', 'Kenya', '$2a$10$Kb1U4f6GHTxXfLFF7FDxjeH7obfb33vdZ1dG94.nmMIXs/fSEZqci');

-- --------------------------------------------------------

--
-- Table structure for table `symptom`
--

CREATE TABLE `symptom` (
  `id` int(11) NOT NULL,
  `name` varchar(65) NOT NULL DEFAULT 'invalid symptom',
  `disease_id` int(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `symptom`
--

INSERT INTO `symptom` (`id`, `name`, `disease_id`) VALUES
(1, 'Bluish/ Gray eye', 1013),
(2, 'Breathing problems', 1011),
(3, 'Closed eyes', 1005),
(4, 'Coughing', 1003),
(5, 'Coughing', 1014),
(6, 'Cyanosis ', 1000),
(7, 'Dehydration', 1006),
(8, 'Depression', 1014),
(9, 'Diarrhea', 1006),
(10, 'Diarrhoea', 1008),
(11, 'Diarrhoea', 1010),
(12, 'Diarrhoea', 1014),
(13, 'Discoloured eggs', 1003),
(14, 'Droopy looking feathers', 1011),
(15, 'Drop in egg production ', 1001),
(16, 'Drowsiness', 1005),
(17, 'Dry unkempt feathers', 1007),
(18, 'Dullness', 1008),
(19, 'Dullness', 1009),
(20, 'Enlarged feather follicles', 1013),
(21, 'Facial swelling', 1001),
(22, 'Greenish diarrhea', 1005),
(23, 'Hanged head', 1005),
(24, 'High mortality', 1006),
(25, 'Hunched up', 1010),
(26, 'Inactivity', 1011),
(27, 'Inappetence', 1001),
(28, 'Increased thirst', 1012),
(29, 'Lack of appetite', 1000),
(30, 'Lack of appetite', 1006),
(31, 'Lethargy', 1012),
(32, 'Listless chicks', 1002),
(33, 'Listlessness', 1007),
(34, 'Mucoid diarrhea', 1002),
(35, 'Nasal discharge', 1001),
(36, 'Nasal discharge', 1014),
(37, 'Oral discharge', 1000),
(38, 'Paralysis', 1008),
(39, 'Paralysis of legs', 1013),
(40, 'Paralysis of neck', 1013),
(41, 'Paralysis of wings', 1013),
(42, 'Picking at own vent', 1002),
(43, 'Picking up food then dropping it', 1004),
(44, 'Poor egg albumen', 1003),
(45, 'Poor growth', 1012),
(46, 'Poor weight gain', 1008),
(47, 'Rapid drop in feeding', 1002),
(48, 'Rapid drop in water consumption', 1002),
(49, 'Rapid weight loss', 1007),
(50, 'Rattling', 1003),
(51, 'Reduced appetite', 1008),
(52, 'Reduced appetite', 1011),
(53, 'Reduced appetite', 1012),
(54, 'Reluctance to eat', 1004),
(55, 'Ruffled feathers', 1000),
(56, 'Ruffled feathers', 1010),
(57, 'Ruffled feathers', 1012),
(58, 'Runny nose', 1011),
(59, 'Sharp drop in egg production', 1003),
(60, 'Sitting in hinged position', 1002),
(61, 'Sleeping with beak touching the floor', 1002),
(62, 'Sneezing', 1001),
(63, 'Sneezing', 1014),
(64, 'Soft egg shell', 1003),
(65, 'Soiled vent feathers', 1002),
(66, 'Swollen wattles', 1001),
(67, 'Unsteady gait', 1002),
(68, 'Watery diarrhea', 1007),
(69, 'Weakness', 1006),
(70, 'Weight loss', 1004),
(71, 'Weight loss', 1010),
(72, 'White/greenish watery mucoid diarrhea ', 1000);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `disease`
--
ALTER TABLE `disease`
  ADD PRIMARY KEY (`disease_id`);

--
-- Indexes for table `farmer_account`
--
ALTER TABLE `farmer_account`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `message`
--
ALTER TABLE `message`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pro`
--
ALTER TABLE `pro`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- Indexes for table `symptom`
--
ALTER TABLE `symptom`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `symptom`
--
ALTER TABLE `symptom`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=73;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
