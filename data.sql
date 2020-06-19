-- MySQL dump 10.17  Distrib 10.3.22-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: animes
-- ------------------------------------------------------
-- Server version	10.3.22-MariaDB-1ubuntu1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `series`
--

DROP TABLE IF EXISTS `series`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `series` (
  `id_serie` int(11) NOT NULL AUTO_INCREMENT,
  `name_serie` varchar(150) NOT NULL,
  `episodes_serie` int(11) NOT NULL DEFAULT 1,
  `status_serie` varchar(45) NOT NULL,
  `type_serie` varchar(45) NOT NULL,
  `studio_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_serie`),
  KEY `Fk_StudioAnime` (`studio_id`),
  CONSTRAINT `Fk_StudioAnime` FOREIGN KEY (`studio_id`) REFERENCES `studios` (`id_studio`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `series`
--

LOCK TABLES `series` WRITE;
/*!40000 ALTER TABLE `series` DISABLE KEYS */;
INSERT INTO `series` VALUES (1,'Death Note',37,'Finished','TV',1),(2,'Shigatsu wa Kimi no Uso',22,'Finished','TV',2),(7,'One Punch Man',12,'Finished','TV',1),(9,'Fairy Tail',175,'Finished','TV',2),(10,'Mobile Suit Gundam Wing',49,'Finished','TV',8),(11,'Slam Dunk',101,'Finished','TV',9),(12,'Inuyasha',167,'Finished','TV',8),(13,'Higurashi no Naku Koro ni',26,'Finished','TV',10);
/*!40000 ALTER TABLE `series` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `studios`
--

DROP TABLE IF EXISTS `studios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `studios` (
  `id_studio` int(11) NOT NULL AUTO_INCREMENT,
  `name_studio` varchar(50) NOT NULL,
  PRIMARY KEY (`id_studio`),
  UNIQUE KEY `name_studio` (`name_studio`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `studios`
--

LOCK TABLES `studios` WRITE;
/*!40000 ALTER TABLE `studios` DISABLE KEYS */;
INSERT INTO `studios` VALUES (2,'A-1 Pictures'),(1,'Madhouse'),(10,'Studio Deen'),(8,'Sunrise'),(9,'Toei Animation');
/*!40000 ALTER TABLE `studios` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-19  9:34:09
