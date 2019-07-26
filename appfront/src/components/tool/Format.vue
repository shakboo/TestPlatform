<template>
  <a-form
    id="components-form-demo-validate-other"
    :form="form"
    @submit="handleSubmit"
  >
        <a-form-item
        disabled
        v-bind="formItemLayout"
        label="转换为"
        >
            <a-select 
            v-decorator="[
            'targetType',
            {rules: [{ required: true, message: '请选择要转换的格式' }]}
            ]">
                <a-select-option value="csv" @click="isTargetXsv=false">
                csv
                </a-select-option>
                <a-select-option value="tsv" @click="isTargetXsv=false">
                tsv
                </a-select-option>
                <a-select-option value="xsv" @click="isTargetXsv=true">
                xsv
                </a-select-option>
                <a-select-option value="txt" @click="isTargetXsv=false">
                txt
                </a-select-option>
            </a-select>
        </a-form-item>
        <a-form-item
            v-if="isTargetXsv"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
            label="xsv分隔符"
            >
            <a-input
                v-decorator="[
                'xsvSplit',
                {rules: [{ required: true, message: '请输入xsv的分隔符' }]}
                ]"
            />
        </a-form-item>
        <a-form-item
            :wrapper-col="{ span: 12, offset: 5 }"
            >
            <input type="file" id="file" required @change="updateSourceXsvSplit" />
        </a-form-item>
        <a-form-item
            v-if="isUploadXsv"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
            label="源文件xsv分隔符"
            >
            <a-input
                v-decorator="[
                'sourceXsvSplit',
                {rules: [{ required: true, message: '请输入源文件的xsv的分隔符' }]}
                ]"
            />
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
    </form>
</template>

<script>
import axios from 'axios'
import Cookies from 'js-cookie';

let csrftoken = Cookies.get('csrftoken');
export default {
  data: () => ({
    uploading: false,
    isTargetXsv: false,
    isUploadXsv: false,
    formItemLayout: {
      labelCol: { span: 5 },
      wrapperCol: { span: 12 },
    },
  }),
  computed: {
      xsvSplit () {
          return "x";
      }
  },
  watch: {
      xsvSplit (val) {
          this.form.setFieldsValue({xsvSplit: val});
      }
  },
  beforeCreate () {
    this.form = this.$form.createForm(this);
  },
  methods: {
      updateSourceXsvSplit (e) {
            const uploadFileName = document.getElementById('file').files[0].name;
            const fileNameArray = uploadFileName.split('.');
            if (fileNameArray[fileNameArray.length - 1] === "xsv") {
                this.isUploadXsv = true;
            } else {
                this.isUploadXsv = false;
            }
      },
      handleSubmit  (e) {
        e.preventDefault();
        this.form.validateFields((err, values) => {
            if (!err) {
                let allowFileArray = new Array("csv", "tsv", "xsv", "txt");
                const uploadFileName = document.getElementById('file').files[0].name;
                const fileNameArray = uploadFileName.split('.');
                if (!allowFileArray.includes(fileNameArray[fileNameArray.length - 1])) {
                    this.$message.error("请上传csv，tsv，xsv，xlsx，txt格式的文件");
                    return false;
                }
                const formData = new FormData();
                this.uploading = true;
                (Object.keys(values)).forEach((key) => {
                    formData.append(key, values[key])
                });
                formData.append('file', document.getElementById("file").files[0]);
                formData.append('fileName', fileNameArray.slice(0, fileNameArray.length - 1).join('.'));
                formData.append('fileType', fileNameArray[fileNameArray.length - 1]);
                axios({
                    url: '/api/v1/tool/format',
                    method: 'post',
                    data: formData,
                    headers: {
                        'X-CSRFToken': csrftoken,
                    },
                    responseType: 'blob',
                }).then((res) => {
                    this.uploading = false;
                    this.$message.success("转换成功");
                    const fileName = fileNameArray.slice(0, fileNameArray.length - 1).join('.') + "." + this.form.getFieldValue("targetType");
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
                        navigator.msSaveBlob(blob, fileName);
                    }
                })
            }
        })
      },
  }
}
</script>