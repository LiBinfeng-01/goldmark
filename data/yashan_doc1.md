# YashanDB 产品描述

### v23.4.1.

### 2025年05月


###### 深圳计算科学研究院 深圳崖山科技有限公司

## 产品简介

### 数据库管理系统发展简史

```
数据库管理系统（Database Management System，DBMS）是一种基础软件，用于对数据进行统一、可靠、高效地管理和组织，保证数据地安全性和完
整性，同时提供高效的数据查询能力。
数据库 （Database）是指按照数据结构来组织、存储和管理数据的集合。
Note :
数据库（Database）也经常被用来代指数据库管理系统（DBMS），仅在两者同时出现时作概念区分。在本手册中，可能存在代指现象。
初代数据库管理系统主要是层次型和网络型，在层次型数据库中数据按照树状结构组织（与文件系统类似），而在网络型数据库中数据被组织成多对多的
网络结构。
1970 年E.F.Codd首次提出关系模型，为数据库系统提供了一种统一的、简洁的数据模型。数据组织方式符合关系模型的数据库称为关系数据库，数据储
存在由行（数据组成的元组）和列（属性）构成的关系表中。与之对应的DBMS则称为关系数据库管理系统（ Relational Database Management
System ， RDBMS ），关系数据库管理系统仍为目前最主流的数据库管理系统。
```
### YashanDB 发展简史

```
崖山数据库系统YashanDB是深圳计算科学研究院自主设计研发的新型数据库管理系统，在经典数据库理论基础上，融入原创的有界计算、近似计算、并
行可扩展和跨模融合计算理论，可满足金融、政企、能源等关键行业对高性能、高并发及高安全性的要求。
YashanDB的发展分为以下阶段：
2013~2018 年 理论证明和奠基
首创提出大数据计算复杂性理论，奠定大数据查询复杂性基础。
原创理论体系：有界计算、增量计算、近似计算、并行计算、跨模融合、逻辑+AI等，奠定理论基础。
A类论文近百篇，荣获数据库顶会最佳论文大满贯，十年时间检验奖。
2019~2022 年 产品和工程落地
深圳计算科学研究院揭牌成立。
YashanDB完成一阶段原创理论、核心技术、自研系统，具备核心系统上线能力。
高端共享集群原型开发完成可行性验证。
经权威机构检测，内核代码自主率100%。
荣获数字中国“十大硬核科技”。
产品亮相CNCC，DeepTech专访报道。
2023 市场复制，关键行业渗透
YashanDB V22.2正式发布，功能、性能、稳定性等全面加强。
YashanDB V23.1发布，高端共享集群、分布式实时数仓、空间数据库三款产品发布。
YashanDB个人版全面开放下载。
亮相第二十五届中国高新技术成果交易会开幕式。
荣获国家工业信息安全发展中心“2023年数字转型自主创新解决方案”。
入选 2023 世界互联网大会领先科技奖收录成果集《科技之魅》。
金融、党政、能源等关键行业渗透。
2024 高端核心 1 ： 1 平替
YashanDB V23.3发布，专为核心场景打造 1 ： 1 平替方案。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
YashanDB企业版全面开放下载。
YashanDB共享集群成功投产，打造政务行业数据库自主创新样板点。
签约中国人民银行数字货币研究所，推动金融行业数字化高质量发展。
获商用密码、EAL4+、网专安全认证，数据库安全能力获权威认可。
荣获地理信息科技进步奖一等奖。
斩获广东省优秀信创产品与解决方案双料荣誉。
```
### YashanDB 部署形态

```
单机（主备）部署（简称：单机部署）
传统的关系型数据库理论与创新的底座引擎技术相结合，适用于集中式事务业务场景，支持主备形态。
共享集群部署
基于共享存储的多活集群，提供计算/存储扩展和金融级高可用能力，适用于高端核心交易场景。
分布式集群部署 （简称：分布式部署）
继承单机能力的原生分布式处理系统，适用于分布式分析业务场景，支持主备形态。
```
### YashanDB 核心特性

```
行式存储 / 列式存储
支持HEAP行存表、TAC列存表、LSC列存表。
支持LIST/RANGE/HASH/INTERVAL分区类型和多种组合二级分区能力。
支持向量化计算。
事务管理
支持完整事务ACID、细粒度锁管理、语句读写一致性，支持读已提交和串行化事务隔离级别、UNDO自管理、多版本并发控制。
高性能查询
提供基于代价和规则的优化器模型。
支持MPP分布式执行方式，具备数据排序、稀疏索引、预读与缓存、数据压缩等存储技术，加上SQL引擎使用了分区剪枝、并行查询、条件下推等特
性，带来高性能查询。
数据复制
支持同步和异步两种复制模式，提供最大保护、最大性能等多种复制策略。
备份恢复
提供数据物理备份和逻辑备份能力，支持全量和二级增量备份，并且支持基于时间点恢复（PITR）功能。
高可用
提供一主多备、级联备高可用能力，支持手动切换和自动选主。
数据库闪回
支持闪回查询、闪回修改以及回收站闪回能力。
通用 SQL 能力
遵循ANSI SQL标准，支持常用SQL语法，提供丰富函数库和数据类型，支持高性能的PLAN算子。
PL
支持存储过程、自定义函数、匿名块、高级包、JOB、触发器等能力。
聚合内存
支持聚合内存（Cohesive Memory）核心技术，用于集群数据库各实例之间协同数据页的读写访问以及各种非数据类的并发控制。
文件系统能力
可直接管理裸盘提供文件系统服务，在共享集群部署时为多节点集群提供并行文件读写能力。
空间数据管理
支持ST_GEOMETRY数据类型，用于存储和访问符合开放地理空间信息联盟（Open Geospatial Consortium，简称OGC）制定的SFA SQL标准的几
何对象。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
在线扩缩容
支持在线扩缩容，以满足客户业务发展需求。
```
### YashanDB 适用场景

#### 在线事务处理

```
YashanDB集中事务体系着眼于下一代基础设施软硬件的发展、以关键行业的核心应用场景作为牵引，提供高性能、高可靠的数据库底座，满足高并发在
线事务处理的业务需求。
极高性能要求的在线交易
此类场景对事务处理的时效性和准确性要求严苛，并要求对数据进行高可靠保护。YashanDB拥有精细化的事务管理能力，强大的存储底座结合高速的
SQL引擎，在保证数据强一致性的基础上实现卓越的事务处理性能。
7*24 小时不间断服务
YashanDB的高可用架构通过多副本日志同步机制可以将数据在多个数据中心之间进行实时同步并持久化，并通过Raft协议的自动选主实现主备在线自
动切换，用户无感知的情况下保证系统的稳定持续运行状态。
企业集中式管控
YashanDB的HTAP混合负载形态，基于"同一份数据，同一个引擎"，同时支持在线实时交易及实时分析两种场景，支持高并发在线吞吐的同时，提供
海量数据的实时在线分析能力。
高可用无缝切换场景
金融、电信、电力等行业对高可用要求较高，常采用Oracle RAC这种依赖共享存储的高可用架构，YashanDB共享集群部署具备同样的高可用架构能
力，当集群中出现单点故障时，会被另外一个节点接管，客户端透明切换，RPO为 0 ，RTO最快只需要十多秒，保障核心业务连续运行不中断。
```
#### 海量数据分析

```
YashanDB分布式分析体系基于有界计算理论的即时分析，专注解决大数据计算效率、数据生产即分析、海量数据存储成本及数据孤岛等数据库瓶颈难题
（3V：Volumn、Velocity、Variety），为用户提供灵活的传统数仓加速、一站式数据仓库解决方案。
海量稳态数据分析
YashanDB的LSC表为列存表，通过列存结构组织数据，支持冷热数据分离，高压缩比的对象存储，通过数据排序，稀疏索引，下推过滤等技术实现海
量数据的高性能查询，同时支持数据写入热数据区提升事务性能，支持热数据区与冷数据区的静默转换和融合查询，主打海量稳态数据的交互式分析
场景。
Ad-Hoc 交互式自助分析场景
对业务进行交互式探索分析，通过高性价比列存引擎、向量化执行引擎、高效的分布式算法等全自研技术达到秒级响应的查询分析体验。
实时个性化推荐场景
基于海量用户历史行为日志数据进行多维分析，支持行为日志的实时入库，并完成历史标签和实时标签的秒级计算，支持通过各类标签组合进行用户
圈选，提高营销成功率。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 产品规格

```
产品规格用于展示用户在使用本产品时，应遵循的一些限制要求，包括如下几部分：
物理规格
逻辑规格
数据类型规格
在本章节及所有产品文档中，在描述数量时：
K表示 1024
M表示1024*
G表示1024*1024*
在描述长度时：
KB表示 1024 字节
MB表示1024*1024字节
GB表示1024*1024*1024字节
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 物理规格

### 数据库

```
规格名称 规格类型 规格值
控制文件数量 最小值 2
控制文件数量 默认值 建库默认不指定时控制文件为 3 个
控制文件数量 最大值 8
数据块大小 最大值 32KB支持8KB^ 、16KB、32KB块大小
数据块大小 默认值 8KB
实例个数 最大值 64
```
### 表空间

```
规格名称 规格类型 规格值
表空间数量 最大值 2048
数据文件数量 单表空间内最大值 64
数据文件数量 最大值 16384
数据文件大小 最小值 1MB
数据文件大小 最大值 2TB
单个表空间Databucket数量 最大值 64
Databucket数量 最大值 4096
Databucket数量 默认值 256
```
### 日志文件

```
规格名称 规格类型 规格值
```
```
日志文件大小 最小值
```
```
6MB
最小值受DB_BLOCK_SIZE，MAX_SESSIONS和REDO_BUFFER_SIZE三个参数的影响，
参考公式：DB_BLOCK_SIZE * MAX_SESSIONS * 8 + REDO_BUFFER_SIZE / 2
日志文件大小 最大值 2047GB
日志Block大小 最小值 512Bytes
日志Block大小 最大值 32KB
日志Block大小 默认值 4KB
日志文件数量 单实例最小值 3
日志文件数量 单实例最大值 256
归档文件数量 最大值 1000000
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

### 主备高可用

```
规格名称 规格类型 规格值
```
```
物理备库 最大备库数量
```
```
单机部署： 32
共享集群部署： 32 （备集群）
分布式部署： 4 （备节点）
```
```
逻辑备库 最大备库数量
```
```
单机部署： 32
共享集群部署：无逻辑备集群
分布式部署：无逻辑备库
```
```
级联备 最大级联数
```
```
单机部署：不限制
共享集群部署：无级联备集群
分布式部署：无级联备
单个备份集文件数量 最大值 17408
同一基线的增量备份数量 最大值 1000
RTO 最大值^30 依赖于自动选主心跳和间隔时间配置秒^
```
### 共享集群部署

```
规格名称 规格类型 规格值
同一套集群的数据库服务器个数 推荐值 2 ~ 4
同一套集群中，单台服务器部署数据库实例的个数 最大值 1
系统盘大小 最小值 1G
LUN路径长度 最大值 31B
```
### 崖山文件系统

```
规格名称 规格类型 规格值
diskgroup数量 全局最大值 512
failuregroup数量 全局最大值 2048
failuregroup数量 单diskgroup内最大值 16
disk数量 全局最大值 65535
disk数量 单diskgroup内最大值 8192
disk数量 单值failuregroup内最大 1024
AU Size 默认值 1M
AU Size 最大值 32M支持^ 1M、4M、8M、16M、32M大小
```
```
连接数 默认值^1024 其中内部连接数为（^ 1 + （实例数 -1） * 2 ）
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
规格名称 规格类型 规格值
连接数 最大值^65535 其中内部连接数为（^ 1 + （实例数 -1） * 2 ）
单个diskgroup的文件个数 最大值 100 万
单个文件大小 最大值 YFS 文件的文件大小上限与 AU size、冗余度有关，YFS 至少支持创建最大 180T 的
```
```
单个disk大小 最大值 4G * Au size默认Au size时，单磁盘最大为^ 4PB
```
```
单个disk大小 最小值 32 * Au size默认Au size时，单磁盘最小为^ 32M
disk路径长度 最大值 31
目录或文件名长度 最大值 31
diskgroup/failuregroup/disk名称长
度 最大值^31
目录和文件绝对路径长度 最大值 255
```
### 分布式部署

```
规格名称 规格类型 规格值
分布式部署 最大部署规模 5MN 8CN 32*5DN
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 逻辑规格

### 对象

```
规格名称 规格类型 规格值
用户数量 最大值 10240 （包括数据库内置用户）
表数量 最大值 不限制
对象名称长度 最大值 64Bytes
角色数量 最大值 9640 （不包括数据库内置角色）
私有临时表数量 单会话内最大值 64
密码长度 最大值 127
全库闪回还原点数量 最大值 8192
```
### 表

```
规格名称 规格类型 规格值
列数 最大值 4096
记录数 最大值 不限制
行长度 最大值 行存：列存：64512Bytes32000KB^
列长度 最大值 8000Bytes
LOB列长度 最大值 不限制
```
### 索引

```
规格名称 规格类型 规格值
索引列数 最大值 32
索引键值长度（含内部格式） 最大值 6000Bytes
单个表的索引数量 最大值 255
索引层数 最大值 24
```
### 访问约束（ AC ）

```
规格名称 规格类型 规格值
列数 最大值 31
列长度 最大值 8000Bytes
单个表的访问约束数量 最大值 255
```
### undo


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
规格名称 规格类型 规格值
undo段数量 最大值 1024
undo段数量 最小值 8
```
### 事务

```
规格名称 规格类型 规格值
事务数量 最大值 不限制
自治事务嵌套层数 最大值 不限制
```
### 序列

```
规格名称 规格类型 规格值
序列值 最大值 1e28 - 1
序列值 最小值 -1e27 + 1
序列步长 最大值 步长绝对值需要小于序列最大值与最小值的差
序列缓存个数 最大值 循环序列的缓存值个数不能超过一个循环的值的个数， 非循环序列无限制
序列缓存个数 最小值 2
```
### 数据分区

```
规格名称 规格类型 规格值
分区数 最大值 分区表1M - 1/索引可以创建的最大分区数量^
```
```
分区列数 最大值^16 间隔分区（^ interval）仅支持单列分区
```
### 语法

```
规格名称 规格类型 规格值
单个SQL语句长度 最大值 2MB - 1（包含空格和特殊字符）
yasql单行字符数 最大值 65534
SQL语句中常量字符串输入长度 最大值 16000
字符串类型输出长度 最大值 65534
SQL语句中参与Join的表数量 最大值 128
存储过程行数 最大值 64K
存储过程参数个数 最大值 4095
自定义函数参数个数 最大值 4095
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
规格名称 规格类型 规格值
自定义高级包元素个数 最大值 1024
窗口函数的分区数量 最大值 1000
窗口函数的ORDER BY列数量 最大值 1000
窗口函数数量 最大值 128
最大投影列长度（包含列存拆分出来的投
影列） 最大值^4096
子查询嵌套数量 最大值 128
并行度数量 最大值 255
物化区单行规格 最大值 63KB
子查询或SELECT子句数量 最大值 170
物化区排序键长度 最大值 20K
```
```
单个SQL语句执行所需stage个数 最大值
```
```
128
单个SQL语句执行最大stage个数受到MAX_PARALLEL_WORKERS配置参数的影响
参考公式：MIN(MAX_PARALLEL_WORKERS, 128)
```
```
聚合函数聚合物化区 最大值
```
```
32KB
多个聚合函数如果返回的结果会如果是变成数据，则会缓存在聚合物化区，该物化区的规
格限制为32K
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 数据类型规格

```
数据类型 类型长度 取值范围 说明
TINYINT 1Bytes -128 ~ 127 小整数值，例如大位数。 boolean，在括号中规定最
SMALLINT 2Bytes -2 (-32,768) ~ 2 - 1 (32,767) /
INTEGER
INT
PLS_INTEGER
```
```
4Bytes -2(2,147,483,647) (-2,147,483,648) ~ 2 - 1
```
```
大整数值。
INT/PLS_INTEGER为INTEGER的别名，
行为完全一致。
```
```
BIGINT 8Bytes
```
-
(-9,223,372,036,854,775,808) ~
2 -
(9,223,372,036,854,775,807)

```
极大整数值。
```
```
FLOAT
BINARY_FLOAT
REAL
```
```
4Bytes
```
```
32 位单精度浮点数
负数：-3.402823E38 ~
-1.401298E-
正数：1.401298E-45 ~
3.402823E38， 0
```
```
表示单精度浮点数，在括号中规定最大位
数，在d参数中规定小数点右侧的最大位
数。
BINARY_FLOAT/REAL为FLOAT的别名，
行为完全一致。
```
```
DOUBLE
BINARY_DOUBLE 8Bytes
```
```
64 位双精度浮点数
负数：
-1.797693134862315807E308 ~
-4.94065645841247E-
正数：4.94065645841247E-
~
1.797693134862315807E308， 0
```
```
表示双精度浮点数。
BINARY_DOUBLE为DOUBLE的别名，行
为完全同DOUBLE。
```
```
NUMBER
DECIMAL
NUMERIC
```
```
20Bytes
```
```
NUMBER(p,s)
p：1 ~ 38
s：-84 ~ 127
```
```
用于精度要求非常高的计算，以二进制存
储，因此实际存储长度、与表达的数据范
围需要换算。
DECIMAL/NUMERIC为NUMBER的别
名，行为完全同NUMBER。
```
```
CHAR(SIZE[CHAR|BYTE])
CHARACTER(SIZE[CHAR|BYTE])
```
```
存储：
1 ~ 8000Bytes
运算：
1 ~ 65534Bytes
```
```
无
```
```
CHARACTER为CHAR的别名，行为完全
同CHAR。
列存不支持CHAR(SIZE[CHAR])的定义。
```
```
NCHAR(SIZE)
```
```
存储：
1 ~ 8000Bytes
运算：
1 ~ 65534Bytes
```
```
无 列存无此类型
```
```
VARCHAR(SIZE[CHAR|BYTE])
CHARACTER
VARYING(SIZE[CHAR|BYTE])
VARCHAR2(SIZE[CHAR|BYTE])
```
```
1 ~ 65534Bytes 无
```
```
CHARACTER VARYING/VARCHAR2为
VARCHAR的别名，行为完全同
VARCHAR。
NVARCHAR(SIZE) 1 ~ 65534Bytes 无 列存无此类型
BLOB 1 ~4G*DB_BLOCK_SIZE 无 无
```
```
CLOB 1 ~4G*DB_BLOCK_SIZE 无 无
```
```
NCLOB 1 ~4G*DB_BLOCK_SIZE 无 列存无此类型
```
```
15 15
31 31
```
```
63
63
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
数据类型 类型长度 取值范围 说明
XMLTYPE 1 ~4G*DB_BLOCK_SIZE 无 列存无此类型
```
```
DATE 8Bytes 1-1-1 00:00:00 ~ 9999-12-3123:59:59 表示YYYY-MM-DD [HH24:MI:SS]
```
```
TIME 8Bytes -838:59:59.999999 ~838:59:59.999999 时分秒 微秒
```
```
TIMESTAMP 8Bytes 1-1-1 00:00:00.000000 ~ 9999-12-31 23:59:59.
```
```
YYYY-MM-DD HH24:MI:SS.FF
TIMESTAMP默认输入微秒精度为 9 位，因
此输入支持最大值为9999-12-
23:59:59.999999999，但输出时微秒只显
示 6 位（其余四舍五入），所以最大值为
9999-12-31 23:59:59.999999，溢出则报
错。
INTERVAL YEAR TO MONTH 4Bytes -178000000-00 ~ 178000000-00 表示时间间隔，年月。
INTERVAL DAY TO SECOND 8Bytes -100000000 00:00:00.000000 ~100000000 00:00:00.000000 表示时间间隔，天时分秒。
```
```
BOOLEAN
```
```
行存：
1Byte
列存：
1Bit
```
```
支持的输入：
非零整数（同BIGINT规格）；
字符串'true'/'false'、'
t'/'f'、'on'/'off'、'yes'/'no'或'0'/'1'，
true/false
```
```
无
```
```
BIT 1 ~ 8Bytes 同BIGINT规格 列存无此类型
```
```
RAW(SIZE)
```
```
存储：
1 ~ 8000Bytes
运算：
1 ~ 65534Bytes
```
```
无 无
```
```
JSON 1 ~ 32MBytes 无 可以解析为32M，JSONJSON对象数据的长度为对象的字符串长度为1-32M1-
```
```
ROWID 16Bytes
```
```
ROWID类型格式为：
dataoid:spaceid:fileid:blockid:dir
dataoid取值范围：0 ~ 2 -
（18,446,744,073,709,551,615）
spaceid取值范围：0 ~ 2 -
（2,047）
fileid取值范围：0 ~ 2-1（ 63 ）
blockid取值范围：0 ~ 2 -
（67,108,863）
dir取值范围：0 ~ 2 -1（4,095）
```
```
列存无此类型
```
```
UROWID 1 ~ 8000Bytes4000Bytes ，默认 无 无
```
```
BOX2D 32Bytes 无 与用于表示ST_GEOMETRST_GEOMETRY相关的空间数据类型，Y的二维边界框。
```
```
ST_GEOMETRY 4GB 无
```
```
与GIS相关的空间数据类型，具体包含
POINT、LINESTRING、POLYGON、
MULTIPOINT、MULTILINESTRING、
MULTIPOLYGON等数据类型。
```
```
64
11
6
26
12
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 兼容性说明

```
本文为您介绍YashanDB与Oracle、MySQL的兼容性，包括与Oracle兼容性说明和与MySQL兼容性说明两个章节。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 与 Oracle 兼容性说明

```
YashanDB（yashan模式）在SQL语法、表达式运算、FILTER CONDITION、数据类型、内置函数、系统视图和PL等基本功能上均与Oracle数据库兼
容，数据库管理和开发人员不需要花费大量的时间去学习新知识，在已交付特性上直接查阅Oracle相关文档，也可流畅地操作使用YashanDB，实现从
Oracle数据库到YashanDB的平滑迁移。
在其他某些功能上YashanDB会与Oracle数据库的表现有所差异或暂时没有进行兼容，可能原因如下：
YashanDB与Oracle数据库的底层架构、产品形态等并不相同。
YashanDB摒弃了一些旧的不符合当前主流业务框架的方法，并增加了自己的特性。
本文将从如下方面具体说明，在单机部署和行式存储模式下，YashanDB对Oracle数据库的兼容情况：
SQL语法
表达式运算
FILTER CONDITION
数据类型
内置函数
PL
系统视图
字符集
SQL引擎
数据库安全
工具兼容
其他兼容
```
### SQL 语法

