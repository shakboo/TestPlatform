<template>

  <!--MySQL-->
  <a-form
    id="components-form-demo-validate-other"
    :form="form"
    @submit="handleSubmit"
    v-if="formType === 'MySQL'"
  >
        <a-form-item
        disabled
        v-bind="formItemLayout"
        label="上传到"
        >
            <a-select 
            default-value="MySQL"
            v-decorator="[
            'selectType',
            {rules: [{ required: true, message: '请选择上传的数据库类型' }]}
            ]">
                <a-select-option value="MySQL" @click="formType='MySQL'">
                MySQL
                </a-select-option>
            </a-select>
        </a-form-item>
        <a-form-item
        label="IP"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-input
                v-decorator="[
                'ip',
                {rules: [{ required: true, message: '请输入地址' }]}
                ]"
            />
        </a-form-item>
        <a-form-item
        label="Port"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-input-number
                v-decorator="[
                'port',
                {rules: [{ required: true, message: '请输入端口' }]}
                ]"
            />
        </a-form-item>
        <a-form-item
        label="用户名"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-input
                v-decorator="[
                'user',
                {rules: [{ required: true, message: '请输入用户名' }]}
                ]"
            />
        </a-form-item>
        <a-form-item
        label="密码"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-input
                v-decorator="[
                'password',
                {rules: [{ required: true, message: '请输入密码' }]}
                ]"
                type="password"
            />
        </a-form-item>
        <a-form-item
        label="数据库名"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-input
                v-decorator="[
                'dbName',
                {rules: [{ required: true, message: '请输入数据库名' }]}
                ]"
            />
        </a-form-item>
        <a-form-item
        label="表名"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-input
                v-decorator="[
                'tableName',
                {rules: [{ required: true, message: '请输入表名' }]}
                ]"
            />
        </a-form-item>
        <!--这个为什么不行阿，真是fuck
        <a-form-item
        v-bind="formItemLayout"
        label="上传文件"
        >
            <a-upload
                :fileList="fileList"
                v-decorator="[
                'uploadFile',
                {rules: [{ required: true, message: '请选择上传的文件'}]}
                ]"
                :remove="handleRemove"
                :beforeUpload="beforeUpload"
                @change="handleChange"
            >
                <a-button>
                <a-icon type="upload" /> Select File
                </a-button>
            </a-upload>
        </a-form-item>
        -->
        <a-form-item
            :wrapper-col="{ span: 12, offset: 5 }"
            >
            <input type="file" id="file" required accept=".csv" />
        </a-form-item>
        <a-form-item
            :wrapper-col="{ span: 12, offset: 5 }"
            >
            <a-button
                type="primary"
                html-type="submit"
                :loading="uploading"
            >
                提交
            </a-button>
        </a-form-item>
    </a-form>
</template>

<script>
import reqwest from 'reqwest'
import Cookies from 'js-cookie';

let csrftoken = Cookies.get('csrftoken');
export default {
  data: () => ({
    fileList: [],
    uploading: false,
    formItemLayout: {
      labelCol: { span: 5 },
      wrapperCol: { span: 12 },
    },
    formType: 'MySQL',
  }),
  beforeCreate () {
    this.form = this.$form.createForm(this);
  },
  methods: {
    beforeUpload(file) {
        return false;
    },
    handleRemove(file) {
        this.fileList = [];
    },
    handleSubmit  (e) {
        e.preventDefault();
        this.form.validateFields((err, values) => {
            if (!err) {
                const formData = new FormData();
                this.uploading = true;
                (Object.keys(values)).forEach((key) => {
                    if (key != "uploadFile") {
                        formData.append(key, values[key])
                    }
                });
                // 同以上fuck
                // formData.append('file', this.fileList[0])
                formData.append('file', document.getElementById("file").files[0]);
                reqwest({
                    url: '/api/v1/tool/upload',
                    method: 'post',
                    processData: false,
                    contentType: false,
                    data: formData,
                    headers: {
                        'X-CSRFToken': csrftoken,
                    },
                    success: (res) => {
                        // this.fileList = [];
                        this.uploading = false;
                        if (res.msg == "success"){
                            this.$message.success('上传成功！');
                        } else {
                            this.$message.error("上传时出现了错误！");
                        }
                    }
                })   
            }
        });
    },
    handleChange(info) {
        let fileList = [...info.fileList];
        fileList = fileList.slice(-1);
        this.fileList = fileList
    },
  }
}
</script>