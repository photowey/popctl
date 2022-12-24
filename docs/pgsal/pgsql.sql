-- 查询 schema 下所有的表
SELECT * FROM pg_tables WHERE schemaname = 'public';
SELECT tableName FROM pg_tables WHERE tableName NOT LIKE 'pg%' AND tableName NOT LIKE 'sql_%' ORDER BY tableName;

-- 查询主键
SELECT
    pg_attribute.attname AS columnName,
    pg_type.typname AS typeName,
    pg_constraint.conname AS pkName
FROM pg_constraint
         INNER JOIN pg_class ON pg_constraint.conrelid = pg_class.oid
         INNER JOIN pg_attribute ON pg_attribute.attrelid = pg_class.oid AND pg_attribute.attnum = pg_constraint.conkey[1]
         INNER JOIN pg_type ON pg_type.oid = pg_attribute.atttypid
WHERE pg_class.relname = 'platform_user'
    AND pg_constraint.contype='p';


-- 查询 表下面所有的字段及其相关属性
SELECT
    a.attnum AS numberz,
    c.relname AS tableName,
    cast(obj_description(relfilenode,'pg_class') as varchar) AS tableComment,
    a.attname AS namez,
    pg_type.typname AS typname,
    pg_type.typlen AS typlen,
    SUBSTRING(format_type(a.atttypid,a.atttypmod) from '\(.*\)') AS formatlength,
    col_description ( a.attrelid, a.attnum ) AS commentz,
    a.attnotnull AS notnullz
FROM
    pg_class AS c,
    pg_attribute AS a
    INNER JOIN pg_type ON pg_type.oid = a.atttypid
WHERE
    c.relname = 'platform_user'
    AND a.attrelid = c.oid
    AND a.attnum > 0;