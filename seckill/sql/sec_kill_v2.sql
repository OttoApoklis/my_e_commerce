-- seckill.goods definition

CREATE TABLE `goods` (
                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
                         `goods_num` varchar(128) DEFAULT NULL COMMENT '商品编号',
                         `goods_name` varchar(128) DEFAULT NULL COMMENT '商品名字',
                         `price` decimal(10,2) NOT NULL COMMENT '商品单价',
                         `pic_url` varchar(128) DEFAULT NULL COMMENT '商品图片',
                         `seller` bigint NOT NULL COMMENT '卖家ID',
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                         `deleted_at` datetime DEFAULT NULL COMMENT '软删除，标记时间可用于恢复',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_goodsnum` (`goods_num`),
                         KEY `idx_seller` (`seller`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品表';


-- seckill.`order` definition

CREATE TABLE `order` (
                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
                         `seller` bigint NOT NULL COMMENT '买方ID',
                         `buyer` bigint NOT NULL COMMENT '卖房ID',
                         `goods_id` bigint NOT NULL COMMENT '商品ID',
                         `goods_num` varchar(128) DEFAULT NULL COMMENT '商品编号',
                         `order_num` varchar(128) DEFAULT NULL COMMENT '订单号',
                         `goods_amount` int DEFAULT NULL COMMENT '商品数量',
                         `price` decimal(10,2) NOT NULL COMMENT '金额',
                         `status` tinyint NOT NULL COMMENT '状态',
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                         `deleted_at` datetime DEFAULT NULL COMMENT '软删除，标记时间可用于恢复',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_ordernum` (`order_num`),
                         KEY `idx_goodsid` (`goods_id`),
                         KEY `idx_seller` (`seller`),
                         KEY `idx_buyer` (`buyer`)
) ENGINE=InnoDB AUTO_INCREMENT=1811 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单表';


-- seckill.quota definition

CREATE TABLE `quota` (
                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
                         `goods_id` bigint DEFAULT NULL COMMENT '商品ID',
                         `num` int NOT NULL COMMENT '限额',
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                         `deleted_at` datetime DEFAULT NULL COMMENT '软删除，标记时间可用于恢复',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_goodsid` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='限额表';


-- seckill.seckill_record definition

CREATE TABLE `seckill_record` (
                                  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
                                  `user_id` bigint NOT NULL COMMENT '用户ID',
                                  `goods_id` bigint NOT NULL COMMENT '商品ID',
                                  `sec_num` varchar(128) DEFAULT NULL COMMENT '秒杀号',
                                  `order_num` varchar(128) DEFAULT NULL COMMENT '订单号',
                                  `price` decimal(10,2) NOT NULL COMMENT '金额',
                                  `status` tinyint NOT NULL COMMENT '状态',
                                  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                                  `deleted_at` datetime DEFAULT NULL COMMENT '软删除，标记时间可用于恢复',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `idx_secnum` (`sec_num`),
                                  UNIQUE KEY `idx_ordernum` (`order_num`),
                                  KEY `idx_userid` (`user_id`),
                                  KEY `idx_userid_create` (`user_id`,`create_time`,`sec_num`,`order_num`,`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='秒杀记录表';


-- seckill.seckill_stock definition

CREATE TABLE `seckill_stock` (
                                 `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
                                 `goods_id` bigint DEFAULT NULL COMMENT '商品ID',
                                 `stock` int NOT NULL COMMENT '库存大小',
                                 `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                                 `deleted_at` datetime DEFAULT NULL COMMENT '软删除，标记时间可用于恢复',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `idx_goodsid` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='秒杀库存表';


-- seckill.user_quota definition

CREATE TABLE `user_quota` (
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
                              `user_id` bigint NOT NULL COMMENT '用户ID',
                              `goods_id` bigint DEFAULT NULL COMMENT '商品ID',
                              `num` int NOT NULL COMMENT '限额',
                              `killed_num` int NOT NULL COMMENT '已经消耗的额度',
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                              `deleted_at` date DEFAULT NULL COMMENT '软删除，标记时间可用于恢复',
                              PRIMARY KEY (`id`),
                              KEY `idx_goodsid` (`goods_id`),
                              KEY `idx_usergoodsid` (`user_id`,`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户限额表';