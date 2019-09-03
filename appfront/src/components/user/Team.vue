<template>
  <div :style="{ padding: '24px', background: '#fff'}">
    <a-table :columns="columns"
    :dataSource="data"
    :pagination="pagination"
    :loading="loading"
    @change="handleTableChange"
    >
      <template slot="username" slot-scope="username">
      {{username}}
      </template>
      <template slot="role" slot-scope="role">
        <a-tag color="blue">{{role}}</a-tag>
      </template>
    </a-table>
  </div>
</template>

<script>
import axios from 'axios';
const columns = [{
  title: '用户名',
  dataIndex: 'username',
  width: '50%',
  scopedSlots: { customRender: 'username' },
}, {
  title: '角色',
  dataIndex: 'role',
  width: '50%',
  scopedSlots: { customRender: 'role' },
}];
export default {
  mounted() {
    this.fetch();
  },
  data () {
    return {
      data: [],
      columns,
      pagination: {},
      loading: false,
    }
  },
  methods: {
    handleTableChange (pagination, filters, sorter) {
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
        this.loading = true;
        axios({
            url: '/api/v1/team',
            method: 'get',
            params: {
                pageSize: 10,
                ...params
            },
        }).then((res) => {
            const pagination = { ...this.pagination };
            pagination.total = res.data.totalCount;
            this.loading = false;
            this.data = res.data.results;
            this.pagination = pagination;
        })
    },
  },
}
</script>