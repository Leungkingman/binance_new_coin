/*
 Navicat MySQL Data Transfer

 Source Server         : bnl
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : 8.218.81.24:3306
 Source Schema         : binance_new_coin

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 20/01/2022 11:39:52
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
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bm_articles
-- ----------------------------
BEGIN;
INSERT INTO `bm_articles` VALUES (93, 'Binance Will List Convex Finance (CVX) and ConstitutionDAO (PEOPLE) in the Innovation Zone', '1640225010894', NULL, NULL);
INSERT INTO `bm_articles` VALUES (94, 'Binance Adds COCOS/TRY, GXS/BNB, LINK/BNB, LUNA/ETH, MDT/BUSD & NULS/BUSD Trading Pairs', '1640272517888', NULL, NULL);
INSERT INTO `bm_articles` VALUES (95, 'Binance Will List Spell Token (SPELL) and TerraUSD (UST)', '1640311117895', NULL, NULL);
INSERT INTO `bm_articles` VALUES (96, 'Binance Adds BNX, KEEP, KLAY, MINA, TRIBE on Cross Margin and FLUX, PEOPLE, REQ, SUN, VOXEL on Isolated Margin', '1640500241671', NULL, NULL);
INSERT INTO `bm_articles` VALUES (100, 'Binance Futures Will Launch Coin-Margined AAVE Perpetual Contracts with Up to 20X Leverage', '1640595065998', NULL, NULL);
INSERT INTO `bm_articles` VALUES (101, 'Binance Will List JOE (JOE)', '1640667703916', NULL, NULL);
INSERT INTO `bm_articles` VALUES (102, 'Binance Futures Will Launch USDT-Margined ROSE Perpetual Contracts with Up to 25X Leverage', '1640778276414', NULL, NULL);
INSERT INTO `bm_articles` VALUES (103, 'Binance Futures Will Launch USDT-Margined ROSE Perpetual Contracts with Up to 25X Leverage', '1640778276421', NULL, NULL);
INSERT INTO `bm_articles` VALUES (104, 'Binance Futures Will Launch USDT-Margined ROSE Perpetual Contracts with Up to 25X Leverage', '1640778276434', NULL, NULL);
INSERT INTO `bm_articles` VALUES (105, 'Binance Adds AGLD, AUDIO, BICO, GTC, JASMY, LPT, QUICK, RNDR on Cross Margin and HIGH, MC, OOKI, SPELL on Isolated Margin', '1640942870848', NULL, NULL);
INSERT INTO `bm_articles` VALUES (106, 'Binance Adds ATOM/ETH, DUSK/BUSD, EGLD/ETH, ICP/ETH, LUNA/BRL, LUNA/UST, NEAR/ETH, ROSE/BNB & VOXEL/ETH Trading Pairs', '1641286609484', NULL, NULL);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=260 DEFAULT CHARSET=utf8;

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
-- Records of bm_exist_coin
-- ----------------------------
BEGIN;
INSERT INTO `bm_exist_coin` VALUES (1, 'FLUX,CVX,JOE');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bm_news_title
-- ----------------------------
BEGIN;
INSERT INTO `bm_news_title` VALUES (1, 'Notice on Converting Delisted Tokens Balances to BUSD (2021-07-02)');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bm_order
-- ----------------------------
BEGIN;
INSERT INTO `bm_order` VALUES (11, '78704040259', 'GALA', '0.11872', '0.012', '', '383.45', '45.523184', '0', '1632076684');
INSERT INTO `bm_order` VALUES (12, '80020351758', 'YGG', '7.4411', '8.9684', '', '1561.05', '11615.98281467', '0', '1632448736');
INSERT INTO `bm_order` VALUES (13, '81491290900', 'FIDA', '7.3', '8.7', '1632968965381', '2068.97', '15080.914489509', '0', '1632968967');
INSERT INTO `bm_order` VALUES (14, '82589642771', 'AGLD', '3.8317', '3.3187', '1633401494813', '4089.38', '15669.1142072801879', '0', '1633401498');
INSERT INTO `bm_order` VALUES (17, '83054776930', 'RAD', '12.025', '0', '1633577073000', '1300.06', '15632.67882539786', '0.00', '1633577073');
INSERT INTO `bm_order` VALUES (18, '83990889524', 'RARE', '1.74460', '1.37601', '1633917838537', '10381.98', '18112.3654587021503', '0', '1633917844');
INSERT INTO `bm_order` VALUES (20, '86733871589', 'CHESS', '7.5662', '5.7984', '1634870452658', '3449.23', '26097.55100409958813687172', '0', '1634870630');
INSERT INTO `bm_order` VALUES (21, '90409803606', 'BNX', '136.855', '101.347', '1636008942920', '140.96', '19291.136824', '0', '1636008948');
INSERT INTO `bm_order` VALUES (22, '90668434978', 'RGT', '0.000', '43.548', '1636078144236', '492.07', '0', '492.07', '1636078151');
INSERT INTO `bm_order` VALUES (23, '92099222187', 'ENS', '55.992', '43.421', '1636513937195', '919.68', '51494.8054291', '0', '1636513939');
INSERT INTO `bm_order` VALUES (24, '93559176641', 'QI', '0.00000', '0.16546', '1636941512440', '215848.46', '0', '215848.46', '1636941515');
INSERT INTO `bm_order` VALUES (27, '96237086883', 'AMP', '0.0000', '0.0511', '1637632767018', '139781.94', '0', '139781.94', '1637632773');
INSERT INTO `bm_order` VALUES (28, '98737625349', 'ALCX', '475.30', '380.24', '1638237552715', '21.04', '10000.312', '0.00', '1638237552');
INSERT INTO `bm_order` VALUES (29, '99470990357', 'MC', '12.439', '9.951', '1638417773329', '1205.91', '15000.31449', '0.00', '1638417780');
INSERT INTO `bm_order` VALUES (30, '100135620211', 'ANY', '16.809', '15.008', '1638583309133', '892.38', '15000.01542', '0', '1638583310');
INSERT INTO `bm_order` VALUES (31, '101857598441', 'BICO', '7.18647', '7.01213', '0', '254.66', '1830.1061104179', '0', '1639015035');
INSERT INTO `bm_order` VALUES (32, '106608182684', 'SPELL', '0.0000000', '0.0231383', '0', '160068.1', '0', '160068.1', '1640311116');
COMMIT;

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

-- ----------------------------
-- Records of bm_user
-- ----------------------------
BEGIN;
INSERT INTO `bm_user` VALUES (1, 'admin', 'xxxxx', 'xxxx', 'xxxxxx', '4000', '', '', '8', '', 1, NULL, 10, '4000', '1000', '8', '20');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