```
YashanDB支持Oracle数据库中主流的SQL语法，其他少数因功能性缺失导致的不兼容将报语法不支持错误，此时应联系我们的技术支持提供变通方案，
SQL语法详细说明请查阅SQL语句。
（ 1 ） DML 类
SELECT
支持大部分查询功能，包括单、多表查询，子查询，内连接，半连接，外连接，分组及聚合，层次查询等
支持UNION、UNION ALL、INTERSECT、MINUS等集合操作
支持如下方式查看执行计划：
EXPLAIN
AUTOTRACE
支持随机抽样查询能力
CTE支持递归功能
INSERT
支持单行、多行插入，同时支持指定分区插入
支持INSERT INTO SELECT语句
支持INSERT ALL语句
支持单表和多表插入
UPDATE
支持单列和多列的更新
支持使用子查询
DELETE
支持单表和多表的删除
支持使用子查询
（ 2 ） DDL 类
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
YashanDB兼容Oracle数据库大部分的对象及对象管理操作，包括：
HEAP表：
CREATE TABLE/CREATE TABLE AS
ALTER TABLE
DROP TABLE
TRUNCATE TABLE
临时表
GLOBAL TEMPORARY TABLE
PRIVATE TEMPORARY TABLE
分区表
支持RANGE/INTERVAL/LIST/HASH类型分区
支持ADD|DROP|TRUNCATE|SPLIT|MERGE PARTITION
支持分区行迁移
支持二级分区
外部表
支持创建目录，CREATE DIRECTORY
支持创建和删除外部表
约束
包括in_line约束/out_of_line约束
支持UNIQUE、PRIMARY KEY、FOREIGN KEY、CHECK、(NOT) NULL类型约束
视图
CREATE VIEW/CREATE OR REPLACE FORCE VIEW
DROP VIEW
视图支持SELECT/INSERT/UPDATE/DELETE
物化视图
CREATE MATERIALIZED VIEW
ALTER MATERIALIZED VIEW
DROP MATERIALIZED VIEW
支持本地物化视图，不支持远程物化视图
BTree索引
包括全局索引和本地（LOCAL）索引
包括唯一（UNIQUE）索引和非唯一索引
包含反向键索引和函数索引
支持对索引REBUILD|UNUSABLE|COALESCE|PARALLEL|RENAME
支持在相同列上创建多个索引
同义词
包括私有同义词和公共（PUBLIC）同义词
序列
包括升序序列和降序序列，可指定CYCLE|NOCYCLE、CACHE|NOCACHE、ORDER|NOORDER
序列支持NEXTVAL和CURRVAL
DBLINK
支持Oracle到YashanDB、YashanDB到Oracle和YashanDB到YashanDB的远程连接
支持PUBLIC和PRIVATE模式的LINK
支持表和视图的连接
支持在YashanDB创建远端对象的同义词、在YashanDB上调用远端存储过程和查看LOB数据
（ 3 ） HINT
YashanDB支持使用HINT并采用了Oracle的HINT语法，可实现join方式、join order、table scan、index scan等的指定能力。关于HINT的详细说明，请查
阅hint。
```
### 表达式运算

```
YashanDB包含了主流的计算框架实现对表达式的运算，此外，Oracle本身由于没有布尔类型（只在其PL中支持）需要使用其他数据类型替代运算，
YashanDB则实现了直接的布尔型表达式运算，详情如下表所示：
表达式运算类型 YashanDB Oracle 数据库
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
表达式运算类型 YashanDB Oracle 数据库
二元运算加法 支持 支持
二元运算减法 支持 支持
二元运算乘法 支持 支持
二元运算除法 支持 支持
二元运算取余 支持 支持
一元运算取反 支持 支持
位运算与 支持 支持
位运算或 支持 支持
位运算异或 支持 支持
Boolean运算 支持 不支持
```
### FILTER CONDITION

```
YashanDB的FILTER CONDITON类型完全兼容Oracle数据库的FILTER CONDITON类型，详情如下表所示：
FILTER CONDITON YashanDB Oracle 数据库
ALL 支持 支持
ANY 支持 支持
AND 支持 支持
EQUAL 支持 支持
NOT EQUAL 支持 支持
EXISTS 支持 支持
NOT EXISTS 支持 支持
GREAT EQUAL 支持 支持
GREATE 支持 支持
IN 支持 支持
NOT IN 支持 支持
IS NULL 支持 支持
IS NOT NULL 支持 支持
LESS 支持 支持
LESS EQUAL 支持 支持
LIKE 支持 支持
NOT LIKE 支持 支持
REG LIKE 支持 支持
NOT REG LIKE 支持 支持
OR 支持 支持
SOME 支持 支持
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

### 数据类型

```
YashanDB目前包含了 27 种数据类型，与Oracle数据库对比情况如下表所示：
数据类型 YashanDB Oracle 数据库
BOOLEAN 支持 不支持
TINYINT 支持 不支持
SMALLINT 支持 支持
INTEGER 支持 支持
BIGINT 支持 不支持
FLOAT 支持 支持
DOUBLE 支持 支持
NUMBER 支持 支持
DATE 支持 支持
TIMESTAMP 支持 支持
TIMESTAMP WITH LOCAL TIME ZONE 支持 支持
TIMESTAMP WITH TIME ZONE 支持 支持
TIME 支持 不支持
INTERVAL YEAR TO MONTH 支持 支持
INTERVAL DAY TO SECOND 支持 支持
CHAR 支持 支持
VARCHAR 支持 支持
NCHAR 支持 支持
NVARCHAR 支持 支持
ST_GEOMETRY 支持 支持
RAW 支持 支持
CLOB 支持 支持
NCLOB 支持 支持
BLOB 支持 支持
BIT 支持 不支持
ROWID 支持 支持
UROWID 支持 支持
CURSOR 支持 支持
JSON 支持 支持
XMLTYPE 支持 支持
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
Note ：
YashanDB的大对象（CLOB/BLOB）数据类型有较强的存储能力，但大对象存取性能不推荐在复杂场景下使用。
YashanDB的XMLTYPE数据类型底层以CLOB进行存储，与Oracle存储方式不同，在函数调用时可能存在差异。
```
### 内置函数

```
YashanDB目前实现了超过 120 个内置函数，具体清单及说明请查阅内置函数。
与Oracle数据库的内置函数对比情况如下表所示：
函数
类型 支持函数列表 YashanDB
```
```
Oracle
数据库
数学
运算
函数
```
```
ABS、ACOS、ASIN、ATAN、ATAN2、COS、COT、 CEIL、DIV、EXP、FLOOR、LOG、LN、MOD、
PI、POW、POWER、ROUND、SIGN、SIN、SINH、STDDEV、STDDEV_POP、STDDEV_SAMP、
SQRT、TAN、TANH、TRUNC、VARIANCE、VAR_POP、VAR_SAMP、......
```
```
支持 支持
```
```
字符
处理
函数
```
```
ASCII、CHR、CONCAT、INITCAP、INSTR、INSTRB、LEFT、LENGTH、LENGTH2、LOWER、LPAD、
LTRIM、NLSSORT、POSITION、REPLACE、RIGHT、RPAD、RTRIM、SPLIT、STRPOS、SUBSTR、
SUBSTRB、SOUNDEX、TRIM、TRANSLATE ... USING、UNISTR、UPPER、empty_LOB、......
```
```
支持 支持
```
```
正则
匹配
函数
```
```
REGEXP_LIKE、REGEXP_COUNT、REGEXP_INSTR、REGEXP_REPLACE、
REGEXP_SUBSTR、...... 支持 支持
```
```
转换
函数
```
```
BIN_TO_NUM、CAST、NUMTODSINTERVAL、NUMTOYMINTERVAL、ROWIDTOCHAR、
TRANSLATE、TO_CHAR、TO_DATE、TO_DSINTERVAL、TO_NUMBER、TO_TIMESTAMP、
TO_YMINTERVAL、......
```
```
支持 支持
```
```
集合
处理
函数
```
```
COALESCE、DECODE、GREATEST、LEAST、NVL、NVL2、...... 支持 支持
```
```
聚集
函数
```
```
AVG、COUNT、GROUP_CONCAT、LISTAGG、MAX、MIN、SUM、MEDIAN、
PERCENTILE_CONT、...... 支持 支持
窗口
函数
```
```
AVG、COUNT、FIRST、FIRST_VALUE、LAST、LAST_VALUE、LEAD、MAX、SUM、RANK、
ROW_NUMBER、...... 支持 支持
系统
函数 SCN_TO_TIMESTAMP、TIMESTAMP_TO_SCN、USERENV、...... 支持 支持
时间
处理
函数
```
```
ADD_MONTHS、CURRENT_TIMESTAMP、EXTRACT、LAST_DAY、NOW、NEXT_DAY、
OVERLAPS、...... 支持 支持
条件
处理
函数
```
```
IF、IFNULL、ISNULL、NULLIF 支持 不支持
```
```
JSON
处理
函数
```
```
JSON、JSON_ARRAY_GET、JSON_ARRAY_LENGTH、JSON_EXISTS、JSON_FORMAT、
JSON_PARSE、JSON_QUERY、JSON_SERIALIZE、JSON_VALUE、...... 支持 支持
随机
函数 RANDOM 支持 不支持
其他
函数 HEXTORAW、SQLCODE、SQLERRM、SYS_CONNECT_BY_PATH、...... 支持 支持
```
### PL

```
YashanDB兼容了Oracle数据库大部分的PL功能，包括：
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
数据类型
流程控制
静态SQL
动态SQL
异常处理
系统定义异常
用户自定义异常
游标
BULK COLLECT
存储过程
过程体加密wrap功能
触发器
支持行级、语句级触发器。
目前仅支持在表上创建触发器，不支持在视图上创建触发器
用户自定义高级包
支持子过程
支持全局变量和TYPE定义
PROCEDURE和FUNCTION支持重载
用户自定义函数
SQL语言的UDF
外置JAVA语言UDF
外置C语言的UDF
用户自定义数据类型
OBJECT TYPE
VARRAY TYPE
TABLE TYPE
支持EXECUTE、UNDER ON对象级的权限控制
JOB
```
### 内置高级包

```
YashanDB兼容了Oracle数据库部分的内置高级包，详情如下表所示：
高级包名称
DBMS_AUDIT_MGMT
DBMS_APPLICATION_INFO
DBMS_CRYPTO
DBMS_DESCRIBE
DBMS_HM
DBMS_IJOB
DBMS_JOB
DBMS_LOB
DBMS_LOCK
DBMS_METADATA
DBMS_MVIEW
DBMS_OUTPUT
DBMS_RANDOM
DBMS_RESOURCE_MANAGER
DBMS_ROWID
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
高级包名称
DBMS_SCHEDULER
DBMS_SESSION
DBMS_SQL
DBMS_STANDARD
DBMS_STATS
DBMS_UTILITY
DBMS_XA
OWA_UTIL
UTL_ENCODE
UTL_FILE
UTL_I18N
UTL_RAW
```
### 系统视图

```
YashanDB兼容Oracle数据库的DBA视图清单如下：（ALL/USER视图兼容清单与DBA视图对应相同）
视图名称
AUDITABLE_SYSTEM_ACTIONS
AUDIT_UNIFIED_ENABLED_POLICIES
AUDIT_UNIFIED_POLICIES
COL
DICT_COLUMNS
DICTIONARY
DICT
DBA_ALL_TABLES
DBA_ARGUMENTS
DBA_AUDIT_MGMT_CLEANUP_JOBS
DBA_AUDIT_MGMT_LAST_ARCH_TS
DBA_COLL_TYPES
DBA_COL_COMMENTS
DBA_CONSTRAINTS
DBA_CONS_COLUMNS
DBA_DATA_FILES
DBA_DB_LINKS
DBA_DEPENDENCIES
DBA_EXTENTS
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
DBA_EXTERNAL_TABLES
DBA_FREE_SPACE
DBA_HISTOGRAMS
DBA_INDEXES
DBA_IND_COLUMNS
DBA_IND_EXPRESSIONS
DBA_IND_PARTITIONS
DBA_IND_STATISTICS
DBA_IND_SUBPARTITIONS
DBA_JOBS
DBA_LOBS
DBA_LOB_PARTITIONS
DBA_LOB_SUBPARTITIONS
DBA_LOG_GROUPS
DBA_MVIEWS
DBA_NESTED_TABLES
DBA_OBJECTS
DBA_OUTLINES
DBA_OUTLINE_HINTS
DBA_PART_COL_STATISTICS
DBA_PART_HISTOGRAMS
DBA_PART_INDEXES
DBA_PART_KEY_COLUMNS
DBA_PART_TABLES
DBA_PROCEDURES
DBA_PROFILES
DBA_RECYCLEBIN
DBA_ROLES
DBA_ROLE_PRIVS
DBA_SCHEDULER_JOBS
DBA_SEGMENTS
DBA_SEQUENCES
DBA_SOURCE
DBA_SUBPARTITION_TEMPLATES
DBA_SUBPART_KEY_COLUMNS
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
DBA_SYNONYMS
DBA_SYS_PRIVS
DBA_TABLES
DBA_TABLESPACES
DBA_TAB_COLS
DBA_TAB_COLUMNS
DBA_TAB_COL_STATISTICS
DBA_TAB_COMMENTS
DBA_TAB_HISTOGRAMS
DBA_TAB_MODIFICATIONS
DBA_TAB_PARTITIONS
DBA_TAB_PRIVS
DBA_TAB_STATISTICS
DBA_TAB_STAT_PREFS
DBA_TAB_SUBPARTITIONS
DBA_TEMP_FILES
DBA_TRIGGERS
DBA_TRIGGER_COLS
DBA_TRIGGER_ORDERING
DBA_TYPES
DBA_TYPE_ATTRS
DBA_TYPE_METHODS
DBA_USERS
DBA_VIEWS
ROLE_SYS_PRIVS
ROLE_TAB_PRIVS
UNIFIED_AUDIT_TRAIL
```
```
YashanDB兼容Oracle数据库的动态视图清单如下：
视图名称
V$2PC_PENDING
V$ARCHIVE_DEST
V$ARCHIVE_DEST_STATUS
V$ARCHIVE_GAP
V$ARCHIVED_LOG
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
V$BUFFER_POOL
V$BUFFER_POOL_STATISTICS
V$CONTROLFILE
V$DATABASE
V$DATAFILE
V$DIAG_INCIDENT
V$DIAG_PROBLEM
V$EVENT_HISTOGRAM
V$EVENT_NAME
V$FIXED_TABLE
V$FIXED_VIEW_DEFINITION
V$HM_CHECK
V$HM_CHECK_PARAM
V$HM_FINDING
V$HM_RUN
V$INSTANCE
V$LOCK
V$LOCKED_OBJECT
V$LOGFILE
V$MYSTAT
V$OPEN_CURSOR
V$OSSTAT
V$PARAMETER
V$PQ_TQSTAT
V$PROCESS
V$PX_SESSION
V$RECOVERY_PROGRESS
V$RESERVED_WORDS
V$ROLLBACK
V$SEGMENT_STATISTICS
V$SEGSTAT
V$SESS_TIME_MODEL
V$SESSION
V$SESSION_EVENT
V$SESSION_WAIT
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
V$SESSTAT
V$SGA
V$SGASTAT
V$SQL
V$SQL_BIND_CAPTURE
V$SQL_PLAN
V$SQL_PLAN_STATISTICS
V$SQLAREA
V$SQLSTATS
V$SQLTEXT
V$STATNAME
V$SYSSTAT
V$SYSTEM_EVENT
V$SYSTEM_PARAMETER
V$SYSTEM_WAIT_CLASS
V$TABLESPACE
V$TEMP_EXTENT_POOL
V$TEMPSEG_USAGE
V$TRANSACTION
V$UNDOSTAT
V$VERSION
```
### 字符集

```
YashanDB支持如下字符集：
ASCII
GBK
UTF8
ISO88591
GB18030
同时，YashanDB支持如下字符集排序方式：
ASCII_GENERAL_CS
ASCII_GENERAL_CI
GBK_GENERAL_CS
GBK_GENERAL_CI
UTF8_GENERAL_CS
UTF8_GENERAL_CI
UTF8_PINYIN_CS
UTF8_PINYIN_CI
ISO88591_GENERAL_CS
ISO88591_GENERAL_CI
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
GB18030_GENERAL_CS
GB18030_GENERAL_CI
GB18030_PINYIN_CS
GB18030_PINYIN_CI
```
### SQL 引擎

```
YashanDB的SQL引擎兼容了Oracle数据库大部分的特性，包括：
支持查询改写
支持预编译语句
支持基于成本的优化器
支持执行计划生成与展示（EXPLAIN）
支持执行计划缓存
支持执行计划快速参数化
支持执行计划绑定
支持Optimizer Hint
支持OutLine
支持like和reglike的模糊匹配能力
```
### 数据库安全

```
权限管理
兼容Oracle数据库的系统级权限
兼容常见的对象权限，支持table、view对象级权限管理及使用
支持授权与移除权限
with admin option
with grant option
支持创建用户自定义角色
支持DBA和PUBLIC系统预定义角色及其权限认证
提供视图查询权限相关内容
身份鉴别
兼容Oracle数据库的密码策略
profile支持用户密码管理
审计
兼容Oracle数据库的统一审计
```
### 工具兼容

```
imp
支持FULL、FROM/TOUSER和TABLES模式的维度导入
支持指定元数据和数据导入
支持覆盖已存在的表，即truncate模式
exp
支持FULL、OWNER和TABLES模式的维度导出
支持指定元数据和数据导出
运维工具
支持 10053 跟踪事件
```
### 其他兼容

```
统计信息
YashanDB的统计信息体系兼容了Oracle数据库大部分的特性，包括：
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
支持收集表（分区表）、索引、列等维度统计信息
支持动态采样功能
支持锁定统计信息
支持收集列直方图统计信息
支持通过JOB配置收集统计信息任务
支持使用高级包DBMS_STATS导入导出统计信息
支持基础统计信息的实时收集
表空间
支持ONLINE/OFFLINE
支持RENAME
可靠性
支持全库闪回、闪回DML操作、闪回查询功能
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 与 MySQL 兼容性说明

```
YashanD提供了MySQL数据库的兼容性功能，但由于不同数据库的底层架构、产品形态等存在差异，导致适配MySQL数据库的某些特性与适配Oracle数
据库的某些特性不能互相兼容。YashanDB通过控制开关的方式，很好地解决了这个问题，从而可以让用户从不同数据库向YashanDB迁移时，无需进行
大量地SQL校对和改写工作。
YashanDB通过语法模式选择决定兼容MySQL适配度，用户在安装数据库时可按需选择：
mysql模式：仅适用于单机部署，在该模式下用户输入的SQL语句将优先按照MySQL的语法体系进行解析，适用于需要密切适配MySQL数据库的使用
场景。
yashan模式：在该模式下使用YashanDB的语法体系。若安装时指定为yashan模式（省略不指定时，默认为yashan模式），安装后无法切换为mysql
模式，此时只能使用扩展支持的部分MySQL特有语法，适用于只需偶尔适配MySQL数据库的使用场景，详细介绍请查阅yashan模式兼容部分MySQL
语法。
```
### mysql 模式

```
在mysql模式下，用户输入的SQL语句将优先按照MySQL的语法体系进行解析，还支持MySQL专有语句，例如show语句、use语句等。
```
#### 控制开关

```
在YashanDB安装过程中，通过yasboot package se gen命令的mode参数可以指定语法模式为mysql模式。
以mysql模式安装后，在数据库OPEN阶段新建的所有会话默认为mysql模式，在NOMOUNT或MOUNT阶段创建的会话则仍为yashan模式，可通过以下
方式查看当前会话的语法模式。
```
```
在mysql模式下，部分数据库管理操作（例如主备切换）无法正常执行，需要先将当前会话切换至yashan模式（不影响其他会话）再执行语句/命令，且
换方式如下：
```
#### 概念介绍

```
名词 yashan 模式（同 Oracle ） mysql 模式
数据库
（Database）
```
```
由表空间、数据文件等组成的完整实例
一套YashanDB环境就是 1 个Database 数据库对象（表、视图等）的逻辑容器
模式
（Schema）
```
```
与用户一一对应，是用户的数据库对象（表、视图等）
的逻辑容器
创建用户时，自动创建 1 个同名schema
```
```
数据库的别名
```
```
用户（User）
```
```
登录账户 + 同名模式拥有者（天然具备该模式下对象的
全部权限）
可被授予其他模式对象的权限
```
```
仅为登录账户
可被授予任意数据库的权限
```
```
角色（Role） 权限的集合
```
```
无
可以使用YashanDB内置的角色（使用时基于yashan语法解析相
关SQL语句）
```
```
-- 查询COMPAT_VECTOR参数值
SHOW PARAMETER COMPAT_VECTOR
name ---------------------------------------------------------------- ------------------------------------------------------------value
----
COMPAT_VECTOR mysql
```
```
-- 切换到yashan模式，再执行相关运维操作
ALTER SESSION SET COMPAT_VECTOR = yashan;
-- 运维操作完成后，再切换回mysql模式继续进行业务相关操作
ALTER SESSION SET COMPAT_VECTOR = mysql;
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

#### 具体兼容项

##### 通信协议兼容

```
以mysql模式安装YashanDB数据库管理系统，会默认开启单独的监听端口，处理通过MySQL协议向YashanDB服务端发起的连接、执行等协议命令。
```
##### 语法兼容

```
在YashanDB的语法体系中，SELECT语句必须包含FROM子句，而MySQL的语法允许SELECT语句不带FROM子句。切换成mysql模式后，执行以下语
句，将返回当前会话的登录用户名，而不是报错：
```
```
当会话的语法模式切换为MySQL后，语法解析、语句执行将按照MySQL 5.7的风格。与YashanDB可能存在以下差异：
词法
例如mysql模式下，反引号（`）表示对象名，单引号（'）表示字符串，双引号（"）默认表示字符串，但在sql_mode中包含ANSI_QUOTES时，表示
对象名。而yashan模式下，双引号表示对象，单引号表示字符串。
对象的概念
例CREATE DATABASE语句，在mysql模式下表示创建一个Schema。
数据类型
例如Bool类型，在mysql模式下表示Tinyint(1)。
字面量的数据类型
字符类型的字面量，在mysql模式下，其数据类型为VARCHAR（与MySQL兼容），在Yashan模式下，类型为CHAR（与Oracle兼容）。
数据定义语言
例如CREATE TABLE语句，在mysql模式下会兼容Engine、Character set等选项，不再支持pctfree等选项。
数据操作语言
例如mysql模式下支持SELECT语句不带FROM子句。
权限
例如执行SELECT FOR UPDATE语句，在mysql模式下需要用户具有表的读权限，以及插入、删除、更新三者之中任意一种权限。而在yashan模式
下，SELECT语句需要读（READ）权限，SELECT FOR UPDATE语句需要单独的SELECT权限。
字符序
mysql模式下，字符类型的比较、排序规则受字符序影响（例如，数据库字符集选择为UTF8MB4时，默认的字符序为UTF8MB$_GENERAL_CI，字符
类型的排序和比较不区分大小写、忽略末尾空格），而在Yashan模式下，则按照二进制的方式比较和排序（区分大小写、末尾空格参与比较）。
Caution :
YashanDB 23.4暂未覆盖MySQL的所有语法及行为，对于暂未覆盖的那部分语句，在mysql模式下执行时仍会按照YashanDB语法进行解析和执
行。
对于MySQL本身不支持的YashanDB功能，这类语句在mysql模式下并不会产生歧义，因此在mysql模式下仍可以按照YashanDB语法进行解析和执
行，例如查询语句的集合操作（MINUS、INTERSECTS）、层次查询（CONNECT BY）。
```
##### 保留字兼容

