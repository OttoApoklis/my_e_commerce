create table `seckill_stock` (
                                   `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                                   `goods_id`                  bigint(20)   comment '商品ID',
                                   `stock`                 int(11)      not null                comment '库存大小',
                                   `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                   `modify_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `idx_goodsid` (`goods_id`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀库存表' ;



create table `seckill_record` (
                                    `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                                    `user_id`                  bigint(20)  not null comment '用户ID',
                                    `goods_id`                  bigint(20) not null  comment '商品ID',
                                    `sec_num`               varchar(128)                         comment '秒杀号',
                                    `order_num`               varchar(128)                         comment '订单号',
                                    `price`                 decimal      not null                comment '金额',
                                    `status`                 int(11)      not null                comment '状态',
                                    `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                    `modify_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `idx_secnum` (`sec_num`),
                                    UNIQUE KEY `idx_ordernum` (`order_num`),
                                    KEY `idx_userid` (`user_id`),
                                    KEY `idx_userid_create` (`user_id`, `create_time`, `sec_num`, `order_num`,`goods_id`) -- 考虑需要根据用户查询订单并按创建时间排序，查询信息包含订单创建时间、商品信息、订单信息和秒杀信息，这样设计可以减少回表
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀记录表' ;


create table `goods` (
                           `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                           `goods_num`               varchar(128)                         comment '商品编号',
                           `goods_name`               varchar(128)                         comment '商品名字',
                           `price`                 decimal      not null                comment '商品单价',
                           `pic_url`               varchar(128)                         comment '商品图片',
                           `seller`                  bigint(20)  not null comment '卖家ID',
                           `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                           `modify_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `idx_goodsnum` (`goods_num`),
                           KEY `idx_seller` (`seller`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '商品表' ;

create table `order` (
                           `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                           `seller`                  bigint(20)  not null comment '买方ID',
                           `buyer`                  bigint(20)  not null comment '卖房ID',
                           `goods_id`                  bigint(20) not null  comment '商品ID',
                           `goods_num`               varchar(128)                         comment '商品编号',
                           `order_num`               varchar(128)                         comment '订单号',
                           `goods_amount`              int(11)                            comment '商品数量',
                           `price`                 decimal     not null                comment '金额',
                           `status`                 int(11)      not null                comment '状态',
                           `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                           `modify_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_goodsid` (`goods_id`),
                           UNIQUE KEY `idx_ordernum` (`order_num`),
                           KEY `idx_seller` (`seller`),
                           KEY `idx_buyer` (`buyer`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '订单表' ;


insert into goods(goods_name, goods_num, price, pic_url, seller) values("redhat", "abc123", 18, "http://", 135);

insert into seckill_stock(goods_id, stock) values(1, 3);

create table `quota` (
                           `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                           `goods_id`                  bigint(20)   comment '商品ID',
                           `num`                 int(11)      not null                comment '限额',
                           `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                           `modify_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `idx_goodsid` (`goods_id`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '限额表' ;

create table `user_quota` (
                                `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                                `user_id`                  bigint(20)  not null comment '用户ID',
                                `goods_id`                  bigint(20)   comment '商品ID',
                                `num`                 int(11)      not null         comment '限额',
                                `killed_num`          int(11)      not null         comment '已经消耗的额度',
                                `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                `modify_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                PRIMARY KEY (`id`),
                                KEY `idx_goodsid` (`goods_id`),
                                KEY `idx_usergoodsid` (`user_id`, `goods_id`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '用户限额表' ;
