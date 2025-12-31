CREATE TABLE IF NOT EXISTS grade (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    name VARCHAR(50) NOT NULL COMMENT '玩家名称',
    score INT NOT NULL COMMENT '游戏分数',
    comment TEXT COMMENT '玩家评论',
    finished DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '完成时间',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id),
    INDEX idx_score (score) COMMENT '分数索引，用于排行榜排序',
    INDEX idx_finished (finished) COMMENT '完成时间索引，用于最近记录查询'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='游戏成绩记录表';

INSERT INTO grade (name, score, comment) VALUES
('张三', 2048, '终于通关了！'),
('李四', 1024, '还不错，继续加油'),
('王五', 512, '第一次玩，挺有趣的'),
('赵六', 4096, '破纪录了！'),
('钱七', 256, '需要多练习'),
('孙八', 8192, '高手在此'),
('周九', 128, '新手报道'),
('吴十', 16384, '无敌是多么寂寞');