```
YashanDB兼容了MySQL数据库大部分保留字，但目前个别保留字仍存在一定区别，包括：
CONNECT
EXCEPT
IF
```
```
SELECT USER();
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
INTERSECT
MINUS
OF
PUBLIC
ROW
ROWS
START
SYSDATE
```
#### 规格差异

```
YashanDB数据库在mysql模式下，绝大多数规格与yashan模式保持一致，以下仅列出与yashan模式下存在差异的规格。
```
##### 对象规格

```
规格名称 规格类型 最大值
user数量 最大值 10240 - 内置user数量 - database/schema数量
database\schema数量 最大值 10240 - 内置user数量 - 普通user数量
user名称长度 最大值 60Bytes
```
##### 数据类型规格

```
详情请查阅数据类型（mysql模式）。
```
#### 功能约束

```
使用mysql模式时，存在如下约束：
约束项 约束行为
部署形态 仅单机部署可选部署为mysql模式。
表类型 mysql模式下仅支持创建行存表。
用户登录 * * 如需使用如需使用yashanyasql以模式下创建的用户登录mysql模式下创建的用户登录YashanDBYashanDB（mysql（模式），则要求客户端支持mysql模式）需要使用双引号将用户名包围。SHA256插件。^
用户删除和
修改 不允许在mysql模式下删除和修改yashan模式下创建的用户。
```
```
对象名称
```
```
lower_case_table_names参数仅支持设置为 1 。
schema、表、视图、表别名的名称按照小写展示，列名按照大写展示。
对象名称匹配按照大小写不敏感的方式匹配。
Binary/Blob
类型文本协
议
```
```
当客户端与服务端字符集不一致时，会按服务端字符集转码，可能与客户端预期不一致。
```
```
字符集 仅支持实例级字符集设置，schema/table/column级的字符集均只为语法兼容。
排序集 仅支持实例级排序集设置，schema/table/column级的排序集均只为语法兼容。
全局变量 全局变量查询和设置仅为语法兼容，除字符集、自动提交外，大部分实际不会生效。
```
```
SQL_MODE
```
```
除ANSI_QUOTES、NO_BACKSLASH_ESCAPES、PIPES_AS_CONCAT、REAL_AS_FLOAT、
PAD_CHAR_TO_FULL_LENGTH外，其它SQL_MODE无论是否设置，均不会对数据库的行为产生影响，实际执行效果与yashan
模式一致。
可执行注释 可执行注释在mysql模式下会被视为注释，对语句执行不产生影响。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

### yashan 模式兼容 MySQL 语法

```
yashan模式是YashanDB的默认语法模式，若安装时指定为yashan模式（省略不指定时，默认为yashan模式），安装后无法切换为mysql模式，此时只能
使用扩展支持的部分MySQL特有语法。
```
#### 控制开关

```
在yashan模式下开启兼容MySQL语法后，可以便捷地使用在Oracle中不可用但在MySQL中可用的功能，不开启时仅此类相冲突的功能不可用，其他正常
功能不受影响。
YashanDB使用配置参数SQL_PLUGIN控制是否开启MySQL兼容的开关，默认值为NONE，即默认不开启MySQL兼容，将参数值设为MySQL表示开启
MySQL语法兼容。
```
```
其中，SCOPE=MEMORY表示仅在当前实例运行期间配置生效；SCOPE=SPFILE表示当前实例运行期间不生效，仅在重启后配置生效；不指定SCOPE
表示配置立即生效，且实例重启后仍生效。
```
#### 影响范围

```
在yashan模式下开启兼容MySQL语法开关后，系统将发生如下变化：
特性 开启前 开启后
多表DELETE 不支持 支持
多表UPDATE 不支持 支持
DELETE table语法 支持 不支持
DELETE table FROM table语法 不支持 支持
```
```
示例（单机HEAP表）
```
```
-- 开启MySQL语法兼容
ALTER SYSTEM SET SQL_PLUGIN = 'MYSQL';
或
ALTER SYSTEM SET SQL_PLUGIN = 'MYSQL' SCOPE = MEMORY;
或
ALTER SYSTEM SET SQL_PLUGIN = 'MYSQL' SCOPE = SPFILE;
-- 关闭MySQL语法兼容
ALTER SYSTEM SET SQL_PLUGIN = 'NONE';
或
ALTER或 SYSTEM SET SQL_PLUGIN = 'NONE' SCOPE = MEMORY;
ALTER SYSTEM SET SQL_PLUGIN = 'NONE' SCOPE = SPFILE;
```
```
-- employees为一张员工信息表，包含如下五条数据
SELECT BRANCH,DEPARTMENT,EMPLOYEE_NO,EMPLOYEE_NAME,SEX,ENTRY_DATE FROM employees;
BRANCH DEPARTMENT EMPLOYEE_NO EMPLOYEE_NAME SEX ENTRY_DATE
------ ---------- ------------- ------------- ----- --------------------------------
0101 000 0101000001 Mask 1 2020 - 09 - 09 22 : 55 : 32
```
01010201 000010 01010000020201010011 John Anna 10 20172022 - - 1208 - - 1410 2222 :: 5555 :: (^3232)
0201 008 0201008003 Jack 1 2021 - 07 - 06 22 : 55 : 32
0101 008 0201008004 Jim 1 2022 - 11 - 18 22 : 55 : 32
-- 创建与employees同构的employees2表
DROP TABLE IF EXISTS employees2;
CREATE TABLE employees2 AS SELECT * FROM employees WHERE 1 = 2 ;
INSERT INTO employees2 VALUES ('0101','008','0201008003','Jim','0',DATE '2021-11-17');
INSERT INTO employees2 VALUES ('0101','000','0101000002','John','1',DATE '2021-11-17');
COMMIT;


###### 深圳计算科学研究院 深圳崖山科技有限公司

#### 具体兼容项

```
本文将从如下方面具体说明，在单机部署和行式存储中，YashanDB的yashan模式对MySQL数据库的兼容情况：
SQL语法
表达式运算
FILTER CONDITION
数据类型
内置函数
PL
系统视图
字符集
SQL引擎
数据库安全
保留字
```
##### SQL 语法

```
YashanDB的yashan模式支持MySQL数据库中主流的SQL语法，其他少数因功能性缺失导致的不兼容将报语法不支持错误，此时应联系我们的技术支持
提供变通方案。
（ 1 ） DML 类
SELECT
支持大部分查询功能，包括单表查询，多表查询，子查询，内连接，半连接，外连接，分组及聚合，层次查询等
支持UNION、UNION ALL等集合操作
支持EXPLAIN方式查看执行计划
支持随机抽样查询能力
INSERT
支持单行，多行插入，同时支持指定分区插入
支持INSERT INTO SELECT语句
支持单表插入
支持INSERT ON DUPLICATE KEY UPDATE语法
UPDATE
```
```
-- 开启前
UPDATE employees,employees2 SET EMPLOYEES.EMPLOYEE_NAME = 'TOM',EMPLOYEES2.EMPLOYEE_NAME = 'TOM' WHERE EMPLOYEES.SEX =
EMPLOYEES2.SEX;
YAS- 04344 multi-table update is not supported
DELETE FROM employees,employees2;
YAS - 04345 multi-table delete is not supported
DELETE employees;
DELETE employees FROM employees;
[ 1 : 18 ]YAS- 04209 unexpected word FROM
-- 开启后
ROLLBACK;
ALTER SYSTEM SET SQL_PLUGIN = 'MYSQL' SCOPE = MEMORY;
UPDATE employees,employees2 SET EMPLOYEES.EMPLOYEE_NAME = 'TOM',EMPLOYEES2.EMPLOYEE_NAME = 'TOM' WHERE EMPLOYEES.SEX =
EMPLOYEES2.SEX;
DELETE FROM employees,employees2;
DELETE employees;
[ 1 : 17 ]YAS- 04209 unexpected word employees
DELETE employees FROM employees;
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
支持单列和多列更新
支持使用子查询
DELETE
支持单表和多表删除
支持使用子查询
（ 2 ） DDL 类
HEAP 表
CREATE TABLE/CREATE TABLE AS
ALTER TABLE
DROP TABLE
TRUNCATE TABLE
临时表
支持临时表的创建与删除
分区表
支持RANGE、LIST、HASH分区
支持ADD|DROP|TRUNCATE PARTITION
约束
包括in_line约束、out_of_line约束
UNIQUE、PRIMARY KEY、FOREIGN KEY、CHECK、(NOT)NULL类型约束
视图
CREATE VIEW
DROP VIEW
视图支持SELECT/INSERT/UPDATE/DELETE
BTree 索引
包括唯一索引和非唯一索引
```
##### 表达式运算

```
YashanDB的yashan模式包含了主流的计算框架实现对表达式的运算，详情如下表所示：
表达式运算类型 YashanDB MySQL
二元运算加法 支持 支持
二元运算减法 支持 支持
二元运算乘法 支持 支持
二元运算除法 支持 支持
二元运算取余 支持 支持
一元运算取反 支持 支持
位运算与 支持 支持
位运算或 支持 支持
位运算异或 支持 支持
```
##### FILTER CONDITION

```
YashanDB的yashan模式中，FILTER CONDITON类型除了ROWNUM外，完全兼容MySQL数据库的FILTER CONDITON类型，详情如下表所示：
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
filter YashanDB MySQL
ALL 支持 支持
AND 支持 支持
ANY 支持 支持
BEWTEEN...AND... 支持 支持
EXISTS 支持 支持
GREAT EQUAL 支持 支持
GREATE 支持 支持
IN 支持 支持
IS NOT NULL 支持 支持
IS NULL 支持 支持
LESS 支持 支持
LESS EQUAL 支持 支持
LIKE 支持 支持
NOT 支持 支持
NOT BEWTEEN...AND... 支持 支持
NOT EQUAL 支持 支持
NOT EXISTS 支持 支持
NOT IN 支持 支持
NOT LIKE 支持 支持
NOT RLIKE 支持 支持
OR 支持 支持
EQUAL 支持 支持
RLIKE 支持 支持
ROWNUM 支持 不支持
SOME 支持 支持
```
##### 数据类型

```
YashanDB的yashan模式实现的数据类型与MySQL数据库对比情况如下表所示：
数据类型 YashanDB MySQL
BOOLEAN 支持 支持
TINYINT 支持 支持
SMALLINT 支持 支持
MEDIUMINT 不支持 支持
INTEGER 支持 支持
BIGINT 支持 支持
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
数据类型 YashanDB MySQL
FLOAT 支持 支持
DOUBLE 支持 支持
NUMBER 支持 不支持
DECIMAL / NUMERIC 不支持 支持
DATE 支持 支持
DATETIME 不支持 支持
TIMESTAMP 支持 支持
TIME 支持 支持
YEAR 不支持 支持
INTERVAL YEAR TO MONTH 支持 不支持
INTERVAL DAY TO SECOND 支持 不支持
CHAR 支持 支持
VARCHAR 支持 支持
RAW 支持 不支持
CLOB 支持 不支持
TINYBLOB 不支持 支持
BLOB 支持 支持
MEDIUMBLOB 不支持 支持
TINYTEXT 不支持 支持
TEXT 不支持 支持
LONGTEXT 不支持 支持
BIT 不支持 支持
ROWID 支持 不支持
JSON 支持 支持
```
##### 内置函数

```
YashanDB的yashan模式实现的内置函数与MySQL数据库的内置函数对比情况如下表所示：
函数类型 支持函数列表 YashanDB MySQL
数学运算函
数
```
```
ABS、ACOS、ASIN、ATAN、ATAN2、AVG、CEIL/CEILING、COS、COT、DIV、FLOOR、MOD、
PI、POW/POWER、RANDOM/RAND、SIGN、SIN、SQRT、TAN、TRUNCATE/TRUNC 支持 支持
```
```
字符处理函
数
```
```
ASCII、BIT_LENGTH、CHR/CHALEASTR 、CHAR_LENGTH/CHARACTER_LENGTH、CONCAT、
CONCAT_WS、FIND_IN_SET、GROUP_CONCAT、INSTR、 LCASE/LOWER、 LEFT、
LENGTH、LPAD、LTRIM、POSITION、OCTET_LENGTH、RIGHT、RPAD、RTRIM、REPLACE、
SUBSTR、SUBSTRING、SUBSTRING_INDEX、TRIM、 UCASE/UPPER
```
```
支持 支持
```
```
正则匹配函
数 REGEXP_LIKE、REGEXP_REPLACE、REGEXP_INSTR、REGEXP_SUBSTR 支持 支持
转换函数 BIN、CAST 支持 支持
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
函数类型 支持函数列表 YashanDB MySQL
聚集函数 AVGSUM、、COUNTVAR_POP、GROUP_CONCA、VAR_SAMP、VAR_SAMPT、MAX、MIN、VARIANCE、STDDEV、STDDEV_POP、STDDEV_SAMP、 支持 支持
窗口函数 FIRST_VALUE、LAG、LAST_VALUE、LEAD、RANK、ROW_NUMBER 支持 支持
时间处理函
数
```
```
CURRENT_TIMESTAMP、DATE、DAYOFWEEK、DATE_FORMAT、DATE_ADD、DATE_SUB、
EXTRACT、LAST_DAY、LOCALTIMESTAMP、NOW、SYSDATE、TIME、TIMESTAMP、
TIMEDIFF、TIMESTAMPDIFF、UTC_TIMESTAMP
```
```
支持 支持
```
```
条件处理函
数 CASE、IF、IFNULL 支持 支持
JSON处理
函数
```
```
JSON、JSON_ARRAY_GET、JSON_ARRAY_LENGTH、JSON_EXISTS、JSON_FORMAT、
JSON_PARSE、JSON_QUERY、JSON_SERIALIZE 支持 不支持
MySQL
information
函数
```
```
BENCHMARK、CHARSET、COERCIBILITY、COLLATION、CONNECTION_ID、
CURRENT_ROLE、CURRENT_USER、DATABASE、ICU_VERSION、ROLES_GRAPHML、
ROW_COUNT、SCHEMA、SESSION_USER、USER、VERSION
```
```
不支持 支持
```
```
MySQL加
密和压缩函
数
```
```
AES_DECRYPT、AES_ENCRYPT、COMPRESS、RANDOM_BYTES、SHA、SHA1、SHA2、
STATEMENT_DIGEST、 STATEMENT_DIGEST_TEXT、UNCOMPRESS、
UNCOMPRESSED_LENGTH
```
```
不支持 支持
```
```
其他函数 BITAND/BIT_ANDLAST_INSERT_ID、、 BITOR/BIT_ORLEAST、SOUNDEX、BITXOR/BIT_XOR、MD5 、COALESCE、ISNULL、GREATEST、 支持 支持
```
##### PL

```
YashanDB的yashan模式兼容了MySQL数据库大部分的PL功能，但语法格式上有一定的区别，包括：
数据类型
流程控制
静态SQL
动态SQL
异常处理
系统定义异常
用户自定义异常
游标
存储过程
触发器
支持行级触发器，但不支持语句级触发器
目前仅支持在表上创建触发器，不支持在视图上创建触发器
用户自定义函数（SQL语言的UDF）
JOB
```
##### 字符集

```
YashanDB的yashan模式目前支持以下字符集：
字符集 YashanDB MySQL
ASCII 支持 支持
GBK 支持 支持
UTF-8 支持 支持
ISO88591 支持 不支持
GB18030 支持 支持
```
```
同时支持如下字符集排序方式：
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
ASCII_GENERAL_CS
ASCII_GENERAL_CI
GBK_GENERAL_CS
GBK_GENERAL_CI
UTF8_GENERAL_CS
UTF8_GENERAL_CI
UTF8_PINYIN_CS
UTF8_PINYIN_CI
ISO88591_GENERAL_CS
ISO88591_GENERAL_CI
GB18030_GENERAL_CS
GB18030_GENERAL_CI
GB18030_PINYIN_CS
GB18030_PINYIN_CI
```
##### SQL 引擎

```
YashanDB的yashan模式下，SQL引擎兼容了MySQL数据库大部分的特性，包括：
支持查询改写
支持预编译语句
支持基于成本的优化器
支持计划生成与展示（EXPLAIN）
支持执行计划缓存
支持执行计划快速参数化
支持Optimizer Hint
支持like和reglike的模糊匹配能力
```
##### 保留字兼容

```
YashanDB的yashan模式兼容了MySQL数据库大部分保留字，但目前个别保留字仍存在一定区别，包括：
CONNECT
EXCEPT
IF
INTERSECT
MINUS
OF
PUBLIC
ROW
ROWS
START
SYSDATE
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## 测试报告

```
本文为您介绍YashanDB在OLTP场景TPC-C模型及OLAP场景TPC-H模型的性能测试数据，包括如下几部分：
TPC-C测试：基于集中式部署的YashanDB模拟TPC-C模型进行测试。
TPC-H测试：基于分布式部署的YashanDB模拟TPC-H模型进行测试。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## TPC-C 测试

```
本章节将介绍在YashanDB单机数据库上运行基于BenchmarkSQL的TPC-C测试的具体操作及相关示例。
```
### TPCC-C 简介

```
TPC-C测试是针对业务处理系统的规范，该规范由TPC（Transaction Processing Performance Council）委员会制定，测试结果主要取决于流量指标和
性价比指标，是最常用的在线事务处理（OLTP）测试基准。
TPC-C测试中模拟了大型商品批发商的交易业务场景，该批发商总共生产销售 100000 种产品，拥有多个负责不同区域的商品仓库，每个仓库需要为 10 个
分销商进行供货，且每个分销商需要为 3000 个客户进行服务，对数据库进行TPC-C测试时可根据服务器及数据库的配置对商品仓库的数量进行调整，从
而模拟出不同的压力场景。
业务场景中主要包含如下 5 类事务：
Stock-Level：用于表示分销商库存状态，库存较低时需要进行补货。
New-Order：用于表示客户提交了一笔新订单，每笔订单平均包括 10 件产品。
Payment：用于表示客户为订单支付费用，更新其账户余额。
Delivery：用于表示根据用户提交的订单进行发货。
Order-Status：用于表查询用户的最近交易状态。
```
### TPC-C 测试工具下载

```
请自行于BenchmarkSQL官网下载BenchmarkSQL5 。
请确保服务器中已有JDK1.8及以上版本的JDK。
```
### TPC-C 测试准备

```
对YashanDB进行TPC-C测试前需对Benchmark SQL5进行配置，使之支持YashanDB数据库：
```
#### 修改 jTPCC.java 文件

```
1. 在操作系统终端执行如下命令并输入密码，切换至root用户。
```
```
2. 执行如下命令，进入tpc-c目录。
```
```
3. 执行如下命令，在vi编辑器中打开/home/yashan/tpc-c/benchmarksql-5.0/src/client/jTPCC.java文件，请注意区分大小写：
```
```
4. 在文件中查找下图内容：
```
```
5. 按 i 进入编辑模式，在dbType = DB_POSTGRES后新增如下内容：
```
$ Password:su root (^)
$ cd tpc-c
$ vi benchmarksql-5.0/src/client/jTPCC.java


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
6. 修改完成后，按 Esc ，输入:wq保存并退出文件编辑。
```
#### 修改 jTPCCConfig.java 文件

```
1. 执行如下命令，在vi编辑器中打开jTPCCConfig.java文件：
```
```
2. 在文件中查找下图内容：
```
```
3. 按 i 进入编辑模式，将该部分内容修改为如下内容：
```
```
4. 修改完成后，按 Esc ，输入:wq保存并退出文件编辑。
```
#### 编译源码

```
执行如下命令进入benchmarksql-5.0目录，并执行ant命令进行编译：
```
```
Note :
如此时返回ant:command not found，可通过执行yum install ant安装ant编译工具。
```
```
else if (iDB.equals("yashandb"))
dbType = DB_YASHANDB;
```
```
$ vi benchmarksql-5.0/src/client/jTPCCConfig.java
```
```
public final static int DB_UNKNOWN = 0 ,
DB_FIREBIRD = 1 ,
DB_ORACLE = 2 ,
DB_POSTGRES = 3 ,
DB_YASHANDB = 4 ;
```
```
$ cd benchmarksql-5.0
$ ant
Buildfile: /home/yashan/tpc-c/benchmarksql-5.0/build.xml
init:
```
compile: [javac (^) ] Compiling 11 source files to /home/yashan/tpc-c/benchmarksql-5.0/build
dist:
[mkdir] Created dir: /home/yashan/tpc-c/benchmarksql-5.0/dist
[jar] Building jar: /home/yashan/tpc-c/benchmarksql-5.0/dist/BenchmarkSQL-5.0.jar
BUILD SUCCESSFUL
Total time: 1 second


###### 深圳计算科学研究院 深圳崖山科技有限公司

#### 创建文件 props.yashandb

