/*
 Navicat Premium Data Transfer

 Source Server         : Local Lokajaya
 Source Server Type    : MariaDB
 Source Server Version : 101100 (10.11.0-MariaDB-log)
 Source Host           : localhost:3306
 Source Schema         : go_auth

 Target Server Type    : MariaDB
 Target Server Version : 101100 (10.11.0-MariaDB-log)
 File Encoding         : 65001

 Date: 26/02/2024 13:00:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for catalogs
-- ----------------------------
DROP TABLE IF EXISTS `catalogs`;
CREATE TABLE `catalogs`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bean` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `price` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of catalogs
-- ----------------------------
INSERT INTO `catalogs` VALUES (1, 'coffe', 'black coffe', '$5');

-- ----------------------------
-- Table structure for distributors
-- ----------------------------
DROP TABLE IF EXISTS `distributors`;
CREATE TABLE `distributors`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `city` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `region` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `country` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of distributors
-- ----------------------------
INSERT INTO `distributors` VALUES (1, 'Arabica Taste', 'Kudus', 'Jawa Tengah', 'indonesia', '085156161510', 'skomfery@gmail.com');

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `file` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `author` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of files
-- ----------------------------
INSERT INTO `files` VALUES (1, 'fery', 'fery.pdf', 'deny fery');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `nama_lengkap` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (2, 'DENY FERY IRAWAN', 'denyfery39@gmail.com', 'fery', '$2a$10$FQoxMiCV3ulR0XGqLI2kLOwKBm9qc77WXxT9jUufI5LyAulcbTpei');
INSERT INTO `users` VALUES (3, 'Coffe Valley', 'coffee@gmail.com', 'admin', '$2a$10$ajw.ZuPVGFavOmlno9540ei71B1HzSBfd0WxirLicH/ERZyn3PsKe');

SET FOREIGN_KEY_CHECKS = 1;
