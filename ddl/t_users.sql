CREATE TABLE `t_users`
(
    `f_id`         int        NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `f_name`       varchar(8) NOT NULL DEFAULT '' COMMENT '用户名',
    `f_del_flag`   tinyint(1) DEFAULT '0' COMMENT '删除标记，false为未删除，true为已删除；默认为false',
    `f_created_at` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `f_updated_at` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`f_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表'