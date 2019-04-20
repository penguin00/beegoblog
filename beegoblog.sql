/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50553
 Source Host           : localhost:3306
 Source Schema         : beegoblog

 Target Server Type    : MySQL
 Target Server Version : 50553
 File Encoding         : 65001

 Date: 20/04/2019 21:32:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `created` datetime NULL DEFAULT NULL,
  `view` bigint(20) NOT NULL DEFAULT 0,
  `topic_time` datetime NULL DEFAULT NULL,
  `topic_count` bigint(20) NOT NULL DEFAULT 0,
  `topic_last_user_id` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `category_created`(`created`) USING BTREE,
  INDEX `category_view`(`view`) USING BTREE,
  INDEX `category_topic_time`(`topic_time`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for reply
-- ----------------------------
DROP TABLE IF EXISTS `reply`;
CREATE TABLE `reply`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tid` bigint(20) NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `content` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT 0,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `content` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `attachment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `view` bigint(20) NOT NULL DEFAULT 0,
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `reply_time` datetime NULL DEFAULT NULL,
  `reply_count` bigint(20) NULL DEFAULT 0,
  `reply_last_user_id` bigint(20) NULL DEFAULT 0,
  `category` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `lables` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `topic_created`(`created`) USING BTREE,
  INDEX `topic_updated`(`updated`) USING BTREE,
  INDEX `topic_view`(`view`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
