<template>
  <div :style="{ padding: '24px', background: '#fff'}">
    <div style="margin-bottom: 16px">
      <a-button
        type="primary"
        @click="addTestcaseModel"
        :disabled="!hasSelected"
      >
        添加用例
      </a-button>
      <a-modal
        title="添加用例"
        :visible="visible"
        @ok="handleSubmit"
        :confirmLoading="confirmLoading"
        @cancel="handleCancel"
        :width="800"
        okText="提交"
        cancelText="取消"
      >
        <a-form 
          :form="form"
          @submit="handleSubmit">
          <a-form-item
            label="模块"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 18 }"
          >
            <a-input 
            v-decorator="[
            'module',
            {rules: [{ required: true, message: '请填写用例模块' }]}
            ]" />
              <a-select-option v-for="option in columns[1].filters"
              :value="option.value">{{ option.text }}</a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item
            label="重要性"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 18 }"
          >
            <a-select
            v-decorator="[
            'importance',
            {rules: [{ required: true, message: '请选择用例的重要性' }]}
            ]">
              <a-select-option v-for="option in columns[2].filters"
              :value="option.value">{{ option.text }}</a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item
            label="用例描述"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 18 }"
          >
            <a-input 
            v-decorator="[
            'describe',
            {rules: [{ required: true, message: '请填写用例描述' }]}
            ]" />
          </a-form-item>
          <a-form-item
            label="测试步骤"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 18 }"
          >
            <a-textarea 
            v-decorator="[
            'step',
            {rules: [{ required: true, message: '请填写测试步骤' }]}
            ]" />
          </a-form-item>
          <a-form-item
            label="预期结果"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 18 }"
          >
            <a-textarea 
            v-decorator="[
            'result',
            {rules: [{ required: true, message: '请填写预期结果' }]}
            ]" />
          </a-form-item>
        </a-form>
      </a-modal>
      <a-button
        type="primary"
        @click="exportTestcase"
        :loading="exportLoading"
      >
        导出用例
      </a-button>
      <a-button
        type="primary"
        @click="importTestcaseModel"
      >
        导入用例
      </a-button>
      <a-modal
        title="上传用例"
        :confirmLoading="importLoading"
        :visible="importModalVisible"
        @ok="handleImportModalSubmit"
        @cancel="handleImportModalCancel"
        okText="提交"
        cancelText="取消"
      >
        <a-upload
        accept=".xlsx"
        :fileList="fileList"
        :remove="handleImportModalRemove"
        :beforeUpload="beforeUpload"
        >
          <a-button>
            <a-icon type="upload" /> 选择文件
          </a-button>
        </a-upload>
      </a-modal>
    </div>
    <a-table :columns="columns"
        :dataSource="data"
        :pagination="pagination"
        :loading="loading"
        :customRow="customRow"
        @change="handleTableChange"
    >
      <template slot="id" slot-scope="id">
      {{ id }}
      </template>
      <template v-for="(col, colIndex) in ['module', 'importance', 'describe', 'step', 'result']"
        :slot="col"
        slot-scope="text, record, index" >
        <div :key="col" v-if="record.editable">
          <a-input
            v-if="col === 'describe' || col === 'module'"
            style="margin: -5px 0"
            :value="text"
            @change="e => handleChange(e.target.value, record.id, col)"
          />
          <a-select
            v-else-if="col === 'importance'"
            style="margin: -5px 0"
            :defaultValue="text"
            @change="e => handleChange(e, record.id, col)"
          >
            <a-select-option v-for="option in columns[colIndex+1].filters"
              :value="option.value">{{ option.text }}</a-select-option>
          </a-select>
          <a-textarea
            v-else-if="col === 'step' || col === 'result'"
            style="margin: -5px 0"
            :value="text"
            :rows="4"
            @change="e => handleChange(e.target.value, record.id, col)"
          />
        </div>
        <div :key="col" v-else><template>{{ text }}</template></div>
      </template>
      <template slot="operation" slot-scope="text, record, index">
        <div class='editable-row-operations'>
          <span v-if="record.editable">
            <a @click="() => save(record.id)">保存</a>
            <a-popconfirm title='确定放弃更改吗？' @confirm="() => cancle(record.id)">
              <a>取消</a>
            </a-popconfirm>
          </span>
          <span v-else>
            <a @click="() => edit(record.id)">编辑</a>
            <a-popconfirm title='确定删除吗？' @confirm="() => handleDelete(record.id)" style="margin-left: 10px"><a>删除</a></a-popconfirm>
          </span>
        </div>
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
  width: '6%',
  scopedSlots: { customRender: 'id' },
}, {
  title: '模块',
  dataIndex: 'module',
  width: '8%',
  filters: [],
  scopedSlots: { customRender: 'module' },
}, {
  title: '等级',
  dataIndex: 'importance',
  filters: [
    { text: '高', value: '高' },
    { text: '中', value: '中' },
    { text: '低', value: '低' },
  ],
  width: '6%',
  scopedSlots: { customRender: 'importance' },
}, {
  title: '用例描述',
  dataIndex: 'describe',
  width: '13%',
  scopedSlots: { customRender: 'describe' },
}, {
  title: '测试步骤',
  dataIndex: 'step',
  width: '32%',
  class: 'comment',
  scopedSlots: { customRender: 'step' },
}, {
  title: '预期结果',
  dataIndex: 'result',
  class: 'comment',
  width: "25%",
  scopedSlots: { customRender: 'result' },
}, {
  title: '操作',
  dataIndex: 'operation',
  scopedSlots: { customRender: 'operation' },
}];

