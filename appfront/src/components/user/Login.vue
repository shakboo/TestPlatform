<template>
    <div id="content">
        <a-form
            id="components-form-demo-normal-login"
            :form="form"
            class="login-form"
            @submit="handleSubmit"
        >
            <a-form-item>
            <a-input
                v-decorator="[
                'username',
                { rules: [{ required: true, message: 'Please input your username!' }] }
                ]"
                placeholder="Username"
            >
                <a-icon
                slot="prefix"
                type="user"
                style="color: rgba(0,0,0,.25)"
                />
            </a-input>
            </a-form-item>
            <a-form-item>
            <a-input
                v-decorator="[
                'password',
                { rules: [{ required: true, message: 'Please input your Password!' }] }
                ]"
                type="password"
                placeholder="Password"
            >
                <a-icon
                slot="prefix"
                type="lock"
                style="color: rgba(0,0,0,.25)"
                />
            </a-input>
            </a-form-item>
            <a-form-item>
            <a-checkbox
                v-decorator="[
                'remember',
                {
                    valuePropName: 'checked',
                    initialValue: true,
                }
                ]"
            >
                Remember me
            </a-checkbox>
            <a
                class="login-form-forgot"
                href=""
            >
                Forgot password
            </a>
            <a-button
                type="primary"
                html-type="submit"
                class="login-form-button"
            >
                Log in
            </a-button>
            Or <a href="javascript:void(0);" @click="gotoRegister">
                register now!
                </a>
            </a-form-item>
        </a-form>
    </div>
</template>

<script>
import axios from 'axios'
import Cookies from 'js-cookie';

import { mapMutations } from 'vuex';

let csrftoken = Cookies.get('csrftoken');
export default {
  beforeCreate () {
    this.form = this.$form.createForm(this);
  },
  methods: {
    ...mapMutations(['changeLogin']),
    handleSubmit (e) {
      e.preventDefault();
      let _this = this;
      this.form.validateFields((err, values) => {
        if (!err) {
          const formData = new FormData();
          formData.append('username', values.username);
          formData.append('password', values.password);
          axios({
            url: '/auth',
            method: 'post',
            data: formData,
            headers: {
                'X-CSRFToken': csrftoken,
            },
            }).then((res) => {
                if (res.data.code === 200) {
                  this.$message.success(res.data.msg);
                } else {
                  this.$message.error(res.data.msg);
                }
                _this.userToken = res.data.data.token;
                _this.changeLogin({ Authorization: _this.userToken, username: values.username });
                let redirect = this.$route.query.redirect;
                redirect ? this.$router.push(redirect) : this.$router.push({name: 'testcase-ml'});
            }).catch(err => { //请求失败就会捕获报错信息
                console.log(err);
                this.$message.error("登录失败");
            })
        } else {
          this.$message.error('用户信息错误');
          return false;
        }
      })
    },
    gotoRegister () {
        this.$router.push({name: 'register'});
    },
  },
};
</script>
<style>
html,body {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
}
#components-form-demo-normal-login .login-form {
    max-width: 500px;
}
#components-form-demo-normal-login .login-form-forgot {
    float: right;
}
#components-form-demo-normal-login .login-form-button {
    width: 100%;
}
#content {
    position:absolute;
    top:50%;
    left:50%;
    transform:translate(-50%,-50%);
}
</style>