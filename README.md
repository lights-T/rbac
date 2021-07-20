### 后台rbac权限

自构建rbac权限后台

### 特点

- 支持多角色用户的CRUD管理，更方便灵活。
- 支持后台结合权限生成菜单，更加隐秘。
- 支持后台菜单排序，更加易于操作。
- 支持三级菜单显示。
- 结合redis缓存，减少性能上的不必要浪费。

### sql

> mysql

```
CREATE TABLE `auth_group` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL COMMENT '身份名称',
  `rules` text NOT NULL COMMENT '权限动作id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unindex_title` (`title`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户权限组表';

CREATE TABLE `auth_group_access` (
  `uid` int(8) unsigned NOT NULL COMMENT '用户id ',
  `group_id` int(8) unsigned NOT NULL COMMENT '身份id',
  `is_delete` tinyint(1) DEFAULT '0' COMMENT '是否删除 1删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE KEY `uid_group_id` (`uid`,`group_id`) USING BTREE,
  KEY `group_id` (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户权限身份关系表';

CREATE TABLE `auth_rule` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `url_path` varchar(80) NOT NULL COMMENT '规则唯一标识路径 模块/方法',
  `title` varchar(20) NOT NULL COMMENT '规则中文名称 要语义化定义，给运营人员使用',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型 1一级菜单 2二级菜单 3三级菜单',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 1删除',
  `pid` smallint(5) NOT NULL DEFAULT '0' COMMENT '父级ID',
  `sort` tinyint(4) NOT NULL DEFAULT '50' COMMENT '排序，用于前端菜单显示，越大显示越靠前',
  `icon` varchar(50) DEFAULT '' COMMENT '菜单图标',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`url_path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户权限动作表';
```

> pg

```
CREATE SEQUENCE IF NOT EXISTS t_auth_group_id_seq;
CREATE TABLE "public"."t_auth_group" (
  "id" int8 NOT NULL DEFAULT nextval('t_auth_group_id_seq'::regclass),
  "title" varchar(50) COLLATE "pg_catalog"."default",
  "rules" text COLLATE "pg_catalog"."default",
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "t_auth_group_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "t_auth_group_unindex_title" UNIQUE ("title")
)
;

ALTER TABLE "public"."t_auth_group" 
  OWNER TO "postgres";

COMMENT ON COLUMN "public"."t_auth_group"."title" IS '身份名称';

COMMENT ON COLUMN "public"."t_auth_group"."rules" IS '权限动作id';

COMMENT ON TABLE "public"."t_auth_group" IS '用户权限组表';


CREATE TABLE "public"."t_auth_group_access" (
  "uid" int8 NOT NULL,
  "group_id" int8 NOT NULL,
  "is_delete" int4 NOT NULL DEFAULT 0,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "t_auth_group_access_unindex_uid_groupId" UNIQUE ("uid", "group_id"),
  CONSTRAINT "t_auth_group_access_unindex_groupId" UNIQUE ("group_id")
)
;

ALTER TABLE "public"."t_auth_group_access" 
  OWNER TO "postgres";

COMMENT ON COLUMN "public"."t_auth_group_access"."uid" IS '用户id';

COMMENT ON COLUMN "public"."t_auth_group_access"."group_id" IS '身份id';

COMMENT ON COLUMN "public"."t_auth_group_access"."is_delete" IS '是否删除 1删除';

COMMENT ON TABLE "public"."t_auth_group_access" IS '用户权限身份关系表';



CREATE SEQUENCE IF NOT EXISTS t_auth_rule_id_seq;
CREATE TABLE "public"."t_auth_rule" (
  "id" int8 NOT NULL DEFAULT nextval('t_auth_rule_id_seq'::regclass),
  "url_path" varchar(50) COLLATE "pg_catalog"."default",
  "title" varchar(20) COLLATE "pg_catalog"."default",
  "type" int4 NOT NULL DEFAULT 1,
  "is_delete" int4 NOT NULL DEFAULT 0,
  "pid" int8 NOT NULL DEFAULT 0,
  "sort" int4 NOT NULL DEFAULT 50,
  "icon" varchar(50) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "t_auth_rule_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "t_auth_group_unindex_urlPath" UNIQUE ("url_path")
)
;

ALTER TABLE "public"."t_auth_rule" 
  OWNER TO "postgres";

COMMENT ON COLUMN "public"."t_auth_rule"."url_path" IS '规则唯一标识路径 模块/方法';

COMMENT ON COLUMN "public"."t_auth_rule"."title" IS '规则中文名称 要逻辑化定义，给运营人员使用';

COMMENT ON COLUMN "public"."t_auth_rule"."type" IS '类型 1一级菜单 2二级菜单 3三级菜单';

COMMENT ON COLUMN "public"."t_auth_rule"."is_delete" IS '是否删除 1删除';

COMMENT ON COLUMN "public"."t_auth_rule"."pid" IS '父级ID';

COMMENT ON COLUMN "public"."t_auth_rule"."sort" IS '排序，大的在前';

COMMENT ON COLUMN "public"."t_auth_rule"."icon" IS '图标';

COMMENT ON TABLE "public"."t_auth_rule" IS '用户权限动作表';


```