export default {
  mounted () {
    var path = this.$route.path;
    this.item = path.split("/")[path.split('/').length - 1];  // 获取路由中项目参数
    this.fetch();
  },
  data () {
    return {
      item: "",  // 当前项目
      data: [],   // 后端返回的数据
      cacheData: [],   // 编辑时缓存的诗句
      pagination: {pageCurrent: 1},   // 分页
      loading: false,   // 用例loading
      selectedRowKey: 0,  // 选中行的id
      selectedElement: [],  // 选中行的[element, index]
      columns,          
      visible: false,        // Modal
      confirmLoading: false,  // Modal
      exportLoading: false,   // export
      importLoading: false,
      importModalVisible: false, // import Modal
      fileList: [],  // import Modal
    }  
  },
  computed: {
    hasSelected() {
      return this.selectedRowKey > 0 || this.pagination.total === 0
    }
  },
  beforeCreate () {
    this.form = this.$form.createForm(this);
  },
  methods: {
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination };
      pager.current = pagination.current;
      this.pagination = pager;

      if (this.selectedRowKey > 0) {
        this.selectedElement[0].style['backgroundColor'] = '#fff';
        this.selectedElement = [];
        this.selectedRowKey = 0;
      };
      
      this.fetch({
        pageSize: pagination.pageSize,
        pageCurrent: pagination.current,
        sortField: sorter.field,
        sortOrder: sorter.order,
        ...filters,
      });

      this.pagination.pageCurrent = pagination.current;
      this.pagination.sortField = sorter.field;
      this.pagination.sortOrder = sorter.order;
      this.pagination.filters = filters
    },

    fetch (params = {}) {
      this.loading = true;
      axios({
        url: '/api/v1/testcase',
        method: 'get',
        params: {
          pageSize: 10,
          item: this.item,
          ...params
        },
      }).then((res) => {
        const pagination = { ...this.pagination };
        pagination.total = res.data.totalCount;
        this.loading = false;
        this.data = res.data.results;
        var modules = res.data.modules
        this.columns[1].filters = [];
        for (var i=0;i<modules.length;i++) {
          this.columns[1].filters.push({"text": modules[i].module, "value": modules[i].module});
        }
        console.log(this.columns);
        this.cacheData = this.data.map(item => ({...item}));
        this.pagination = pagination;
      });
    },
    handleChange (value, key, column) {
      const newData = [...this.data];
      const target = newData.filter(item => key === item.id)[0];
      if (target) {
        target[column] = value;
        this.data = newData;
      }
    },
    edit (key) {
      const newData = [...this.data];
      const target = newData.filter(item => key === item.id)[0];
      if (target) {
        target.editable = true;
        this.data = newData;
      }
    },
    save (key) {
      const newData = [...this.data];
      const formData = new FormData();
      const target = newData.filter(item => key === item.id)[0];
      formData.append('id', target.id);
      formData.append('module', target.module);
      formData.append('importance', target.importance);
      formData.append('describe', target.describe);
      formData.append('step', target.step);
      formData.append('result', target.result);
      formData.append('item', this.item);
      if (target) {
        axios({
          url: '/api/v1/testcase',
          method: 'put',
          data: formData,
        }).then((res) => {
          this.$message.success('修改成功');
          delete target.editable;
          this.data = newData;
          this.cacheData = newData.map(item => ({ ...item }));
        });
      }
    },
    handleDelete (key) {
      const newData = [...this.data];
      const target = newData.filter(item => key === item.id)[0];
      if (target) {
        axios({
          url: 'api/v1/testcase',
          method: 'delete',
          params: {
            'id': target.id,
            'item': this.item, 
          },
        }).then((res) => {
          this.$message.success("删除成功");
          // 如果是一页的最后一条数据，删除后跳转到上一页
          if (this.pagination.total % 10 === 1) {
            this.pagination.pageCurrent -= 1;
          }
          this.fetch({
            pageCurrent: this.pagination.pageCurrent,
            sortOrder: this.pagination.sortOrder,
            sortField: this.pagination.sortField,
            ...this.pagination.filters,
          });
        })
      }
    },
    cancle (key) {
      const newData = [...this.data];
      const target = newData.filter(item => key === item.id)[0];
      if (target) {
        Object.assign(target, this.cacheData.filter(item => key === item.id)[0]);
        delete target.editable;
        this.data = newData;
      }
    },
    customRow (record, index) {
      return {
        on: {
          click: (e) => {
            // 选中元素
            if (e.target.tagName != "TR") {
              e = e.target.parentNode;
              while (e.tagName != "TR") {
                e = e.parentNode;
              };
              if (this.selectedElement.length > 0) { 
                if (index == this.selectedElement[1]) {
                  e.style['backgroundColor'] = "#fff";
                  this.selectedElement = [];
                  this.selectedRowKey = 0;
                } else {
                  e.style['backgroundColor'] = "#87CEFB";
                  this.selectedElement[0].style['backgroundColor'] = "#fff";
                  this.selectedRowKey = record.id;
                  this.selectedElement = [e, index];
                }    
                
              } else {
                e.style['backgroundColor'] = "#87CEFB"; 
                this.selectedRowKey = record.id;
                this.selectedElement = [e, index];
              }
            } else {
              if (this.selectedElement.length > 0) {
                if (index == this.selectedElement[1]) {
                  e.target.style['backgroundColor'] = "#fff";
                  this.selectedElement = [];
                  this.selectedRowKey = 0;
                } else {
                  e.target.style['backgroundColor'] = "#87CEFB";
                  this.selectedElement[0].style['backgroundColor'] = "#fff";
                  this.selectedRowKey = record.id;
                  this.selectedElement = [e.target, index];
                } 
              } else {
                e.target.style['backgroundColor'] = "#87CEFB"; 
                this.selectedRowKey = record.id;
                this.selectedElement = [e.target, index];
              } 
            }        
          },
        }
      }
    },
    addTestcaseModel () {
      this.visible = true;
    },
    importTestcaseModel () {
      this.importModalVisible = true;
    },
    handleCancel () {
      this.visible = false;
    },
    handleImportModalCancel () {
      this.importModalVisible = false;
    },
    handleSubmit (e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          const formData = new FormData();
          formData.append('preId', this.selectedRowKey);
          formData.append('module', values.module);
          formData.append('importance', values.importance);
          formData.append('describe', values.describe);
          formData.append('step', values.step);
          formData.append('result', values.result);
          formData.append('item', this.item);
          this.confirmLoading = true;
          axios({
            url: '/api/v1/testcase',
            method: 'post',
            data: formData,
          }).then((res) => {
            this.$message.success('添加成功');
            this.confirmLoading = false;
            this.visible = false;
            this.fetch({
              pageCurrent: this.pagination.pageCurrent,
              sortOrder: this.pagination.sortOrder,
              sortField: this.pagination.sortField,
              ...this.pagination.filters,
            });
            // 提交后清除modal里的内容
            this.form.clearField('module');
            this.form.clearField('importance');
            this.form.clearField('describe');
            this.form.clearField('step');
            this.form.clearField('result');
          });
        } else {
          this.$message.error("添加失败");
          return false;
        }
      })
    },
    exportTestcase () {
      this.exportLoading = true;
      axios({
        url: '/api/v1/testcase/export',
        method: 'get',
        params: {
          "item": this.item,
        },
        responseType: 'blob',
      }).then((res) => {
        this.exportLoading = false;
        this.$message.success('导出成功');
        var fileName = "测试用例.xlsx";
        if (this.item === 'ml') {
          fileName = "机器学习平台" + fileName;
        } else if (this.item === 'sj') {
          fileName = "数据治理平台" + fileName;
        }
        if ('download' in document.createElement('a')) { // 非IE下载
          const elink = document.createElement('a');
          elink.download = fileName;
          elink.style.display = 'none';
          elink.href = URL.createObjectURL(res.data);
          document.body.appendChild(elink);
          elink.click();
          URL.revokeObjectURL(elink.href) ;// 释放URL 对象
          document.body.removeChild(elink);
        } else { // IE10+下载
          navigator.msSaveBlob(res.data, fileName);
        }
      });
    },
    handleImportModalRemove(file) {
      const index = this.fileList.indexOf(file);
      const newFileList = this.fileList.slice();
      newFileList.splice(index, 1);
      this.fileList = newFileList
    },
    beforeUpload(file) {
      this.fileList = [file]
      return false;
    },
    handleImportModalSubmit (e) {
      e.preventDefault();
      if (this.fileList.length == 0) {
        this.$message.error("请选择上传文件");
        return false;
      }
      const formData = new FormData();
      formData.append('file', this.fileList[0]);
      formData.append('item', this.item);
      this.importLoading = true;
      axios({
        url: '/api/v1/testcase/import',
        method: 'post',
        data: formData
      }).then((res) => {
        this.importLoading = false;
        this.fileList = [];
        this.importModalVisible = false;
        this.fetch({
          pageCurrent: this.pagination.pageCurrent,
          sortOrder: this.pagination.sortOrder,
          sortField: this.pagination.sortField,
          ...this.pagination.filters,
        });
        this.$message.success("用例上传成功");
      })
    }
  }
}
</script>

<style>
.comment{
  white-space: pre-wrap;
}
.editable-row-operation a {
  margin-right: 8px;
}
.ant-table-tbody>tr.ant-table-row-hover:not(.ant-table-expanded-row)>td, .ant-table-tbody>tr:hover:not(.ant-table-expanded-row)>td, .ant-table-thead>tr.ant-table-row-hover:not(.ant-table-expanded-row)>td, .ant-table-thead>tr:hover:not(.ant-table-expanded-row)>td {
  background: #87CEFB !important;
}
</style>
