<script>

import axios from "axios";
import {DocumentOutline as DocumentIcon} from '@vicons/ionicons5'
export default {
  props: {
    shareCode: String,
  },
  data() {
    return {
      shareCode_: this.shareCode,
      note_: this.note,
    }
  },
  methods: {
    downloadNote: function () {
      const blob = new Blob([this.note_], {type:"text/plain;charset=utf-8"})
      let url = window.URL.createObjectURL(blob);
      // 生成一个a标签
      let link = document.createElement("a");
      link.style.display = "none";
      link.href = url;
      link.download = "note_" + this.shareCode + ".txt";
      document.body.appendChild(link);
      link.click();
    }
  },
  components: {
    DocumentIcon
  },
  mounted() {
    axios
        .get( this.BASE_URL + "n/" + this.shareCode)
        .then((response) => {
          // console.log(response)
          this.note_ = response.data.note;
        })
        .catch(function (error) { // 请求失败处理
          // console.log(error);
        });

  }
}
</script>

<template>
  <div id="container1">
    <n-card title="Note 笔记" style="margin-bottom: 16px; border-radius: 15px">
      <n-space vertical id="space1">
        <n-h1>{{ this.shareCode_ }}</n-h1>
        <n-button strong secondary round type="info" style="float: right; margin: 1rem" @click="downloadNote">
          <n-icon><document-icon /></n-icon>Download 下载笔记
        </n-button>
        <n-p>
          {{ this.note_ }}
        </n-p>
      </n-space>
    </n-card>

  </div>
</template>

<style scoped>
#container1 {
  padding: 5rem;
  border-radius: 2rem;
  max-width: 50vw;
  max-height: 55vh;
  margin: 10vh auto 0;
}
</style>