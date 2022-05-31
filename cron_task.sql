/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 127.0.0.1:3306
 Source Schema         : cron_task

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 31/05/2022 15:06:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父集',
  `orders` int(11) NOT NULL DEFAULT '0' COMMENT '排序\n',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'icon',
  `uri` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地址',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `is_category` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为栏目\n',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='栏目表';

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` VALUES (1, 0, 500, '主页展示', 'fa-bar-chart', '', '2022-03-25 10:50:08', '2022-03-31 18:34:56', 1);
INSERT INTO `admin_menu` VALUES (2, 0, 100, '用户管理', 'fa-tasks', '', '2022-03-25 10:50:08', '2022-03-25 10:50:08', 1);
INSERT INTO `admin_menu` VALUES (3, 2, 3, '用户列表', 'fa-users', 'user/list', '2022-03-25 10:50:08', '2022-03-25 10:50:08', 1);
INSERT INTO `admin_menu` VALUES (4, 2, 4, '角色列表', 'fa-user', 'roles/list', '2022-03-25 10:50:08', '2022-03-25 10:50:08', 1);
INSERT INTO `admin_menu` VALUES (6, 2, 6, '菜单列表', 'fa-bars', 'menu/list', '2022-03-25 10:50:08', '2022-03-25 10:50:08', 1);
INSERT INTO `admin_menu` VALUES (7, 8, 0, '任务日志', 'fa-cloud', 'log/list', '2022-03-25 10:50:08', '2022-04-02 14:34:16', 1);
INSERT INTO `admin_menu` VALUES (8, 0, 200, '任务管理', 'fa-list-ul', 'tasks', '2022-03-25 17:19:27', '2022-04-02 14:35:22', 1);
INSERT INTO `admin_menu` VALUES (15, 3, 0, '添加页面', '', 'user/add', NULL, NULL, 0);
INSERT INTO `admin_menu` VALUES (17, 3, 0, '编辑页面', '', 'user/edit', NULL, NULL, 0);
INSERT INTO `admin_menu` VALUES (19, 3, 0, '删除用户', '', 'user/del', NULL, NULL, 0);
INSERT INTO `admin_menu` VALUES (21, 4, 0, '角色添加', '', 'roles/add', '2022-04-01 17:59:32', '2022-04-01 17:59:32', 0);
INSERT INTO `admin_menu` VALUES (22, 4, 0, '角色编辑', '', 'roles/edit', '2022-04-01 17:59:57', '2022-04-01 17:59:57', 0);
INSERT INTO `admin_menu` VALUES (23, 4, 0, '角色删除', '', 'roles/del', '2022-04-01 18:00:41', '2022-04-01 18:00:41', 0);
INSERT INTO `admin_menu` VALUES (24, 4, 0, '权限添加', '', 'roles/roles', '2022-04-01 18:01:19', '2022-04-01 18:01:19', 0);
INSERT INTO `admin_menu` VALUES (25, 6, 0, '菜单添加', '', 'menu/add', '2022-04-01 18:20:14', '2022-04-01 18:23:48', 0);
INSERT INTO `admin_menu` VALUES (26, 6, 0, '菜单编辑', '', 'menu/edit', '2022-04-01 18:21:51', '2022-04-01 18:21:51', 0);
INSERT INTO `admin_menu` VALUES (27, 6, 0, '删除菜单', '', 'menu/del', '2022-04-01 18:22:23', '2022-04-01 18:22:23', 0);
INSERT INTO `admin_menu` VALUES (37, 8, 20, '任务列表', 'fa-tasks', 'task/list', '2022-04-02 14:25:26', '2022-04-02 14:26:26', 1);
INSERT INTO `admin_menu` VALUES (38, 37, 0, '添加任务', 'fa-adjust', 'task/add', '2022-04-02 14:27:01', '2022-04-02 14:27:01', 0);
INSERT INTO `admin_menu` VALUES (39, 37, 0, '编辑任务', 'fa-adjust', 'task/edit', '2022-04-02 14:27:52', '2022-04-02 14:27:52', 0);
INSERT INTO `admin_menu` VALUES (40, 37, 0, '删除任务', 'fa-adjust', 'task/del', '2022-04-02 14:28:29', '2022-04-02 14:28:29', 0);
INSERT INTO `admin_menu` VALUES (41, 37, 0, '开启任务', 'fa-adjust', 'task/start', '2022-04-02 14:29:48', '2022-04-20 16:13:42', 0);
INSERT INTO `admin_menu` VALUES (42, 1, 0, '首页展示', 'fa-adjust', 'page/welcome', '2022-04-02 15:16:15', '2022-04-02 15:32:41', 1);
INSERT INTO `admin_menu` VALUES (43, 8, 8, '任务类型', 'fa-adjust', 'group/list', '2022-04-06 16:53:47', '2022-04-06 16:53:47', 1);
INSERT INTO `admin_menu` VALUES (44, 43, 0, '分类添加', 'fa-adjust', 'group/add', '2022-04-06 16:54:24', '2022-04-06 16:54:24', 0);
INSERT INTO `admin_menu` VALUES (45, 43, 0, '分类编辑', 'fa-adjust', 'group/edit', '2022-04-06 16:54:54', '2022-04-06 16:54:54', 0);
INSERT INTO `admin_menu` VALUES (46, 43, 0, '分类删除', 'fa-adjust', 'group/del', '2022-04-06 16:55:22', '2022-04-06 16:55:22', 0);
INSERT INTO `admin_menu` VALUES (47, 7, 0, '日志详情', 'fa-adjust', 'log/detail', '2022-04-22 10:23:54', '2022-04-22 10:23:54', 0);
COMMIT;

-- ----------------------------
-- Table structure for admin_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_permissions`;
CREATE TABLE `admin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`),
  UNIQUE KEY `admin_permissions_slug_unique` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_permissions