```
1. 执行如下命令，进入/home/yashan/tpc-c/benchmarksql-5.0/run目录：
```
```
2. 执行如下命令，通过vi编辑器创建文件props.yashandb：
```
```
3. 按 i 进入编辑模式，将如下内容新增至文件中（可根据实际环境进行修改）：
```
```
其中：
参数 含义
db 数据库，须与上述修改文件中添加的数据库名称相同
driver 驱动程序文件，须使用YashanDB的JDBC驱动
conn 连接描述符，格式为：conn=jdbc:yasdb://ip:port/database_name
user 数据库用户
password 数据库用户的密码
warehouses 用于指定测试中商品仓库的数量，通常测试场景为100~1000仓
loadWorkers 数据导入的并发数，通常设置为CPU线程总数的2~4倍
terminals 业务运行的并发数，通常设置为CPU线程总数的2~4倍
runTxnsPerTerminal 每个会话运行的固定事务数量，通常配置为 0 ，以runMins限制测试时间
runMins 用于指定测试运行的时间，通常建议压测时间为10~30分钟
limitTxnsPerMin 每秒事务数上限，压测场景下通常设置为 0
terminalWarehouseFixed 会话和仓库的绑定模式，通常设置为true
```
```
$ cd /home/yashan/tpc-c/benchmarksql-5.0/run
```
```
$ vi props.yashandb
```
```
db=yashandb
driver=com.yashandb.jdbc.Driver
conn=jdbc:yasdb://localhost:1688/yashandb
user=sys
password=sys
```
warehousesloadWorkers= (^10) = 2
terminals= 10
runTxnsPerTerminal= 0
runMins= 5
limitTxnsPerMin= 0
terminalWarehouseFixed=true
newOrderWeight= 45
paymentWeight= 43
orderStatusWeight= 4
deliveryWeight= 4
stockLevelWeight= 4
resultDirectoryosCollectorScript=my_result_%tY-%tm-%td_%tH%tM%tS=./misc/os_collector_linux.py
osCollectorInterval= 1


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
参数 含义
newOrderWeight
paymentWeight
orderStatusWeight
deliveryWeight
stockLevelWeight
```
```
五种事务的占比，所有值的和须为 100
标准的业务比率配置为45:43:4:4:4
```
```
resultDirectory 指定存储测试结果的目录。%tY、%tm、%td、%tH、%tM和%tS是时间格式化参数
osCollectorScript 指定操作系统收集器的脚本路径。通常用于收集操作系统资源使用情况的统计信息，例如况、磁盘I/O等 CPU使用率、内存使用情
osCollectorInterval 指定操作系统收集器脚本运行间隔（单位：秒）
```
```
4. 修改完成后，按 Esc ，输入:wq保存并退出文件编辑。
```
#### 修改文件 funcs.sh

```
1. 执行如下命令，在vi编辑器中打开funcs.sh文件：
```
```
2. 按 i 进入编辑模式，将文档内容替换为如下内容：
```
```
$ vi /home/yashan/tpc-c/benchmarksql-5.0/run/funcs.sh
```
```
# ----
# $ 1 is the properties file
# ----
PROPS=$ 1
if [! - f ${PROPS} ] ; then
echo "${PROPS}: no such file" >& 2
exit 1
fi
# ----
# getProp()
#
# # --Get-- a config value from the properties file.
function getProp()
{
grep "^${1}=" ${PROPS} | sed -e "s/^${1}=//"
}
# ----
# getCP()
#
# Determine the CLASSPATH based on the database system.
# ----
function setCP()
{ case "$(getProp db)" in
firebird)
cp="../lib/firebird/*:../lib/*"
;;
oracle)
cp="../lib/oracle/*"
if [! - z "${ORACLE_HOME}" - a -d ${ORACLE_HOME}/lib ] ; then
cp="${cp}:${ORACLE_HOME}/lib/*"
fi
cp="${cp}:../lib/*"
;;
postgres)
cp ;;= "../lib/postgres/*:../lib/*"
yashandb)
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
3. 修改完成后，按 Esc ，输入:wq保存并退出文件编辑。
```
#### 添加 yashandb java connector 驱动

```
执行此步骤前须确保已上传YashanDB JDBC驱动，本例中驱动包所在路径为/home/yashan/yashandb_jdbc/。
执行如下命令添加驱动：
```
#### 修改 runDatabaseBuild.sh 文件

```
1. 执行如下命令，在vi编辑器中打开runDatabaseBuild.sh文件：
```
```
2. 在文件中查找下图内容：
```
```
3. 按 i 进入编辑模式，将该部分内容修改为如下内容：
```
```
4. 修改完成后，按 Esc ，输入:wq保存并退出文件编辑。
```
### TPC-C 测试运行

#### 部署 YashanDB 数据库

```
最佳性能数据会因为测试环境的CPU、内存、IO、网络条件差异而不同，为了让YashanDB在测试环境上达到最佳的TPC-C性能表现，需要根据测试环境
的配置进行性能调优，数据库性能调优的详细信 息请查阅YashanDB性能调优文档。
TPC-C测试调优主要分为参数配置调优和建库配置调优：
数据库参数配置调优
在TPC-C测试场景下主要关注缓存大小与分区、IO参数等性能参数的配置。
以 1000 仓、 256 并发的测试场景为例，推荐配置以下性能调优参数：
```
```
cp="../lib/yashandb/*:../lib/*"
;;
esac
myCP=".:${cp}:../dist/*"
export myCP
}
```
# # --Make-- (^) sure that the properties file does have db= and the value
# is a database, we support.
# ----
case "$(getProp db)" in
firebird|oracle|postgres|yashandb)
;;
"") echo "ERROR: missing db= config option in ${PROPS}" >& 2
exit 1
;;
*) echo "ERROR: unsupported database type 'db=$(getProp db)' in ${PROPS}" >& 2
exit 1
;;
esac
$ mkdir -p /home/yashan/tpc-c/benchmarksql-5.0/lib/yashandb/
$ cp /home/yashan/yashandb_jdbc/yashandb-jdbc-1.5-SNAPSHOT.jar /home/yashan/tpc-c/benchmarksql-5.0/lib/yashandb/
$ vi /home/yashan/tpc-c/benchmarksql-5.0/run/runDatabaseBuild.sh
AFTER_LOAD="indexCreates foreignKeys buildFinish"


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
数据库配置参数的详细说明请查阅配置参数。
数据库建库配置调优
Note :
安装部署后，YashanDB将默认创建一个初始数据库，可根据实际需求删除初始数据库（DROP DATABASE）后自定义创建数据库（不适用于标准版
或企业版的分布式部署）。
创建数据库时主要关注以下方面：
redo文件的大小：redo文件配置太小，在压测场景下容易出现redo追尾的情况，严重影响数据库性能。通过V$SYSTEM_EVENT视图中的“checkpoint
completed”的等待事件判断是否出现redo追尾的情况。
DATA文件的大小：预占数据空间可以避免数据库运行过程中空间动态扩展对性能的影响，因此为了最佳性能表现，需要初始化足够大的表空间。
在实际使用过程中，也应及时调整redo文件、DATA文件等相关配置，具体可查阅文件管理。
如果存储条件允许，请将redo文件与数据文件分盘部署，以减少两者间的IO争用，以 1000 仓、 256 并发的测试场景为例，推荐的建库语句如下：
```
```
数据库配置参数和建库配置调优的详细介绍请查阅数据库配置调优。
```
#### 清理 TPC-C 数据

```
在/home/yashan/tpc-c/benchmarksql-5.0/run目录中，执行如下命令进行数据清理：
```
```
# Data buffer用于数据块的缓存，其大小会影响数据访问的缓存命中率。建议将规划内存的80%配置为数据缓存区，如果缓存区的大小大于总数据大小可取得最佳性
能。
DATA_BUFFER_SIZE= 200 G
# Data buffer的分区数，在大并发的测试场景下，将Data buffer分区可以降低缓存区的锁冲突。
_DATA_BUFFER_PARTS= 8
# VM Buffer用于保存例如order/group by等数据运算的中间结果，当VM空间不足时会产生内存与SWAP表空间的换入，影响数据库性能表现。
# 因此需要配置合理的VM Buffer大小以避免换入的产生。通过视图V$VM中的SWAPPED_OUT_BLOCKS字段可以获取SWAP的次数，当为 0 时可以取到最佳性能。
```
VM_BUFFER_SIZE# VM Buffer的分区数，在大并发的测试场景下，将= 25 G (^) VM buffer可以降低缓存区的锁冲突。
VM_BUFFER_PARTS= 8
# 全局大页内存区大小，在大并发的测试场景下，需要提高配置以避免资源不足。
LARGE_POOL_SIZE= 1 G
# 全局执行内存区大小，在大并发的测试场景下，需要提高配置以避免资源不足。
WORK_AREA_POOL_SIZE= 2 G
# undo数据的保留时间，用于一致性读或者数据闪回。对于类似TPC-C小事务的场景下，降低undo数据的保留时间可以提高undo分配效率，提升数据库性能。
UNDO_RETENTION= 15
# 会话级CURSOR的数量，在大并发的测试场景下提高配置可以消除全局CURSOR的竞争。
_SESSION_RESERVED_CURSORS= 64
# 增量Checkpoint的时间间隔，后台脏块刷盘会与redo刷盘产生IO争用，在DATA和redo同盘部署的场景下，降低Checkpoint频率可以提升数据库性能。
CHECKPOINT_TIMEOUT= 900
# 指定触发checkpoint的从恢复点到当前redo日志刷盘点的redo大小间隔，通常配置为redo文件总大小的一半。
CHECKPOINT_INTERVAL= 10 G
# 共享内存池的大小，提高配置以避免SQL缓存或者元数据缓存的频繁失效。
SHARE_POOL_SIZE= 2 G
CREATE DATABASE tpcc LOGFILE(
'/data1/redo1' size 20 G BLOCKSIZE 512 ,
'/data1/redo2' size 20 G BLOCKSIZE 512 ,
'/data1/redo3' size 20 G BLOCKSIZE 512 ,
'/data1/redo4' size 20 G BLOCKSIZE 512 ,
'/data1/redo5' size 20 G BLOCKSIZE 512 ,
'/data1/redo6' size 20 G BLOCKSIZE 512 ,
'/data1/redo7' size 20 G BLOCKSIZE 512 ,
'/data1/redo8' size 20 G BLOCKSIZE 512 ,
'/data1/redo9' size 20 G BLOCKSIZE 512 ,
'/data1/redo10' size 20 G BLOCKSIZE 512 )
UNDO TABLESPACE DATAFILE '/data2/undo' size 10 G
SWAP SYSTEM TABLESPACETABLESPACE TEMPFILE DATAFILE '/data2/swap''/data2/system' size size 10 G (^5) G
SYSAUX TABLESPACE DATAFILE '/data2/sysaux' size 5 G
DEFAULT TABLESPACE DATAFILE '/data2/users' size 300 G;


###### 深圳计算科学研究院 深圳崖山科技有限公司

#### TPC-C 数据装载

```
目录中执行如下命令进行数据导入：
```
```
$ ./runDatabaseDestroy.sh props.yashandb
# ------------------------------------------------------------
# Loading SQL file ./sql.common/tableDrops.sql
# ------------------------------------------------------------
drop table bmsql_config;
drop table bmsql_new_order;
drop table bmsql_order_linedrop table bmsql_oorder; ;
drop table bmsql_history;
drop table bmsql_customer;
drop table bmsql_stock;
drop table bmsql_item;
drop table bmsql_district;
drop table bmsql_warehouse;
drop sequence bmsql_hist_id_seq;
```
```
$ ./runDatabaseBuild.sh props.yashandb
# ------------------------------------------------------------
# Loading SQL file ./sql.common/tableCreates.sql
# ------------------------------------------------------------
create table bmsql_config (
cfg_name varchar( 30 ) primary key,
cfg_value varchar( 50 )
);
create table bmsql_warehouse (
w_id integer not null,w_ytd decimal(12,2),
w_tax decimal(4,4),
w_name varchar( 10 ),
w_street_1 varchar( 20 ),
w_street_2 varchar( 20 ),
w_city varchar( 20 ),
w_state char( 2 ),
w_zip char( 9 )
);
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

#### 运行 TPC-C 测试

```
执行如下语句进行TPC-C测试：
$ ./runBenchmark.sh props.yashandb
20 :50:04,561 [main] INFO jTPCC : Term-00,
20 :50:04,563 [main] INFO jTPCC : Term-00, +-------------------------------------------------------------+
20 :50:04,563 [main] INFO jTPCC : Term-00, BenchmarkSQL v5.0
20 :50:04,563 [main] INFO jTPCC : Term-00, +-------------------------------------------------------------+
20 :50:04,563 [main] INFO jTPCC : Term-00, (c) 2003 , Raul Barbosa
20 :50:04,563 [main] INFO jTPCC : Term-00, (c) 2004 -2016, Denis Lussier
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
测试结果中的tpmC值表示每分钟内系统处理的新订单个数，即系统最大吞吐量。tpmC值常作为性能指标，值越高表示数据库性能越好。
```
```
20 :50:04,565 [main] INFO jTPCC : Term-00, (c) 2016 , Jan Wieck
20 :50:04,565 [main] INFO jTPCC : Term-00, +-------------------------------------------------------------+
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## TPC-H 测试（分布式）

```
本章节将介绍如何在YashanDB分布式数据库上进行TPC-H测试。
```
### TPC-H 简介

```
TPC-H（商业智能计算测试）是美国交易处理效能委员会（TPC，Transaction Processing Performance Council）组织制定的用来模拟决策支持类应用
的一个测试集。这种商业测试可以全方位评测系统的整体商业计算综合能力，对厂商的要求更高，同时也具有普遍的商业实用意义，目前在银行信贷分析
和信用卡分析、电信运营分析、税收分析、烟草行业决策分析中都有广泛的应用。
TPC-H基准测试由TPC-D（由TPC于 1994 年制定的标准，用于决策支持系统方面的测试基准）发展而来的。TPC-H用3NF实现了一个数据仓库，共包含 8
个基本关系，其主要评价指标是各个查询的响应时间，即从提交查询到结果返回所需时间。
```
### 环境要求

```
JDK：用于JDBC方式连接数据库，建议OpenJDK 1.8.0_412及以上版本，安装操作请查阅OpenJDK官网 。
make：用于编译TPC-H数据生成工具，建议GUN Make 3.82及以上版本，安装操作请查阅GUN Make官网 。
gcc：用于编译TPC-H数据生成工具，建议GCC 10.2.1及以上版本，安装操作请查阅GCC官网 。
YashanDB JDBC驱动：用于JDBC方式连接数据库，建议1.6.1及以上版本，安装操作请查阅安装YashanDB JDBC驱动。
Python：用于执行TPC-H测试脚本，建议Python 3.6.8及以上版本，安装操作请查阅Python官网 。
JPype1：用于执行TPC-H测试脚本的依赖库，建议1.3.0及以上版本，安装操作请查阅PyPI官网 。
TPC-H Tool：用于生成TPC-H测试相关数据的工具，建议3.0.0版本，安装操作请查阅TPC-H官网 。
YashanDB数据库：用于TPC-H测试的数据库，建议23.2.8.100及后续补丁版本，安装操作请查阅分布式部署。TPC-H测试的参数推荐特性暂时只在
23.2大版本受支持。
```
### 分布式数据库部署

```
需准备 4 台硬件规格一致的服务器，且建议服务器硬件规格不低于16C32G，否则会影响TPC-H SF100的性能测试结果。如果是TPC-H SF1000的性能测
试，建议服务器硬件规格不低于32C256G。另外需手动调整集群部署文件为CN节点设置USERS_DATASPACE_SCALE_OUT_FACTOR建库参数值，推
荐设置为CPU核数的两倍且不超过 128 。
服务器要求以及分布式各节点/节点组规划如下表所示。
服务器名称 服务器 IP 节点 / 节点组 内存 CPU 磁盘
服务器 1 192.168.1.1 MN：1-1（主），CN：2-1（主） 64G 16 核 HDD
服务器 2 192.168.1.2 DN：3-1（主） 64G 16 核 SSD
服务器 3 192.168.1.3 DN：4-1（主） 64G 16 核 SSD
服务器 4 192.168.1.4 DN：5-1（主） 64G 16 核 SSD
```
```
若服务器的硬件规格不一致，需手动设置CN节点的并行度以使DN节点充分利用CPU资源，建议值为DN节点所在服务器的CPU核数。
```
### 运行 TPC-H 测试

#### 数据库参数调优

```
YashanDB在分布式部署下支持TPC-H测试的参数推荐，通过参数推荐高级包来设置调优后的参数。
如需进行TPC-H SF100的基准测试，可执行如下命令应用调优参数，执行成功后还需重启数据库使配置生效。
```
```
Note :
TPC-H场景推荐参数作用范围仅限于TPC-H测试查询阶段，建议在数据导入、统计信息收集等准备步骤完成后再执行命令进行参数调优。
```
```
EXEC DBMS_PARAM.OPTIMIZE(True, 'LSC', scene=>'TPCH', scale_factor=> 100 );
SELECT DBMS_PARAM.SHOW_RECOMMEND() FROM dual;
EXEC DBMS_PARAM.APPLY_RECOMMEND();
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
参数推荐仅在硬件规格足够的情况下使TPC-H测试的性能最优。例如，TPC-H SF100的场景下，DN节点至少需要16C32G的规格。
```
#### 安装 TPC-H 工具

```
1. 下载 TPC-H Tool。
2. 解压文件并进入指定路径。
```
```
3. 在tpcd.h头文件中增加YashanDB编译选项。
```
```
4. 基于编译模板文件修改CC、DATABASE、MACHINE、WORKLOAD 等参数定义生成编译文件。
```
```
5. 编译文件。
```
#### 生成数据

```
可按需通过指定参数生成TPC-H规范的任意数据规模，例如1G、10G、100G、1T等，本文以100G为例。
```
#### 新建测试用户和测试表

