-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 12, 2022 at 06:00 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `article`
--

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `title`, `content`, `category`, `created_date`, `updated_date`, `status`) VALUES
(6, 'Steam Deck is Slowly, But Surely Adding Windows Support', 'The Steam Deck has added Windows drivers, meaning Valve’s handheld gaming PC will be able to boot a Windows operating system. With some caveats, of course. Currently, the Steam Deck uses a custom version of SteamOS, but Valve has said in the past that Win', 'Tech', '2022-03-11 16:07:40', '2022-03-11 16:07:40', 'Publish'),
(7, 'Elden Ring Has Had a Huge Launch in the US', 'Elden Ring is officially the United States\' best-selling game of 2022 so far, with a huge launch according to the latest NPD data. FromSoftware\'s latest is the second best launch of the last twelve months, falling short of only Call of Duty: Vanguard. Tha', 'XBOX', '2022-03-11 16:09:25', '2022-03-11 16:09:25', 'Publish'),
(8, 'Nintendo Becomes the Latest Gaming Giant to Stop Sales in Russia', 'Nintendo has halted all shipments to Russia for the foreseeable future, but stopped short of publicly supporting Ukraine. In a statement provided to IGN, the gaming giant said the rapidly changing circumstances around exporting and selling its products in', 'Nintendo', '2022-03-11 16:10:43', '2022-03-11 16:10:43', 'Publish'),
(13, 'Sony PlayStation Suspends Software and Hardware Sales in Russia', 'Sony has shared a statement with IGN saying that it will be suspending all software and hardware sales in Russia due to its recent and ongoing invasion of Ukraine. Sony Interactive Entertainment (SIE) joins the global community in calling for peace in Ukr', 'PlayStation', '2022-03-12 00:59:04', '2022-03-12 04:53:27', 'Publish'),
(14, 'Kirby And the Forgotten Land Demo Available Now For Nintendo Switch', 'Kirby and the Forgotten Land has received a surprise demo, Nintendo announced today. The demo features the first three levels from Kirby\'s initial world, offering players a taste of the adventure that some are comparing to The Last of Us and The Legend of', 'Nintendo', '2022-03-12 02:09:37', '2022-03-12 02:16:04', 'Draft'),
(15, 'A God of War TV Series Is In the Works at Amazon Prime Video', 'Amazon Prime Video will potentially expand their video game TV series lineup with none other than a God of War live-action series. According to Deadline, Prime Video is in talks with PlayStation to adapt God of War as a live-action TV series. The Expanse ', 'PlayStation', '2022-03-12 04:15:56', '2022-03-12 04:15:56', 'Publish');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
