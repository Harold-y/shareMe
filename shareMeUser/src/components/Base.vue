
<template>
  <div id="cover">

  </div>
  <n-space vertical size="large" >
    <n-layout>
      <n-layout-header id="header1">
        <img src="../assets/logo2.png" id="img1"/>
        <n-gradient-text :size="20" :gradient="{
          from: 'rgb(255 255 255)',
          to: 'rgb(245 245 245)'
        }" id="shareMeText">ShareME</n-gradient-text>
        <n-menu id="menu1" mode="horizontal" :options="menuOptions" />
      </n-layout-header>
    </n-layout>
  </n-space>

</template>

<style>
body {
  background-image: url('../assets/DSC05891.jpg');
  background-size: cover;
  position: relative;
  height: 100vh;
  width: 100vw;
  background-position: center;
  background-attachment: fixed;
  background-repeat: no-repeat;
  top: 0;
}
#header1 {
  font-size: 20px;
  background-color: rgba(0, 0, 0, 0.16);
  position: fixed;
}
#img1 {
  width: 50px;
  border-radius: 30px;
  vertical-align: middle;
  margin-top: 0.2rem;
  margin-left: 0.2rem;
}
#shareMeText {
  vertical-align:middle; margin-left: 5px
}
.n-menu .n-menu-item-content .n-menu-item-content-header {
  color: white
}
#menu1 {
  vertical-align:middle;
}

</style>

<script lang="ts">
import router from "../router"
import { defineComponent, h, ref, Component } from 'vue'
import { NIcon } from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import {
  AddCircleOutline as AddIcon,
  DocumentsOutline as FileIcon,
  UnlinkOutline as LinkIcon,
  PencilOutline as WriteIcon,
  HomeOutline as HomeIcon
} from '@vicons/ionicons5'
import {RouterLink} from "vue-router";
function renderIcon (icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions: MenuOption[] = [
  {
    label: () =>
        h(
            RouterLink,
            {
              to: {
                name: 'home',
                params: {
                  lang: 'zh-CN'
                }
              }
            },
            { default: () => 'Home 主页' }
        ),
    key: 'home',
    icon: renderIcon(HomeIcon)
  },
  {
    label: 'Upload 添加',
    key: 'dance-dance-dance',
    icon: renderIcon(AddIcon),
    children: [
      {
        type: 'group',
        label: 'File 文件',
        key: 'file',
        children: [
          {
            label: () =>
                h(
                    RouterLink,
                    {
                      to: {
                        name: 'upload',
                        params: {
                          'type': 'file'
                        }
                      }
                    },
                    { default: () => 'Upload 上传' }
                ),
            key: 'fileUp',
            icon: renderIcon(FileIcon),
          }
        ]
      },
      {
        type: 'group',
        label: 'Other 其它',
        key: 'other',
        children: [
          {
            label: () =>
                h(
                    RouterLink,
                    {
                      to: {
                        name: 'upload',
                        params: {
                          'type': 'URL'
                        }
                      }
                    },
                    { default: () => 'URL 链接' }
                ),
            key: 'url',
            icon: renderIcon(LinkIcon),
          },
          {
            label: () =>
                h(
                    RouterLink,
                    {
                      to: {
                        name: 'upload',
                        params: {
                          type: 'note'
                        }
                      }
                    },
                    { default: () => 'Note 笔记' }
                ),
            key: 'note',
            icon: renderIcon(WriteIcon),
          }
        ]
      },
    ]
  }
]

export default defineComponent({
  setup () {
    return {
      activeKey: undefined,
      menuOptions
    }
  }
})
</script>