```
1. 新建测试用户并使用测试用户连接数据库。
```
```
2. 新建测试表。
```
```
$ unzip TPC-H_Tools_v3.0.0.zip
$ cd TPC-H_Tools_v3.0.0/dbgen
```
```
$ vi tpcd.h
#ifdef YASHANDB
#define GEN_QUERY_PLAN ""
#define START_TRAN ""
#define END_TRAN "COMMIT"
#define SET_OUTPUT ""
#define SET_ROWCOUNT "limit %d;\n"
#define SET_DBASE ""
#endif
```
```
$ cp makefile.suite Makefile
$ vi Makefile
CC = gcc
DATABASE = YASHANDB
MACHINE = LINUX
WORKLOAD = TPCH
```
```
$ make
```
```
$ ./dbgen -s 100
```
$ $ mkdirmv *.tbl /data/tpch/SF100 -p /data/tpch/SF100 (^)
CREATE USER regress IDENTIFIED BY regress;
GRANT dba TO regress;
conn regress/regress;
DROP TABLE IF EXISTS region;
CREATE TABLE region (
R_REGIONKEY INTEGER NOT NULL,
R_NAME CHAR( 25 ) NOT NULL,
R_COMMENT VARCHAR( 152 )
) ORGANIZATION LSC
ORDER BY (R_REGIONKEY);


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
DROP TABLE IF EXISTS nation;
CREATE TABLE nation (
N_NATIONKEY INTEGER NOT NULL,
N_NAME CHAR( 25 ) NOT NULL,
N_REGIONKEY INTEGER NOT NULL,
N_COMMENT VARCHAR( 152 )
```
)ORDER ORGANIZATION LSC BY (N_NATIONKEY (^) );
DROP TABLE IF EXISTS supplier;
CREATE TABLE supplier (
S_SUPPKEY INTEGER NOT NULL,
S_NAME CHAR( 25 ) NOT NULL,
S_ADDRESS VARCHAR( 40 ) NOT NULL,
S_NATIONKEY INTEGER NOT NULL,
S_PHONE CHAR( 15 ) NOT NULL,
S_ACCTBAL DECIMAL( 15 , 2 ) NOT NULL,
S_COMMENT VARCHAR( 101 ) NOT NULL
) ORGANIZATION LSC
ORDER BY (S_SUPPKEY)
PARTITION BY HASH(S_SUPPKEY) PARTITIONS AUTO;
DROP TABLE IF EXISTS part;
CREATE TABLE part (
P_PARTKEY INTEGER NOT NULL,
P_NAME VARCHAR( 55 ) NOT NULL,
P_MFGR CHAR( 25 ) NOT NULL ENCODING DICTIONARY(RLE),
P_BRAND CHAR( 10 ) NOT NULL ENCODING DICTIONARY(RLE),
P_TYPE VARCHAR( 25 ) NOT NULL ENCODING DICTIONARY(RLE),
P_SIZE INTEGER NOT NULL,
P_CONTAINER CHAR( 10 ) NOT NULL ENCODING DICTIONARY(RLE),
P_RETAILPRICE P_COMMENT DECIMALVARCHAR(( 1523 ,) 2 ) NOTNOT NULLNULL,
) ORGANIZATION LSC
ORDER BY (P_PARTKEY)
PARTITION BY HASH(P_PARTKEY) PARTITIONS AUTO;
DROP TABLE IF EXISTS partsupp;
CREATE TABLE partsupp (
PS_PARTKEY INTEGER NOT NULL,
PS_SUPPKEY INTEGER NOT NULL,
PS_AVAILQTY INTEGER NOT NULL,
PS_SUPPLYCOST DECIMAL( 15 , 2 ) NOT NULL,
PS_COMMENT VARCHAR( 199 ) NOT NULL
)ORDER ORGANIZATION LSC BY (PS_SUPPKEY (^) , PS_PARTKEY)
PARTITION BY HASH(PS_PARTKEY) PARTITIONS AUTO;
DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
C_CUSTKEY INTEGER NOT NULL,
C_NAME VARCHAR( 25 ) NOT NULL,
C_ADDRESS VARCHAR( 40 ) NOT NULL,
C_NATIONKEY INTEGER NOT NULL,
C_PHONE CHAR( 15 ) NOT NULL,
C_ACCTBAL DECIMAL( 15 , 2 ) NOT NULL,
C_MKTSEGMENT CHAR( 10 ) NOT NULL ENCODING DICTIONARY(RLE),
C_COMMENT ) ORGANIZATION LSCVARCHAR ( 117 ) NOT NULL
ORDER BY (C_CUSTKEY)
PARTITION BY HASH(C_CUSTKEY) PARTITIONS AUTO;
DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
O_ORDERKEY INTEGER NOT NULL,
O_CUSTKEY INTEGER NOT NULL,
O_ORDERSTATUS CHAR( 1 ) NOT NULL,
O_TOTALPRICE DECIMAL( 15 , 2 ) NOT NULL,
O_ORDERDATE DATE NOT NULL,


###### 深圳计算科学研究院 深圳崖山科技有限公司

#### 准备查询语句

```
TPC-H测试包含 22 条查询语句，需将其保存为.yml文件并与执行脚本放在相同路径，供执行脚本解析使用，文件名以tpch.yml为例。
```
```
O_ORDERPRIORITY CHAR( 15 ) NOT NULL ENCODING DICTIONARY(RLE),
O_CLERK CHAR( 15 ) NOT NULL ENCODING DICTIONARY(RLE),
O_SHIPPRIORITY INTEGER NOT NULL,
O_COMMENT VARCHAR( 79 ) NOT NULL
) ORGANIZATION LSC
ORDER BY(O_ORDERDATE, O_ORDERKEY)
PARTITION BY HASH(O_ORDERKEY) PARTITIONS AUTO;
```
(^) DROP TABLE IF EXISTS lineitem;
CREATE TABLE lineitem (
L_ORDERKEY INTEGER NOT NULL,
L_PARTKEY INTEGER NOT NULL,
L_SUPPKEY INTEGER NOT NULL,
L_LINENUMBER INTEGER NOT NULL,
L_QUANTITY DECIMAL( 15 , 2 ) NOT NULL,
L_EXTENDEDPRICE DECIMAL( 15 , 2 ) NOT NULL,
L_DISCOUNT DECIMAL( 15 , 2 ) NOT NULL,
L_TAX DECIMAL( 15 , 2 ) NOT NULL,
L_RETURNFLAG CHAR( 1 ) NOT NULL,
L_LINESTATUS CHAR( 1 ) NOT NULL,
L_SHIPDATE DATE NOT NULL,
L_COMMITDATE DATE NOT NULL,
L_RECEIPTDATE DATE NOT NULL,
L_SHIPINSTRUCT CHAR( 25 ) NOT NULL ENCODING DICTIONARY(RLE),
L_SHIPMODE CHAR( 10 ) NOT NULL ENCODING DICTIONARY(RLE),
L_COMMENT VARCHAR( 44 ) NOT NULL
) ORGANIZATION LSC
ORDER BY(L_SHIPDATE, L_ORDERKEY)
PARTITION BY HASH(L_ORDERKEY) PARTITIONS AUTO;
queries Q1: |:
-- Q1: Pricing Summary Report Query
select
l_returnflag,
l_linestatus,
sum(l_quantity) as sum_qty,
sum(l_extendedprice) as sum_base_price,
sum(l_extendedprice * (1 - l_discount)) as sum_disc_price,
sum(l_extendedprice * (1 - l_discount) * (1 + l_tax)) as sum_charge, avg(l_quantity) as avg_qty,
avg(l_extendedprice) as avg_price,
avg(l_discount) as avg_disc,
count(*) as count_order
from
lineitem
where
l_shipdate <= date '1998-12-01' - interval '90' day
group by
l_returnflag,
l_linestatus
order by
l_returnflag, l_linestatus
Q2: |
-- Q2: Minimum Cost Supplier Query
select
s_acctbal,
s_name,
n_name,
p_partkey,
p_mfgr,
s_address,
s_phone,


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
s_comment
from
part,
supplier,
partsupp,
nation,
region
```
where p_partkey = ps_partkey (^)
and s_suppkey = ps_suppkey
and p_size = 15
and p_type like '%BRASS'
and s_nationkey = n_nationkey
and n_regionkey = r_regionkey
and r_name = 'EUROPE'
and ps_supplycost = (
select
min(ps_supplycost)
from
partsupp,
supplier,
nation,
region
where
p_partkey = ps_partkey
and s_suppkey = ps_suppkey
and s_nationkey = n_nationkey
and n_regionkey = r_regionkey
and r_name = 'EUROPE'
)
order by
s_acctbal desc,
n_name, s_name, (^)
p_partkey
limit 100
Q3: |
-- Q3: Shipping Priority Query
select
l_orderkey,
sum(l_extendedprice*(1-l_discount)) as revenue,
o_orderdate,
o_shippriority
from
customer,
orders, lineitem (^)
where
c_mktsegment = 'BUILDING'
and c_custkey = o_custkey
and l_orderkey = o_orderkey
and o_orderdate < date '1995-03-15'
and l_shipdate > date '1995-03-15'
group by
l_orderkey,
o_orderdate,
o_shippriority
order by
revenue desc, o_orderdate
limit 10
Q4: |
-- Q4: Order Priority Checking Query
select
o_orderpriority,
count(*) as order_count
from
orders
where
o_orderdate >= date '1993-07-01'


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
and o_orderdate < date '1993-07-01' + interval '3' month
and exists (
select
*
from
lineitem
where
```
l_orderkey = o_orderkey and l_commitdate < l_receiptdate (^)
)
group by
o_orderpriority
order by
o_orderpriority
Q5: |
-- Q5: Local Supplier Volume Query
select
n_name,
sum(l_extendedprice * (1 - l_discount)) as revenue
from
customer,
orders,
lineitem,
supplier,
nation,
region
where
c_custkey = o_custkey
and l_orderkey = o_orderkey
and l_suppkey = s_suppkey
and c_nationkey = s_nationkey
and s_nationkey = n_nationkey
and n_regionkey = r_regionkey and r_name = 'ASIA'
and o_orderdate >= date '1994-01-01'
and o_orderdate < date '1994-01-01' + interval '1' year
group by
n_name
order by
revenue desc
Q6: |
-- Q6: Forecasting Revenue Change Query
select
sum(l_extendedprice*l_discount) as revenue
from
lineitem where
l_shipdate >= date '1994-01-01'
and l_shipdate < date '1994-01-01' + interval '1' year
and l_discount between 0.06 - 0.01 and 0.06 + 0.01
and l_quantity < 24
Q7: |
-- Q7: Volume Shipping Query
select
supp_nation,
cust_nation,
l_year,
sum(volume) as revenue
from ( select (^)
n1.n_name as supp_nation,
n2.n_name as cust_nation,
extract(year from l_shipdate) as l_year,
l_extendedprice * (1 - l_discount) as volume
from
supplier,
lineitem,
orders,
customer,
nation n1,


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
nation n2
where
s_suppkey = l_suppkey
and o_orderkey = l_orderkey
and c_custkey = o_custkey
and s_nationkey = n1.n_nationkey
and c_nationkey = n2.n_nationkey
```
and ( (n1.n_name = 'FRANCE' and n2.n_name = 'GERMANY') (^)
or (n1.n_name = 'GERMANY' and n2.n_name = 'FRANCE')
)
and l_shipdate between date '1995-01-01' and date '1996-12-31'
) as shipping
group by
supp_nation,
cust_nation,
l_year
order by
supp_nation,
cust_nation,
l_year
Q8: |
-- Q8: National Market Share Query
select
o_year,
sum(case
when nation = 'BRAZIL'
then volume
else 0
end) / sum(volume) as mkt_share
from (
select
extract(year from o_orderdate) as o_year, l_extendedprice * (1-l_discount) as volume, (^)
n2.n_name as nation
from
part,
supplier,
lineitem,
orders,
customer,
nation n1,
nation n2,
region
where
p_partkey = l_partkey and s_suppkey = l_suppkey (^)
and l_orderkey = o_orderkey
and o_custkey = c_custkey
and c_nationkey = n1.n_nationkey
and n1.n_regionkey = r_regionkey
and r_name = 'AMERICA'
and s_nationkey = n2.n_nationkey
and o_orderdate between date '1995-01-01' and date '1996-12-31'
and p_type = 'ECONOMY ANODIZED STEEL'
) as all_nations
group by
o_year
order by o_year (^)
Q9: |
-- Q9: Product Type Profit Measure Query
select
nation,
o_year,
sum(amount) as sum_profit
from (
select
n_name as nation,
extract(year from o_orderdate) as o_year,


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
l_extendedprice * (1 - l_discount) - ps_supplycost * l_quantity as amount
from
part,
supplier,
lineitem,
partsupp,
orders,
nation where
s_suppkey = l_suppkey
and ps_suppkey = l_suppkey
and ps_partkey = l_partkey
and p_partkey = l_partkey
and o_orderkey = l_orderkey
and s_nationkey = n_nationkey
and p_name like '%green%'
) as profit
group by
nation,
o_year
order by
nation,
o_year desc
Q10: |
-- Q10: Returned Item Reporting Query
select
c_custkey,
c_name,
sum(l_extendedprice * (1 - l_discount)) as revenue,
c_acctbal,
n_name,
c_address,
```
c_phone, c_comment (^)
from
customer,
orders,
lineitem,
nation
where
c_custkey = o_custkey
and l_orderkey = o_orderkey
and o_orderdate >= date '1993-10-01'
and o_orderdate < date '1993-10-01' + interval '3' month
and l_returnflag = 'R'
and c_nationkey = n_nationkey group by
c_custkey,
c_name,
c_acctbal,
c_phone,
n_name,
c_address,
c_comment
order by
revenue desc
limit 20
# Q11中{fraction:.10f}需要根据SF进行替换，替换值为0.0001/SF，详情请查阅TPC-H官方文档
(^) -- Q11: Important Stock Identification QueryQ11: | (^)
select
ps_partkey,
sum(ps_supplycost * ps_availqty) as value
from
partsupp,
supplier,
nation
where
ps_suppkey = s_suppkey
and s_nationkey = n_nationkey


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
and n_name = 'GERMANY'
group by
ps_partkey having
sum(ps_supplycost * ps_availqty) > (
select
sum(ps_supplycost * ps_availqty) * {fraction:.10f}
from
```
partsupp, supplier, (^)
nation
where
ps_suppkey = s_suppkey
and s_nationkey = n_nationkey
and n_name = 'GERMANY'
)
order by
value desc
Q12: |
-- Q12: Shipping Modes and Order Priority Query
select
l_shipmode,
sum(case
when o_orderpriority = '1-URGENT'
or o_orderpriority = '2-HIGH'
then 1
else 0
end) as high_line_count,
sum(case
when o_orderpriority <> '1-URGENT'
and o_orderpriority <> '2-HIGH'
then 1
else 0
end) as low_line_count from
orders,
lineitem
where
o_orderkey = l_orderkey
and l_shipmode in ('MAIL', 'SHIP')
and l_commitdate < l_receiptdate
and l_shipdate < l_commitdate
and l_receiptdate >= date '1994-01-01'
and l_receiptdate < date '1994-01-01' + interval '1' year
group by
l_shipmode
order by l_shipmode (^)
# Q13与TPC-H原生SQL有所区别，请以本文为准
Q13: |
-- Q13: Customer Distribution Query
select
c_count,
count(*) as custdist
from (
select
c_custkey,
count(o_orderkey) as c_count
from
customer left outer join orders on c_custkey = o_custkey
and o_comment not like '%special%requests%'
group by
c_custkey
)as c_orders
group by
c_count
order by
custdist desc,
c_count desc
Q14: |


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
-- Q14: Promotion Effect Query
select
100.00 * sum(case
when p_type like 'PROMO%'
then l_extendedprice*(1-l_discount)
else 0
end) / sum(l_extendedprice * (1 - l_discount)) as promo_revenue
```
from lineitem, (^)
part
where
l_partkey = p_partkey
and l_shipdate >= date '1995-09-01'
and l_shipdate < date '1995-09-01' + interval '1' month
Q15: |
-- Q15: Top Supplier Query
select
s_suppkey,
s_name,
s_address,
s_phone,
total_revenue
from
supplier,
revenue0
where
s_suppkey = supplier_no
and total_revenue = (
select
max(total_revenue)
from
revenue0
) order by (^)
s_suppkey
Q16: |
-- Q16: Parts/Supplier Relationship Query
select
p_brand,
p_type,
p_size,
count(distinct ps_suppkey) as supplier_cnt
from
partsupp,
part
where p_partkey = ps_partkey (^)
and p_brand <> 'Brand#45'
and p_type not like 'MEDIUM POLISHED%'
and p_size in (49, 14, 23, 45, 19, 3, 36, 9)
and ps_suppkey not in (
select
s_suppkey
from
supplier
where
s_comment like '%Customer%Complaints%'
)
group by p_brand, (^)
p_type,
p_size
order by
supplier_cnt desc,
p_brand,
p_type,
p_size
Q17: |
-- Q17: Small-Quantity-Order Revenue Query
select


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
sum(l_extendedprice) / 7.0 as avg_yearly
from
lineitem,
part
where
p_partkey = l_partkey
and p_brand = 'Brand#23'
and p_container = 'MED BOX' and l_quantity < (
select
0.2 * avg(l_quantity)
from
lineitem
where
l_partkey = p_partkey
)
Q18: |
-- Q18: Large Volume Customer Query
select
c_name,
c_custkey,
o_orderkey,
o_orderdate,
o_totalprice,
sum(l_quantity)
from
customer,
orders,
lineitem
where
o_orderkey in (
select
l_orderkey from
lineitem
group by
l_orderkey having
sum(l_quantity) > 300
)
and c_custkey = o_custkey
and o_orderkey = l_orderkey
group by
c_name,
c_custkey,
o_orderkey,
```
o_orderdate, o_totalprice (^)
order by
o_totalprice desc,
o_orderdate
limit 100
Q19: |
-- Q19: Discounted Revenue Query
select
sum(l_extendedprice * (1 - l_discount)) as revenue
from
lineitem,
part
where ( (^)
p_partkey = l_partkey
and p_brand = 'Brand#12'
and p_container in ('SM CASE', 'SM BOX', 'SM PACK', 'SM PKG')
and l_quantity >= 1 and l_quantity <= 1 + 10
and p_size between 1 and 5
and l_shipmode in ('AIR', 'AIR REG')
and l_shipinstruct = 'DELIVER IN PERSON'
)
or
(


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
p_partkey = l_partkey
and p_brand = 'Brand#23'
and p_container in ('MED BAG', 'MED BOX', 'MED PKG', 'MED PACK')
and l_quantity >= 10 and l_quantity <= 10 + 10
and p_size between 1 and 10
and l_shipmode in ('AIR', 'AIR REG')
and l_shipinstruct = 'DELIVER IN PERSON'
```
) or (^)
(
p_partkey = l_partkey
and p_brand = 'Brand#34'
and p_container in ('LG CASE', 'LG BOX', 'LG PACK', 'LG PKG')
and l_quantity >= 20 and l_quantity <= 20 + 10
and p_size between 1 and 15
and l_shipmode in ('AIR', 'AIR REG')
and l_shipinstruct = 'DELIVER IN PERSON'
)
Q20: |
-- Q20: Potential Part Promotion Query
select
s_name,
s_address
from
supplier,
nation
where
s_suppkey in (
select
ps_suppkey
from
partsupp
where ps_partkey in ( (^)
select
p_partkey
from
part
where
p_name like 'forest%'
)
and ps_availqty > (
select
0.5 * sum(l_quantity)
from
lineitem where
l_partkey = ps_partkey
and l_suppkey = ps_suppkey
and l_shipdate >= date '1994-01-01'
and l_shipdate < date '1994-01-01' + interval '1' year
)
)
and s_nationkey = n_nationkey
and n_name = 'CANADA'
order by
s_name
Q21: |
-- Q21: Suppliers Who Kept Orders Waiting Query select
s_name,
count(*) as numwait
from
supplier,
lineitem l1,
orders,
nation
where
s_suppkey = l1.l_suppkey
and o_orderkey = l1.l_orderkey


###### 深圳计算科学研究院 深圳崖山科技有限公司

#### 导入数据

```
and o_orderstatus = 'F'
and l1.l_receiptdate > l1.l_commitdate
and exists (
select
*
from
lineitem l2
```
where l2.l_orderkey = l1.l_orderkey (^)
and l2.l_suppkey <> l1.l_suppkey
)
and not exists (
select
*
from
lineitem l3
where
l3.l_orderkey = l1.l_orderkey
and l3.l_suppkey <> l1.l_suppkey
and l3.l_receiptdate > l3.l_commitdate
)
and s_nationkey = n_nationkey
and n_name = 'SAUDI ARABIA'
group by
s_name
order by
numwait desc,
s_name
limit 100
Q22: |
-- Q22: Global Sales Opportunity Query
select
cntrycode, count(*) as numcust, (^)
sum(c_acctbal) as totacctbal
from (
select
substring(c_phone from 1 for 2) as cntrycode,
c_acctbal
from
customer
where
substring(c_phone from 1 for 2) in
('13','31','23','29','30','18','17')
and c_acctbal > (
select avg(c_acctbal) (^)
from
customer
where
c_acctbal > 0.00
and substring (c_phone from 1 for 2) in
('13','31','23','29','30','18','17')
)
and not exists (
select
*
from
orders where
o_custkey = c_custkey
)
) custsale
group by
cntrycode
order by
cntrycode


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
进入数据库安装目录，执行如下命令导入数据。
Note :
导数阶段建议调整参数DATA_BUFFER_SIZE=30G、VM_BUFFER_SIZE=20G、SCOL_DATA_BUFFER_SIZE=1G，参数随测试使用的数据量增大适
当增加。
```
#### 创建 Q15 所需视图

#### 将热数据转为冷数据

```
$ vi load_data.sh
#! /bin/bash
YASDB_HOME='./' # 替换为实际安装数据库的路径
CN_NODE_IP='127.0.0.1' # 替换为实际的CN节点IP地址
DATA_PATH='/data/tpch/SF100' # 替换为实际的TPC-H数据存放路径
$YASDB_HOME/bin/yasldr regress/regress@$CN_NODE_IP:1688 control_text="'load data
options(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '$DATA_PATH/region.tbl' without
embedded fields terminated by '|' append into table region(R_REGIONKEY, R_NAME, R_COMMENT)'" senders= 9
$YASDB_HOME/bin/yasldr regress/regress@$CN_NODE_IP:1688 control_text="'load data
options(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '$DATA_PATH/nation.tbl' without
embedded fields terminated by '|' append into table nation(N_NATIONKEY, N_NAME, N_REGIONKEY, N_COMMENT)'" senders= 9
$YASDB_HOME/bin/yasldr regress/regress@$CN_NODE_IP:1688 control_text="'load data
options(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '$DATA_PATH/supplier.tbl'
without embedded fields terminated by '|' append into table supplier(S_SUPPKEY, S_NAME, S_ADDRESS, S_NATIONKEY, S_PHONE,
S_ACCTBAL, S_COMMENT)'" senders= 9
$YASDB_HOME/bin/yasldr regress/regress@$CN_NODE_IP:1688 control_text="'load data
options(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '$DATA_PATH/part.tbl' without
embedded fields terminated by '|' append into table part(P_PARTKEY, P_NAME, P_MFGR, P_BRAND, P_TYPE, P_SIZE, P_CONTAINER,
```
P_RETAILPRICE, P_COMMENT)'"$YASDB_HOME/bin/yasldr regress/regress@ senders= (^9) $CN_NODE_IP:1688 control_text="'load data
options(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '$DATA_PATH/partsupp.tbl'
without embedded fields terminated by '|' append into table partsupp(PS_PARTKEY, PS_SUPPKEY, PS_AVAILQTY, PS_SUPPLYCOST,
PS_COMMENT)'" senders= 9
$YASDB_HOME/bin/yasldr regress/regress@$CN_NODE_IP:1688 control_text="'load data
options(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '$DATA_PATH/customer.tbl'
without embedded fields terminated by '|' append into table customer(C_CUSTKEY, C_NAME, C_ADDRESS, C_NATIONKEY, C_PHONE,
C_ACCTBAL, C_MKTSEGMENT, C_COMMENT)'" senders= 9
$YASDB_HOME/bin/yasldr regress/regress@$CN_NODE_IP:1688 control_text="'load data
options(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '$DATA_PATH/orders.tbl' without
embedded fields terminated by '|' append into table orders(O_ORDERKEY, O_CUSTKEY, O_ORDERSTATUS, O_TOTALPRICE, O_ORDERDATE,
O_ORDERPRIORITY, O_CLERK, O_SHIPPRIORITY, O_COMMENT)'" senders= 9
$YASDB_HOMEoptions(DEGREE_OF_PARALLELISM=16,TRIM=NOTRIM,ENABLE_BULK=FALSE,DECODER_THREAD_TIMES=7) infile '/bin/yasldr regress/regress@$CN_NODE_IP:1688 control_text="'load data (^) $DATA_PATH/lineitem.tbl'
without embedded fields terminated by '|' append into table lineitem(L_ORDERKEY, L_PARTKEY, L_SUPPKEY, L_LINENUMBER,
L_QUANTITY, L_EXTENDEDPRICE, L_DISCOUNT, L_TAX, L_RETURNFLAG, L_LINESTATUS, L_SHIPDATE, L_COMMITDATE, L_RECEIPTDATE,
L_SHIPINSTRUCT, L_SHIPMODE, L_COMMENT)'" senders= 9
$ chmod +x load_data.sh
$ ./load_data.sh
CREATE VIEW revenue0 (supplier_no, total_revenue) AS
SELECT
l_suppkey,
SUM(l_extendedprice * ( 1 - l_discount))
FROM
lineitem
WHERE
l_shipdate >= DATE '1996-01-01'
(^) GROUPAND BY l_shipdate < DATE '1996-01-01' + INTERVAL '3' month
l_suppkey;


###### 深圳计算科学研究院 深圳崖山科技有限公司

#### 收集统计信息

```
执行如下语句开始收集统计信息。
Note :
收集统计信息阶段建议调整参数DATA_BUFFER_SIZE=30G、VM_BUFFER_SIZE=20G、SCOL_DATA_BUFFER_SIZE=1G、
COLUMNAR_VM_BUFFER_SIZE=10G，参数随测试使用的数据量增大适当增加。
```
#### 执行查询

```
ALTER SYSTEM SET _enable_alter_slice=true scope=memory;
ALTER SYSTEM SET DATA_TRANSFORMER_ENABLED=true scope=memory;
ALTER TABLE region ALTER SLICE ALL STABLE;
ALTER TABLE nation ALTER SLICE ALL STABLE;
ALTER TABLE supplier ALTER SLICE ALL STABLE;
ALTER TABLE part ALTER SLICE ALL STABLE;
```
ALTERALTER TABLETABLE partsupp customer ALTERALTER SLICE SLICE ALLALL STABLE STABLE;; (^)
ALTER TABLE orders ALTER SLICE ALL STABLE;
ALTER TABLE lineitem ALTER SLICE ALL STABLE;
ALTER TABLE region ALTER SLICE ALL COMPACT;
ALTER TABLE nation ALTER SLICE ALL COMPACT;
ALTER TABLE supplier ALTER SLICE ALL COMPACT;
ALTER TABLE part ALTER SLICE ALL COMPACT;
ALTER TABLE partsupp ALTER SLICE ALL COMPACT;
ALTER TABLE customer ALTER SLICE ALL COMPACT;
ALTER TABLE orders ALTER SLICE ALL COMPACT;
ALTER TABLE lineitem ALTER SLICE ALL COMPACT;
ALTER TABLE region ALTER SLICE ALL CLEAN;
ALTER TABLE nation ALTER SLICE ALL CLEAN;
ALTER TABLE supplier ALTER SLICE ALL CLEAN;
ALTER TABLE part ALTER SLICE ALL CLEAN;
ALTER TABLE partsupp ALTER SLICE ALL CLEAN;
ALTER TABLE customer ALTER SLICE ALL CLEAN;
ALTER TABLE orders ALTER SLICE ALL CLEAN;
ALTER TABLE lineitem ALTER SLICE ALL CLEAN;
ANALYZE SCHEMA regress ESTIMATE_PERCENT 1 PARALLEL_DEGREE 16 METHOD_OPTION 'FOR ALL COLUMNS SIZE AUTO' GRANULARITY 'GLOBAL';
$ vi tpch.py
#! /usr/bin/python
import time
import jpype
import yaml
jpype.addClassPath('./lib/yasdb-jdbc-{版本号}.jar') # 替换路径为真实的JDBC驱动路径
jpype.startJVM(jpype.getDefaultJVMPath())
url = 'jdbc:yasdb://127.0.0.1:1688/test' # 127.0.0.1替换成实际CN节点的IP地址
conn = jpype.java.sql.DriverManager.getConnection(url, 'regress', 'regress')
with open('./tpch.yml', "r") as f:
queries = yaml.safe_load(f)
totalCost = 0
print("----------TPC-H Benchmark Begin----------")
for name, sql in queries["queries"].items():
if name == "Q11":
sql = sql.format(fraction=0.0001 / 100 ) # 这里的 100 需根据实际测试的数据规模调整
status = 'SUCCESS'
start = time.time()
try:
prep_stmt = conn.prepareStatement(sql)


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
等待测试结束，查看测试结果。
```
```
prep_stmt.execute()
except Exception:
status = 'ERROR'
finally:
elapsed = time.time() - start
print(f"{name:>3} exec status：{status:>7}，elapsed：{elapsed:>7.3f}s")
totalCost += elapsed
printjpype.shutdownJVM("----------TPC-H Benchmark END({:.3f}s)------------"() .format(totalCost))
```
```
$ python tpch.py
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## YashanDB 标准版与企业版功能差异

#### 版本描述

```
标准版：YashanDB面向小规模用户推出的商业版本，价格适中，除多模数据类型、高级安全能力等企业级功能外，该版本包含YashanDB数据库所有
基础核心能力，支持单机主备部署形态，配套完整数据迁移和监控运维工具，可以为政府或中小企业提供支撑其业务所需的基本能力。
企业版：YashanDB面向大规模用户推出的商业版本，该版本包含YashanDB数据库完整核心能力，支持PB级海量数据存储和大量的并发用户，支持多
模数据类型、高级安全能力，支持单机（主备）部署、共享集群部署以及分布式部署形态，配套完整数据迁移和监控运维工具，可以满足支撑各类企
业应用。
```
#### 版本差异

```
类目 功能项 标准版 企业版
产品形态 版本名称 Standard Edition Enterprise Edition
```
```
产品形态 部署形态 单机部署
```
```
单机部署
共享集群部署
分布式部署
物理规格 最大连接数 1024 无限制
物理规格 最大存储容量 无限制 无限制
```
```
物理规格 硬件平台
```
```
x86
ARM
龙芯
```
```
x86
ARM
龙芯
物理规格 使用时间 无限制 无限制
基础功能 单表最大行数 无限制 无限制
基础功能 单表最大列数 4096 4096
基础功能 基本数据类型 支持 支持
基础功能 内置高级包 支持 支持
基础功能 并行查询 支持 支持
基础功能 闪回查询 支持 支持
基础功能 存储过程调试功能 支持 支持
基础功能 外部函数扩展 支持 支持
基础功能 DBLINK 支持 支持
基础功能 物化视图 支持 支持
基础功能 诊断包 支持 支持
核心组件 驱动 支持 支持
安全能力 通讯加密 支持 支持
安全能力 存储加密 支持 支持
安全能力 三权分立 支持 支持
安全能力 审计 支持 支持
高阶功能 分区表 一级分区 二级分区
高阶功能 列存储 支持 支持
高阶功能 JSON类型 不支持 支持
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
类目 功能项 标准版 企业版
高阶功能 GIS 不支持 支持
配套工具 数据迁移工具 支持 支持
配套工具 运维监控工具 支持 支持
配套工具 开发者工具 支持 支持
产品服务 软件升级 支持 支持
产品服务 专家服务 支持 支持
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

## Release Notes

### 版本信息

```
版本号：v23.4.1.100
发版日期： 2025 年 5 月 8 日
```
### 版本定位

```
YashanDB v23.4版本定位为长期服务版本（LTS），新版本围绕金融核心业务系统建设产品能力，目标对等 1 比 1 平替Oracle，全面增强Oracle兼容性，
构建金融行业高可用能力、实现库级闪回秒级恢复等竞争力。同时发布mysql语法模式功能，提升性能，配套完善的数据迁移、监控运维和开发者工具，
支持市场项目规模销售和批量复制。
```
### 特性更新

```
YashanDB v23.4版本继承自YashanDB v23.3.1版本全部发布的功能，同时新增了以下关键特性：
```
#### SQL 引擎增强

##### SQL 功能增强

```
支持TIMESTAMP时区类型：新增TIMESTAMP时区数据类型，存储日期和时间信息、时区偏移量或时区名称，用于需要准确的时间、支持跨国业务的
场景。
支持JSON_VALUE函数：新增JSON_VALUE函数，实现从JSON数据中提取单个标量值（如字符串、数字、布尔值）。
INSTR函数支持CLOB类型：支持在CLOB类型数据中搜索指定的字符串，并返回其位置，提升处理大文本数据提高效率。
扩展投影列字段的别名长度：在v23.4版本前，投影列字段和别名长度最大为20B，在使用*查询和未指定别名的场景下字段名较长容易被截断和重名造
成不好体验。该版本针对此场景做了能力增强，扩展投影列和别名长度至64B，满足更大的使用范围，提高易用性。
```
##### PL 功能增强

```
PACKAGE支持PROCEDURE和FUNCTION重载功能：支持存储过程和函数的重载功能，可以在同一个包内定义多个同名的存储过程或函数，但它们
的参数列表需要数量、数据类型或者顺序不同，YashanDB会依据调用时传入的实际参数，来判断具体调用哪个重载版本。主要用于相似功能不同参
数、开发兼容不同版本的场景诉求。提高代码可读性与可维护性提升，及增强开发人员使用的灵活性。
支持PIPELINED管道函数：通过PIPE ROW函数将数据一行一行地返回给调用者，不需要等待所有数据处理完才返回结果，在SQL调用所需的行数少
且计算量大场景下可以提高性能，同时简化代码。
游标支持FOR UPDATE功能：PL新增支持游标for update功能，用于对查询结果集中的行进行锁定，阻止其他事务对这些行进行修改或删除操作，配
合CURRENT OF子句，在游标操作里自动精确地定位到当前游标所在行，从而对该行数据进行更新或删除操作。主要解决数据一致性要求高和逐行处
理的场景。
PL新增MOD取模用法：Oracle的PL/SQL语法中不支持使用%操作符作取模操作，只能使用MOD语法。YashanDB新增PL/SQL的MOD语法，此语法实
现的是取模操作，同操作符“%”的作用一致，丰富使用场景，增加兼容性。
PL解析机制优化：在该版本之前的静态SQL和显式游标等功能中，PL语句与SQL语句混合使用时，数据库会优先进行PL语句块变量匹配，该机制会导
致迁移Oracle代码时产生大量兼容性相关的修改工作量。为了提升Oracle语义兼容性，当前版本通过去除预编译动作，优先匹配SQL语句中的表列资
源，其次再进行匹配PL语句块变量。
```
##### SQL 性能优化

```
优化规则增强
常量参数化与基于参数的计划管理：针对统计信息波动或环境变化易引发的执行计划劣化、性能抖动甚至宕机风险，该版本智能执行计划管理采用
父子游标两级缓存机制，通过SQL标准化与常量参数化合并相似查询，有效降低缓存池的内存损耗。在常见拼接SQL场景下，软解析内存消耗可有
效降低80%以上。
支持COUNT(*)算子下推：SQL中投影列为count(*)、count(1)常量计算单表全量数据的场景下，支持count流程直接下推到存储层，减少SQL层与存
储层交互次数，节省交互时间，提升性能。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
CTE查询重写支持计划共享与物化区共享：在使用CTE查询且存在多个子查询的场景下，通过优化器按照物化/代价计算/内联的方式优化CTE，实现
子查询物化共享，减少CTE执行次数，加快查询效率。随着CTE的使用变多，性能的优化效果更为明显。
多列IN支持走索引优化：在多列在IN的场景下，例如(col1, col2, col3) IN ((value1_1, value1_2, value1_3), (value2_1, value2_2, value2_3))，
YashanDB针对此场景对优化器进行优化，使用col1, col2, col3上的组合索引、单列索引生成计划进行评估，按照选择最低cost的索引和扫描方式执
行，避免选择代价大的索引和扫描方式的情况，提升执行效率和性能。
子查询中排序算子优化：通过对优化器完善，在父查询投影列只有汇聚和常量，且FROM后子查询中有排序算子的场景下，实现对子查询的排序算
子的消除，即不关心子查询的返回结果是否有序，如SELECT COUNT(*) FROM (SELECT * FROM T1 ORDER BY 1)不需要对T1表进行排序直接
COUNT表T1，减少排序的动作，提升了执行效率。
LIKE变量||'%'时支持索引：当Filter为like变量||'%'时，在23.4版本之前执行计划为全表扫，在表数据量大且返回数据少的场景下耗时长，该
版本实现在此场景下支持走Filter字段的索引，加快查询速度。
执行性能提升
汇聚函数带DISTINCT的算法优化：针对聚集函数带DISTINCT的场景，增加HASH去重算法，在选择无需排序的语句场景下，随着数据的重复度提
升，相较于历史版本性能提升明显。
多行子查询执行优化：在结果集为多行的子查询场景下，不需要重复执行，通过子查询缓存复用机制，有效降低大数据量子查询的性能开销，相较
于历史版本，复杂子查询性能提升 3 倍以上，充分满足高并发OLTP与复杂分析混合负载需求。
支持SORT GROUPBY STOPKEY算子：当SQL外层有rownum或者limit语法只返回少数行时，且内层存在大表的group的时候，支持SORT GROUP
BY STOPKEY算子可以加速SQL的执行效率。
针对PL语言的FORALL INSERT数据处理中，通过FORALL批量插入技术，使标量类型写入性能普遍提升85%，部分类型如CLOB类型写入速度大幅
优于Oracle；UDT类型的写入速度提升在30%以上。
DETERMINISTIC函数调用优化：该版本通过对调用的函数是deterministic的进行执行调整，输入为常量情况下，直接改写成常量；输入为绑定参
数，参与优化和执行的动态常量特性缓存；使同场景性能提升明显，优于Oracle。
带UDT的全表扫描支持并行：该版本新增支持带Geometry列的表在全表扫的时候支持并行，且部分GIS函数在计算中也支持并行能力，提高了GIS
函数和对象的处理性能。
DBLink性能优化
DBLink的交互时延、交互次数、查询投影优化、内部协议编解码和OCI驱动解码优化：通过对交互时延、交互次数和内部协议编解码优化，使单次
各个步骤的消耗和时长减少，最终达到性能优化的效果，通过此次优化，使YashanDB使用DBLink访问Oracle在指定的场景下性能提升数倍。
DBlink支持收集及使用统计信息：YashanDB通过DBLink访问Oracle时，支持对远端表的统计信息进行查询和汇总，优化器从而在SQL中存在
DBLink时可以更准确地生成执行计划，提升复杂查询中存在DBLink场景的查询性能。
```
##### GIS 能力增强

```
在v23.4版本之前，CREATE TABLE语法创建GIS表时只支持ST_GEOMETRY类型，该版本实现了Geometry列支持声明指定子类型、SRID，例如
geometry(point, 4326)，兼容了PostGIS的语法和行为，同时支持Geometry_Columns视图，可查询Geometry数据的子类型、SRID、维度，提高
易用性。
支持ST_MakeValid函数，在不丢弃顶点的情况下，把无效的Geometry对象转换成有效的Geometry对象。
支持ST_CollectionExtract函数，从一个（Multi）Geometry对象中，找出并返回指定类型的Geometry对象。
支持ST_Centroid函数， 用于计算几何图形的质心（几何中心）。
支持ST_PointOnSurface函数，返回一个位于geometry实例内部的任意点。
支持ST_Transform函数，将传入的Geometry对象的坐标参考系转换成指定的坐标参考系并返回。
```
#### 存储引擎增强

##### 存储能力增强

```
支持全库闪回：基于闪回日志快照点（Marker）技术，通过FLASHBACK DATABASE语法子句，实现全库秒级回滚至任意时间点。主要用于系统升级/
迁移回滚、安全事件响应、测试环境重置、数据审计和分析、数据清洗回退、逻辑错误修复等故障场景。
Ystream解析DDL支持附加主键信息：YashanDB将主键信息组装成JSON字符串，写入redo日志，由YStream解析成字符串，实现使用Ystream时可以
解析出主键信息，丰富使用场景。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
ALTER TABLE支持UPDATE GLOBAL INDEXES功能：ALTER TABLE通过新增UPDATE GLOBAL INDEXES语法子句，实现了在对分区表进行
DROP、SPLIT、TRUNCATE等分区操作时，同时更新相关的全局索引，避免了因分区操作导致的索引失效问题，保证了查询的正确性和性能。还简
化了管理，无需手动重建索引。
支持合并分区：通过新增MERGE PARTITIONS语法子句，支持将相邻的分区合并为一个分区，从而减少分区的数量。合并分区有助于对数据进行整
合，特别是当数据不再需要按照原有的分区规则进行存储时。合并分区时，删除与所选分区相对应的本地索引分区，并将与合并分区相对应的本地索
引分区标记为UNUSABLE。可以手动重建索引维护或者使用UPDATE GLOBAL INDEXES子句更新全局索引。可以有效减少分区数量可以降低分区表
的管理复杂度，减少维护工作量。
支持表空间备份恢复：新增BACKUP/RESTORE TABLESPACE语法子句，实现对表空间备份进行备份恢复能力，可以有效减少备份恢复时间，提供
灵活的备份恢复策略，同时支持DBA_DATAFILE_BACKUPSET视图查看表空间备份集信息。该版本备份恢复表空间仅支持单机部署上。
该版本默认部署时打开回收，即回收站开关RECYCLEBIN_ENABLED默认为ON，用于误操作的恢复。
统计信息能力增强
支持基础统计信息的实时收集（DML实时更新统计信息）：新增全局参数OPTIMIZER_REAL_TIME_STATISTICS控制是否开启实时收集统计信
息，默认为FALSE关闭，设置为TRUE开启后，对在表进行DML/DDL的时实时收集统计信息，对基础的统计信息（例如row count、min/max等值）
可以在SQL执行过程中实现实时修改。可以提升统计信息的准确度，也可以让优化器避免使用默认统计信息，造成性能损耗。
DBMS_STATS高级包gather_schema_stats函数增加options参数：在23.4版本之前，使用DBMS_STATS.GATHER_SCHEMA_STATS按SCHEMA
收集统计信息时会全部重新对象收集一次，在部分对象已存在统计信息的场景下效率低。该版本通过新增设置options能力，支持了按照SCHEMA维
度时也可按需收集表维度的诉求，默认时只收集SCHEMA下满足条件的表。提高了统计信息的收集效率。
支持默认采样比例auto_sample_size：通过新增DBMS_STAT.AUTO_SAMPLE_SIZE常量来指定自动采样率，在使用
GATHER_DATABASE_STATS、GATHER_INDEX_STATS、GATHER_SCHEMA_STATS和GATHER_DATABASE_STATS收集统计信息时生效，
此参数不生效于动态采样。增加了Oracle的兼容性，提升客户的易用性。
统计信息收集进度展示：新增视图DBA_OPTSTAT_OPERATIONS，记录执行gather_table_stats, gather_index_stats,、gather_schema_stats、
gather_database_stats存储过程的操作记录，包含使用DBMS_STATS高级包和数据库级别执行的统计操作的历史记录；新增视图
DBA_OPTSTAT_OPERATION_TASKS，记录统计信息父操作中要处理的目标对象和信息，该视图会将DBA_OPTSTAT_OPERATIONS中的父操作
分解为对统计信息收集对象的子操作，对象包含表、索引或分区。在手动触发和定时任务等触发统计信息收集任务时，通过这 2 个视图查询任务的进
度及状态、起始时间等信息。
```
##### 性能优化

```
支持SESSION级BLOCK CACHE能力：在访问全局资源时，需要通过加锁控制并发。对于读而言虽然请求的是共享锁，但是加共享锁通常都会有Spin
Lock的竞争，在并发较高的情况下，Spin Lock冲突会很大。通过增加SESSION级别的Block Lock Cache使资源本地化，降低全局资源锁冲突，提高
了并发读的性能。同时增加多个系统参数控制每个SESSION Block Lock Cache的数量，增加动态视图V$SESSION_LOCK_CACHES查看每个
SESSION的Cache情况，使用户可以在不同高并发场景下灵活性地进行调整。
INSERT INTO SELECT场景的IO性能优化：在INSERT INTO SELECT操作且插入的表具有多索引的场景下，YashanDB在该版本对索引事务内部机制
进行了优化，减少小数据量时REDO的刷盘IO次数和redo落盘等待时间，提升该场景下的插入性能且持平Oracle。
支持设置表并行：新增支持CREATE/ALTER TABLE配置表的并行度，表设置表并行度后，查询全表扫描时默认使用并行扫描。同时可以使用
DBA_TABLES视图的degree字段可以查看当前的并行情况。在处理大规模数据查询的场景可以有效提升查询性能。
创建/检查约束支持并行：在大数据量场景下，表存在约束时，检查关联条件的完整性时都会进行大量的数据扫描，造成极大的耗时。因此在创建/检查
约束全表扫描的场景下，增加并行处理，提升处理性能。
```
##### 高可用增强

```
一主一备自定义配置条件后自动切换能力增强：在单机一主一备部署形态下，可以使用基于yasom进程的仲裁选主功能保障业务连续性，该版本支持
置条件故障切换，通过新增配置参数FAILOVER_HEALTH_CONDITION和FAILOVER_ERROR_CONDITION，分别代表在特定异常时（包含归档磁盘
满，redo刷盘失败，dbwr刷脏页失败等）和在主库向客户端抛出某些错误码时的切换条件值，在触发条件后主库立即关闭实现自动切换。
```
#### 共享集群产品能力

##### 多地多中心高可用能力

```
主备集群支持在主集群所有实例故障场景下自动failover，提升产品故障感知和切换能力，降低运维复杂性和提升产品易用性。
备集群所有实例支持可读，提升了在读写分离场景下，备集群的读性能和稳定性。
容灾集群支持集群单实例在本地磁盘部署，降低容灾场景的部署和管理成本。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
共享集群IO脑裂保护增强：在v23.3版本，共享集群提供了基于在途IO保护算法的IO Fencing方法，但是在途IO保护算法会一定程度上延长集群RTO，
当存储设备IO性能差或IO卡顿现象频发时，还需要用户调大共享集群的DISK_HB_KEEP_ALIVE参数配置值以确保算法的可靠性。为了更好的解决此
场景问题，在v23.4版本提供了基于硬件IO保护方式，即就是基于SCSI等协议持久预留的IO Fencing方法利用共享存储设备普遍支持的SCSI持久预留
命令为共享集群构建IO栅栏，具有绝对的安全性和更优的RTO指标。
```
##### 功能优化完善

```
新增支持在线扩容：在实际业务中，随着业务增长现有数据库节点规格性能可能不满足业务现状，需要进行节点扩容来确保业务运行。在v23.4版本，
共享集群支持在线扩容数据库节点，解决随着实际业务增长需要在线扩展数据库节点业务诉求。
在线DDL能力增强：在v23.4版本共享集群提供在线create/rebuild索引以及在线shrink table等能力，实现降低运维复杂性和提升产品易用性。
新增支持GIS能力：在v23.4版本新增GIS能力，通过内置地理空间数据处理引擎，支持复杂空间查询，满足智慧城市、物流追踪等场景需求。
新增备集群支持备份：在v23.4之前版本，共享集群备份仅支持在主集群备份，但是在实际业务中主集群业务会比较大，执行备份会影响主集群运行性
能和稳定性，因此在v23.4版本支持备集群支持备份，来降低在主集群备份带来影响，提升系统稳定性。
支持NOLOGGING能力：支持在yasldr、CREATE/ALTER INDEX、LOAD DATA等使用场景可以显示指定启用NOLOGGING能力，因为过程中不产生
REDO日志，从而提升在相关场景性能。
YFS能力提升：当前YFS管理上不能将数据分散在多个磁盘上，会存在IO负载不均衡，导致性能较差。在v23.4版本YFS提供了条带化能力，支持按照
不同文件设置条带化来平衡磁盘组中所有磁盘的负载和减少I/O延迟，例如DATA_STRIPING、REDO_STRIPING等参数设置不同数据文件是否启用
YFS条带化功能，提升整体的性能。
```
##### 性能优化

```
国产性能优化：v23.4版本共享集群针对鲲鹏、海光等国产化CPU多NUMA场景下进行了深度优化和探索，在鲲鹏CPU和TPCC场景下测试，计算节点
从1 -> 2 -> 4节点的扩展比达0.8以上。同时提供完整的调优手册方便使用者查阅和参考，具体请查阅TPC-C性能调优。
UNDO数据亲和性（Affinity）：共享集群中每个实例都有自己的UNDO表空间，由于共享集群各个实例需要Cache Fusion来进行一致性访问等，为降
低UNDO管理在Cache Fusion场中的性能损耗，在v23.4版本提供 UNDO 亲和性能力，通过将特定事务或数据操作绑定到特定实例的UNDO表空间，从
而降低跨实例的锁竞争和全局缓存（Global Cache）传输开销，来提升系统整体性能和稳定性。
```
#### 分布式产品能力

##### 列存储能力增强

```
CREATE TABLE AS SELECT支持BULKLOAD：通过在create table as后的select加上/ +bulkload /的hint，实现列存LSC表在使用CTAS语法时使用
bulkload功能，对LSC表进行批量插入，可以减少redo的产生，提高插入速度。
支持OUTLINE LOB字段长度的计算：支持对LSC表的outline lob字段长度进行计算，包含LENGTH、CHAR_LENGTH、LENGTHB等函数对字符和字
节进行长度统计。
INSERT / +BULKLOAD / INTO支持在AUTOCOMMIT ON的情况下执行：当前列存LSC表insert / +bulkload / into要求必须在autocommit off的情况下执
行，需要在脚本或者SQL中手动加上事务提交，不然可能会导致部分未结束事务。该版本通过新增enable_bulkload_auto_commit参数，实现对insert
/ +bulkload / into的自动提交控制，使用户在脚本中使用bulkload特性时更加灵活，代码更简洁。
INSERT INTO SELECT和CTAS支持OUTLINE LOB：INSERT INTO SELECT和CREATE TABLE AS SELECT为日常运维和业务系统中常用操作，在
v23.4版本前，列存不支持使用INSERT INTO SELECT和CREATE TABLE AS SELECT语句对超过 32000 字节的LOB列的操作。在该版本实现了此功
能，满足了数仓场景下的日常运维操作习惯，简化了业务逻辑，提高灵活性。
支持DUAL和列表混合查询（采用列执行引擎）：在v23.4版本之前，不支持DUAL表跟列存表进行JOIN等混合查询，该版本实现了该功能，满足更多
的业务使用场景。
CREATE TABLE AS SELECT支持重试：当前在insert select阶段，如果遇到资源不足，并没有重试，直接报错返回给客户端了。该版本通过对内存配
额、会话资源和相关报错等信息判断，实现对在特定场景下create table as select失败时可以进行自动重试，无需用户干预，最大重试时间为 30 秒，如
在 30 秒重试不成功则失败。
支持SYS_GUID函数：在分布式系统和特定需要保证唯一性的诉求下，实现通过sys_guid函数生成全局唯一标识符，可用于LSC列存生成主键数据和
数据同步等业务场景。
支持TABLESPACE SET的空间回收功能：分布式部署中通过支持ALTER TABLESPACE xxx SHRINK SPACE [KEEP xxx]语法功能，实现
TABLESPACE SET空间压缩和指定压缩后的表空间大小，增加了表空间集的数据膨胀的维护手段，提高了用户的资源利用率。
```
##### 性能优化


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
支持并行创建备库：通过支持BUILD DATABASE [to standby （*）] [PARALLELISM integer]语法功能，实现多线程并发创建备库，同时优化创建备库
时LSC表Slice文件的扫描和读取的性能，大幅度提升了分布式备库构建速度。
```
##### 分布式运维能力增强

```
新增V$XFMR_STAT显示列式存储用于转换生成稳态数据的所有还未执行完毕的xfmr任务信息。
新增与V$XFMR_HIS_STAT视图，显示列式存储用于转换生成稳态数据的所有已经执行完毕的xfmr任务信息。包含展示XFMR任务的资源占用，包括
内存，配额，换入换出，执行时间等信息，可用于分析正在执行中与已执行完毕的历史XFMR任务的行为是否符合预期以及是否影响性能。
PQ_POOL_SIZE内存使用情况可观测性：PQ_POOL是主要用于指定并行线程间数据传递使用的内存池，常用的操作有分布式下跨节点的行列执行、
共享集群下的global view查询。该版本把PQ_POOL_SIZE参数从隐藏参数修改为正式参数，供用户自行配置，同时在多个动态视图新增PQ_POOL相
关的指标，提高了易用性，且方便获取诊断信息。
LSC表支持空间的及时回收与复用：该版本对LSC表的空间管理进行优化，在申请空间时，优先使用可回收利用的空间，减少膨胀；在空间不足时，
后台线程支持强制复用，保证业务的成功；采用更加积极的后台空间回收策略，执行效率更加高效。
```
#### MySQL 兼容

```
具体使用请查阅标注为（mysql模式）的相关文档。
```
##### 新增 MySQL 数据类型

```
支持BINARY数据类型，用于存储二进制字符串，可以存储任意类型的二进制数据，如图像、音频、视频、文件，将数据以字节的形式进行存储，不进
行字符集转换和编码处理，能够准确地保存原始数据的每个字节。可以避免字符集转换和编码。
支持FLOAT和DOUBLE数据类型： 4 字节单精度近似的数值数据，最大 23 位精度； 8 字节双精度近似的数值数据，最大 53 位精度。
支持存储和计算可变长度的字符VARCHAR数据类型，最大支持存储 65535 字节，存储字符个数等于小于 65535 个字符，受字符集影响。
支持无符号数据类型，都表示为正数：
TINYINT UNSIGNED：用于存储非常小的整数，占用 1 个字节。可以表示 0 到 255 之间的整数。
SMALLINT UNSIGNED：存储较小的整数，占用 2 个字节。能表示 0 到 65535 之间的整数。
MEDIUMINT UNSIGNED：存储中等大小的整数，占用 3 个字节。取值范围是 0 到 16777215 。
INTEGER/INT UNSIGNED：存储整数，占用 4 个字节。可以表示 0 到 4294967295 的整数。
BIGINT UNSIGNED：用于存储较大的整数，占用 8 个字节。取值范围是 0 到 18446744073709551615 。
DECIMAL UNSIGNED：用于存储精确的小数，DECIMAL[(M[,D])] [UNSIGNED]中M表示总位数（精度），D表示小数点后的位数（小数位数）。
FLOAT UNSIGNED：用于存储浮点数，占用 4 个字节。可以表示非负的单精度浮点数。
DOUBLE UNSIGNED：用于存储浮点数，占用 8 个字节。表示非负的双精度浮点数。
支持YEAR、DATE和TIME时间数据类型，YEAR类型用于表示 4 个字符的年份值。DATE类型用于表示具有日期部分但没有时间部分的值，以“YYYY-
MM-DD”格式检索和显示DATE值。支持的范围是“1000-01-01”到“9999-12-31”。TIME类型以'hh:mm:ss'格式（或对于较大的小时值，则
为'hhh:mm:ss'格式）检索和显示TIME值。TIME值的范围可以从'-838:59:59'到'838:59:59'。
```
##### 新增 MySQL 内置函数

```
CRC32函数计算循环冗余校验值并返回 32 位无符号值。如果参数为NULL，则结果为NULL。
RAND()函数返回 0 到 1 的随机数值，RAND(N)时返回固定的值。
TRUNCATE函数返回截断为小数点后D位的数字X。如果D为 0 ，则结 果没有小数点或小数部分。D可以为负数，导致值X小数点左边的D位数字变为
零。如果X或D为NULL，则该函数返回NULL。
FROM_BASE64函数把Base64编码的字符串解码为原始的二进制数据。
LCASE函数用于将字符串中的所有字符转换为小写形式。
LOCATE函数可以在一个字符串中查找另一个子字符串的位置。
REVERSE函数将输入字符串的字符顺序进行反转，也就是让字符串逆序排列。
SPACE函数根据输入的参数值，创建一个由相应数量空格字符构成的字符串。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
UCASE函数用于将字符串中的所有字符转换为大写形式。
UNHEX函数将十六进制字符串转换为对应的二进制字符串。
ANY_VALUE函数用于在处理GROUP BY子句时使用的一个工具，主要用于解决SQL标准要求SELECT投影列中的非聚合列必须全部包含在GROUP
BY子句中的问题。
SLEEP函数主要功能是让当前执行的线程暂停指定的时间。
VALUES函数主要在INSERT ... ON DUPLICATE KEY UPDATE语句里，VALUES()函数用于引用插入行的值。当插入的数据遇到唯一键冲突时，ON
DUPLICATE KEY UPDATE子句会执行更新操作，此时可以使用VALUES()函数获取原本要插入的值，从而在更新时使用这些值。
CHARSET函数主要用于返回字符串表达式所使用的字符集名称。
COLLATION函数用于返回给定字符串表达式所使用的排序规则（Collation）名称。
CURRENT_USER函数是一个系统函数，其主要功能是返回当前执行SQL语句的用户和主机信息。
DATABASE函数函数会返回正在使用的数据库名称。
SCHEMA()函数与DATABASE()函数功能基本相同，主要用于返回当前MySQL会话所使用的数据库的名称。
SESSION_USER函数用于返回当前会话中连接到数据库的用户名。
SYSTEM_USER函数用于返回当前与数据库建立会话的用户的用户名。
VERSION()函数的主要功能是返回当前数据库的版本信息。
VALIDATE_PASSWORD_STRENGTH函数给定一个表示明文密码的参数，此函数返回一个整数来指示密码的强度，如果参数为NULL，则返回
NULL。返回值范围从 0 （弱）到 100 （强）。
时间日期函数
ADDDATE函数对日期或日期时间值进行加法运算，也就是在指定的日期或日期时间上添加一定的时间间隔，从而得到一个新的日期或日期时间
值。
CURDATE函数返回当前执行该函数时所在系统的日期，返回值的格式为YYYY-MM-DD。
ADDTIME函数用于对时间或日期时间值加上一个时间间隔，从而得到一个新的时间或日期时间结果。
CURRENT_DATE函数会返回当前数据库所在系统的日期，其返回值格式为YYYY-MM-DD。
CURTIME函数会返回当前数据库所在系统的时间，返回值的格式通常为HH:MM:SS。
FROM_UNIXTIME函数将Unix时间戳转换为日期时间格式。
TO_DAYS函数用于将一个日期或日期时间值转换为从公元 0 年 1 月 1 日开始到指定日期所经过的天数。
TO_SECONDS函数功能是把一个日期或日期时间值转换为从公元 0 年 1 月 1 日00:00:00开始到指定日期或日期时间所经过的秒数。
CURRENT_TIME函数用于返回当前数据库的时间，格式为HH:MM:SS，表示小时、分钟和秒。
DATE_SUB函数的主要功能是从指定的日期或日期时间值中减去一个时间间隔，从而得到一个新的日期或日期时间值。
DATEDIFF函数主要用于计算两个日期之间相差的天数。
DAY函数是一个日期和时间处理函数，其主要功能是从给定的日期或日期时间值中提取出对应的天数部分。
DAYOFMONTH函数用于从指定的日期或日期时间值里提取该日期是所在月份的第几天。
HOUR函数是用于处理时间和日期时间数据的函数，它的主要功能是从给定的时间或日期时间值里提取小时部分。
MICROSECOND函数用于从给定的时间或日期时间值里提取微秒部分。
MINUTE函数主要用于从给定的时间或日期时间值里提取分钟部分。
MONTH函数用于从给定的日期或日期时间值中提取月份信息。
SECOND函数用于从给定的时间或日期时间值里提取秒数部分。
SUBDATE函数的主要功能是从指定的日期或日期时间值中减去一个时间间隔，进而得到一个新的日期或日期时间值。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
TIME_FORMAT函数用于将时间或日期时间值按照指定的格式进行格式化输出。
TIME_TO_SEC函数的主要功能是将一个时间值转换为对应的秒数。
WEEK函数的主要作用是从给定的日期或日期时间值中提取周信息。
YEAR函数是从给定的日期或日期时间值中提取年份信息。
GET_FORMAT函数用于返回预定义的日期、时间或日期时间格式字符串。
QUARTER函数用于从给定的日期或日期时间值中提取对应的季度信息。
CONVERT_TZ函数的主要功能是将一个日期时间值从一个时区转换到另一个时区。
DAYNAME函数用于返回给定日期对应的星期几的英文名称。
DAYOFYEAR函数用于返回给定日期是该年的第几天，其返回值范围是 1 到 366 。
FROM_DAYS函数的主要功能是将一个天数数值转换为对应的日期。
MONTHNAME函数用于返回给定日期对应的月份英文名称。
SEC_TO_TIME函数用于将以秒为单位的时间值转换为时间格式HH:MM:SS。
SUBTIME函数用于从给定的时间或日期时间值中减去一个时间间隔，进而得到新的时间或日期时间值。
TIMESTAMPADD函数的主要作用是在指定的日期或日期时间值上添加一个时间间隔，从而得到一个新的日期或日期时间值。
UTC_DATE函数用于返回当前的UTC（协调世界时）日期。
UTC_TIME函数的作用是返回当前的UTC（协调世界时）时间。
WEEKDAY函数用于返回给定日期对应的星期几，不过它返回的是用数字表示的结果，且以 0 开始计数。
WEEKOFYEAR函数的主要功能是返回给定日期是该年的第几周。
YEARWEEK函数用于将日期转换为年份和周数的组合形式，它能方便地对数据按年份和周进行分组统计、筛选等操作。
```
##### 新增 MySQL 系统视图

```
PROCESSLIST视图用于展示当前MySQL服务器中正在执行的线程信息。
SESSION_STATUS视图用于展示当前会话的各种状态信息。
SESSION_VARIABLES系统视图展示当前会话的系统变量及其对应的值。
GLOBAL_VARIABLES视图展示数据库的全局系统变量及其当前值。
```
##### 其他兼容功能

```
权限：支持grant/revoke对象级权限和全局系统权限，并且支持一个SQL可以grant/revoke授权/撤销多个权限。
通过使用0x开头和x''方式，支持表示十六进制字面量，可以用于插入、比较、计算和表示二进制数等场景，丰富了YashanDB的兼容性，提高可读性和
易用性。
字符集和字符序
新增支持GBK、GB18030字符集和gbk_bin、gb18030_bin排序。
支持按collation排序和比较字符串，当前支持包含ASCII、GBK、UTF8、ISO88591、UTF16和GB18030对应字符集的相应字符序。
sql_mode
PAD_CHAR_TO_FULL_LENGTH：默认情况下，检索CHAR列值时会修剪尾随空格。如果启用了PAD_CHAR_TO_FULL_LENGTH，则不会进行修
剪，检索到的CHAR值会填充至其完整长度。
NO_FIELD_OPTIONS：在SHOW CREATE TABLE的输出中不打印MySQL特定的列选项。该字段目前只用于兼容，可以设置但不生效。
NO_KEY_OPTIONS：不要在SHOW CREATE TABLE的输出中打印MySQL特定的索引选项。该字段目前只用于兼容，可以设置但不生效。
NO_TABLE_OPTIONS：不要在SHOW CREATE TABLE的输出中打印MySQL特定的表选项。该字段目前只用于兼容，可以设置但不生效。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
支持同时执行多条SQL语句，SQL之间使用分号隔开，返回多个SQL执行结果。
支持的mysqldump工具功能增强
支持全库/部分库/全表/部分表/表结构数据备份。
支持多个参数项。
```
#### 安全能力

```
支持表加密能力：通过CREATE TABLE支持table_encryption_clause子句功能，实现对指定表的加密，当插入或更新数据时会被截获、加密，然后用
加密后的格式保存，被查询时自动解密。加密和解密过程对用户透明。
支持密钥管理：新增能力二级密钥体系，通过内部机制（系统表或数据文件）管理DEK、外部钱包管理MEK，实现主密钥及数据加密密钥在使用时按
需生成。同时外部钱包支持创建/删除、权限管理等能力，帮忙主密钥实现轮转，更新密钥后对旧密钥进行作废，无法还原密钥等能力。
支持防篡改能力：通过HIST_CHECK高级包提供了一组内置的存储过程/函数，用于开启/关闭指定表的防篡改功能，校验指定表是否发生过篡改。
支持可配置的密码强度：在23.4版本之前，YashanDB开启密码强度后，规则是固定的，密码长度为8 - 64位和密码必须同时包含数字、大写字母、小
写字母和特殊字符。为了满足不同用户的安全准则和能覆盖更多的业务场景诉求，该版本YashanDB通过配置参数，实现对密码长度的控制，增加对密
码管理的灵活性。
YashanDB该版本支持审计和记录通过MySQL客户端登录登出的用户信息，完善了审计信息，且登录登出也是问题诊断的一个重要信息，提高安全能
力。
```
#### 运维诊断

```
支持执行计划选择诊断功能（ 10053 事件）：通过支持设置 10053 事件，例如'10053 TRACE NAME CONTEXT FOREVER, LEVEL 1'，记录SQL生成
执行计划过程，例如索引选择，JOIN顺序，算子选择等信息。用于跟踪SQL语句成本计算的内部事件，记载CBO模式下优化器如何计算SQL成本，生
成相应的执行计划。同时支持设置生成TRACE日志的路径和大小，提高易用性。详细使用介绍请查阅追踪。
支持使用SQL_ID绑定执行计划：通过新增CREATE OUTLINE outline_name ON sql_id USING HINT hint语法，给指定的SQL_ID的SQL创建
OUTLINE，让SQL固定为指定的HINT的指引的执行计划，解决无需业务进行SQL更改，也能到计划绑定的目的。提高了灵活性。
AWR快照存储优化和支持设置TOPSQL的数量：通过dbms_awr新增modify_snapshot_settings函数，支持exec
dbms_awr.modify_snapshot_settings(topnsql=>100)设置AWR的TOPSQL输出行数，默认为 30 ，可以更灵活展示AWR的TOP SQL的数据。
AWR功能存储优化：在v23.4版本前，清理AWR数据是通过DELETE操作删除的，在删除大量数据时，耗费时间较长。该版本对AWR能力增强，通过
把对应的系统表调整为分区表，使在清理历史数据的时候变成DROP分区操作，提高清理效率。
V$SQL_BIND_CAPTURE视图支持记录SQL的绑定变量信息：在v23.4版本之前，YashanDB无法获取SQL的绑定变量记录，该版本通过
V$SQL_BIND_CAPTURE视图新增字段实时记录SQL绑定参数具体得值信息，方便用于诊断SQL问题。
新增V$TEMPSEG_USAGE视图记录临时段的使用情况，包含在执行排序、哈希连接及执行索引创建、统计信息等操作时由数据库自动创建存储中间
结果。以及存储会话创建会话级全局临时表和私有临时表的存储空间。可用于性能监控与调优和故障排查。
SESSION视图要展示客户端连接的软件版本、协议版本：为了丰富诊断信息和和易用性，该版本在V$SESSION视图新增多个字段，包含
CLIENT_PROTOCAL_VERSION、CLIENT_VERSION、CLIENT_DRIVER、MACHIE和TERMINAL等，展示更多的客户端信息。同时登录过程中，
支持打印客户端软件/协议版本信息的日志。
V$SESSION对DDL语句支持查询：V$SESSION视图新增DTEXT字段，用于展示会话正在执行的DDL语句文本，超过 200 字节的部分截断，当会话不
在执行DDL语句时，该字段为NULL。
新增DICT_COLUMNS、DICTIONARY和DICT视图，DICT是DICTIONARY的同义词，DICTIONARY视图记录了系统视图与系统同义词信息，相当于
ALL_VIEWS和ALL_SYNONYMS中展示的内置view和同义；DICT_COLUMNS记录了上述DICTIONARY视图的列信息。
DBA_TABLES新增INITIAL_EXTENT、NEXT_EXTENT、MIN_EXTENTS和MAX_EXTENTS字段，代表表的初始化extent的大小、下一个分配extent
的大小，单位是字节。新增MIN_EXTENTS和MAX_EXTENTS字段代表段中允许的最小/大扩展数；分区表的这些字段对应的值为NULL。可用于监控
表的存储空间使用情况、预测未来的空间增长趋势、定位空间不足的问题和空间资源规划等场景。
支持ROLE_SYS_PRIVS和ROLE_TAB_PRIVS视图：新增ROLE_SYS_PRIVS视图用于显示授予角色的系统权限信息，ROLE_TAB_PRIVS视图用于
显示授予角色的对象权限信息。
支持DBA_EXTENTS视图：DBA_EXTETS支持查看每一个Extent所属的对象信息、Extent的起始位置和大小。 通过该视图可以查到某个数据文件的高
水位线上的Extent属于哪个对象方便分析高水位，从而对相应的对象进行空间回收从而Shrink表空间减少膨胀。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
支持DBA_TAB_HISTOGRAMS视图，查看数据库表的直方图信息：新增DBA_TAB_HISTOGRAMS视图会记录对表列收集统计信息时生成的直方图数
据，这些信息有助于优化器更为精准地估算查询结果集的大小，进而生成更高效的执行计划。
安装时支持安装包自校验功能
在使用OM安装部署数据库时，通过特定算法，对安装包进行计算值与使用公钥解密的值对比一致判断是否被篡改，如被篡改则安装强制终止，同时
也支持yasboot支持命令直接校验指定安装包是否被篡改。提高易用性和安全性。
针对资源要求不高或环境资源小的场景下，YashanDB通过对DBMS_PARAM包进行优化，推荐参数时降低默认值及去除对性能和可用性无影响的参
数资源设置，从而达到使用更小的内存配置达到更好数据库参数配置的优化效果，结合新增指定内存大小部署方式，使用户在小资源场景下也能体
验相应好的性能；同时增加引导式安装部署和前置校验的能力，提高易用性。
```
#### 驱动

##### JDBC

```
支持JDK 1.6/1.7版本：该版本之前，JDBC驱动依赖JDK 1.8及以上版本，为了匹配和兼容更多的客户场景，通过发布新的JDBC驱动版本，YashanDB
实现支持JDK 1.6和1.7版本的能力。
支持geoserver 2.23.3/2.25.5版本的方言包：为了满足不同场景的诉求，YashanDB支持GeoServer 2.23.3/2.25.5版本，通过提供gt-jdbc-yashandb-xxx
方言包和YashanDB JDBC驱动，用户可在GeoServer进行配置后，选择YashanDB为数据库源，可进行访问读取矢量数据，发布图层等操作。
支持standbyLoadBalance能力：JDBC连接串支持配置驱动多个IP及standby关键字，实现自动识别出备节点并轮询连接到配置IP上，同时达到负载均
衡效果。
PreparedStatment后支持自适应DDL变更：在v23.4版本之前，YashanDB的JDBC接口在获取PreparedStatement以后，缓存起来。此时改变
PreparedStatment对应表的结构，例如alter table modify col将col从varchar2(30)改成varchar2(32)，使用PrepareStatement.exeucteQuery，就会出
YAS-04007 Message：result set metadata changed异常。该版本通过对驱动和服务端交互方案进行优化后，实现了JDBC能支持元数据变化之后继续
执行PreparedStatment语句的能力。
JDBC支持mapDataToTimestamp接口：新增mapDataToTimestamp接口，在使用getString获取一个底层为date类型的数据时，实现了转换成
java.sql.Timestamp类型的能力，提高兼容性。
```
##### OCI

```
支持LOB函数：实现通过OCI驱动提供LOB数据类型实现对大型数据的存储和检索的功能。
支持事务函数：实现通过OCI驱动设置和控制事务的功能。
支持数字函数：实现通过OCI驱动提供OCI NUMBER与各种数值类型转换的功能。
支持字符串函数：实现通过OCI驱动提供处理字符串和管理内存的功能。
```
##### ODBC

```
支持SQLExecDirectW函数，允许应用程序直接发送SQL命令到数据库（宽字符版本，它的后缀W表示使用宽字符(Unicode)版本），一步完成语句准
备和执行，支持SELECT、INSERT、UPDATE等SQL语句。
支持SQLMoreResults函数，用于处理多结果集的场景。检查是否还有更多结果集可以处理，移动到下一个结果集。
支持GUID数据类型，是全局唯一标识符，通常用于标识数据库中的对象，通常以字符串形式表示，例如{12345678-1234-1234-1234-
1234567890AB}。
支持取单列数据多次调用SQLGetData接口：ODBC支持取单列数据多次调用SQLGetData接口：通过PHP pdo_odbc方式调用ODBC驱动取单列数据
时，是循环调用SQLGetData接口取数固定 256 ，单列数据超过 256 需要支持循环调用。该版本YashanDB新增此功能，实现处理大字段数据能力，避免
一次性加载大量数据导致内存溢出，提高了处理性能和灵活性。
SQLColAttributeW支持获取属性SQL_DESC_BASE_TABLE_NAME和SQL_DESC_AUTO_UNIQUE_VALUE。
SQLFetchScroll支持入参为SQL_FETCH_FIRST。
```
##### Python

```
支持loadbalance负载均衡：Python驱动支持loadBalance负载均衡，在高可用主备和负载均衡场景下，可以配置多个IP/PORT，应用将自动识别出连
接数最少的节点并连接，实现负责均衡的效果。
支持django orm框架：Python驱动支持Django ORM框架，为3.2.25版本，使开发人员轻松切换切换到YashanDB，而无需修改大量的代码，提高了代
码的可移植性，提升开发效率。
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

##### 其他驱动能力增强

```
适配Entity Framework 6框架：支持Entity Framework 6框架设置通过ADO.NET驱动连接YashanDB数据库，使用户可以使用EF6框架的功能开发
YashanDB，可以有效提高开发效率和增强可维护性。
支持Perl驱动：兼容了Perl语言中和数据库进行通讯的标准接口DBI模块，包含一系列的方法、变量和常量，实现可使用Perl驱动连接YashanDB数据库
进行SQL查询、处理结果集、事务管理等操作，相对于其他驱动，具备高效处理文本、易于学习和使用等优势。
支持.NET2.0驱动：新增支持.NET2.0驱动，实现基于VB.NET语言对.NET2.0平台开发的 32 位应用项目，丰富了支持的应用场景。
支持GO驱动和GORM框架：YashanDB支持GO驱动并适配GORM框架，支持开发人员使用Golang语言对YashanDB数据库进行开发，同时支持
GORM框架，使得开发人员可以使用Go代码轻松实现数据库的增删改查操作，无需编写复杂的SQL语句。
C驱动支持设置SQL执行超时时间：通过新增语句句柄属性设置yacSetStmtAttr接口，设置当前stmt的SQL执行超时时间，可以有效避免因长时间等待
而导致的连接失败或查询超时问题。
```
#### 工具能力

```
yasboot支持主备滚动升级：在23.4.1版本之前，只支持从23.2.0.0及以上版本向版本号前三位相同但第四位不同的新版本进行滚动升级。该版本支持
了从23.4.0.0及以上版本向版本号前三位不完全相同的新版本进行滚动升级，通过先把备库切换成逻辑备库，然后升级逻辑备库，完成之后做主备切
换，再重新把原主库搭建成物理备库，有必要再回切。YashanDB支持使用运维工具yasboot一键式完成滚动升级操作。减少了升级维护时间和提高易
用性。
yasboot支持重建备库：在23.4版本之前，运维工具yasboot仅在部署时支持主备部署，但无法在备库出现故障重建备库。该版本通过新增yasboot
node build相关命令，实现一键式重建备库的能力，提高易用性。
yasrman新增恢复命令中指定mapfile路径，可通过该映射文件中的路径映射信息恢复对应的数据文件至指定位置，解决目标环境的磁盘分配空间与原
来备份环境差异过大的场景，提高易用性。
yasql支持@@功能和spool数值开头文件：新增SPOOL命令支持以数字开头作为输出文件名，用于将yasql会话中的输出内容保存到文件中，丰富了使
用场景；当前版本之前，使用yasql支持@符号调用操作系统上的SQL文件执行，该版本增强次功能，通过@@实现调用当前脚本的路径作为相对路径
的基准点SQL脚本文件，方便嵌套脚本，适合处理多层级的脚本调用。
yasldr支持导入GIS数据多个SRID：在23.4版本之前，yasldr导入CSV格式的GIS数据时，一次导入只能使用参数指定一个SRID，不满足一个表存在多
个SRID的使用场景，该版本通过对yasldr增加对Geometry列函数计算，实现导入不同行不同的SRID，提高导入GIS数据灵活性。
```
### 兼容性变更

#### 产品规格

```
规格名称
```
```
规
格
类
型
```
```
变
更
类
型
```
```
描述 产品形态
```
```
全库闪回还原点数
量
```
```
最
大
值
```
```
新
增^8192 单机部署
```
```
yasql单行字符数
```
```
最
大
值
```
```
变
更 由^32000 修改为^65534
```
```
单机部署、共享集
群部署、分布式部
署
```
```
TIME类型
```
```
取
值
范
围
```
```
变
更
```
```
由00:00:00.000000 ~ 23:59:59.999999修改为-838:59:59.999999 ~
838:59:59.999999
```
```
单机部署、共享集
群部署、分布式部
署
```
```
VARCHAR类型
```
```
最
大
长
度
```
```
变
更 由32000Bytes修改为65534Bytes
```
```
单机部署、共享集
群部署、分布式部
署
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
规格名称
```
```
规
格
类
型
```
```
变
更
类
型
```
```
描述 产品形态
```
```
RAW类型
```
```
最
大
长
度
```
```
变
更 由8000Bytes修改为65534Bytes
```
```
单机部署、共享集
群部署、分布式部
署
```
```
语法模式（Yashan
模式或MySQL模
式）
```
```
切
换
方
式
```
```
变
更
```
```
由允许在会话中双向切换
修改为：
部署时通过EMPTY_STRING_AS_NULL参数指定语法模式，且部署为Yashan模式
后不允许切换到MySQL模式，部署为MySQL模式后可切换到Yashan模式
```
```
单机部署
```
#### 视图 & 系统表

```
视图名称
```
```
变
更
类
型
```
```
描述 产品形态
```
```
GV$AUD_UNIFIED
DV$AUD_UNIFIED
```
```
新
增 显示审计信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$BUFFER_POOL_PART_STATISTICS
V$BUFFER_POOL_PART_STATISTICS
```
```
新
增 显示数据缓存区的统计信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$DBLINK_XACT_STAT
V$DBLINK_XACT_STAT
```
```
新
增
```
```
显示所有database link在沙箱进程yex_server上链接组及对应链接的
的相关统计信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$ENCRYPTION_KEYS
V$ENCRYPTION_KEYS
```
```
新
增 显示主密钥描述属性
```
```
单机部署、共享集
群部署
GV$ENCRYPTION_WALLE
V$ENCRYPTION_WALLE
```
```
新
增 显示有关钱包状态和透明数据加密（TDE）的钱包位置的信息
```
```
单机部署、共享集
群部署
GV$EVENT_HISTOGRAM
DV$EVENT_HISTOGRAM
V$EVENT_HISTOGRAM
```
```
新
增
```
```
显示各等待事件从建库至今的等待次数、最大等待时间和总等待时间
的直方图
```
```
单机部署、共享集
群部署、分布式部
署
GV$EVENT_NAME
V$EVENT_NAME
```
```
新
增 显示等待事件的信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$FILESTAT
V$FILESTAT
```
```
新
增 显示已完成的物理读写次数、以文件级别完成的单块和多块I/O总数
```
```
单机部署、共享集
群部署、分布式部
署
GV$FLASHBACK_DATABASE_LOG
V$FLASHBACK_DATABASE_LOG
```
```
新
增 显示全库闪回功能的相关信息 单机部署
GV$FLASHBACK_DATABASE_LOGFILE
V$FLASHBACK_DATABASE_LOGFILE
```
```
新
增 显示全库闪回所有日志文件的信息 单机部署
GV$GCS_RESOURCE
V$GCS_RESOURCE
```
```
新
增 显示共享集群数据页面资源情况 共享集群部署
GV$GLS_RESOURCE
V$GLS_RESOURCE
```
```
新
增 显示共享集群数据锁资源情况 共享集群部署
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
```
```
变
更
类
型
```
```
描述 产品形态
```
```
GV$GRC_AFFINITY_POLICY
V$GRC_AFFINITY_POLICY
```
```
新
增 显示资源情况
```
```
单机部署、共享集
群部署、分布式部
署
GV$GTID_RESOURCE
V$GTID_RESOURCE
```
```
新
增 显示共享集群全局事务ID资源情况 共享集群部署
GV$MEX_AREA
V$MEX_AREA
```
```
新
增 显示MEX内存池AREA部分的信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$MEX_BASE
V$MEX_BASE
```
```
新
增 显示MEX内存池BASE部分的信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$MEX_EDEN
V$MEX_EDEN
```
```
新
增 显示MEX内存池EDEN部分的信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$MEX_POOL_MODULE
V$MEX_POOL_MODULE
```
```
新
增 显示MEX内存池各个功能模块的信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$MEX_SESSION
V$MEX_SESSION
```
```
新
增 显示各个会话中MEX内存池的使用情况
```
```
单机部署、共享集
群部署、分布式部
署
GV$MEX_SESSION_MODULE
V$MEX_SESSION_MODULE
```
```
新
增 显示各个会话中MEX内存池各个功能模块的使用情况
```
```
单机部署、共享集
群部署、分布式部
署
GV$PROCEDURE_STATS
V$PROCEDURE_STATS
```
```
新
增 显示PL Pool中自定义对象占用的内存信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$PX_STAGE
V$PX_STAGE
```
```
新
增 显示STAGE信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$RESTORE_POINT
V$RESTORE_POINT
```
```
新
增 显示全库闪回所有还原点的信息 单机部署
GV$RSRC_SESSION_INFO
V$RSRC_SESSION_INFO
```
```
新
增 显示会话使用资源相关信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$SEGSTAT
V$SEGSTAT
```
```
新
增 显示segment级别统计信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$SQLCOMMAND
V$SQLCOMMAND
```
```
新
增 显示所有SQL命令相关信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$XFMR_HIS_STAT
V$XFMR_HIS_STAT
```
```
新
增
```
```
显示列式存储用于转换生成稳态数据的所有已经执行完毕的xfmr任务
信息
```
```
单机部署、分布式
部署
GV$YFS_MEMORY_POOL
V$YFS_MEMORY_POOL
```
```
新
增 显示当前节点中YFS内存池（MEMORY POOL）信息 共享集群部署
GV$TABLESPACE_SET
V$TABLESPACE_SET
```
```
新
增 显示当前节点的所有表空间集信息 分布式部署
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
```
```
变
更
类
型
```
```
描述 产品形态
```
```
GV$TEMPSEG_USAGE
V$TEMPSEG_USAGE
```
```
新
增 显示临时属性的segment信息
```
```
单机部署、共享集
群部署、分布式部
署
GV$XFMR_STAT
V$XFMR_STAT
```
```
新
增
```
```
显示列式存储用于转换生成稳态数据的所有还未执行完毕的xfmr任务
信息
```
```
单机部署、分布式
部署
V$DF_RECOVERY_PROGRESS 新增 显示表空间文件回放进度汇总信息 单机部署
```
```
DBA_TSS_TABLESPACES 新增 显示所有归属于表空间集的表空间 分布式部署
```
```
DBA_OPTSTAT_OPERATION_TASKS 新增 显示使用作的历史记录DBMS_STATS高级包和数据库级别执行的统计信息收集操 单机部署、共享集群部署
```
```
DBA_OPTSTAT_OPERATIONS 新增 图显示使用操作的历史记录DBMS_STATS高级包和数据库级别执行的统计信息收集 单机部署、共享集群部署
```
```
GEOMETRY_COLUMN_TYPE_MODE 新增 显示当前用户可访问的所有对子类型有限制条件的类型的列信息 ST_GEOMETRY
```
```
单机部署、共享集
群部署部署、分布
式部署
GV$ARCHIVE_DEST_STATUS
V$ARCHIVE_DEST_STATUS
```
```
变
更
```
```
* 新增DEST_NAME字段：表示参数名称
* 新增DB_UNIQUE_NAME字段：表示备库的名称
```
```
单机部署、共享集
群部署部署、分布
式部署
GV$ARCHIVED_LOG
V$ARCHIVED_LOG
```
```
变
更
```
```
* 新增STATUS字段：表示归档的状态
* 新增ARCHIVED字段：表示在线redo是否被归档
* 新增DELETED字段：表示归档文件是否被删除
```
```
单机部署、共享集
群部署、分布式部
署
GV$BUFFER_CONTROL
V$BUFFER_CONTROL
```
```
变
更 新增OBJ字段：当前页面所属的对象ID
```
```
单机部署、共享集
群部署、分布式部
署
GV$BUFFER_POOL_STATISTICS
V$BUFFER_POOL_STATISTICS
```
```
变
更
```
```
* 新增NAME字段：表示缓冲区的名称，当前固定为DEFAULT
* 新增SET_MSIZE字段：表示缓冲区内可设置的最大的可容纳数据块
数量
```
```
单机部署、共享集
群部署、分布式部
署
DV$COLUMNAR_STAGE_QUOTA
GV$COLUMNAR_STAGE_QUOTA
V$COLUMNAR_STAGE_QUOTA
```
```
变
更 新增BULK_SIZE字段：表示每批次记录行数
```
```
单机部署、分布式
部署
```
```
GV$DATABASE
V$DATABASE
```
```
变
更
```
```
* 新增DBID字段：表示数据库ID
* 新增NAME字段：表示数据库名称
* 新增CREATED字段：表示建库时间
* 新增RESETLOGS_CHANGE#字段：表示open resetlogs时的系统
修改序列号（SCN）
* 新增FLASHBACK_ON字段：表示是否开启全库闪回功能
```
```
单机部署、共享集
群部署、分布式部
署
```
```
GV$DATAFILE
V$DATAFILE
```
```
变
更 新增FILE#字段：表示数据文件的编号
```
```
单机部署、共享集
群部署、分布式部
署
GV$DATATYPE
V$DATATYPE
```
```
变
更 修改MAX_SIZE字段类型由SMALLINT变更为INTEGER
```
```
单机部署、共享集
群部署、分布式部
署
GV$DICT_CACHE
V$DICT_CACHE
```
```
变
更
```
```
新增NOLOGGING_INSTANCE_ID字段：表示当前字典缓存对象开启
nologging的实例ID，若对象非nologging表，默认为-1
```
```
单机部署、共享集
群部署、分布式部
署
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
```
```
变
更
类
型
```
```
描述 产品形态
```
```
DV$DICT_CURSOR
GV$DICT_CURSOR
V$DICT_CURSOR
```
```
变
更
```
```
新增CONTEXT字段：表示游标相关的上下文信息，为预留字段，默
认为空
```
```
单机部署、共享集
群部署、分布式部
署
GV$GRC_RESOURCE
V$GRC_RESOURCE
```
```
变
更
```
```
* 新增OBJ字段，当前页面所属的对象ID
* 修改TYPE字段类型由TINYINT变更为INTEGER 共享集群部署
GV$INSTANCE
V$INSTANCE
```
```
变
更 新增DATABASE_STATUS字段，表示数据库的状态
```
```
单机部署、共享集
群部署、分布式部
署
GV$INSTANCE_RECOVERY
V$INSTANCE_RECOVERY
```
```
变
更 新增PHASE字段，表示当前故障恢复的阶段
```
```
单机部署、共享集
群部署、分布式部
署
GV$LOGFILE
V$LOGFILE
```
```
变
更
```
```
* 新增ARCHIVED字段：表示redo文件是否被归档
* 新增HEALTH字段：表示数据库加载时redo文件的健康度
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DV$RSRC_CONSUMER_GROUP
GV$RSRC_CONSUMER_GROUP
V$RSRC_CONSUMER_GROUP
```
```
变
更
```
```
* 新增SESSION_SPA_MAX_USE_MEM字段：表示会话内存使用最
大值
* 新增CONCURRENCY_LIMIT字段：表示资源计划配置最大并发数
* 新增CONCURRENCY_NUM字段：表示当前正在执行的资源密集型
SQL数量
* 新增EXECUTION_WAITERS字段：表示等待执行的资源密集型
SQL
* 新增REQUESTS字段：表示资源组内已经执行的资源密集型SQL数
量
* 新增QUEUE_NUMBER字段：表示排队叫号值
* 新增CONCURRENCY_LIMIT_HIT字段：表示同时执行资源密集型
SQL数量达到上限次数
* 新增QUEUED_TIME字段：表示资源组内所有会话累计等待执行时
间
* 新增QUEUE_TIMEOUTS字段：表示资源组内会话累计等待调度超
时次数
* 新增PARALLEL_DOWNGRADE_TIMES字段：表示资源组内出现
并行资源降级次数
```
```
单机部署、共享集
群部署、分布式部
署
```
```
GV$SESSION
V$SESSION
```
```
变
更
```
```
* 新增RETRY_CNT字段：表示当前命令的重试次数
* 新增RETRY_INFO字段：表示当前命令的重试错误码
* 新增EXEC_STATUS字段：表示执行状态
```
```
单机部署、共享集
群部署、分布式部
署
GV$TABLESPACE
V$TABLESPACE
```
```
变
更
```
```
新增ENCRYPT_ALGO字段：表示表空间加密算法，如果是非加密表
空间，该字段值为空 共享集群部署
DV$REPLICATION
V$REPLICATION
```
```
变
更
```
```
新增TRIGGER_COND_FAILOVER字段：表示主库是否触发了条件
故障切换 分布式部署
```
```
V$MYSQL_COLLATION 变更 新增YashanDBYAS_COLLA字符序TION_IDID 字段：表示MySQL字符序对应的
```
```
单机部署、共享集
群部署、分布式部
署
V$YSTREAM_STAT 变更 新增CAPTURE_LFN字段：表示已解析的redo日志LFN 单机部署、共享集群部署
```
```
DBA_AUDIT_MGMT_CLEANUP_JOBS 变更
```
```
修改START_DATE、END_DATE、LAST_START_DATE、
NEXT_RUN_DATE 字段类型由TIMESTAMP变更为TIMESTAMP
WITH TIME ZONE
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_DATA_FILES 变更 * * 新增新增ONLINE_STAUTOEXTENSIBLEATUS字段：表示文件状态字段：表示文件是否可以自动扩展^
```
```
单机部署、共享集
群部署、分布式部
署
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
视图名称
```
```
变
更
类
型
```
```
描述 产品形态
```
```
DBA_INDEXES 变更 新增对 ORPHANED_ENTRIES字段：表示全局索引是否包含孤儿键值
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_JOBS 变更 修改变更为LAST_DATIMESTTEAMP WITH TIME ZONE、THIS_DATE、NEXT_DATE字段由TIMESTAMP
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_OUTLINES 变更 新增SQL_ID字段：表示SQL_ID语法OUTLINE对应的SQL语句ID值
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_RSRC_PLAN_DIRECTIVES 变更 新增资源密集型CONCURRENCY_LIMITSQL的数量 字段：表示资源使用组内允许同时执行
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_SCHEDULER_JOBS 变更
```
```
修改LAST_DATE、THIS_DATE、LAST_START_DATE、
NEXT_RUN_DATE字段由TIMESTAMP变更为TIMESTAMP WITH
TIME ZONE
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_SEGMENTS 变更 新增SEGMENT_SUBTYPE字段：表示segment子类型
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_TABLES 变更
```
```
* 新增DEGREE字段：表示表的并行度
* 新增INITIAL_EXTENT字段：表示segment的首个extent大小，单位
为bytes
```
```
单机部署、共享集
群部署、分布式部
署
```
```
DBA_TABLESPACES 变更 新增BIGFILE字段：字段兼容，无实际意义
```
```
单机部署、共享集
群部署、分布式部
署
```
```
UNIFIED_AUDIT_TRAIL 变更
```
```
* 新增YLS_POLICY_NAME字段：表示行访问控制策略名称
* 新增YLS_GRANTEE字段：表示应用行访问控制策略的用户
* 新增YLS_MAX_READ_LABEL字段：表示用户的最大读标签
* 新增YLS_MAX_WRITE_LABEL字段：表示用户的最大写读标签
* 新增YLS_MIN_WRITE_LABEL字段：表示用户的最小写读标签
* 新增YLS_STRING_LABEL字段：表示标签内容
* 新增YLS_LABEL_COMPONENT_TYPE字段：表示组件类型
* 新增YLS_LABEL_COMPONENT_NAME字段：表示组件名称
* 新增RLS_INFO字段：表示执行语句中表关联的行访问控制策略名
称
```
```
单机部署、共享集
群部署、分布式部
署
```
#### 系统配置参数

```
配置参数 修改类型 描述 产品形态
DB_FLASHBACK_FILE_DEST_SIZE 新增 全库闪回所产生的日志文件能占据的最大空间限制 单机部署
DB_FLASHBACK_RETENTION_TARGET 新增 全库闪回所产生的日志文件的最大保留时长 单机部署
WALLET_ROOT 新增 指定钱包的根路径 单机部署、共享集群部署
BUDDY_INSTANCE_SCAN_INTERVAL 变更 默认值从64M修改为128M 共享集群部署
BUDDY_INSTANCE_SCAN_TIMEOUT 变更 默认值从 10 修改为 60 共享集群部署
```
### 升级说明


###### 深圳计算科学研究院 深圳崖山科技有限公司

```
产品形态产品形态 升级说明升级说明 升级方式升级方式
```
```
单机部署
```
```
支持从v22.2版本升级到v23.4版本
支持从v23.1版本升级到v23.4版本
支持从v23.2版本升级到v23.4版本
支持从v23.3版本升级到v23.4版本
```
```
离线升级
```
```
共享集群部署
```
```
支持从v23.1版本升级到v23.4版本
支持从v23.2版本升级到v23.4版本
支持从v23.3版本升级到v23.4版本
```
```
离线升级
```
```
分布式部署
```
```
支持从v23.1版本升级到v23.4版本
支持从v23.2版本升级到v23.4版本
支持从v23.3版本升级到v23.4版本
```
```
离线升级
```
### 周边配套

```
YashanDB v23.4版本推荐使用的平台工具版本如下：
组件 配套版本
Yashandb JDBC驱动程序（JDBC） 1.9.31.9.3-jre6^
Yashandb ODBC驱动程序（ODBC） 23.4.1.100
Yashandb C驱动程序（C） 23.4.1.100
Yashandb C#驱动程序（ADO.NET） 1.6.1
Yashandb Python驱动程序（Python） 1.2.0
YashanDB 客户端（yasql） 23.4.1.100
YashanDB 监控运维工具（YCM） 23.4.1.0
YashanDB 迁移平台工具（YMP） 23.4.1.0
YashanDB 开发者工具 23.4.1.0
```
### 版本约束

```
约束项 产品形态
1 增加节点时，主备集群所有节点需处于正常可用状态，增加节点后数量不超过原有规模约束 单机部署
2 删除节点操作时，不支持直接删除主节点，并且不支持将主备集群全部节点删除 单机部署
3 增加节点时，集群所有节点需处于正常可用状态。不可组内无主节点时扩容 分布式部署
4 删除节点操作时， 不支持直接删除主节点，并且不支持将DN组/MN全部节点删除 分布式部署
5 统计信息收集过程中，如果存在CN故障，则需要恢复CN后重新收集 分布式部署
6 扩容OUTLINECN后，需要重新收集统计信息，如果有使用OUTLINE，则需要在收集统计信息完成后，重建 分布式部署
7 insert into select from外部表，只支持insert写入复制表的行表 分布式部署
8 drop userDBLINK对象失败或部分成功，则该用户不允许登录，也不允许创建该用户下的表、视图、索引、AC以及 分布式部署
```
```
9 过程体中的SQL语句存在语法错误时，返回的错误消息的Position不准确 单机部署、共享集群部署、分布式部署
```

###### 深圳计算科学研究院 深圳崖山科技有限公司

```
约束项 产品形态
10 执行误 blob转字符串时，如blob不是合法的utf8编码，行存引擎会忽略不合法的字符，列存引擎会返回错 单机部署、共享集群部署、分布式部署
```
```
11 SQLMAP原始语句及映射语句长度不得超过 32000 字节 单机部署、共享集群部署、分布式部署
```


