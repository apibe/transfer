/*
 ONLINE-MAP-MAKING Premium Data Transfer

 Source Server         : {SERVER_NAME}
 Source Server Type    : {SERVER_TYPE}
 Source Host           : {HOST}
 Source Catalog        : {DATABASE}
 Source Table          : {TABLE}
 Source Schema         : {SCHEMA}

 Date: {02/01/2006 15:04:05}
*/
-- ----------------------------
-- Before create for {TABLE}
-- ----------------------------
{BEFORE_CREATE}

-- ----------------------------
-- Table structure for {TABLE}
-- ----------------------------
CREATE TABLE t_test ( YOU text, I text);

-- ----------------------------
-- After create for {TABLE}
-- ----------------------------
{AFTER_CREATE}

-- ----------------------------
-- Before insert for {TABLE}
-- ----------------------------
{BEFORE_INSERT}

-- ----------------------------
-- Records of {TABLE}
-- ----------------------------
INSERT INTO t_test ('YOU','YOU','I') VALUE ('1','2');
INSERT INTO t_test ('YOU') VALUE ('1');
INSERT INTO t_test ('YOU','YOU','I') VALUE ('1','2');

-- ----------------------------
-- After insert for {TABLE}
-- ----------------------------
{AFTER_INSERT}
