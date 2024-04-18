-- +goose Up
PRAGMA foreign_keys = false;


DROP TABLE IF EXISTS "USER";
CREATE TABLE "USER" (
  "id" integer NOT NULL,
  "USER" TEXT NOT NULL,
  "PASS" text NOT NULL,
  PRIMARY KEY ("id")
);


INSERT INTO "USER" VALUES (1, 'venus', '4f98416a6d70405114960fdcef0bc3e5');

DROP TABLE IF EXISTS "burplog";
CREATE TABLE "burplog" (
  "id" INTEGER NOT NULL,
  "time" text,
  "clientIP" text,
  "statusCode" text(3),
  "reqMethod" TEXT(5),
  "reqUri" TEXT,
  "full_message" TEXT,
  PRIMARY KEY ("id")
);


DROP TABLE IF EXISTS "gobylog";
CREATE TABLE "gobylog" (
  "id" INTEGER NOT NULL,
  "time" text,
  "clientIP" text,
  "statusCode" text(3),
  "reqMethod" TEXT(5),
  "reqUri" TEXT,
  "full_message" TEXT,
  PRIMARY KEY ("id")
);


DROP TABLE IF EXISTS "log";
CREATE TABLE "log" (
  "id" INTEGER NOT NULL,
  "time" text,
  "clientIP" text,
  "statusCode" text(3),
  "reqMethod" TEXT(5),
  "reqUri" TEXT,
  "full_message" TEXT,
  PRIMARY KEY ("id")
);


DROP TABLE IF EXISTS "mysqllog";
CREATE TABLE "mysqllog" (
  "id" INTEGER NOT NULL,
  "time" TEXT,
  "msg" TEXT,
  PRIMARY KEY ("id")
);


DROP TABLE IF EXISTS "pot";
CREATE TABLE "pot" (
  "id" INTEGER NOT NULL,
  "name" text NOT NULL,
  "pottype" TEXT NOT NULL,
  "state" integer NOT NULL,
  "url" TEXT NOT NULL,
  "configid" integer NOT NULL,
  PRIMARY KEY ("id")
);


INSERT INTO "pot" VALUES (1, 'BurpSuite蜜罐', '反制/命令执行', 0, 'burplog', 1001);
INSERT INTO "pot" VALUES (2, 'VPN蜜罐', '诱导/反制', 0, 'vpnlog', 1002);
INSERT INTO "pot" VALUES (3, 'Goby蜜罐', '反制/命令执行', 0, 'gobylog', 1003);
INSERT INTO "pot" VALUES (4, 'Mysql蜜罐', '文件读取', 0, 'mysqllog', 1004);

DROP TABLE IF EXISTS "pot_config";
CREATE TABLE "pot_config" (
                              "configid" INTEGER(4) NOT NULL,
                              "port" integer(5) NOT NULL,
                              "payload" text NOT NULL,
                              "fileexists" integer(1) NOT NULL DEFAULT 0,
                              "username" TEXT NOT NULL,
                              "password" TEXT NOT NULL DEFAULT '',
                              "filelist" TEXT NOT NULL DEFAULT '',
                              "ip" TEXT NOT NULL DEFAULT '',
                              PRIMARY KEY ("configid")
);

INSERT INTO "pot_config" VALUES (1004, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO "pot_config" VALUES (1003, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO "pot_config" VALUES (1002, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO "pot_config" VALUES (1001, 0, 0, 0, 0, 0, 0, 0);

DROP TABLE IF EXISTS "vpnlog";
CREATE TABLE "vpnlog" (
  "id" INTEGER NOT NULL,
  "time" text,
  "clientIP" text,
  "statusCode" text(3),
  "reqMethod" TEXT(5),
  "reqUri" TEXT,
  "full_message" TEXT,
  PRIMARY KEY ("id")
);
