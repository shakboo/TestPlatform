#coding=utf-8

import csv
import pymysql
import pandas
import sys

class saveCsvToMysql():

    def __init__(self, ip, port, user, password):
        self.config = dict(
            host= ip,
            port = int(port),
            user= user,
            password=password,
            cursorclass=pymysql.cursors.DictCursor
        )

    def main(self, dbName, tableName, f):

        # 建立连接
        conn = pymysql.Connect(**self.config)

        # 自动确认commit
        conn.autocommit(1)

        # 设置光标
        cursor = conn.cursor()

        # 读取csv
        f = pandas.read_csv(f, encoding='utf8')

        # 创建database
        cursor.execute('CREATE DATABASE IF NOT EXISTS {0}'.format(dbName))
        # 连接database
        conn.select_db(dbName)
        # 创建table
        cursor.execute('DROP TABLE IF EXISTS {0}'.format(tableName))
        cursor.execute('CREATE TABLE {0}({1}) default charset=utf8'.format(tableName, self._make_table_sql(f)))
        # 提取数据转list
        values = f.values.tolist()
        s = ','.join(['%s' for _ in range(len(f.columns))])
        # 批量插入数据
        cursor.executemany('INSERT INTO {0} VALUES ({1})'.format(tableName, s), values)

        cursor.close()
        conn.close()

    def _make_table_sql(self, f):
        columns = f.columns.tolist()
        print(columns)
        types = f.ftypes
        make_table = []
        for item in columns:
            if 'int' in types[item]:
                char = item + ' INT'
            elif 'float' in types[item]:
                char = item + ' FLOAT'
            elif 'object' in types[item]:
                char = item + ' VARCHAR(255)'
            elif 'datetime' in types[item]:
                char = item + ' DATETIME'
            make_table.append(char)
        return ','.join(make_table)

if __name__ == '__main__':
    s = saveCsvToMysql(*(sys.argv[1:5]))
    s.main(*(sys.argv[5:]))