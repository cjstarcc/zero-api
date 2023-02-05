create table if not exists category
(
    id          smallint unsigned auto_increment comment '分类id'
    primary key,
    parentid    smallint    default 0                 not null comment '父类别id当id=0时说明是根节点,一级类别',
    name        varchar(50) default ''                not null comment '类别名称',
    status      tinyint     default 1                 not null comment '类别状态1-正常,2-已废弃',
    create_time datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
    )
    comment '商品类别表' charset = utf8mb4;

create table if not exists product
(
    id          bigint unsigned auto_increment comment '商品id'
    primary key,
    cateid      smallint unsigned default '0'               not null comment '类别Id',
    name        varchar(100)      default ''                not null comment '商品名称',
    subtitle    varchar(200)      default ''                not null comment '商品副标题',
    images      varchar(1024)     default ''                not null comment '图片地址,逗号分隔',
    detail      varchar(1024)     default ''                not null comment '商品详情',
    price       decima l(20, 2)    default 0.00              not null comment '价格,单位-元保留两位小数',
    stock       int               default 0                 not null comment '库存数量',
    status      int               default 1                 not null comment '商品状态.1-在售 2-下架 3-删除',
    create_time datetime          default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime          default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
    )
    comment '商品表' charset = utf8mb4;

create index ix_cateid
    on product (cateid);

create index ix_update_time
    on product (update_time);

create table if not exists user
(
    id          bigint unsigned auto_increment comment '用户ID'
    primary key,
    username    varchar(50)  default ''                not null comment '用户名',
    password    varchar(50)  default ''                not null comment '用户密码，MD5加密',
    phone       varchar(20)  default ''                not null comment '手机号',
    question    varchar(100) default ''                not null comment '找回密码问题',
    answer      varchar(100) default ''                not null comment '找回密码答案',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint uniq_phone
    unique (phone),
    constraint uniq_username
    unique (username)
    )
    comment '用户表' charset = utf8mb4;

create index ix_update_time
    on user (update_time);

