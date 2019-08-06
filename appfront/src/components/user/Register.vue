<template>
    <div id="content">
        <a-form
        :form="form"
        @submit="handleSubmit"
        >
            <a-form-item
            v-bind="formItemLayout"
            label="Username"
            >
            <a-input
            v-decorator="[
            'username',
            { rules: [{ required: true, message: 'Please input your username!' }] }
            ]"
            placeholder="Username"
            >
            </a-form-item>
            <a-form-item
            v-bind="formItemLayout"
            label="Password"
            >
            <a-input
            v-decorator="[
            'password',
            {
                rules: [{
                required: true, message: 'Please input your password!',
                }, {
                validator: validateToNextPassword,
                }],
            }
            ]"
            type="password"
            />
            </a-form-item>
            <a-form-item
            v-bind="formItemLayout"
            label="Confirm Password"
            >
                <a-input
                v-decorator="[
                'confirm',
                {
                    rules: [{
                    required: true, message: 'Please confirm your password!',
                    }, {
                    validator: compareToFirstPassword,
                    }],
                }
                ]"
                type="password"
                @blur="handleConfirmBlur"
                />
            </a-form-item>
            <a-form-item v-bind="tailFormItemLayout">
                <a-button
                    type="primary"
                    html-type="submit"
                >
                    Register
                </a-button>
            </a-form-item>
        </a-form>
    </div>
</template>

<script>
import axios from 'axios'
import Cookies from 'js-cookie';

let csrftoken = Cookies.get('csrftoken');
export default {
    data () {
        return {
            confirmDirty: false,
            formItemLayout: {
                labelCol: { span: 24 },
                wrapperCol: { span: 24 },
            },
            tailFormItemLayout: {
                wrapperCol: {
                xs: {
                    span: 24,
                    offset: 0,
                    },
                sm: {
                    span: 16,
                    offset: 8,
                    },
                },
            },
        };
    },

    beforeCreate() {
        this.form = this.$form.createForm(this);
    },

    methods: {
        handleConfirmBlur (e) {
            const value = e.target.value;
            this.confirmDirty = this.confirmDirty || !!value;
        },
        compareToFirstPassword (rule, value, callback) {
            const form = this.form;
            if (value && value !== form.getFieldValue('password')) {
                callback('Two passwords that you enter is inconsistent');
            } else {
                callback()
            }
        },
        validateToNextPassword (rule, value, callback) {
            const form = this.form;
            if (value && this.confirmDirty) {
                form.validateFields(['confirm'], { force: true });
            }
            callback();
        },
        handleSubmit (e) {
            e.preventDefault();
            this.form.validateFieldsAndScroll((err, values) => {
                if (!err) {
                    const formData = new FormData();
                    formData.append('username', values.username);
                    formData.append('password', values.password);
                    axios({
                        url: '/register',
                        method: 'post',
                        data: formData,
                        headers: {
                            'X-CSRFToken': csrftoken,
                        },
                    }).then((res) => {
                        this.$message.success(res.data.msg);
                        this.$router.push({name: 'login'});
                    }).catch(err => {
                        this.$message.error("注册失败");
                    })
                } else {
                    this.$message.error('注册信息错误');
                    return false;
                }
            })
        },
    },
    
}
</script>