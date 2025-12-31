# 数据库设计文档

## 概述

本数据库为2048小游戏排行榜系统设计，支持成绩记录、排行榜查询和最近记录查询功能。

## 表结构设计

### 1. 主表：grade_records（成绩记录表）

| 字段名 | 数据类型 | 约束 | 说明 |
|--------|----------|------|------|
| id | BIGINT | PRIMARY KEY, AUTO_INCREMENT | 主键ID |
| name | VARCHAR(50) | NOT NULL | 玩家名称 |
| score | INT | NOT NULL | 游戏分数 |
| comment | TEXT | NULL | 玩家评论 |
| finished | DATETIME | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 完成时间 |
| created_at | DATETIME | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME | NOT NULL, DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |

### 2. 视图：rank_list（排行榜视图）

按分数降序排列的排行榜视图，包含排名位置。

### 3. 视图：recent_records（最近记录视图）

按完成时间降序排列的最近记录视图。

## 索引设计

### 主键索引
- `PRIMARY KEY (id)` - 主键索引

### 辅助索引
- `idx_score (score)` - 分数索引，用于排行榜排序
- `idx_finished (finished)` - 完成时间索引，用于最近记录查询  
- `idx_name (name)` - 玩家名称索引
- `idx_score_finished (score DESC, finished DESC)` - 复合索引，优化排序查询

## 查询示例

### 1. 获取排行榜（按分数排序）
```sql
SELECT id, name, score, comment, finished
FROM grade_records 
ORDER BY score DESC 
LIMIT 50;
```

### 2. 获取最近记录（按时间排序）
```sql
SELECT id, name, score, comment, finished
FROM grade_records 
ORDER BY finished DESC 
LIMIT 50;
```

### 3. 插入新记录
```sql
INSERT INTO grade_records (name, score, comment) 
VALUES (?, ?, ?);
```

### 4. 获取玩家最高分数
```sql
SELECT MAX(score) as max_score
FROM grade_records 
WHERE name = ?;
```

### 5. 获取分数统计
```sql
SELECT 
    COUNT(*) as total_records,
    AVG(score) as avg_score,
    MAX(score) as max_score,
    MIN(score) as min_score
FROM grade_records;
```

## 性能优化

### 索引优化
- 为常用查询字段建立索引
- 使用复合索引优化排序查询
- 定期分析和优化表

### 查询优化
- 使用LIMIT限制返回记录数
- 避免全表扫描
- 使用覆盖索引减少回表

## 数据维护

### 定期清理（可选）
可以定期清理过期的记录，例如只保留最近1年的数据：
```sql
DELETE FROM grade_records 
WHERE finished < DATE_SUB(NOW(), INTERVAL 1 YEAR);
```

### 数据备份
建议定期备份数据库，特别是成绩记录表：
```bash
mysqldump -u username -p minigame grade_records > grade_records_backup.sql
```

## 扩展设计

### 未来可能的扩展
1. **玩家表**：单独维护玩家信息，支持用户注册登录
2. **游戏记录表**：记录每局游戏的详细过程
3. **分类表**：支持不同游戏模式或难度等级
4. **统计表**：预计算的统计数据，提高查询性能

### 字段扩展建议
- 添加 `game_mode` 字段支持不同游戏模式
- 添加 `difficulty` 字段支持不同难度等级
- 添加 `duration` 字段记录游戏时长
- 添加 `moves` 字段记录移动步数

## 安全配置

### 数据库用户权限
建议创建专用数据库用户，只授予必要的权限：
```sql
CREATE USER 'minigame_user'@'%' IDENTIFIED BY 'strong_password';
GRANT SELECT, INSERT, UPDATE, DELETE ON minigame.grade_records TO 'minigame_user'@'%';
FLUSH PRIVILEGES;
```

### 连接安全
- 使用SSL/TLS加密连接
- 限制数据库访问IP
- 定期更新数据库密码