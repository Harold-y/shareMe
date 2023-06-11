<script>
import axios from "axios";
import router from "../router/index.js";

export default {
  data() {
    return {
      shareCode: null,
      warningShow: false,
      errors: null
    }
  },
  methods: {
    queryShareCode: function ()
    {
      this.shareCode = this.shareCode.toUpperCase()
      axios
          .get( this.BASE_URL+ "get/" + this.shareCode)
          .then((response) => {
            // console.log(response)

            const data = response.data;
            if (data.files) {
              router.push('/f/' + this.shareCode)
            }
            if (data.link) {
              window.open(data.link, '_blank');
            }
            if (data.note) {
              router.push('/n/' + this.shareCode)
            }
            this.shareCode = ""
          })
          .catch(function (error) { // 请求失败处理
            // console.log(error);
            this.warningShow = true
            this.errors = error
            setTimeout(() => {this.warningShow = false},6000);
          });
    }
  }
}

</script>

<template>
  <div id="container1">
    <n-alert title="Failed 失败！" type="error" style="border-radius: 15px; margin-bottom: 1rem" v-show="warningShow">
      {{ errors }}
    </n-alert>
    <n-card title="Home 主页" style="margin-bottom: 16px; border-radius: 15px">
      <n-space vertical id="space1">


        <n-input size="large" round placeholder="ShareCode 分享代码" clearable="clearable" maxlength="4" v-model:value="shareCode"/>
        <div style="margin-top: 3rem">
          <n-button color="#74D2E7" id="bt1" round="round" size="large" @click="queryShareCode">
            Query 查询
          </n-button>
        </div>
      </n-space>
    </n-card>

  </div>

</template>



<style scoped>
#bt1 {
  display: block;
  margin: 0 auto;
}
#container1 {
  padding: 5rem;
  border-radius: 2rem;
  max-width: 50vw;
  max-height: 55vh;
  margin: 10vh auto 0;
}
</style>
<script setup>
</script>
