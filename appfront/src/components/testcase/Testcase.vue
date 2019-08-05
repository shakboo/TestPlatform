<template>
    <div :style="{ padding: '24px', background: '#fff', textAlign: 'center' }">
        <a-table :columns="columns"
            :dataSource="data"
            :pagination="pagination"
            :loading="loading"
            @change="handleTableChange"
        >
            <template slot="id" slot-scope="id">
            {{ id }}
            </template>
        </a-table>
    </div>
</template>

<script>
import axios from 'axios';

const columns = [{
  title: '编号',
  dataIndex: 'id',
  sorter: true,
  width: '5%',
  scopedSlots: { customRender: 'id' },
}, {
  title: '模块',
  dataIndex: 'module',
  filters: [
    { text: '数据集', value: '数据集' },
    { text: '实验', value: '实验' },
    { text: '模型', value: '模型' },
    { text: '在线服务', value: '在线服务' },
    { text: '用户', value: '用户' },
    { text: '帮助文档', value: '帮助文档' },
  ],
  width: '10%',
}, {
  title: '用例描述',
  dataIndex: 'describe',
  width: '20%',
}, {
  title: '测试步骤',
  dataIndex: 'step',
  width: '40%',
  class: 'comment',
}, {
  title: '预期结果',
  dataIndex: 'result',
  class: 'comment',
}];

export default {
  mounted () {
    this.fetch();
  },
  data () {
    return {
      data: [],
      pagination: {},
      loading: false,
      columns,
    }
  },
  methods: {
    handleTableChange (pagination, filters, sorter) {
      console.log(pagination);
      const pager = { ...this.pagination };
      pager.current = pagination.current;
      this.pagination = pager;
      this.fetch({
        pageSize: pagination.pageSize,
        pageCurrent: pagination.current,
        sortField: sorter.field,
        sortOrder: sorter.order,
        ...filters,
      });
    },

    fetch (params = {}) {
      this.loading = true
      axios({
        url: '/api/v1/testcase',
        method: 'get',
        data: {
          pageSize: 10,
          ...params
        },
      }).then((res) => {
        const pagination = { ...this.pagination };
        // Read total count from server
        pagination.total = res.data.totalCount;
        // pagination.total = 10000;
        this.loading = false;
        this.data = res.data.results;
        this.pagination = pagination;
      });
    }
  },
}
</script>

<style>
.comment{
  white-space: pre-wrap;
}
</style>
