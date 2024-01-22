-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: db_kredit
-- ------------------------------------------------------
-- Server version	11.1.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customer_limits`
--

DROP TABLE IF EXISTS `customer_limits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer_limits` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_id` bigint(20) unsigned DEFAULT NULL,
  `limit` bigint(20) DEFAULT NULL,
  `period` bigint(20) DEFAULT NULL,
  `used_limit` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_customer_limits_deleted_at` (`deleted_at`),
  KEY `idx_customer_limits_customer_id` (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer_limits`
--

LOCK TABLES `customer_limits` WRITE;
/*!40000 ALTER TABLE `customer_limits` DISABLE KEYS */;
INSERT INTO `customer_limits` VALUES (1,'2024-01-20 23:56:27.000','2024-01-21 20:31:00.852',NULL,1,100000,1,210422),(2,'2024-01-20 23:56:27.000','2024-01-21 20:31:00.854',NULL,1,200000,2,210422),(3,'2024-01-20 23:56:27.000','2024-01-21 20:31:00.855',NULL,1,500000,3,109588),(4,'2024-01-20 23:56:27.000','2024-01-21 20:31:00.857',NULL,1,700000,6,102500);
/*!40000 ALTER TABLE `customer_limits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `email` longtext DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `nik` longtext DEFAULT NULL,
  `full_name` longtext DEFAULT NULL,
  `legal_name` longtext DEFAULT NULL,
  `born_place` longtext DEFAULT NULL,
  `born_at` datetime(3) DEFAULT NULL,
  `salary` bigint(20) DEFAULT NULL,
  `photo_ktp` longtext DEFAULT NULL,
  `photo_selfie` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `risk_id` bigint(20) DEFAULT NULL,
  `risk` longtext DEFAULT NULL,
  `interest` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_customers_deleted_at` (`deleted_at`),
  KEY `idx_customers_risk_id` (`risk_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'test@gmail.com','0896556564645','$2a$10$hdNzwmG0V4G9ogVXwhPy1O9lE1z14bWE26QfTYyirgwdWUAi990ge','33234354353432','testing','testing','test place','2024-01-20 23:56:27.000',5000000,NULL,NULL,NULL,NULL,NULL,NULL,'LOW',5);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `installments`
--

DROP TABLE IF EXISTS `installments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `installments` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `transaction_id` bigint(20) DEFAULT NULL,
  `installment_amount` bigint(20) DEFAULT NULL,
  `installment_date` datetime(3) DEFAULT NULL,
  `installment_period` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_installments_deleted_at` (`deleted_at`),
  KEY `idx_installments_transaction_id` (`transaction_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `installments`
--

LOCK TABLES `installments` WRITE;
/*!40000 ALTER TABLE `installments` DISABLE KEYS */;
INSERT INTO `installments` VALUES (1,'2024-01-21 19:31:43.258','2024-01-21 19:31:43.258',NULL,5,3544,'2024-01-21 19:31:43.256',1);
/*!40000 ALTER TABLE `installments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `partners`
--

DROP TABLE IF EXISTS `partners`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `partners` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `address` longtext DEFAULT NULL,
  `client_key` longtext DEFAULT NULL,
  `client_secret` longtext DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `api_key` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_partners_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `partners`
--

LOCK TABLES `partners` WRITE;
/*!40000 ALTER TABLE `partners` DISABLE KEYS */;
INSERT INTO `partners` VALUES (1,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,'Partner A','Address','42255ae3bb6866a8d77f105caa25204201af42cb',NULL,1,'42255ae3bb6866a8d77f105caa25204201af42cb');
/*!40000 ALTER TABLE `partners` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `risk_limits`
--

DROP TABLE IF EXISTS `risk_limits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `risk_limits` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `risk_id` bigint(20) DEFAULT NULL,
  `limit` bigint(20) DEFAULT NULL,
  `period` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_risk_limits_deleted_at` (`deleted_at`),
  KEY `idx_risk_limits_risk_id` (`risk_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `risk_limits`
--

LOCK TABLES `risk_limits` WRITE;
/*!40000 ALTER TABLE `risk_limits` DISABLE KEYS */;
INSERT INTO `risk_limits` VALUES (1,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,1,100000,1),(2,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,1,200000,2),(3,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,1,500000,3),(4,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,1,700000,6),(5,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,2,1000000,1),(6,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,2,1200000,2),(7,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,2,1500000,3),(8,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,2,2000000,6);
/*!40000 ALTER TABLE `risk_limits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `risks`
--

DROP TABLE IF EXISTS `risks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `risks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `risk` longtext DEFAULT NULL,
  `interest` double DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_risks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `risks`
--

LOCK TABLES `risks` WRITE;
/*!40000 ALTER TABLE `risks` DISABLE KEYS */;
INSERT INTO `risks` VALUES (1,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,'LOW',5,1),(2,'2024-01-20 23:56:27.000','2024-01-20 23:56:27.000',NULL,'MEDIUM',10,1);
/*!40000 ALTER TABLE `risks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `customer_id` bigint(20) DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `admin_fee` bigint(20) DEFAULT NULL,
  `otr` bigint(20) DEFAULT NULL,
  `installment_amount` bigint(20) DEFAULT NULL,
  `interest` float DEFAULT NULL,
  `asset_name` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `installment_period` bigint(20) DEFAULT NULL,
  `transaction_date` datetime(3) DEFAULT NULL,
  `partner_id` bigint(20) DEFAULT NULL,
  `token` varchar(191) DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `order_id` varchar(191) DEFAULT NULL,
  `txn_id` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_transactions_deleted_at` (`deleted_at`),
  KEY `idx_transactions_customer_id` (`customer_id`),
  KEY `idx_transactions_partner_id` (`partner_id`),
  KEY `idx_transactions_token` (`token`),
  KEY `idx_transactions_order_id` (`order_id`),
  KEY `idx_transactions_txn_id` (`txn_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,0,'089538768687',50000,500000,0,0,'jacket','2024-01-20 23:33:01.597','2024-01-21 00:35:10.151',NULL,0,'2024-01-20 23:33:01.595',0,NULL,'CANCELLED','Test','018d2938-721a-716b-a4c7-ef978b910ab6'),(3,0,'089538768687',50000,500000,0,0,'jacket','2024-01-21 00:35:10.152','2024-01-21 08:37:24.428',NULL,0,'2024-01-21 00:35:10.151',0,NULL,'CANCELLED','Test','018d2971-56c7-7103-b549-d49adbc1a75f'),(4,0,'089538768687',50000,500000,0,0,'jacket','2024-01-21 08:37:24.430','2024-01-21 17:48:19.924',NULL,0,'2024-01-21 08:37:24.429',1,NULL,'CANCELLED','Test','018d2b2a-d74d-7317-b2a7-ce7f3785a6f2'),(5,1,'089538768687',500,10000,3544,5,'jacket','2024-01-21 17:48:19.924','2024-01-21 19:15:43.849',NULL,5,'2024-01-21 17:48:19.924',1,NULL,'COMPLETE','Test','018d2d23-3a54-7cec-a410-c6cd13d57a25'),(6,1,'089538768687',0,100000,50417,5,'jacket','2024-01-21 19:55:10.829','2024-01-21 20:28:37.022',NULL,2,'2024-01-21 19:55:10.829',1,NULL,'COMPLETE','Test2','018d2d97-5c6d-7441-b91c-2a235a64cb30'),(7,1,'089538768687',0,100000,17084,5,'jacket','2024-01-21 20:30:17.303','2024-01-21 20:31:00.866',NULL,6,'2024-01-21 20:30:17.302',1,NULL,'COMPLETE','Test3','018d2db7-80d6-7690-b62b-cbaa5ad776cd');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'db_kredit'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-23  3:35:08
