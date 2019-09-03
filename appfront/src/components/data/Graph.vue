<template>
    <a-form
    :form="form"
    @submit="handleSubmit">
        <a-form-item
        label="节点数"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-input-number
            :min="1"
            :max="10000"
            v-decorator="[
                'node',
                {rules: [{ required: true, message: '请输入节点数!' }]}
            ]"
            />
        </a-form-item>
        <a-form-item
        label="类型"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 12 }"
        >
            <a-select
                v-decorator="[
                'graphType',
                {rules: [{ required: true, message: '请选择生成图的类型!' }]}
                ]"
                placeholder="Select a option"
            >
                <a-select-option value="unweight">
                无权无向图
                </a-select-option>
                <a-select-option value="weight">
                有权无向图
                </a-select-option>
            </a-select>
        </a-form-item>
        <a-form-item
            :wrapper-col="{ span: 12, offset: 5 }"
            >
            <a-button
                type="primary"
                html-type="submit"
                :loading="uploading"
            >
                生成
            </a-button>
        </a-form-item>
    </a-form>
</template>

<script>
import Cookies from 'js-cookie';
import axios from 'axios';

let csrftoken = Cookies.get('csrftoken');
export default {
    data () {
        return {
        formLayout: 'horizontal',
        form: this.$form.createForm(this),
        uploading: false,
        };
    },
    methods: {
        handleSubmit(e) {
            e.preventDefault();
            this.form.validateFields((err, values) => {
                if(!err) {
                    console.log('Received values of form: ', values);
                    this.uploading = true;
                    axios({
                        url: '/api/v1/data/graph',
                        method: 'get',
                        headers: {
                            'X-CSRFToken': csrftoken,
                            'Accept': 'application/octet-stream'
                        },
                        params: {
                            ...values,
                        },
                        responseType: 'blob',
                    }).then((res) => {
                        this.uploading = false;
                        const fileName = 'result.csv';
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
                    })
                }
            });
        }
    }
}
</script>