-- ----------------------------
BEGIN;
INSERT INTO `admin_permissions` VALUES (1, 'All permission', '*', '', '*', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (2, 'Dashboard', 'dashboard', 'GET', '/', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (3, 'Login', 'auth.login', '', '/auth/login\r\n/auth/logout', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (4, 'User setting', 'auth.setting', 'GET,PUT', '/auth/setting', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (5, 'Auth management', 'auth.management', '', '/auth/roles\r\n/auth/permissions\r\n/auth/menu\r\n/auth/logs', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_menu`;
CREATE TABLE `admin_role_menu` (
  `role_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_menu` VALUES (2, 1, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (2, 9, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (2, 2, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (2, 13, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (2, 5, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (2, 7, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (2, 8, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (2, 12, '2022-03-30 18:07:21', '2022-03-30 18:07:21');
INSERT INTO `admin_role_menu` VALUES (1, 1, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 42, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 2, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 3, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 15, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 17, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 19, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 4, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 21, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 22, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 23, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 24, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 6, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 25, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 26, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 8, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 7, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 47, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 37, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 38, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 39, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 40, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 41, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 43, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 44, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 45, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
INSERT INTO `admin_role_menu` VALUES (1, 46, '2022-04-22 10:29:53', '2022-04-22 10:29:53');
COMMIT;

-- ----------------------------
-- Table structure for admin_role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permissions`;
CREATE TABLE `admin_role_permissions` (
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `admin_role_permissions_role_id_permission_id_index` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_permissions
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_permissions` VALUES (1, 1, NULL, NULL);
INSERT INTO `admin_role_permissions` VALUES (1, 2, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_users`;
CREATE TABLE `admin_role_users` (
  `role_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `admin_role_users_role_id_user_id_index` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_users` VALUES (1, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`),
  UNIQUE KEY `admin_roles_slug_unique` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
BEGIN;
INSERT INTO `admin_roles` VALUES (1, 'Administrator', 'administrator', '2022-03-19 13:33:51', '2022-03-19 13:33:51');
INSERT INTO `admin_roles` VALUES (2, '管理员', '管理员', '2022-03-24 17:40:18', '2022-03-24 17:40:18');
INSERT INTO `admin_roles` VALUES (3, '普通用户', '普通用户', '2022-04-02 11:45:55', '2022-04-02 11:45:55');
COMMIT;

-- ----------------------------
-- Table structure for admin_task
-- ----------------------------
DROP TABLE IF EXISTS `admin_task`;
CREATE TABLE `admin_task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '分组ID',
  `task_name` varchar(50) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '任务名称',
  `task_type` char(4) NOT NULL COMMENT '任务类型 GET  POST',
  `iplong` int(11) NOT NULL DEFAULT '0' COMMENT 'ip转的的int',
  `http_url` varchar(512) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '执行url\n',
  `host` varchar(16) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '请求host',
  `description` varchar(200) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '任务描述',
  `cron_spec` varchar(100) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '时间表达式',
  `concurrent` tinyint(1) NOT NULL DEFAULT '1' COMMENT '同一个任务是否允许并行执行',
  `command` varchar(512) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '参数',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0停用 1启用',
  `notify` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通知设置',
  `notify_email` varchar(256) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '通知人列表',
  `timeout` smallint(6) NOT NULL DEFAULT '0' COMMENT '超时设置',
  `execute_count` int(11) NOT NULL DEFAULT '0' COMMENT '累计执行次数',
  `prev_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上次执行时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_group_id` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of admin_task
-- ----------------------------
BEGIN;
INSERT INTO `admin_task` VALUES (4, 1, 1, '测试任务', 'GET', 181410010, 'http://www.baidu.com', '', '', '*/10 * * * * *', 0, '', 1, 1, 'yinglun@aa.com', 0, 156, 1651130390, '2022-04-28 15:19:50', '2022-04-28 15:19:50');
INSERT INTO `admin_task` VALUES (5, 1, 2, '新建任务名称', 'GET', 181410010, 'http://www.baidu.com', '', '', '*/10 * * * * *', 0, '', 1, 0, '', 0, 157, 1653980760, '2022-05-31 15:06:00', '2022-05-31 15:06:00');
COMMIT;

-- ----------------------------
-- Table structure for admin_task_group
-- ----------------------------
DROP TABLE IF EXISTS `admin_task_group`;
CREATE TABLE `admin_task_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `group_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '说明',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_task_group
-- ----------------------------
BEGIN;
INSERT INTO `admin_task_group` VALUES (1, 1, '发布系统', '', '2022-04-06 17:20:04');
INSERT INTO `admin_task_group` VALUES (2, 1, '消息系统', '', '2022-04-06 17:30:55');
COMMIT;

-- ----------------------------
-- Table structure for admin_task_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_task_log`;
CREATE TABLE `admin_task_log` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '任务ID',
  `httpcode` int(5) NOT NULL DEFAULT '0' COMMENT 'http状态码',
  `output` mediumtext NOT NULL COMMENT '任务输出',
  `error` text NOT NULL COMMENT '错误信息',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `process_time` int(11) NOT NULL DEFAULT '0' COMMENT '消耗时间/毫秒',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_task_id` (`task_id`,`created_at`),
  KEY `user_id` (`user_id`,`task_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8323 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin_task_log
-- ----------------------------
BEGIN;
INSERT INTO `admin_task_log` VALUES (8238, 1, 5, 200, '[]', '无', 0, 56, '2022-05-31 14:52:00', '2022-05-31 14:52:00');
INSERT INTO `admin_task_log` VALUES (8239, 1, 5, 200, '[]', '无', 0, 600, '2022-05-31 14:52:10', '2022-05-31 14:52:10');
INSERT INTO `admin_task_log` VALUES (8240, 1, 5, 200, '[]', '无', 0, 32, '2022-05-31 14:52:20', '2022-05-31 14:52:20');
INSERT INTO `admin_task_log` VALUES (8241, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:52:30', '2022-05-31 14:52:30');
INSERT INTO `admin_task_log` VALUES (8242, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:52:40', '2022-05-31 14:52:40');
INSERT INTO `admin_task_log` VALUES (8243, 1, 5, 200, '[]', '无', 0, 22, '2022-05-31 14:52:50', '2022-05-31 14:52:50');
INSERT INTO `admin_task_log` VALUES (8244, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:53:00', '2022-05-31 14:53:00');
INSERT INTO `admin_task_log` VALUES (8245, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:53:10', '2022-05-31 14:53:10');
INSERT INTO `admin_task_log` VALUES (8246, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:53:20', '2022-05-31 14:53:20');
INSERT INTO `admin_task_log` VALUES (8247, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:53:30', '2022-05-31 14:53:30');
INSERT INTO `admin_task_log` VALUES (8248, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:53:40', '2022-05-31 14:53:40');
INSERT INTO `admin_task_log` VALUES (8249, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:53:50', '2022-05-31 14:53:50');
INSERT INTO `admin_task_log` VALUES (8250, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:54:00', '2022-05-31 14:54:00');
INSERT INTO `admin_task_log` VALUES (8251, 1, 5, 200, '[]', '无', 0, 9, '2022-05-31 14:54:10', '2022-05-31 14:54:10');
INSERT INTO `admin_task_log` VALUES (8252, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:54:20', '2022-05-31 14:54:20');
INSERT INTO `admin_task_log` VALUES (8253, 1, 5, 200, '[]', '无', 0, 26, '2022-05-31 14:54:30', '2022-05-31 14:54:30');
INSERT INTO `admin_task_log` VALUES (8254, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:54:40', '2022-05-31 14:54:40');
INSERT INTO `admin_task_log` VALUES (8255, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:54:50', '2022-05-31 14:54:50');
INSERT INTO `admin_task_log` VALUES (8256, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:55:00', '2022-05-31 14:55:00');
INSERT INTO `admin_task_log` VALUES (8257, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:55:10', '2022-05-31 14:55:10');
INSERT INTO `admin_task_log` VALUES (8258, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:55:20', '2022-05-31 14:55:20');
INSERT INTO `admin_task_log` VALUES (8259, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:55:30', '2022-05-31 14:55:30');
INSERT INTO `admin_task_log` VALUES (8260, 1, 5, 200, '[]', '无', 0, 49, '2022-05-31 14:55:40', '2022-05-31 14:55:40');
INSERT INTO `admin_task_log` VALUES (8261, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:55:50', '2022-05-31 14:55:50');
INSERT INTO `admin_task_log` VALUES (8262, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:56:00', '2022-05-31 14:56:00');
INSERT INTO `admin_task_log` VALUES (8263, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:56:10', '2022-05-31 14:56:10');
INSERT INTO `admin_task_log` VALUES (8264, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:56:20', '2022-05-31 14:56:20');
INSERT INTO `admin_task_log` VALUES (8265, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:56:30', '2022-05-31 14:56:30');
INSERT INTO `admin_task_log` VALUES (8266, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:56:40', '2022-05-31 14:56:40');
INSERT INTO `admin_task_log` VALUES (8267, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:56:50', '2022-05-31 14:56:50');
INSERT INTO `admin_task_log` VALUES (8268, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:57:00', '2022-05-31 14:57:00');
INSERT INTO `admin_task_log` VALUES (8269, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:57:10', '2022-05-31 14:57:10');
INSERT INTO `admin_task_log` VALUES (8270, 1, 5, 200, '[]', '无', 0, 5, '2022-05-31 14:57:20', '2022-05-31 14:57:20');
INSERT INTO `admin_task_log` VALUES (8271, 1, 5, 200, '[]', '无', 0, 5, '2022-05-31 14:57:30', '2022-05-31 14:57:30');
INSERT INTO `admin_task_log` VALUES (8272, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:57:40', '2022-05-31 14:57:40');
INSERT INTO `admin_task_log` VALUES (8273, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:57:50', '2022-05-31 14:57:50');
INSERT INTO `admin_task_log` VALUES (8274, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:58:00', '2022-05-31 14:58:00');
INSERT INTO `admin_task_log` VALUES (8275, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:58:10', '2022-05-31 14:58:10');
INSERT INTO `admin_task_log` VALUES (8276, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:58:20', '2022-05-31 14:58:20');
INSERT INTO `admin_task_log` VALUES (8277, 1, 5, 200, '[]', '无', 0, 9, '2022-05-31 14:58:30', '2022-05-31 14:58:30');
INSERT INTO `admin_task_log` VALUES (8278, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 14:58:40', '2022-05-31 14:58:40');
INSERT INTO `admin_task_log` VALUES (8279, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:58:50', '2022-05-31 14:58:50');
INSERT INTO `admin_task_log` VALUES (8280, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:59:00', '2022-05-31 14:59:00');
INSERT INTO `admin_task_log` VALUES (8281, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:59:10', '2022-05-31 14:59:10');
INSERT INTO `admin_task_log` VALUES (8282, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:59:20', '2022-05-31 14:59:20');
INSERT INTO `admin_task_log` VALUES (8283, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 14:59:30', '2022-05-31 14:59:30');
INSERT INTO `admin_task_log` VALUES (8284, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:59:40', '2022-05-31 14:59:40');
INSERT INTO `admin_task_log` VALUES (8285, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 14:59:50', '2022-05-31 14:59:50');
INSERT INTO `admin_task_log` VALUES (8286, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:00:00', '2022-05-31 15:00:00');
INSERT INTO `admin_task_log` VALUES (8287, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:00:10', '2022-05-31 15:00:10');
INSERT INTO `admin_task_log` VALUES (8288, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 15:00:20', '2022-05-31 15:00:20');
INSERT INTO `admin_task_log` VALUES (8289, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 15:00:30', '2022-05-31 15:00:30');
INSERT INTO `admin_task_log` VALUES (8290, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:00:40', '2022-05-31 15:00:40');
INSERT INTO `admin_task_log` VALUES (8291, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:00:50', '2022-05-31 15:00:50');
INSERT INTO `admin_task_log` VALUES (8292, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 15:01:00', '2022-05-31 15:01:00');
INSERT INTO `admin_task_log` VALUES (8293, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:01:10', '2022-05-31 15:01:10');
INSERT INTO `admin_task_log` VALUES (8294, 1, 5, 200, '[]', '无', 0, 10, '2022-05-31 15:01:20', '2022-05-31 15:01:20');
INSERT INTO `admin_task_log` VALUES (8295, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:01:30', '2022-05-31 15:01:30');
INSERT INTO `admin_task_log` VALUES (8296, 1, 5, 200, '[]', '无', 0, 9, '2022-05-31 15:01:40', '2022-05-31 15:01:40');
INSERT INTO `admin_task_log` VALUES (8297, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 15:01:50', '2022-05-31 15:01:50');
INSERT INTO `admin_task_log` VALUES (8298, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:02:00', '2022-05-31 15:02:00');
INSERT INTO `admin_task_log` VALUES (8299, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:02:10', '2022-05-31 15:02:10');
INSERT INTO `admin_task_log` VALUES (8300, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:02:20', '2022-05-31 15:02:20');
INSERT INTO `admin_task_log` VALUES (8301, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:02:30', '2022-05-31 15:02:30');
INSERT INTO `admin_task_log` VALUES (8302, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:02:40', '2022-05-31 15:02:40');
INSERT INTO `admin_task_log` VALUES (8303, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:02:50', '2022-05-31 15:02:50');
INSERT INTO `admin_task_log` VALUES (8304, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:03:00', '2022-05-31 15:03:00');
INSERT INTO `admin_task_log` VALUES (8305, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:03:10', '2022-05-31 15:03:10');
INSERT INTO `admin_task_log` VALUES (8306, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:03:20', '2022-05-31 15:03:20');
INSERT INTO `admin_task_log` VALUES (8307, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:03:30', '2022-05-31 15:03:30');
INSERT INTO `admin_task_log` VALUES (8308, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:03:40', '2022-05-31 15:03:40');
INSERT INTO `admin_task_log` VALUES (8309, 1, 5, 200, '[]', '无', 0, 9, '2022-05-31 15:03:50', '2022-05-31 15:03:50');
INSERT INTO `admin_task_log` VALUES (8310, 1, 5, 200, '[]', '无', 0, 8, '2022-05-31 15:04:00', '2022-05-31 15:04:00');
INSERT INTO `admin_task_log` VALUES (8311, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:04:10', '2022-05-31 15:04:10');
INSERT INTO `admin_task_log` VALUES (8312, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:04:20', '2022-05-31 15:04:20');
INSERT INTO `admin_task_log` VALUES (8313, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:04:30', '2022-05-31 15:04:30');
INSERT INTO `admin_task_log` VALUES (8314, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:04:40', '2022-05-31 15:04:40');
INSERT INTO `admin_task_log` VALUES (8315, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:04:50', '2022-05-31 15:04:50');
INSERT INTO `admin_task_log` VALUES (8316, 1, 5, 200, '[]', '无', 0, 9, '2022-05-31 15:05:00', '2022-05-31 15:05:00');
INSERT INTO `admin_task_log` VALUES (8317, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:05:10', '2022-05-31 15:05:10');
INSERT INTO `admin_task_log` VALUES (8318, 1, 5, 200, '[]', '无', 0, 6, '2022-05-31 15:05:20', '2022-05-31 15:05:20');
INSERT INTO `admin_task_log` VALUES (8319, 1, 5, 200, '[]', '无', 0, 9, '2022-05-31 15:05:30', '2022-05-31 15:05:30');
INSERT INTO `admin_task_log` VALUES (8320, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:05:40', '2022-05-31 15:05:40');
INSERT INTO `admin_task_log` VALUES (8321, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:05:50', '2022-05-31 15:05:50');
INSERT INTO `admin_task_log` VALUES (8322, 1, 5, 200, '[]', '无', 0, 7, '2022-05-31 15:06:00', '2022-05-31 15:06:00');
COMMIT;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `name` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '真实名称',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `mobile` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号码：',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '权限id',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态 1：正常用户；0 ：禁用',
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `remarks` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ----------------------------
-- Records of admin_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_users` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '隔壁老王', '1', '13522164134', 1, 1, 'yinglun@aa.com', '', '2022-03-19 13:33:51', '2022-04-24 11:29:16');
INSERT INTO `admin_users` VALUES (3, 'diaop', 'e10adc3949ba59abbe56e057f20f883e', '阿坝', '', '1234345555', 2, 1, 'yinglun@aa.com', 'fdgdsfdasf', '2022-03-23 18:56:03', '2022-03-21 18:56:03');
INSERT INTO `admin_users` VALUES (5, 'diaopdsfs', 'e10adc3949ba59abbe56e057f20f883e', '精编', '', '1234345555', 1, 1, 'yinglun@aa.com', 'dsafdsafdcafesdacx', '2022-03-23 11:51:39', '2022-03-23 11:51:39');
INSERT INTO `admin_users` VALUES (7, 'redsfacxc', 'e10adc3949ba59abbe56e057f20f883e', 'ddd', '', '1234345555', 2, 1, 'yinglun@aa.com', 'dsfdczx', '2022-03-23 11:55:32', '2022-04-01 16:54:25');
INSERT INTO `admin_users` VALUES (8, 'xiaomign', 'e10adc3949ba59abbe56e057f20f883e', '', '', '1234345555', 2, 1, 'yinglun@aa.com', 'rrr', '2022-04-02 11:45:04', '2022-04-02 11:52:42');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
