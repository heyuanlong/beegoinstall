

CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `keywords` varchar(255) DEFAULT '' COMMENT '关键词',
  `content` text COMMENT '正文',
  `author` varchar(50) NOT NULL DEFAULT '' COMMENT '作者',
  `time` varchar(19) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  `count` int(11) NOT NULL DEFAULT '0' COMMENT '阅读次数',
  `classifyID` int(11) DEFAULT '0',
  `status` int(4) NOT NULL DEFAULT '0' COMMENT '状态: 0草稿，1已发布',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8 COMMENT='文章';


CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户';

insert into users(username,password) values("klmmx","14e1b600b1fd579f47433b88e8d85291");   --123456

CREATE TABLE IF NOT EXISTS `classify`(
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '名称',
  UNIQUE KEY(`name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章分类';

insert into classify(name) values('歪里'),('山羊'),('暖锋'),('骆驼'),('文竹'),('七宝'),('玄衣'),('以衡'),('顶菇'),('豆芽'),('龙象');
insert into classify(name) values('default');
update classify set id=0 where name='default';

CREATE TABLE IF NOT EXISTS `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `articleid` int(11) NOT NULL,
  `content` text COMMENT '',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名字',
  `time` varchar(19) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '时间',
  PRIMARY KEY (`id`),
  KEY `articleid` (`articleid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论';
