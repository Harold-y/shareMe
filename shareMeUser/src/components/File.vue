<script>

import axios from "axios";
import router from "../router/index.js";
import { DownloadOutline as DownloadIcon, EyeOutline as EyeIcon, ImageOutline as ImageIcon, DocumentTextOutline as TextIcon, Code as CodeIcon, FileTrayOutline as OtherIcon } from '@vicons/ionicons5'
export default {
  props: ['shareCode'],
  data() {
    return {
      shareCode_: this.shareCode,
      files: this.note,
      imgTypes: ['jpg', 'jpeg', 'gif', 'webp', 'tiff', 'png'],
      textTypes: ['txt', 'md', 'csv', 'doc', 'docs', 'docx', 'pdf'],
      programTypes: ['html', 'py', 'go', 'vue', 'java', 'cpp', 'c', 'json', 'yaml', 'r', 'rmd', 'ipynb', 'sql'],
      allTypes: ['jpg', 'jpeg', 'gif', 'webp', 'tiff', 'png', 'txt', 'md', 'csv', 'doc', 'docs', 'docx', 'html',
        'py', 'go', 'vue', 'java', 'cpp', 'c', 'json', 'yaml', 'r', 'rmd', 'ipynb', 'sql', 'pdf'],
      showModal: false,
      previewType: null,
      previewImageSrc: null,
      prevText: null,
      prevTextBr: null,
    }
  },
  components: {
    DownloadIcon,
    EyeIcon,
    ImageIcon,
    TextIcon,
    CodeIcon,
    OtherIcon
  },
  methods: {

    prevView: async function (fileName) {
      let blob
      let url
      axios
          .get(this.BASE_URL + "download/" + this.shareCode_ + "/" + fileName, {responseType: "blob"})
          .then(async (response) => {
            blob = new Blob([response.data]);
            url = window.URL.createObjectURL(blob);

            const extension = fileName.split('.').at(-1).toLowerCase()
            if (this.imgTypes.includes(extension)) {
              this.previewType = "img"
              this.previewImageSrc = url
              this.showModal = true
            } else if (this.textTypes.includes(extension)) {
              if (extension === 'pdf') {
                this.previewType = "pdf"
                blob = new Blob([response.data], { type: "application/pdf" });
                url = window.URL.createObjectURL(blob);
                window.open(url, '_blank');
              } else {
                this.previewType = "text"
                this.prevText = await blob.text()
                this.prevTextBr = this.prevText.replace(/\n/g, '<br>')
                this.showModal = true
              }
            } else if (this.programTypes.includes(extension)) {
              this.previewType = "program"
              this.prevText = await blob.text()
              this.showModal = true

            } else {
              this.previewType = "No Preview Available 无法预览该文件"
              this.showModal = true
            }
          })
          .catch(function (error) { // 请求失败处理
            // console.log(error);
          });


    },
    download: function (fileName) {
      axios
          .get( this.BASE_URL + "download/" + this.shareCode_ + "/" + fileName, {responseType: "blob"})
          .then((response) => {
            let url = window.URL.createObjectURL(new Blob([response.data]));
            // 生成一个a标签
            let link = document.createElement("a");
            link.style.display = "none";
            link.href = url;
            link.download = fileName;
            document.body.appendChild(link);
            link.click();
          })
          .catch(function (error) { // 请求失败处理
            // console.log(error);
          });
    },
    downloadZIP: async function () {
      await axios
          .get( this.BASE_URL + "download/"+ "all/" + this.shareCode_ , {responseType: "blob"})
          .then((response) => {
            let url = window.URL.createObjectURL(new Blob([response.data]));
            // 生成一个a标签
            let link = document.createElement("a");
            link.style.display = "none";
            link.href = url;
            link.download = "target.tar.gz";
            document.body.appendChild(link);
            link.click();
          })
          .catch(function (error) { // 请求失败处理
            // console.log(error);
          });
    }
  },computed: {
  },
  mounted() {
    axios
        .get( this.BASE_URL + this.shareCode)
        .then((response) => {
          // console.log(response)
          this.files = response.data.files;
        })
        .catch(function (error) { // 请求失败处理
          // console.log(error);
        });

  }
}
</script>

<template>
  <div id="container1">
    <n-modal v-model:show="showModal">
      <n-card
          style="width: 40vw"
          title="Preview 预览"
          :bordered="false"
          size="huge"
          role="dialog"
          aria-modal="true"
      >
        <template #header-extra>
          {{ previewType }}
        </template>
        <n-image
            width="200"
            :src="previewImageSrc"
            v-show="previewType === 'img'"
        />
        <n-p v-show="previewType === 'text' || previewType === 'program'" v-html="prevTextBr">

        </n-p>
      </n-card>
    </n-modal>
    <n-card title="File 文件" style="margin-bottom: 16px; border-radius: 15px">
      <n-space vertical id="space1">
        <n-h3>{{ this.shareCode_ }}</n-h3>
        <n-table striped>
          <thead>
          <tr>
            <th>File Name 文件名称</th>
            <th>Extension 拓展名</th>
            <th>Action 操作</th>
          </tr>
          </thead>
          <tbody v-for="file in files" id="files">
          <tr>
            <td>{{ file.split('.').at(0) }}</td>
            <td>{{ file.split('.').at(-1) }}
              <n-icon style="vertical-align: middle" v-if="imgTypes.includes(file.split('.').at(-1).toLowerCase())"><image-icon ></image-icon></n-icon>
              <n-icon style="vertical-align: middle" v-if="textTypes.includes(file.split('.').at(-1).toLowerCase())"><text-icon ></text-icon></n-icon>
              <n-icon style="vertical-align: middle" v-if="programTypes.includes(file.split('.').at(-1).toLowerCase())"><code-icon ></code-icon></n-icon>
              <n-icon style="vertical-align: middle" v-if="!allTypes.includes(file.split('.').at(-1).toLowerCase())"><other-icon ></other-icon></n-icon>
            </td>
            <td>
              <n-button circle size="small" @click="download(file)">
                <template #icon>
                  <n-icon><download-icon /></n-icon>
                </template>
              </n-button>
              <n-button circle size="small" @click="prevView(file)" style="margin-left: 10px" v-if="allTypes.includes(file.split('.').at(-1).toLowerCase())">
                <template #icon>
                  <n-icon><eye-icon /></n-icon>
                </template>
              </n-button>
            </td>
          </tr>
          </tbody>
        </n-table>
        <n-button strong secondary round type="info" style="margin-top: 1rem" @click="downloadZIP">
          <n-icon><download-icon /></n-icon>Pack Download 打包下载
        </n-button>
      </n-space>
    </n-card>

  </div>
</template>

<style scoped>
#files {
  cursor: pointer;
}
#container1 {
  padding: 5rem;
  border-radius: 2rem;
  max-width: 50vw;
  max-height: 55vh;
  margin: 10vh auto 0;
}
</style>