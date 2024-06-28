/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : binance_new_coin

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 28/06/2024 09:38:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bm_articles
-- ----------------------------
DROP TABLE IF EXISTS `bm_articles`;
CREATE TABLE `bm_articles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `title_type` varchar(100) DEFAULT NULL,
  `publish_time` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for bm_binance_coin
-- ----------------------------
DROP TABLE IF EXISTS `bm_binance_coin`;
CREATE TABLE `bm_binance_coin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `coins` text NOT NULL,
  `coin_number` int NOT NULL,
  `create_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=260 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for bm_domain
-- ----------------------------
DROP TABLE IF EXISTS `bm_domain`;
CREATE TABLE `bm_domain` (
  `id` int NOT NULL AUTO_INCREMENT,
  `domain` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for bm_exist_coin
-- ----------------------------
DROP TABLE IF EXISTS `bm_exist_coin`;
CREATE TABLE `bm_exist_coin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `coin` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for bm_mx_order_id
-- ----------------------------
DROP TABLE IF EXISTS `bm_mx_order_id`;
CREATE TABLE `bm_mx_order_id` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` varchar(255) NOT NULL,
  `create_time` varchar(0) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for bm_news_title
-- ----------------------------
DROP TABLE IF EXISTS `bm_news_title`;
CREATE TABLE `bm_news_title` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for bm_order
-- ----------------------------
DROP TABLE IF EXISTS `bm_order`;
CREATE TABLE `bm_order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `coin` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `price` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `find_price` varchar(100) NOT NULL,
  `find_time` varchar(100) NOT NULL,
  `amount` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `total` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `left` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Table structure for bm_user
-- ----------------------------
DROP TABLE IF EXISTS `bm_user`;
CREATE TABLE `bm_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `mx_access_key` varchar(255) DEFAULT NULL,
  `mx_secret_key` varchar(255) DEFAULT NULL,
  `buy_usdt` varchar(255) DEFAULT NULL,
  `profit` varchar(0) DEFAULT NULL,
  `loss` varchar(0) DEFAULT NULL,
  `slippage` varchar(255) DEFAULT NULL,
  `next_slippage` varchar(255) DEFAULT NULL,
  `worker_count` int DEFAULT '1',
  `inform_time` int DEFAULT NULL,
  `queue_time_gap` int DEFAULT NULL,
  `request_time_gap` varchar(255) DEFAULT NULL,
  `new_request_time_gap` varchar(255) DEFAULT NULL,
  `long_profit` varchar(100) DEFAULT NULL,
  `short_profit` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
