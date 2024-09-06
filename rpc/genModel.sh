#生成表名
tables=$2
#表生成的genModel目录
modelDir=./models

#数据库配置
host=127.0.0.1
port=3306
dbname=$1
username=root
passwd=123456

echo "开始创建库 $dbname 的表：$2"

goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modelDir}" --style=goZero -c