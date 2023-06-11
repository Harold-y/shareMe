<script>
import { ArchiveOutline as ArchiveIcon } from '@vicons/ionicons5'
import axios from "axios";
import {useMessage} from "naive-ui";

export default {
  props: {
    type: String
  },
  data() {
    return {
      theType: this.type,
      noteVal: null,
      status: null,
      shareCode: null,
      shareFiles: new FormData(),
      successShow: false,
      warningShow: false,
      errors: null,
      URLVal: null
    }
  },
  components: {
    ArchiveIcon
  },
  methods: {
    changeVal: function (val) {
      this.theType = val
    },
    addShareFile: function ({file}) {
      this.shareFiles.append("fileList", file.file)
    },
    shareFile: function () {
      axios
          .post( this.BASE_URL+ "add", this.shareFiles, {
            headers: {
              'Content-Type': 'multipart/form-data',
            }
          })
          .then((response) => {
            // console.log(response)
            this.shareCode = response.data.shareCode
            this.successShow = true
            setTimeout(() => {this.successShow = false},10000);
            this.$refs['fileUploader'].clear()
          })
          .catch(function (error) { // 请求失败处理
            // console.log(error);
            this.warningShow = true
            this.errors = error
            setTimeout(() => {this.warningShow = false},6000);
          });
    },
    shareURL: function () {
      axios
          .post(this.BASE_URL + "u/add", {
            URL: this.URLVal
          })
          .then((response) => {
            // console.log(response)
            this.shareCode = response.data.shareCode
            this.successShow = true
            setTimeout(() => {this.successShow = false},10000);
            this.URLVal = ""
          })
          .catch(function (error) { // 请求失败处理
            this.warningShow = true
            this.errors = error
            setTimeout(() => {this.warningShow = false},6000);
          });
    },
    shareNote: function () {
      axios
          .post(this.BASE_URL + "n/add", {
            'Note': this.noteVal
          })
          .then((response) => {
            // console.log(response)
            this.shareCode = response.data.shareCode
            this.successShow = true
            setTimeout(() => {this.successShow = false},10000);
            this.noteVal = ""
          })
          .catch(function (error) { // 请求失败处理
            this.warningShow = true
            this.errors = error
            setTimeout(() => {this.warningShow = false},6000);
          });

    }
  },
  watch: {
    type: function(newVal, oldVal) { // watch it
      this.theType = newVal
    }
  }
}
</script>

<template>
  <div id="container1">
    <n-space vertical id="space1" >
      <n-alert title="Success 成功！" type="success" style="border-radius: 15px; margin-bottom: 1rem" v-show="successShow">
        shareCode 分享代码: {{ shareCode }}
      </n-alert>

      <n-alert title="Failed 失败！" type="error" style="border-radius: 15px; margin-bottom: 1rem" v-show="warningShow">
        {{ errors }}
      </n-alert>

      <n-card title="Upload 上传" style="margin-bottom: 16px; border-radius: 15px">
        <n-tabs type="line" :default-value="this.theType" :value="this.theType" :on-update:value="changeVal" animated>
          <n-tab-pane name="file" tab="File 文件">
            <n-upload
                ref="fileUploader"
                multiple
                default-upload: false
                :custom-request="addShareFile"
                max="15"
            >
              <n-upload-dragger>
                <div style="margin-bottom: 12px">
                  <n-icon size="48" :depth="3">
                    <archive-icon />
                  </n-icon>
                </div>
                <n-text style="font-size: 16px">
                  Click or Drag File to This Area <br/>点击或者拖动文件到该区域来上传
                </n-text>
                <n-p depth="3" style="margin: 8px 0 0 0">
                  请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码
                </n-p>
              </n-upload-dragger>
            </n-upload>
            <n-button color="#74D2E7" id="bt1" round="round" size="large" style="margin-top: 1rem" @click="shareFile">
              Upload 上传
            </n-button>
          </n-tab-pane>

          <n-tab-pane name="URL" tab="URL 链接">
            <n-input round placeholder="URL 链接，begin with http:// or https:// " v-model:value="URLVal"/>
            <n-button color="#74D2E7" id="bt1" round="round" size="large" style="margin-top: 1rem" @click="shareURL">
              Upload 上传
            </n-button>
          </n-tab-pane>
          <n-tab-pane name="note" tab="Note 笔记">
            <n-input
                v-model:value="noteVal"
                type="textarea"
                placeholder="Note"
                autosize="true"
            />
            <n-button color="#74D2E7" id="bt1" round="round" size="large" style="margin-top: 1rem" @click="shareNote">
              Upload 上传
            </n-button>
          </n-tab-pane>
        </n-tabs>
      </n-card>
    </n-space>

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